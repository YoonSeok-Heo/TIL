package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", MyBenchLogger(), func(c *gin.Context) {
		c.String(http.StatusOK, "benchmark")
	})

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "login")
	})
	authorized.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusOK, "submit")
	})
	authorized.POST("/read", func(c *gin.Context) {
		c.String(http.StatusOK, "read")
	})

	// nested group
	testing := authorized.Group("testing")
	// visit 0.0.0.0:8080/testing/analytics
	testing.GET("/analytics", func(c *gin.Context) {
		c.String(http.StatusOK, "analytics")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//샘플 변수 설정
		c.Set("example", "12345")

		// Request 이전

		c.Next()

		// Request 이후

		latency := time.Since(t)

		log.Print(latency)

		// 송신할 상태 코드에 접근
		status := c.Writer.Status()
		log.Println(status)
	}
}
