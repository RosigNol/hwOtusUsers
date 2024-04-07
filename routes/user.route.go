package routes

import (
	"github.com/gin-gonic/gin"
	"otusHWUsers/controllers"
)

type UserRoutes struct {
	userController controllers.UserController
}

func NewUserRoutes(userController controllers.UserController) UserRoutes {
	return UserRoutes{userController}
}

func (rc *UserRoutes) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("/user")
	router.POST("", rc.userController.SignUpUser)
	router.DELETE(":id", rc.userController.DeleteUser)
	router.GET(":id", rc.userController.GetUser)
	router.PUT(":id", rc.userController.UpdateUser)
	router.GET("/list", rc.userController.GetUsers)
}
