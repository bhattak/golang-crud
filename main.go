package main

import (
	"go-db/database"
	"go-db/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.SetUp()
	var router = gin.Default()

	router.GET("/", service.HomeHandler)
	router.GET("/users", service.UsersHandler)
	router.POST("/add", service.AddHandler)
	router.GET("/user/:id", service.UserHandler)
	router.PUT("/update/:id/", service.UpdateUser)
	router.DELETE("/delete/:id", service.DeleteUser)

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
