package service

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/golangAPI-gin/pojo"
	"log"
	"net/http"
)


//Get userList
func FindAllUsers(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, userList)
	users := pojo.FindAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func FindUserById(ctx *gin.Context) {
	user := pojo.FindUserById(ctx.Param("id"))
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	ctx.JSON(http.StatusOK, user)
}

func PostUser(ctx *gin.Context) {
	user := pojo.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error: " + err.Error())
		return
	}
	//userList = append(userList, user)
	newUser := pojo.CreateUser(user)
	ctx.JSON(http.StatusOK, newUser)
}


//delete user
func DeleteUser(ctx *gin.Context) {
/*
	#string到int
	int,err:=strconv.Atoi(string)
	#string到int64
	int64, err := strconv.ParseInt(string, 10, 64)
	#int到string
	string:=strconv.Itoa(int)
	#int64到string
	string:=strconv.FormatInt(int64,10)
*/
	user := pojo.DeleteUser(ctx.Param("id"))
	if ! user {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	ctx.JSON(http.StatusOK, "Successfully")
}

func PutUser(ctx *gin.Context)  {
	user := pojo.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(ctx.Param("id"), user)
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}

	ctx.JSON(http.StatusOK, user)
}





