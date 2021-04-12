package routers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIndex(t *testing.T) {
	router := gin.Default()
	router.GET("/index.html", Index)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/index.html", nil)
	router.ServeHTTP(w, req)
}

func TestNotFoundError(t *testing.T) {
	router := gin.Default()
	router.GET("/400.html", NotFoundError)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/400.html", nil)
	router.ServeHTTP(w, req)
}

func TestInternalServerError(t *testing.T) {
	router := gin.Default()
	router.GET("/500.html", InternalServerError)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/500.html", nil)
	router.ServeHTTP(w, req)
}
