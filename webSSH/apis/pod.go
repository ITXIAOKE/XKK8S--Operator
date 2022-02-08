package apis

import (
	"gin-client-go.com/webSSH/service"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

func GetPods(c *gin.Context){
	namespaceName := c.Param("namespaceName")
	pods,err:=service.GetPods(namespaceName)
	if err!=nil{
		klog.Fatal(err)
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusOK,pods)
}


func ExecContainer (c *gin.Context){
	namespaceName:=c.Param("namespaceName")
	podName:=c.Param("podName")
	containerdName:=c.Param("containerName")
	method:=c.DefaultQuery("action","sh")
	err:=service.WebSSH(namespaceName,podName,containerdName,method,c.Writer,c.Request)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
		return
	}
}
