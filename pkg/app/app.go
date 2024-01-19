package app

import (
	"fmt"
	"go_gin_framework/pkg/response"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func ListApp(c *gin.Context)  {
	example := c.MustGet("example")
	fmt.Println("example: ", example)
}


func CreateApp(c *gin.Context)  {

	type InputParams struct {
		ClusterId uint   `json:"cluster_id" form:"cluster_id" binding:"required"`
		Namespace string `json:"namespace" form:"namespace" binding:"required"`
	}

	type Output struct {
		Name string `json:"name"`
	}

	var input InputParams

	err := c.ShouldBind(&input)
	if err != nil {
		log.Error(err)
		res := response.ErrorResponse("name is required")
		// 返回结果
		c.JSON(res.Status, res)
		return
	}
	output := &Output{
		Name: "app",
	}

	if output.Name == "" {
		res := response.ErrorResponse("name is required")
		// 返回结果
		c.JSON(res.Status, res)
	}
	// 获取格式化结果
	res := response.SuccessResponse(output)
	// 返回结果
	c.JSON(res.Status, res)
}


func CommentList(c *gin.Context)  {
	example := c.MustGet("example")
	fmt.Println("example: ", example)
}

