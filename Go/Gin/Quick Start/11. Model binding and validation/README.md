# Model binding and validation

모델 바인딩

gin에서 리퀘스트로 넘어오는 데이터를 바인딩하기 위한 과정이 있다.

1. 바인딩할 객체(?)를 struct형식으로 만들어준다.
    ```go
    type Login struct {
        User     string `form:"user" json:"user" xml:"user"  binding:"required"`
        Password string `form:"password" json:"password" xml:"password" binding:"required"`
    }
    ```
    - struct는 string타입으로 만들어진다.
    - form, json, xml 형식으로 받는다. 
    - binding 옵션에 required가 들어가면 필수값이다.
    - **근데 대소문자를 어케 구별하는지 모르겠다..? 수정해봐도 구별을 안한다. 원래 안하는건가?**~~~~

2. 바인딩한 객체를 받아온다.
    ```go
    if err := c.ShouldBindJSON(&json); err != nil {
    ```
    - c.SuouldBindJSON -> JSON은 바인딩하는 메서드이다.

## 바인딩 메서드 정리
- Type - Must bind
    - Methods - `Bind`, `BindJSON`, `BindXML`, `BindQuery`, `BindYAML`, `BindHeader`, `BindTOML`
    - Behavior - These methods use `MustBindWith` under the hood. If there is a binding error, the request is aborted with `c.AbortWithError(400, err).SetType(ErrorTypeBind)`. This sets the response status code to 400 and the Content-Type header is set to text/plain; charset=utf-8. Note that if you try to set the response code after this, it will result in a warning `[GIN-debug]` `[WARNING] Headers were already written. Wanted to override status code 400 with 422.` If you wish to have greater control over the behavior, consider using the ShouldBind equivalent method.
- Type - Should bind
    - Methods - `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery`, `ShouldBindYAML`, `ShouldBindHeader`, `ShouldBindTOML`,
    - Behavior - These methods use `ShouldBindWith` under the hood. If there is a binding error, the error is returned and it is the developer's responsibility to handle the request and error appropriately.
    


## code

### Go
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "JSON you are logged in"})
	})

	// Example for binding XML (
	//  <?xml version="1.0" encoding="UTF-8"?>
	//  <root>
	//    <user>manu</user>
	//    <password>123</password>
	//  </root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

```

## JSON request and response

![image](https://user-images.githubusercontent.com/113662725/227789836-e2cd9165-6ae4-44cb-8931-ff06282af0ae.png)
