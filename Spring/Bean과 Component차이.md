# Bean과 Component의 차이

업무를 하면서 새로만든 클래스를 빈으로 등록하려고 @Bean을 추가하는데 자꾸 에러가 발생했다.

혹시 몰라서 @Component를 추가하니 정상적으로 실행된다.

둘의 차이점은 무엇일까?

[https://youngjinmo.github.io/2021/06/bean-component/](https://youngjinmo.github.io/2021/06/bean-component/)를 참고하여 작성했습니다.

---

## IoC 컨테이너

빈과 컴포넌트를 알기전에 IoC를 간단하게 알아봐야한다.

**스프링에서의 제어권은 IoC컨테이너에 있다. 그래서 IoC(Inversion Of Control), 제어의 역전이라고 한다.**

말그대로 개발자가 객체를 제어하는 것이 아닌 **스프링에서 객체들을 제어**하는것이다. 

스프링에서 객체들을 제어하려면 빈(Bean)으로 등록되어야 하는데, 과거에는 xml에 모두 적어주었으나 **요즘엔 어노테이션으로 간단하게 등록**할 수 있다.

---

## 빈(Bean) 등록

스프링 MVC에서는 @Controller, @Service, @Repository 등으로 빈을 등록할 수 있으며, 

다른 방법으로는 @Bean과 @Component를 이용해 스프링 컨테이너에 빈으로 등록하는 방법이 있다.

> @Controller, @Service, @Repository도 확인해보면 내부에 @Component가 있는걸 확인할 수 있다.
> 
> ```java
> @Target({ElementType.TYPE})
> @Retention(RetentionPolicy.RUNTIME)
> @Documented
> @Component
> public @interface Service {
>   @AliasFor(
>       annotation = Component.class
>   )
>   String value() default "";
> }
> ```

---

## Bean과 Component의 차이?

```@Bean```은 메소드 레벨에서 선언하며, 반환되는 객체를 개발자가 수동으로 빈으로 등록하는 어노테이션이다.

```java
@Configuration
public class AppConfig {
   @Bean
   public MemberService memberService() {
      return new MemberServiceImpl();
   }
}
```

```@Component```는 클래스 레벨에서 선언하며, 스프링이 런타임시에 컴포넌트스캔을 하여 자동으로 빈을 찾아 등록하는 어노테이션이다.

```java
@Component
public class Utility {
   // ...
}
```

개발자가 컨트롤이 불가능한 외부 라이브러리를 빈으로 등록하고 싶을 경우는 ```@Bean```을 사용하고, 개발자가 직접 컨트롤이 가능한 클래스는 ```@Component```를 사용한다고 한다.

---

## 결론

| Bean                           |Component|
|--------------------------------|---|
| 메소드에서 사용                       |클래스에서 사용|
| 개발자가 컨트롤이 불가능한<br>외부 라이브러리에 사용 |개발자가 직접 컨트롤이 가능한<br>내부 클래스에 사용|