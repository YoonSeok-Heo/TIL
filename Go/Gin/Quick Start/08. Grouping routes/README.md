# Grouping routes

Grouping route 기능이다.

v1, v2로 나뉘어서 엔드포인트를 2그룹으로 나누었다.


## code

```go
func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 login")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 submit")
		})
		v1.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 read")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 login")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 submit")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 read")
		})
	}

	router.Run(":8080")
}
```

### v1

![image](https://user-images.githubusercontent.com/113662725/226328526-bd797616-4569-4383-98a1-c008d6de3b30.png)

### v2

![image](https://user-images.githubusercontent.com/113662725/226328617-c486dfca-a578-4c86-82a0-046ec5793b76.png)
