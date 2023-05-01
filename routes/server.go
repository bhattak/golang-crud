package routes

import (
	// "go-db/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()

	router.GET("/", HomeHandler)

	router.Run("localhost:8002")
}
