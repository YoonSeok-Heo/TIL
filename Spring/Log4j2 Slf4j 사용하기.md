# Log4j2 Slf4j 사용하기


### 테스트 환경

- windows 11
- java 17
- Gradle 7.6.1
- Spring boot 3.1
- lombok 1.18.26

### [Slf4j가 무엇인지 확인하기](https://github.com/YoonSeok-Heo/TIL/blob/main/Spring/Slf4j.md)

## Log4j2가 성능

Spring Boot에서 기본적으로 사용하는 logback보다 성능이 상당히 좋다.

### [Log4j2 성능 확인하기](https://logging.apache.org/log4j/2.x/performance.html)

---

## 적용하기

### build.gradle

spring boot에서 기본적으로 사용하는 log Framwork를 없에야한다. **충돌이 일어나기 때문이다**.
> 기본 framework가 아니더라도 다른 log Framwork를 사용한다면 종속성을 제거해주어야 한다.

아래 코드를 확인하고 해당하는 부분만 추가하면 된다.

```
configurations {
    // logback 제거
	all {
		exclude group: 'org.springframework.boot', module: 'spring-boot-starter-logging'
	}
}

dependencies {
    // log4j2 추가
	implementation "org.springframework.boot:spring-boot-starter-log4j2";
}
```

### src/main/resources/application.yml

yml파일에 로깅정보파일 경로를 입력할 것이다.

```yaml
logging:
  config: classpath:log4j2/log4j2-local.xml
```

### src/main/resources/log4j2/log4j2-local.xml

로깅 설정파일

없다면 만들어주면 된다. 그리고 원하는대로 수정하면 된다!

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Configuration status="INFO">
    <Appenders>
        <Console name="console" target="SYSTEM_OUT">
            <PatternLayout pattern="[%-5level] %d{yyyy-MM-dd HH:mm:ss.SSS} [%t] %c{1} - %msg%n"/>
        </Console>
    </Appenders>
    <Loggers>
        <Root level="debug" additivity="false">
            <AppenderRef ref="console"/>
        </Root>
    </Loggers>
</Configuration>
```