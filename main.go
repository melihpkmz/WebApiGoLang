package main

import (
	"MongoApi/Controllers"
	"MongoApi/Middleware"
	"MongoApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router *gin.Engine



func main(){

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Book server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "book.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	Router = gin.New()
	Router.Use(Middleware.CustomMiddleware)
	v1 := Router.Group("/api")
	{
		book := v1.Group("/book")
		book.GET("",Controllers.GetBooks)
		book.POST("",Controllers.CreateBook)
	}
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.Run(":5000")
}


