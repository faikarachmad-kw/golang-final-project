package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.DELETE("/", middlewares.Authentication(), controllers.UserDelete)
		userRouter.PUT("/", middlewares.Authentication(), controllers.UserUpdate)
	}

	photoRouter:=r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/",controllers.CreatePhoto)
		photoRouter.GET("/",controllers.GetPhoto)
		photoRouter.PUT("/:photoID",middlewares.Authorization("photo"),controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID",middlewares.Authorization("photo"),controllers.DeletePhoto)
	}

	commentRouter:=r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/",controllers.CreatePhoto)
		commentRouter.GET("/",controllers.GetPhoto)
		commentRouter.PUT("/:commentID",middlewares.Authorization("comment"),controllers.UpdatePhoto)
		commentRouter.DELETE("/:commentID",middlewares.Authorization("comment"),controllers.DeletePhoto)
	}

	socmedRouter:=r.Group("/socialmedias")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.POST("/",controllers.CreateSocmed)
		socmedRouter.GET("/",controllers.GetSocmed)
		socmedRouter.PUT("/:commentID",middlewares.Authorization("socmed"),controllers.UpdateSocmed)
		socmedRouter.DELETE("/:commentID",middlewares.Authorization("socmed"),controllers.DeleteSocmed)

	}
	
	return r
}