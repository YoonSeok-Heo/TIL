# Intro
처음 들어본다 하지만 성능에 영향을 끼친다고 한다 공부해보자


## Spring Cache Abstraction
- 애플리케이션에 "투명하게(transparently)" 캐시를 넣어주는 기능
- 메소드, 클래스에 적용 가능
- 캐시 인프라는 스프링 부트 자동 설정으로 세팅되고, 프로퍼티로 관리 가능

### 캐시가 투명하게(transparently) 자리잡는다?
- 데이터를 통신하는 시스템 쌍방이 캐시의 존재를 모른다는 의미
- "캐시가 있건 없건, 시스템의 기대 동작은 동일해야 한다."
- 캐시의 목표: 오로지 "성능"
- 캐시의 개념과 목적에 부합하는 성질이자, 조건


### 캐시는 왜쓰지? 반복잡업에서 사용한다.
- 잘 바뀌지 않는 정보를 외부 저장소에서 반복적으로 읽어온다면
- 기대값이 어차피 같다면
- 캐싱해서 성능 향상, I/O 감소

### 사용방법!!

![image](https://user-images.githubusercontent.com/51642448/155523427-6565c233-5a68-45ea-80c4-82f465b7f9e2.png)
간단하게 등록한다.

![image](https://user-images.githubusercontent.com/51642448/155523497-c97ead12-7e37-4518-9a9a-6bc540cf5865.png)
@Cacheable()를 원하는 메소드에 넣어준다.<br/>
값을 넣어주면 알아서 캐싱을 해준다. 


### 캐싱에서 생각해야 하는 것들
- 무엇을 캐시할까?
- 얼마나 오랫동안 캐시할까?
- 언제 캐시를 갱신할까?

### 주요 캐시 어노테이션
- @EnableCaching : 캐시 활성화
- @Cacheable : 캐시 등록
- @CacheEvict : 캐시 삭제
- @CachePut : 캐시 갱신
