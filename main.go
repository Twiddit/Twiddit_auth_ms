package main

import (
	"github.com/Twiddit/Twiddit_auth_ms/controllers"
	"github.com/Twiddit/Twiddit_auth_ms/initializers"
	"github.com/Twiddit/Twiddit_auth_ms/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.POST("/validate", middleware.RequireAuth, controllers.Validate)
	router.POST("/logout", middleware.RequireAuth, controllers.Logout)
	router.POST("/change_password", middleware.PasswordControl, controllers.ChangePassword)
	router.Run()
}
