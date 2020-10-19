package Services

import (
	"MongoApi/DatabaseSettings"
	"MongoApi/Middleware"
	"MongoApi/Models"
	"MongoApi/Utilities"
	"context"
	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetAllBooks(books *[]Models.Book,c *gin.Context){
	collection := DatabaseSettings.Connect().Database("BookDb").Collection("Books")
	collection.Find(context.Background(),bson.D{})
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	errCursor := cursor.All(context.Background(),books)
	if errCursor != nil{
		panic(errCursor)
	}
	c.JSON(http.StatusOK,books)

}

func InsertOneBook(reqBook *Models.RequestBook,c *gin.Context){
	var book Models.Book
	book.ID = guid.New().String()
	book.Author = reqBook.Author
	book.Title = reqBook.Title
	collection := DatabaseSettings.Connect().Database("BookDb").Collection("Books")

	v := validator.New()
	errValidation := v.Struct(reqBook)

	if errValidation != nil {
		validationErrors := errValidation.(validator.ValidationErrors)
		for _, val := range validationErrors{
			if val.ActualTag() == "required"{
				if val.Field() == "Author"{
					Utilities.Logger(Utilities.LogType.String(1),"User tried to create empty book with author")
					c.Error(Middleware.AuthorEmptyException())
				}
				if val.Field() == "Title"{
					Utilities.Logger(Utilities.LogType.String(1),"User tried to create empty book with title")
					c.Error(Middleware.TitleEmptyException())
				}
			}
		}
	}else{
		_, err := collection.InsertOne(context.Background(), book)
		if  err != nil {
			panic(err)
		}else{
			c.JSON(200,book)
		}
	}

}
