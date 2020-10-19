package Tests

import (
	"MongoApi/Models"
	"MongoApi/Services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGetBooks(t *testing.T){
	assert := assert.New(t)
	var Book []Models.Book
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Services.GetAllBooks(&Book,c)
	assert.NotNil(Book)
	assert.Equal(http.StatusOK,w.Code,"Get OK")
}
