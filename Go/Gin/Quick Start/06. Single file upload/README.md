# Single file upload

단일 파일 업로드

## code

```go
func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		
		c.SaveUploadedFile(file, filepath.Join("E:\\workspace\\Go\\main", file.Filename))

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
```

![image](https://user-images.githubusercontent.com/113662725/226177755-76153fec-e814-45d2-8edc-90f118073a3c.png)
