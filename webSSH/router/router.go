package router

import (
	"gin-client-go.com/webSSH/apis"
	"gin-client-go.com/webSSH/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	middleware.InitMiddler(r)
	r.GET("/ping", apis.Ping)
	r.GET("/namespace", apis.GetNamespaces)
	r.GET("/pod", apis.GetPods)
	r.GET("/namespace/:namespaceName/pod/:podName/container/:containerName", apis.ExecContainer)
}
