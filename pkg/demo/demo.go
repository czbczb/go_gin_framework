package demo

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func GetEnvDemo (c *gin.Context) {
	mode := os.Getenv("apiType")
	fmt.Println("apiType: ", mode)
	ginmode := os.Getenv("GIN_MODE")
	fmt.Println("ginmode: ", ginmode)


	mode = gin.Mode()
	fmt.Println("gin.mode: ", mode)
}
func ListApp (c *gin.Context) {
	mode := os.Getenv("apiType")
	fmt.Println("apiType: ", mode)
	ginmode := os.Getenv("GIN_MODE")
	fmt.Println("ginmode: ", ginmode)


	mode = gin.Mode()
	fmt.Println("gin.mode: ", mode)
}