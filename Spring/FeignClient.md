# FeignClient
#HTTP통신

chatGPT api를 이용해서 데이터를 받아오는 작업을 진행하는데 처음 접하게 되었다. 상사분의 코드를 보면서 말이다. 그리고 좀 더 디테일하게 사용하기 위해 공부를 해보자.

이 포스트는 [<공식문서>](https://docs.spring.io/spring-cloud-openfeign/docs/current/reference/html/)를 바탕으로 작성했다.

---

# 1. Declarative REST Client: Feign

## 1.1. 의존성 추가

### 1.1.1. Maven

Spring Cloud만 의존성을 추가해도 내부에 feign이 들어있는거 같았다. 이건 다시 확인 해봐야함.

```xml
<!-- https://mvnrepository.com/artifact/org.springframework.cloud/spring-cloud-starter-openfeign -->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-openfeign</artifactId>
    <version>3.1.7</version>
</dependency>
```

### 1.1.2. Gradle
```gradle
// https://mvnrepository.com/artifact/org.springframework.cloud/spring-cloud-starter-openfeign
implementation group: 'org.springframework.cloud', name: 'spring-cloud-starter-openfeign', version: '3.1.7'
```

## 1.2. 기본 세팅

프로젝트에서 Feign을 사용하려면 스프링 부트에 @EnbleFeignClients를 추가해 주어야 한다. (config파일을 설정하면 안넣어도 된다고 하는데 확인을 다시 해봐야 겠다.)

해당 어노테이션은 하위 클래스를 탐색하면서 @FeignClient를 찾아 구현체를 생성한다고 한다.

> 어쩐지 @EnbleFeignClients가 없을 때 인터페이스가 빈등록이 안되어 있다고 진행이 안돼서 상당히 난감했다.

**예제 코드를 한번 보자!**

```java
@SpringBootApplication
@EnableFeignClients
public class Application {

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

}
```

> Tip !
> @EnableFeignClients를 사용할때 basePackages를 명시하면 탐색 범위를 설정할 수 있다고한다
>> ex. @EnableFeignClients(basePackages = "com.example.clients")
> 

## 1.3. 인터페이스 생성

```java
@FeignClient("stores")
public interface StoreClient {
    @RequestMapping(method = RequestMethod.GET, value = "/stores")
    List<Store> getStores();

    @RequestMapping(method = RequestMethod.GET, value = "/stores")
    Page<Store> getStores(Pageable pageable);

    @RequestMapping(method = RequestMethod.POST, value = "/stores/{storeId}", consumes = "application/json")
    Store update(@PathVariable("storeId") Long storeId, Store store);

    @RequestMapping(method = RequestMethod.DELETE, value = "/stores/{storeId:\\d+}")
    void delete(@PathVariable Long storeId);
}
```

위 소스를 보면 기존 컨트롤러랑 상당히 비슷한 모습을 볼 수 있다.

### 1.3.1. @FeignClient()

다양한 사용법이 있지만 기본적인 역할을 확인해 보자 "stores"라는 값은 FeignClient의 이름을 뜻한다.

url속성에는 호출할 api의 url을 설정할 수 있고, configuration속성은 FeignClient의 설정 정보가 셋팅된 클래스를 넣어줄 수 있다.

```java
@FeignClient(name = "store", url = "localhost:8080/", configuration = StoreFeignClientConfiguration.class)
```

