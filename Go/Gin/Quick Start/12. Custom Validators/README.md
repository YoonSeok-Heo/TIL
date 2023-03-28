# Custom Validators

커스텀 밸리데이션

바인딩한 데이터 검증 과정을 수정할 수 있다.

[validator docs](https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme) < 검증에 대한 값은 여기 정리되어 있다. 보면서 찾자.

1. 바인딩 struct생성
    ```go
    type Booking struct {
    	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	    CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
    }
    ```
    - bookabledate는 직접 생성한 검증이다. 
    - gtfield=CheckIn은 CheckIn보다 커야한다는 뜻이다.
   
2. bookableDate - custom validator
    ```go
    var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	    date, ok := fl.Field().Interface().(time.Time)
	    if ok {
    		today := time.Now()
    		if today.After(date) {
    			return false
    		}
    	}
    	return true
    }
    ```
   - 시작일이 오늘날짜보다 빠를경우 false를 반환하고 이상없으면 true를 반환(예약일이 과거가 아니라는 뜻)

3. main
    ```go
    func main() {
        route := gin.Default()
    
        if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
            v.RegisterValidation("bookabledate", bookableDate)
        }
    
        route.GET("/bookable", getBookable)
        route.Run(":8085")
    }
    ```
    - if문이 정확이 이해가 가질 않는다. 어떤걸 의미하는 것인가?? 더 공부해야겠다.
    - getBookable메서드를 실행시킨다.!


## code

### Go
```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
```

## request and response

### 정상적인 값 넣을경우

![image](https://user-images.githubusercontent.com/113662725/228295644-34c0b8c4-915f-4408-9255-842b4fd4f560.png)

### CheckOut값이 CheckIn보다 빠를경우

![image](https://user-images.githubusercontent.com/113662725/228296440-4503cead-3221-4268-93ef-3df9dbf050a8.png)

### 예약일이 과거일 경우

![image](https://user-images.githubusercontent.com/113662725/228298171-1f2a4549-e19e-4713-8970-480344f4bae2.png)