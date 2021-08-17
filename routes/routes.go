package routes

import (
	"net/http"

	"github.com/ashwins93/social-way-be/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	R  *gin.Engine
	DB *gorm.DB
}

func NewService(r *gin.Engine, db *gorm.DB) *Service {
	return &Service{R: r, DB: db}
}

func (s *Service) SetupV1Routes() {
	v1Router := s.R.Group("/api/v1")

	userRouter := v1Router.Group("/users")

	userRouter.GET("", func(c *gin.Context) {
		var users []model.User
		s.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})
}
