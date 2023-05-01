package main

import (
	"go-db/model"
	"go-db/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	model.SetUp()
	var router = gin.Default()

	router.GET("/", routes.HomeHandler)
	router.GET("/users", routes.UsersHandler)
	router.POST("/add", routes.AddHandler)
	router.GET("/user/:id", routes.UserHandler)
	router.PUT("/update/:id/", routes.UpdateUser)
	router.DELETE("/delete/:id", routes.DeleteUser)

	router.Run("localhost:8002")
}

/*
	******API ENDPOINTS******
	curl localhost:8002/
	curl localhost:8002/users
	curl -X POST localhost:8002/add -H 'Content-Type:application/json' -d '{"name":"hash","email":"hash@hash.hash","password":"hash"}'
	curl localhost:8002/user/2
	curl -X PUT localhost:8002/update/2?pwd=newpwd
	curl -X DELETE localhost:8002/delete/2

*/
