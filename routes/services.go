package routes

import (
	"fmt"
	"go-db/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello there !!!"})
}

func UserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, e := model.GetUser(id)
	fmt.Println("ERROR :::", e.Error())
	if e != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": e.Error()})
		return
	}
	c.IndentedJSON(http.StatusFound, user)
}

func AddHandler(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	errr := model.CreateUser(&user)
	if errr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": errr.Error()})
		return
	}
	c.IndentedJSON(http.StatusFound, user)
}

func UsersHandler(c *gin.Context) {
	userList, e := model.GetAllUser()
	if e != nil {
		panic(e)
	}
	c.IndentedJSON(http.StatusFound, userList)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	pwd, ok := c.GetQuery("pwd")
	if !ok {
		c.IndentedJSON(http.StatusFound, gin.H{"message": "Query param not found"})
		return
	}
	errr := model.UpdateUser(id, pwd)
	if errr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not update user"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errr := model.DeleteUser(id); errr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not delete user"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})

}
