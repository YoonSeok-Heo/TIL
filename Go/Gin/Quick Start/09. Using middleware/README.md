# Using middleware

middelware 사용법

미들웨어는 양 쪽에 연결되며 데이터를 주고받을 수 있도록 중간역할을 한다고 한다.

[gin-gonic](https://github.com/gin-gonic/gin/blob/master/docs/doc.md)을 따라하면서 느낀건데 request를 처리하기전에 사전작업을 하는 모습이 Spring에 filter정도를 닮은거 같다. 이건 그냥 내 생각이다.


## 보기.

### gin.New() vs gin.Default()
Gin Engine을 생성할 때 gin.Default()는 기본적으로 Logger라우터와 Recovery라우터가 들어있다.

그와 반대로 gin.New()로 생성할 경우 라우터가 들어있지 않은 상태로 Engine 오브젝트가 반환된다.

```go
gin.New()
gin.Default()
```

### Custom router
```go
r.GET("/benchmark", MyBenchLogger(), func(c *gin.Context) {
    c.String(http.StatusOK, "benchmark")
})
```
다른 엔드포인트를 작성할때와 다르게 파라미터가 한개 더 들어가있다. MyBenchLogger는 라우터를 구현한 것이다.
```go
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
```
함수처럼 구현해서 사용한다. 

여기서 c.Next()를 기준으로 Request이전과 이후로 나누어지는데 이부분을 보고 Spring에서 Filter와 비슷하다고 생각했다.





## code

### Go
```go
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
		log.Println(status)~~~~
	}
}
```

### 라우터 확인

![image](https://user-images.githubusercontent.com/113662725/226780888-ad519d58-220e-456e-97a4-8c8cb7995927.png)

위 코드를 실행하면 /benchmark의 handlers만 4인것을 확인할 수 있다. 한개 더 추가했기 때문이다. 

### 로그
![image](https://user-images.githubusercontent.com/113662725/226781209-591e6951-8cbd-4fe9-b87d-209d75069306.png)

MyBenchLogger에 log를 찍어둔 부분이 출력된 내용이다.