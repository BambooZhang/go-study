package main

import (
	"coin.merchant/routers"
	"common/server"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	routers.InitRouters(app)
	app.Run(iris.Addr(server.Config.App.Port))
}
