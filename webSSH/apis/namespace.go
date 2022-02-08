package apis

import (
	"gin-client-go.com/webSSH/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNamespaces(c *gin.Context) {
	namespaces, err := service.GetNamespaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, namespaces)
}
