package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "otusHWUsers/db/sqlc"
	"otusHWUsers/utils"
	"otusHWUsers/models"
	"github.com/google/uuid"
)

type UserController struct {
	db *db.Queries
}

func NewUserController(db *db.Queries) *UserController {
	return &UserController{db}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(db.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(currentUser)}})
}

func (uc *UserController) UpdateMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(db.User)

	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := currentUser.ID

	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.UpdateUserParams{
		ID :       id,
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  hashedPassword,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := uc.db.UpdateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success update", "data": gin.H{"user": user}})
}

func (ac *UserController) SignUpUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.CreateUserParams{
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  hashedPassword,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := ac.db.CreateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success sign up", "data": gin.H{"user": user}})
}

func (ac *UserController) DeleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, "Wrong id")
		return
	}

	err = ac.db.DeleteUser(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success delete"})
}

func (ac *UserController) GetUser(ctx *gin.Context) {
	println(ctx.Param("id"))
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, "Wrong id")
		return
	}


	user, err := ac.db.GetUserById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success get user", "data": gin.H{"user": user}})
}

func (ac *UserController) UpdateUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))

	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.UpdateUserParams{
		ID :       id,
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  hashedPassword,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := ac.db.UpdateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success update", "data": gin.H{"user": user}})
}

func (ac *UserController) GetUsers(ctx *gin.Context) {

	args := &db.ListUsersParams{
		Limit: 100,
		Offset: 0,
	}

	items, err := ac.db.ListUsers(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success get user", "data": gin.H{"user": items}})
}
