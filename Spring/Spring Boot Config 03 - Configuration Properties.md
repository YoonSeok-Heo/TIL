---
layout: post
title: "[Spring-Boot] Spring-Boot 설정을 아라보자_03(Configuration Properties)"
tags: 
  - Configuration Properties
  - SpringFramework
categories:
  - Spring boot
  - Spring Framework
---

## Intro
Configuration Properties를 알아보자~!


## 1. Configuration Properties
각종 설정값을 외부로 분리해낸 것
- 서로 다른 환경에서도 사용할 수 있음
- 애플리케이션을 새로 컴파일하지 않고 설정값을 바꿀 수 있음
- 종류
  - Java properties file
  - YAML
  - environment variable
  - command-line argument

### 1.1. 외부 설정의 우선 순위
1. 디폴트 프로퍼티
2. @Configuration 클래스에 @PropertySource 로 정의된 것
3. 설정 파일: application.properties
4. RandomValuePropertySource
5. OS 환경변수
6. 자바 시스템 프로퍼티: System.getProperties()
7. JNDI 속성: java:comp/env
8. ServletContext - 초기 파라미터
9. ServletConfig - 초기 파라미터
10. SPRING_APPLICATION_JSON 안의 프로퍼티들
11. Command-line arguments
12. 테스트에 사용된 프로퍼티들
13. @TestPropertySource
14. Devtools 글로벌 세팅: $HOME/.config/spring-boot


### 1.2. 설정 파일(Config data)의 우선순위
1. JAR 패키지 안의 application.properties, application.yaml
2. JAR 패키지 안의, 프로파일이 지정된 파일: application-{profile}.properties
3. JAR 패키지 밖의 파일
4. JAR 패키지 밖의, 프로파일이 지정된 파일

### 1.3. 설정 파일(Config data)의 위치(우선순위)
1. classpath<br/>
1.1. classpath:/<br/>
1.2. classpath:/config
2. 현재 디렉토리
2.1. ./<br/>
2.2. ./config<br/>
2.3. ./config/child

#### 설정이 중복 될 경우
- 순서대로 읽기 때문에 우선 순위가 낮은 설정이 남게 된다.
- 한마디로 처음 설정은 뒤에 설정에 덮어쓰기가 된다.


## 2. 설정 파일(Config data)을 읽는 방법
- @Value
- Environment
- @ConfigurationProperties


### 2.1. @value
- SpEL 로 프로퍼티명을 표현
- type-safe 하지 않음
- 인스턴스화 이후에 주입하므로, final 쓸 수 없음
- 생성자 안에서 보이지 않음 (대안: @PostConstruct)
- Relaxed binding 지원 (kebab-case only)
- meta-data 없음 (javadoc은 적용 가능)

```java
@Component
public class MyBean {
  @Value("${name}")
  private String name;

  // ...
}
```
#### 참고사항
- 생성자 주입을 하지 않고 필드에 바로 주입했을 경우 값이 정상적으로 들어가지 않는다. 인스턴스가 생성된 후에 값이 들어가기 때문에 잘 확인해야한다.


### 2.2. Environment
- 애플리케이션 컨텍스트에서 꺼내오는 방법
- Environment 빈을 가져오는 방법
- 눈에 잘 안들어옴

- Environment 클래스를 이용해서 가져온다
```java
Environment environment;
environment.getProperty{"{속성이름}"}
```


### 2.3. ConstrunctorBinding
- Immutable한 프로퍼티를 구현할 수 있는 방법 - 추천
  

## 결론 
- @ConfigurationProperties 를 이용해서
- 상수처럼, immutable 하고 type-safe 하고 명확한 프로퍼티를 만들어 사용하자