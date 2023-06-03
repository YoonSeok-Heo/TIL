---
layout: post
title: "[Spring-Boot] Spring-Boot 설정을 아라보자_01(SpringBootApplication)"
tags: 
  - Configuration Properties
categories:
  - Spring boot
  - Spring Framework
---

## Intro
강의에서 SpringBootApplication 잘사용하면 스프링의 신이라 했다. 그래서 정리한다.



## SpringBootApplication

Application.java 에서 @SpringBootApplication은 자동으로 설정을 해주기 위한 어노테이션으로
<br/>
org.springframework.boot.autoconfigure 패키지에 들어 있다. 

SpringBootApplication을 열어보게 되면 스프링의 고유 어노테이션 3개가 있다. 
```java
@SpringBootConfigurationm
@EnableAutoConfiguration
@ComponentScan 
```
세 가지가 모두 들어있다. (위 어노테이션들은 나중에 설명)

SpringBootApplication는 크게 두 단계로 읽힌다. <br/>

1. @ComponentScan을 통해서 클래스에 @Component라는 어노테이션이 붙어있으면 수집을 한다.
@Controller, @Service, @Repositorty 등 개발에 자주 사용하는 어노테이션 모두 해당된다. 

2. 두번째는 @EnableAutoConfiguration은 spring-boot-autoconfigure/META-INF/spring.fatories안에 들어있는 자동 설정들을 **조건에 따라서** 등록 한다.(조건은 나중에 다시 정리)


@ComponentScan에서 첫번째로 빈들을 등록하고 나머지 빈들은 @EnableAudtoConfiguration을 통해 빈을 등록하는 과정을 거친다.


