package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz returns response body "OK"
func Healthz(context *gin.Context) {
	context.String(http.StatusOK, "OK")
}

// Readyz returns response body "OK"
func Readyz(context *gin.Context) {
	context.String(http.StatusOK, "OK")
}
