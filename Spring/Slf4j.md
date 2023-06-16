# Slf4j

@Slf4j 어노테이션을 사용하면서 문득 궁금해졌다.

### 로그 프레임워크는 어떤게 있을까?

구글링 해보면 로깅 프레임워크는 Log4j, logback, Log4j2, slf4j정도로 나오고 있다.
(4개 외에 자바에 내장되어있는것도 있고 다양하게 있다. 왼쪽의 4개는 주로 사용하는 프레임워크이다.)


### Slf4j

**Simple Logging Facade for Java**의 약자이다. 
이름에서 확인할 수 있듯이 다양한 로깅 프레임워크(log4j2, logback 등)에 대한 인터페이스 역할을 하는 라이브러리라고 보면된다.

인터페이스이기 때문에 Slf4j는 단독으로 사용하지 않고 로깅 프레임워크와 함께 사용해야한다.

즉. 인터페이스를 사용하여 로깅을 구현하게 된다면 추후에 필요로 의해 로깅 프레임워크를 변경해도 코드의 변경없이 가능하다.(로깅 프레임워크를 수정한다고해서 코드를 수정할 필요가 없다)

---

## SLF4J의 동작 과정

### SLF4J Bridging Modules

1. 다른 로깅 API로의 Logger 호출을 SLF4J 인터페이스로 리다이렉트하여 SLF4J API가 대신 처리할 수 있도록 일종의 어댑터 역할을 하는 라이브러리
2. log4j-over-slf4j, jcl-over-slf4j, jul-to-slf4j

### SLF4J API(인터페이스)
1. 로깅에 대한 추상 레이어를 제공합니다. 즉 추상 메서드를 제공해요.
2. 하지만 위에서 봤듯이 이 dependency만 추가하면 binding 이 없다는 에러를 만나게 돼요. 추상 클래스이기 때문에 단독적으로 사용할 수가 없는 것이죠.
3. slf4j-api

### SLF4J Binding
1. SLF4J 인터페이스를 Logging Framework와 연결하는 어댑터 역할을 하는 라이브러리
2. SLF4J API를 구현한 클래스에서 Binding으로 연결될 Logger의 API를 호출해요. 하나에 API에는 하나의 Binding만을 두어야 해요.
3. logback: logback-classic

### Logging Framework
1. logback: logback-core



slf4j는 컴파일 시 Logging Framework를 선택할 수 있다는 장점이 있어요.
만일 이를 변경하고자 한다면, 위에서 Logging Framework 와 Binding을 바꾸면 돼요.