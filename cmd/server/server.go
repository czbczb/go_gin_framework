package main

import (
	"fmt"
	"go_gin_framework/pkg/app"
	"go_gin_framework/pkg/demo"
	"go_gin_framework/pkg/rateLimit"
	"net/http"
	"regexp"
  "time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	// gin.Default比gin.New 多了日志记录、错误处理、Gzip 压缩，
	// gin.New()是一个空白的引擎实例
  r := gin.Default()


  r.Use(getCorsMiddleware())

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

	r.GET("/demo", demo.GetEnvDemo)


  appGroup := r.Group("app")

  tokenBucket := rateLimit.NewTokenBucket(20, time.Second)
  appGroup.Use(tokenBucket.RateLimitMiddleware())

  appGroup.Use(getTokenMiddleware())
  appGroup.GET("/list", app.ListApp)
  appGroup.GET("/create", app.CreateApp)
  appGroup.GET("/detail/comment", app.CommentList)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func getTokenMiddleware() gin.HandlerFunc {
  return func (ctx *gin.Context) {
    accessToken := ctx.GetHeader("Authorization")
    // Set example variable
    ctx.Set("example", "12345")


    fmt.Println(accessToken)
  
    // before request 
    ctx.Next()
    // after request
  }
  
}


func getCorsMiddleware() gin.HandlerFunc {
  return func (c *gin.Context) {
    // set cors
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization", "demo")
    corsConfig.AllowOriginFunc = func(origin string) bool {
      _, err := regexp.MatchString(`demo.com`, origin)
      if err != nil {
        log.Errorln(err)
      }
      return true
    }
    corsConfig.AllowCredentials = true
  }
}