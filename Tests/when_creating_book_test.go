package Tests

import (
	"MongoApi/Models"
	"MongoApi/Services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)


func TestCreatingBook(t *testing.T){
	assert := assert.New(t)
	var Book Models.RequestBook

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Book.Title = "Abc"
	Book.Author = "AbcTest"

	Services.InsertOneBook(&Book,c)
	assert.Equal(200,w.Code,"Book is created, succesfully")

}

