package handler

import (
	"net/http"

	"github.com/Hdeee1/be-go-restaurant-management/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	services	domain.UserServices
}

type registerResponse struct {
	ID			int		`json:"id"`
	FullName	string	`json:"full_name"`
	Phone		string	`json:"phone"`
	Role		string	`json:"role"`
}

type loginResponse struct {
	FullName		string	`json:"full_name"`
	Phone			string	`json:"phone"`
	AccessToken		string	`json:"access_token"`
	RefreshToken	string	`json:"refresh_token"`
}

func NewUserHandler(router *gin.RouterGroup, services domain.UserServices) {
	handler := &UserHandler{services: services}

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var req domain.RegisterRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.Register(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, registerResponse{
		ID: user.ID,
		FullName: user.FullName,
		Phone: user.Phone,
		Role: user.Role,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {}