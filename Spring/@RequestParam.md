# @RequestParam

> 컨트롤러에서 데이터를 받으면서 이상한짓을 많이했다.. 분명 예전에도 겪었던 일인데 또 실수해서 상기시키기 위해 정리좀 해야겠다.

@RequestParam은 컨트롤러에서 파라미터 값을 받을 때 사용하는 어노테이션이다.

우리가 Get메소드를 사용할 때 크게 두 가지의 형태로 파라미터를 전달한다.

```1) http://localhost:8080?name=heo```

```2) http://localhost:8080/heo```

둘다 heo라는 인자를 서버로 보낼 수 있다.

@RequestParam은 전자의 ```?name=heo```으로 넘어온 값을 처리해주는 어노테이션이다.


## 사용법

### 1) 파라미터 이름을 명시하지 않고 사용

```
호출 url
http://localhost:8080?name=heo
```
```java
@GetMapping
public String requestParamTest(@RequestParam String name){
    return name;
}
```

받으려고 하는 변수명(@RequestParam String name)과 보내는 파라미터명(?name=heo)이 같아야한다. 

변수명이 같지 않을경우 다음과 같이 BadRequest를 뱉는다.
```console
{
    "timestamp": "2023-05-28T14:09:05.378+00:00",
    "status": 400,
    "error": "Bad Request",
    "path": "/"
}
```

---

### 2) 파라미터 이름을 명시하고 사용
```
호출 url
http://localhost:8080?name=heo
```
```java
@GetMapping
public String requestParamTest(@RequestParam("name") String name){
    return name;
}
```

매핑시켜줄 변수를 적어주는 방법이다. url에서 name으로 들어오는 값이 변수 name으로 들어가게 된다.

이해가 안갈지도 모르니 다른 예를 보자.

```
호출 url
http://localhost:8080?name=heo
```
```java
@GetMapping
public String requestParamTest(@RequestParam("name") String eman){
    return eman;
}
```
```console
리턴값
heo
```

위와 다른점은 url의 name과 파라미터 변수명 eman가 달라도 매핑값을 정해주면 된다는 것이다.

이것또한 매핑해주는 값과 url값이 일치해야한다. 꼭 신경쓰자.

---

## 속성

RequestParam은 몇 개의 속성을 넣을 수 있다. 

### 1) required

말그대로 필수 값인지 확인해준다. 기본값은 true로 되어 있다.

```java
@GetMapping
public String requestParamTest(@RequestParam(value = "name", required = false) String eman){
    return eman;
}
```

위 코드 처럼 required의 값을 false로 해두면 위에서 발생한 매핑이 안되어서 생기는 Bad Request를 막을 수 있다.

### 2) defaultValue

기본값을 넣어줄 수 있다. 예를들어 매핑 값이 잘못들어가 값이 없게 된다면 기본으로 정해주도록 할 수 있다.

```java
@GetMapping
public String requestParamTest(@RequestParam(value = "name", required = false, defaultValue = "default-heo") String eman){
    return eman;
}
```

url파라미터가 틀려서 값이 안들어가게 된다면 기본적으로 "default-heo"라는 값이 들어가게 된다. 

---

## 많은 파라미터를 한번에 받는다면?

파라미터가 많아지면 일일히 다 적어줘야하는건가???

그건 아니다. Map형태로 간단히 받을 수 있다.

```java
@GetMapping
public HashMap<String, String> requestParamTest(@RequestParam HashMap<String, String> paramMap){
    return paramMap;
}
```

따로 매핑해주는 작없은 없고 맵만 파라미터로 받아주면 된다.

```
호출 url
http://localhost:8080?lsat-name=heo&fitst-name=ys
```
```console
결과 return값
{
    "lsat-name": "heo",
    "fitst-name": "ys"
}
```

그런데 이런방식으로는 잘 쓰지 않는다. **커멘드 객체**방식을 주로 사용한다.

이건 다음에 설명


