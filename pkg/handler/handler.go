package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/polyk005/magazin/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up", h.SignUp)
	// 	auth.POST("/sign-in", h.SignIn)
	// }

	return router
}
