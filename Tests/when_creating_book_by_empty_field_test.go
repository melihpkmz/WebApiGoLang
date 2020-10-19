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

func TestCreatingBookWithEmptyField(t *testing.T){
	assert := assert.New(t)
	var Book Models.RequestBook

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Book.Title = "Abc"

	Services.InsertOneBook(&Book,c)
	assert.Equal(http.StatusBadRequest,w.Code,"Throw empty field, succesfully")

}
