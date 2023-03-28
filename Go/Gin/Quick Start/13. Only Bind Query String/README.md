# Only Bind Query String

ShouldBindQuery는 무조건 Query의 파라미터를 가져온다.

1. route.Any
   ```go
   route.Any("/testing", startPage)
   ```
   - Any는 GET, POST 등 따로 지정하지않고 모든 메서드를 받는다.
  ![image](https://user-images.githubusercontent.com/113662725/228302037-766f1d2b-faad-420a-a20c-d76656129930.png)

2. ShouldBindQuery()
   - Query의 값을 바인딩한다.


## code

### Go
```go
package main

import (
   "log"
   "net/http"
   "time"

   "github.com/gin-gonic/gin"
)

type Person struct {
   Name       string    `form:"name"`
   Address    string    `form:"address"`
   Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
   CreateTime time.Time `form:"createTime" time_format:"unixNano"`
   UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
   route := gin.Default()
   route.GET("/testing", startPage)
   route.Run(":8085")
}

func startPage(c *gin.Context) {
   var person Person
   // If `GET`, only `Form` binding engine (`query`) used.
   // If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
   // See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
   if c.ShouldBind(&person) == nil {
      log.Println(person.Name)
      log.Println(person.Address)
      log.Println(person.Birthday)
      log.Println(person.CreateTime)
      log.Println(person.UnixTime)
   }

   c.String(http.StatusOK, "Success")
}
```

## JSON request and response

![image](https://user-images.githubusercontent.com/113662725/228302691-faf964da-2ce0-4e04-808a-76975fc315fa.png)