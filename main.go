// package main

// import (
// 	"fmt"
// 	"go-crud/controllers"
// 	"go-crud/initializers"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDB()
// }

// func main() {
// 	r := gin.Default()
// 	r.POST("/posts", controllers.PostsCreate)
// 	r.PUT("/posts/:id", controllers.PostUpdate)
// 	r.GET("/posts", controllers.PostsIndex)
// 	r.GET("/posts/:id", controllers.PostById)
// 	r.DELETE("/posts/:id", controllers.PostDelete)

// 	fmt.Println("time", time.Now().String())

// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID
	newUUID := uuid.New()

	// Convert UUID to a string
	uuidString := newUUID.String()

	fmt.Printf("Generated UUID: %s\n", uuidString)

	// Parse a UUID from a string
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
	} else {
		fmt.Printf("Parsed UUID: %s\n", parsedUUID.String())
	}
}
