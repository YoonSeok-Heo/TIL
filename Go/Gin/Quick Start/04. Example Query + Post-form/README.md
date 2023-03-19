# Another example: query + post form

query랑 post-form을 같이 쓰는 경우이다. 

## code

### http request
```http request
POST /post?id=happyjohn&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
```

### Go
```go
func main() {
  router := gin.Default()

  router.POST("/post", func(c *gin.Context) {

    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
  })
  router.Run(":8080")
}
```

### http response
```json
{
  "id": "happyjohn",
  "message": "this_is_great",
  "name": "john",
  "page": "1"
}
```

![image](https://user-images.githubusercontent.com/113662725/226172533-dfac1240-c142-4db8-8623-7e45cf17e8ae.png)

