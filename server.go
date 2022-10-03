package main

import (
	"github.com/adewumi0550/pro_work/m/v2/config"
	"github.com/adewumi0550/pro_work/m/v2/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	authController controllers.AuthController = controllers.AuthControllerImpl()
)

func main() {

	defer config.CloseDatabase(db)
	r := gin.Default()
	authRoutes := r.Group("user")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
