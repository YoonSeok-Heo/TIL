# Bind Uri

Header값을 바인딩한다.

그럼 헤더 따로 바디 따로 해야하나?? 궁금해

## code

### Go
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		h := testHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusOK, err)
		}

		fmt.Printf("%#v\n", h)
		c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})

	r.Run()

	// client
	// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
	// output
	// {"Domain":"music","Rate":300}
}
```

## request and response

![image](https://user-images.githubusercontent.com/113662725/228310435-f6132190-ef1c-4802-82ab-ced4e2ed6c47.png)