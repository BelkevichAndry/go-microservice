package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/BelkevichAndry/go-microservice/shippy-user-service/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func createUser(ctx context.Context, service micro.Service, user *proto.User) error {
	client := proto.NewUserService("shippy.service.user", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	fmt.Println("Response: ", rsp.User)

	return nil
}

func main() {
	service := micro.NewService()
	service.Init(micro.Action(func(c *cli.Context) error {
		name := c.String("name")
		email := c.String("email")
		company := c.String("company")
		password := c.String("password")

		ctx := context.Background()
		user := &proto.User{
			Name:     name,
			Email:    email,
			Company:  company,
			Password: password,
		}

		if err := createUser(ctx, service, user); err != nil {
			log.Println("error creating user: ", err.Error())
			return err
		}

		return nil
	}))
}
