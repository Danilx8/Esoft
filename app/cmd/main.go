package main

import (
	"esoft/app/api/route"
	"esoft/app/bootstrap"
	"esoft/docs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a Esoft server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	app, err := bootstrap.App()
	if err != nil {
		fmt.Printf("Get error while start app: %v", err)
		return
	}

	e := gin.Default()

	route.Setup(app.DB, e)

	e.Run(fmt.Sprintf(":%s", app.Config.Server.Port))
}
