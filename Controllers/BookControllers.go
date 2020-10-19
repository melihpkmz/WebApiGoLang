package Controllers

import (
	"MongoApi/Models"
	"MongoApi/Services"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {

	var reqBook Models.RequestBook
	c.BindJSON(&reqBook)
	Services.InsertOneBook(&reqBook,c)
	/*err := Services.InsertOneBook(&reqBook)
	if err !=nil{
		c.JSON(http.StatusBadRequest,Middleware.AuthorEmptyException())
	}else{
		c.JSON(http.StatusCreated,reqBook)
	}*/
}

func GetBooks(c *gin.Context) {
	var books []Models.Book
	Services.GetAllBooks(&books,c)


}
