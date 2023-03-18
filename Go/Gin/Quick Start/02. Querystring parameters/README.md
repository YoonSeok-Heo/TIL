# Parameters in path

## code
```go
func main() {
  router := gin.Default()

  router.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })
  router.Run(":8080")
}
```

## 01

```go
router.GET("/welcome", func(c *gin.Context) {
  firstname := c.DefaultQuery("firstname", "Guest")
  lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

  c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
})
```

![image](https://user-images.githubusercontent.com/113662725/226117547-28f699d7-3556-4a3a-aa74-91a54a8d6fc0.png)

|구분| 값                        |
|---|--------------------------|
|request| 127.0.0.1:8080/welcome?firstname=Jane&lastname=Doe |
|response| Hello Jane Doe    |

- c.DefaultQuery의 첫번째 인자는 파라미터값을 가져오고 값이 들어오지 않을경우 2번째 인자가 기본값으로 들어간다.
- c.Query파라미터의 값을 가져온다.

