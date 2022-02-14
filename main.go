package main

import (
	"log"

	iris "github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"data": "hello world",
		})
	})

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app
}

func main() {
	log.Fatal(newApp().Run(iris.Addr(":8080")))
}
