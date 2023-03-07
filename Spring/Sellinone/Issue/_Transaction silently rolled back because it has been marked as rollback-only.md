# Transaction silently rolled back because it has been marked as rollback-only
#Transactional annotation #Spring boot #JPA

## 1. 발생 현상
회원가입 로직을 만들고 테스트를 하는 도중 경험해보지 못한 현상을 확인했다.

에러가 발생하는 소스를 try/catch문으로 분명 묶어 두었는데 다른 에러가 발생하는 것을 확인했다.

**※ 에러는 회원가입시 중복된 계정이름으로 가입을 시도했을때 primary key에 걸려서 회원가입이 되지 않을 때 발생했다.**
```console
2023-03-07T23:29:58.495+09:00 ERROR 18016 --- [nio-8080-exec-1] o.h.engine.jdbc.spi.SqlExceptionHelper   : (conn=247) Duplicate entry 'user' for key 'PRIMARY'
2023-03-07T23:29:58.510+09:00 ERROR 18016 --- [nio-8080-exec-1] c.w.sellinone.service.user.UserService   : UserService join: could not execute statement; SQL [n/a]; constraint [PRIMARY]
2023-03-07T23:29:58.513+09:00 ERROR 18016 --- [nio-8080-exec-1] c.w.sellinone.service.user.UserService   : org.springframework.orm.jpa.vendor.HibernateJpaDialect.convertHibernateAccessException(HibernateJpaDialect.java:270)
2023-03-07T23:29:58.523+09:00 ERROR 18016 --- [nio-8080-exec-1] o.a.c.c.C.[.[.[/].[dispatcherServlet]    : Servlet.service() for servlet [dispatcherServlet] in context with path [] threw exception [Request processing failed: org.springframework.transaction.UnexpectedRollbackException: Transaction silently rolled back because it has been marked as rollback-only] with root cause

org.springframework.transaction.UnexpectedRollbackException: Transaction silently rolled back because it has been marked as rollback-only
	at org.springframework.transaction.support.AbstractPlatformTransactionManager.processCommit(AbstractPlatformTransactionManager.java:752) ~[spring-tx-6.0.5.jar:6.0.5]
	at org.springframework.transaction.support.AbstractPlatformTransactionManager.commit(AbstractPlatformTransactionManager.java:711) ~[spring-tx-6.0.5.jar:6.0.5]
...
...
```
위 로그를 확인해보면 try/catch문으로 묶어둔 부분은 ERROR가 정상적으로 로그가 찍혔는데 예상한 리턴값이 반환되지 않았다.

### 예상한 API 리턴값
```json
{
    "status": 400,
    "message": "회원 가입 실패",
    "result": {}
}
```

### 발생한 API 리턴값
```json
{
    "timestamp": "2023-03-07T14:29:58.547+00:00",
    "status": 500,
    "error": "Internal Server Error",
    "path": "/user"
}
```

return 값이 다르다보니 위 문제를 해결하기로 마음먹었다.

## 2. 발생 원인
간단하게 설명하면(사실 아직 정확히 이해를 못함...) Transation이 중첩(?) 되어서 발생하는 에러이다.

회원가입 로직이 있는 Service 클래스에 @Transactional 어노테이션이 있다. 그리고 save메서드를 호츨하는 Repository도 JpaRepository를 구현하고 있다.
(JpaRepository구현체는 transactional어노테이션이 달려있다.)

Service -> 외부 트랜잭션

repository.save(T) -> 내부 트랜잭션

내부 트랜잭션이 외부 트랜잭션 자식으로 들어가는데 자식 트랜잭션에서 예외가 발생하면 부모에서 롤백이 발생한다.

## 3. 해결 방법
Transactional Propagation을 정리해야겠다.