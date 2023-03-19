# Another example: query + post form

query, form-data에 맵형식으로 데이터를 넣을 수 있다. 

이런 형식은 처음본다. 다른언어에도 있는지 확인해봐야겠다. 

## code

### http request
```http request
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
```

### Go
```go
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	router.Run(":8080")
}
```

### http response
```json
{
  "ids": {
    "a": "1234",
    "b": "hello"
  },
  "names": {
    "first": "thinkerou",
    "second": "tianou"
  }
}
```

### query
![image](https://user-images.githubusercontent.com/113662725/226172804-b9b772ad-5d05-4148-9151-38309afa7302.png)

### form-data
![image](https://user-images.githubusercontent.com/113662725/226172813-1e2a1f42-0f30-429f-8841-90d3f2e6b448.png)