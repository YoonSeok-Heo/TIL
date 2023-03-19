# Multipart Urlencoded Form

form-data에 데이터를 담아 보낼 때 사용하는 구문

주로 이미지나 파일 데이터를 넘길 때 주로 사용했다.

## code
```go
func main() {
  router := gin.Default()

  router.POST("/form_post", func(c *gin.Context) {
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")
	

    c.JSON(http.StatusOK, gin.H{
      "status":  "posted",
      "message": message,
      "nick":    nick,
    })
  })
  router.Run(":8080")
}
```

![image](https://user-images.githubusercontent.com/113662725/226171767-547a2e6d-89eb-4f59-9e8a-1678bdbfd921.png)

- body -> form-data에 데이터를 담아 request를 보낸다.
- c.PostForm, c.DefaultPostForm을 사용해 데이터를 받는다.

