# Connection Pool

Go에서 DB를 사용하고 있는데 Connection을 일일히 Close해주고있다... 이러면 안될거 같은데
<br>항상 써오던 Spring에서는 DB정보만 넘겨주면 DataSource 인터페이스를 사용해 커넥션풀을 관리해준다!!
<br>그런데?? Golang에서는 어떻게 될까?. 그래서 알아보기로 했따.

## Connection??

Connection이란? Application과 Database간 통신을 할 수 있는 수단이라고 보면 된다.
<br>DB서버와 연결을 맺고 끊음을 반복하면서 Connection을 관리하지만, 그 과정들의 성능에 영향을 끼치게 된다.
<br>이를 보완할 수 있는 방법이 Connection Pool이다.

## Connection Pool?

커넥션 풀은 DB와 연결된 Connection들을 미리 만들어 놓고 이를 Pool로 관리하는 것이다.
<br>즉, 필요할 때마다 커넥션 풀의 커넥션을 사용하고 반환하는 기법이다.
<br>이처럼 미리 만들어 놓은 커넥션을 이용하면 Connection에 필요한 비용을 줄일 수 있다.

