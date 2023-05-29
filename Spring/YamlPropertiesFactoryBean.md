# YamlPropertiesFactoryBean

DTO클래스에 스프링 설정 값을 기본값으로 넣어주고 싶었는데(생성자를 통해서) @Value나 @PropertySource와 같은 기능들은 클래스가 빈으로 등록되어야만 사용할 수 있었다.
내부 값이 계속해서 바뀌는 DTO의 특성상 빈으로 등록하는건 좋지 못한 방법이라고 생각해서 설정 값을 가져올만한 클래스를 알아보았다.

## 사용법

### 구현한 클래스 코드부터 보자.

```java
public class YamlPropertyFactory {

    public Map<Object, Object> getObject(){

        YamlPropertiesFactoryBean factory = new YamlPropertiesFactoryBean();
        factory.setResources(new ClassPathResource("application.yml"));
        Properties properties = factory.getObject();

        return new HashMap<Object, Object>(properties);
    }
}
```

1. YamlPropertiesFactoryBean을 생성.
2. setResource메서드를 이용해 factory객체에 Resource값을 넣어준다.
3. getObject를 이용해 읽어들인 Resource를 Hashtable형태로 반환한다.(Properties객체가 Hashtable을 상속받은 클래스이다.)
4. Hashtable형태는 자주 사용하지 않고 내부적으로 스레드세이프할 필요가 없으니 HashMap으로 리턴한다.

> ```ClassPathResource```클래스를 타고타고 올라가다보면 ```Resource```를 구현한 것을 확인할 수 있다.
> 
> ```new ClassPathResource({String path})``` 생성자 파라미터에 들어간 경로 값을 읽어 ```Resource```로 넘겨준다

---

### 사용한 클래스를 보자.

```java
public class TestDto{
    
    private String suffix;
    
    public TestDto(){
        YamlPropertyFactory factory = new YamlPropertyFactory();
        this.suffix = (String) factory.getObject().get("test.suffix");
    }
}
```
```yaml
application.yml

test:
  suffix: .yaml
```

String변수 suffix에 application.yml의 test.suffix의 값을 생성하면서 넣으려고 한 것이다.(Default값처럼 자동으로 들어가게끔)

이작업을 하려고 포스트를 작성해봄

---

# 결론

### 이렇게 사용해도 되는지 모르겠다.

우선 가장 큰 문제점은 TestDto가 **자주 생성된다면 문제가 발생할 것이다.** 

TestDto 기본 생성자에 YamlPropertyFactory객체를 생성하는 코드가 있고,
그 안에는 YamlPropertiesFactoryBean객체를 생성해 리소스 파일을 읽는 작업을 하게 된다. 
DTO한번 생성하는데 너무 많은 일을 하는것 같다....

그럼 싱글톤으로 생성해 둘까?? 생각만해봄 뭐가 맞는지는 정확히 모르겠다.