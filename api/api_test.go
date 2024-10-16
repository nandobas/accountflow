package api_test

import (
	"accountflow/api"
	"accountflow/environment"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type testApiSuite struct {
	suite.Suite
	api    api.Service
	router *gin.Engine
}

func TestApi(t *testing.T) {
	suite.Run(t, new(testApiSuite))
}

func (t *testApiSuite) SetupTest() {

	environment.InitEnv("../.env")
	t.api = *api.NewService()
	t.api.Config()
	t.router = t.api.Engine
}

func (t *testApiSuite) TearDownTest() {

}

func (t *testApiSuite) TestApi_Test() {

	// Arrange
	url := "/test"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rw := httptest.NewRecorder()

	// Action
	t.router.ServeHTTP(rw, req)

	t.Equal(http.StatusOK, rw.Code)

	var data string
	err := json.Unmarshal(rw.Body.Bytes(), &data)
	if err != nil {
		t.T().Fatalf("Erro ao processar a resposta: %v", err)
	}
	t.Equal("Hello!", data)

}
