# Parameters in path

## code
```go
func main() {
  router := gin.Default()

  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  router.POST("/user/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/user/:name/*action" // true
    c.String(http.StatusOK, "%t", b)
  })

  router.GET("/user/groups", func(c *gin.Context) {
    c.String(http.StatusOK, "The available groups are [...]")
  })

  router.Run(":8080")
}
```

## 01

```go
router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
})
```

![image](https://user-images.githubusercontent.com/113662725/226116988-d27c04f1-2f33-4424-a3ba-ec338a09ef91.png)

|구분| 값                        |
|---|--------------------------|
|request| 127.0.0.1:8080/user/john |
|response| Hello john     |

- :name 부분에 홍길동이 들어가고 c.Param을 이용해 받는다.
- 위 핸들러는 '/user/john' 은 사용가능하나 '/user/john/'은 사용이 불가능하다(마지막에 '/'가 중요하다.)

## 02

```go
router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
})
```

![image](https://user-images.githubusercontent.com/113662725/226117006-f0c823eb-e7e6-4818-908b-477a56da8459.png)

|구분|값|
|---|---|
|request|127.0.0.1:8080/user/john/send|
|response|john is /send|

- *action이 의미하는걸 정확히 모르겠다 알아봐야겠음.(수정)
- action변수에 "/send"가 들어간다 '/'까지포함(?)
- '/user/john/'과 '/user/john/send' 매칭한다. 
- '/user/john'과 매칭되는 라우터가 없으면 '/user/john/'으로 redirect(한글로 뭐라해야하지)한다. 

## 03
```go
router.POST("/user/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/user/:name/*action" // true
    c.String(http.StatusOK, "%t", b)
})
```

![image](https://user-images.githubusercontent.com/113662725/226117092-581a9ca7-0193-434b-8bc8-a9920d48d031.png)

|구분| 값                                  |
|---|------------------------------------|
|request| POST 127.0.0.1:8080/user/john/send |
|response| true                               |

