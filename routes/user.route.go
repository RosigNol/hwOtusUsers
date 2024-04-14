package routes

import (
	"github.com/gin-gonic/gin"
	"otusHWUsers/controllers"
	"otusHWUsers/middleware"
	db "otusHWUsers/db/sqlc"
)

type UserRoutes struct {
	userController controllers.UserController
	db             *db.Queries
}

func NewUserRoutes(userController controllers.UserController, db *db.Queries) UserRoutes {
	return UserRoutes{userController, db}
}

func (rc *UserRoutes) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("/user")
	router.POST("", rc.userController.SignUpUser)
	router.DELETE(":id", rc.userController.DeleteUser)
	router.GET(":id", rc.userController.GetUser)
	router.PUT(":id", rc.userController.UpdateUser)
	router.GET("/list", rc.userController.GetUsers)
	router.GET("/me", middleware.DeserializeUser(rc.db), rc.userController.GetMe)
	router.PUT("/me", middleware.DeserializeUser(rc.db), rc.userController.UpdateMe)
}
