# Multiple file upload

멀티 파일 업로드

같은 값에다 넣어서 반복문을 돌려준다.!

## code

```go
func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload"]

		for _, file := range files {
			log.Println(file.Filename)

			c.SaveUploadedFile(file, filepath.Join("E:\\workspace\\Go\\main", file.Filename))
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}
```

![image](https://user-images.githubusercontent.com/113662725/226178547-8c8c9c84-6138-40d7-aa5e-f61478fd6772.png)

