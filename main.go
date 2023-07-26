package main

import (
	"fmt"
	"go-crud/controllers"
	"go-crud/initializers"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostById)
	r.DELETE("/posts/:id", controllers.PostDelete)

	fmt.Println("time", time.Now().String())

	r.Run() // listen and serve on 0.0.0.0:8080
}
