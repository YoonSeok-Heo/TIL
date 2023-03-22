# How to write log file

로그파일 생성

gin Engine 오브젝트를 생성하기 전에 설정을 변경하는 과정을 거친다.
```go
gin.DisableConsoleColor()
f, _ := os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)
```



## code

### Go
```go
func main() {
  // Disable Console Color, you don't need console color when writing the logs to file.
  gin.DisableConsoleColor()

  // Logging to a file.
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f)

  // Use the following code if you need to write the logs to file and console at the same time.
  // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

   router.Run(":8080")
}
```

## 로그 파일

![image](https://user-images.githubusercontent.com/113662725/226960099-b2f700de-7dd5-48d0-8b01-30735210eaa7.png)
