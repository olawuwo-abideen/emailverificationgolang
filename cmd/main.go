package main

import (
	"github.com/olawuwo-abideen/emailverificationgolang/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app := api.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8000")
}
