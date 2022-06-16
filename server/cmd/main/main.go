/*
Package main gin-nuxt-server

Swagger Reference: https://github.com/swaggo/swag#declarative-comments-format

Usage：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./cmd/main/main.go --output ./internal/app/swagger
	swag init --generalInfo cmd/gin-nuxt-server/main.go --output ./internal/app/swagger
	swag init --generalInfo ./main.go --output ../../internal/app/swagger
*/
package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mlsjla/gin-nuxt/server/internal/app"
	"github.com/mlsjla/gin-nuxt/server/pkg/logger"
)

// Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "8.1.0"

// @title gin-nuxt-server
// @version 8.1.0
// @description RBAC scaffolding based on GIN + GORM + CASBIN + WIRE.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
// @contact.name lanya
// @contact.email valueme@qq.com
func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "gin-nuxt-server"
	app.Version = VERSION
	app.Usage = "RBAC scaffolding based on GIN + GORM + CASBIN + WIRE."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "Run http server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Value:    "./configs/config.toml",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "Casbin model configuration(.conf)",
				Value:    "./configs/model.conf",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "menu",
				Usage:    "Initialize menu's data configuration(.yaml)",
				Value:    "./configs/menu.yaml",
				Required: false,
			},
			&cli.StringFlag{
				Name:  "www",
				Usage: "Static site directory",
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetModelFile(c.String("model")),
				app.SetWWWDir(c.String("www")),
				app.SetMenuFile(c.String("menu")),
				app.SetVersion(VERSION))
		},
	}
}
