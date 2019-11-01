package grpchttp

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4/middleware"

	"google.golang.org/grpc"
	"v2ray.com/core/app/proxyman/command"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/common/uuid"
	"v2ray.com/core/proxy/vmess"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo/v4"
	handleService "v2ray.com/core/app/proxyman/command"
	statsService "v2ray.com/core/app/stats/command"
)

type Email struct {
	Email string `json:"email"`
}
type UserRequest struct {
	Emails []Email `json:"emails"`
	Tag    string  `json:"tag"`
}
type UserAddResponse struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	AlterID int    `json:"alterId"`
	Error   string `json:"error"`
}
type UserRemoveResponse struct {
	Email   string `json:"email"`
	Removed bool   `json:"removed"`
	Error   string `json:"error"`
}
type Fail struct {
	Error string `json:"error"`
}
type DataResponse struct {
	Email string
	Down  string
}

func HttpRouter() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	e.POST("/users", func(c echo.Context) error {
		var ur UserRequest
		if err := c.Bind(&ur); err != nil {
			return c.JSON(400, Fail{"创建用户参数不对，请检查"})
		}
		if len(ur.Emails) == 0 || ur.Tag == "" {
			return c.JSON(400, Fail{"创建用户参数不对，请检查"})
		}
		conn, err := grpc.DialContext(context.Background(), os.Getenv("grpchost"), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return c.JSON(400, Fail{"创建用户失败：" + err.Error()})
		}
		client := handleService.NewHandlerServiceClient(conn)
		ure := make([]UserAddResponse, len(ur.Emails))
		for k, v := range ur.Emails {
			u := protocol.NewID(uuid.New())
			_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
				Tag: ur.Tag,
				Operation: serial.ToTypedMessage(
					&command.AddUserOperation{
						User: &protocol.User{
							Email: v.Email,
							Account: serial.ToTypedMessage(&vmess.Account{
								Id:      u.String(),
								AlterId: 64,
							}),
						},
					}),
			})
			ure[k].Email = v.Email
			if err != nil {
				ure[k].Error = err.Error()
				continue
			}
			ure[k].AlterID = 64

			ure[k].Id = u.String()
		}
		return c.JSON(200, ure)
	})

	e.DELETE("/users", func(c echo.Context) error {
		var ur UserRequest
		if err := c.Bind(&ur); err != nil {
			return c.JSON(400, Fail{"删除用户参数不对，请检查"})
		}
		if len(ur.Emails) == 0 || ur.Tag == "" {
			return c.JSON(400, Fail{"删除用户参数不对，请检查"})
		}
		conn, err := grpc.DialContext(context.Background(), os.Getenv("grpchost"), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return c.JSON(400, Fail{"删除用户失败：" + err.Error()})
		}
		urs := make([]UserRemoveResponse, len(ur.Emails))
		client := handleService.NewHandlerServiceClient(conn)
		for k, v := range ur.Emails {
			_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
				Tag:       ur.Tag,
				Operation: serial.ToTypedMessage(&command.RemoveUserOperation{Email: v.Email}),
			})
			urs[k].Email = v.Email
			if err != nil {
				urs[k].Removed = false
				urs[k].Error = err.Error()
				continue
			}

			urs[k].Removed = true
		}

		return c.JSON(200, urs)
	})
	e.GET("/datas/:email", func(c echo.Context) error {
		var d DataResponse
		conn, err := grpc.DialContext(context.Background(), os.Getenv("grpchost"), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return c.JSON(400, Fail{"获取用户流量数据失败：" + err.Error()})
		}
		client := statsService.NewStatsServiceClient(conn)
		em := c.Param("email")
		r := &statsService.QueryStatsRequest{Pattern: em + ">>>traffic>>>downlink", Reset_: false}
		resp, err := client.QueryStats(context.Background(), r)
		if err != nil {
			return c.JSON(400, Fail{"获取用户流量失败：" + err.Error()})
		}
		s := proto.MarshalTextString(resp)
		fmt.Println(s)
		if len(s) < 1 {
			return c.JSON(400, Fail{"获取用户流量失败：" + err.Error()})
		}
		sv := regexp.MustCompile(`value:.*`).FindString(s)
		v := strings.Split(sv, ":")
		d.Email = em
		if len(v) > 1 {
			d.Down = v[1]
		}
		return c.JSON(200, d)
	})
	e.Start(":1323")
}
