---
layout: post
title: "[Spring-Boot] Spring-Boot 설정을 아라보자_02(@Component vs Configuration)"
tags: 
  - Configuration Properties
categories:
  - Spring boot
  - Spring Framework
---

## Intro
Component vs Configuration ~



## 빈을 만드는 방법

### 1. @Component
```java
@Component
public class BubbleSort <T extends Comparable>{
    @Override
    public List<T> sort(List<T> list){
      List<T> output = new ArrayList<>(list)
      ...
    }
}
```

### 2. @Configuration
```java
@Configuration
public class Config {

  @Bean
  public Sort<String> bubbleSort()  {
    return new BubbleSort<>();
  }
}
```
1. 보편적으로 빈을 생성하는 방법이다 
2. 프록시 빈 형태로 빈을 생성한다.

### 3. @Bean
```java
@Component
public class Config {

  @Bean
  public Sort<String> bubbleSort(){
    return new BubbleSort<>();
  }
}
```
1. Lite Mode(경량 빈) 빈 등록 <br/>
기본적으로 빈을 만들게 되면 프록시 빈으로 만든다.
하지만 Lite Mode는 프록시 빈으로 등록하지 않는다.
** 프록시 빈은 따로 알아보자**

2. 기본적인 빈들은 AOP에서 사용하는 프록시 빈으로 생성을 한다. 하지만 경량빈은 프록시 빈으로 생성하지 않기 때문에 여러 번을 호출하면 여러개의 인스턴스가 생성된다.

3. 일반적으로 AOP의 기능을 사용하려면 프록시 빈으로 생성해야 한다.

#### 왜 Lite Mode인가? 
일반적으로 프록시 빈을 만들 때 좀더 오랜 시간이 소요 된다.
더빠르게 만들어지기 때문에 Lite Mode이다



## Component vs Bean?

### @Component
- Class-level annotation(클래스에 등록한다)
- 등록하려는 빈의 클래스 소스가 편집 가능한 경우 사용
- auto-detection에 걸림

### @Bean
- method-level annotaion(메소드에 등록한다)
- 좀 더 읽기 쉬움
- 인스턴스화 하는 코드가 수동으로 작성됨(ex. new를 사용해 인스턴스를 만든다.)
- 빈의 인스턴스 코드와 클래스 정의가 분리된 구조
- 따라서 외부 라이브러리, 써드 파티 클래스도 빈으로 등록 가능
  

## Component: Stereotype Annotaions
컴포넌트에 해당하는 스테레오 타입 어노테이션들
- @Controller
- @SErvice
- @Repository

## @Configuration
이 클래스는 각종 빈 **설정**을 담고 있다.
<br/>
결국 @Configuration 안에도 @Component가 들어있다.

1. @SpringBootApplication 이 컴퓨넌트 스캔을 통해 @Configuration을 찾아냄
2. 안의 빈 설정(메소드)을 읽어서 스프링 컨테이너에 등록
3. 필요한 곳에 주입
4. 또는 각종 스프링 인터페이스의 구현에 함께 활용


## 결론 !!!

어노테이션이 의도에 맞게 사용되었는지 잘 봐주자

- 빈 설정은 @Configuration, 클래스 등록은 @Component
- 정확한 목적을 모르고 쓰면 "잘 모르겠는데 어쨌든 돌아가요" 시간폭탄
- IDE가 노란 줄 그어주면
  - 잘 돌아가네
  - 수정 할 곳이 있다. 

