package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"otusHWUsers/controllers"
	"otusHWUsers/middleware"
	db "otusHWUsers/db/sqlc"
)

type AuthRoutes struct {
	authController controllers.AuthController
	db             *db.Queries
}

func NewAuthRoutes(authController controllers.AuthController, db *db.Queries) AuthRoutes {
	return AuthRoutes{authController, db}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("", AuthMiddleware())
	router.GET("/logout", middleware.DeserializeUser(rc.db), rc.authController.LogoutUser)
}

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
				token, err := ctx.Cookie("access_token")
				if err != nil {
					fmt.Println("No access_token")
					ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Not auth"})
				} else {
					ctx.Header("X-Token", token)
					ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "You auth"})
				}
		}
}
