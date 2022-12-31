# MySql

- 가장 인기가 많은 오픈소스 관계형 데이터 베이스 (오픈소스이기 때문에 소스를 확인할 수 있다)
  - http://github.com/mysql/mysql-server
  - https://dev.mysql.com/doc/internals/en/guided-tour.html
  - https://dev.mysql.com/doc/dev/mysql-server/latest/
- 높은 접근성과 낮은 비용
- 네이버, 카카오, 토스 등 대부분의 국내 IT기업에서 가장 많이 사용됨
- sql안식 표준을 지키고 있다.

---
## MySQL 서버

![image](https://user-images.githubusercontent.com/113662725/210131221-a1cc4264-ecfd-424b-ad7b-465d78fcfccf.png)

### 1. MySQL 엔진

판단과 명령을 하는 두뇌와 같은 역할

![image](https://user-images.githubusercontent.com/113662725/210131390-f3c53a1f-a84f-4d64-96d2-6023cc4121e3.png)

#### 쿼리파서
- SQL을 파싱하여 Syntax Tree를 만듬
- 이 과정에서 문법 오류 검사가 이루어짐

#### 전처리기
- 쿼리파서에서 만든 Tree를 바탕으로 전처리 시작
- 테이블이나 컬럼 존재 여부, 접근권한 등 Semantic 오류 검사

#### 옵티마이저
- 쿼리를 처리하기 위한 여러 방법들을 만들고, 각 방법들의 비용정보와 테이블의 통계정보를 이용해 비용을 산정
- 테이블 순서, 불필요한 조건 제거, 통계정보를 바탕으로 전략을 결정 (실행 계획 수립)
- 옵티마이저가 어떤 전략을 결정하느냐에 따라 성능이 많이 달라진다.
- 가끔씩 성능이 나쁜 판단을 해 개발자가 힌트를 사용해 도움을 줄 수 있다.

#### 쿼리 실행기
- 옵티마이저가 결정한 계획대로 스토리지 엔진에 요청하는 역할
- Handler API를 이용

**참고**
- 쿼리파서, 전처리기는 컴파일 과정과 매우 유사하다.
- 하지만 SQL은 프로그래밍 언어처럼 컴파일 타임때 검증 할 수 없어 매번 구문 평가를 진행





### 2. 스토리지 엔진

판단과 명령을 수행하는 팔과 다리 역할

- 디스크에서 데이터를 가져오거나 저장하는 역할
- MySQL 스토리지 엔진은 플러그인 형태로 Handler API만 맞춘다면 직접 구현해서 사용할 수 있다.
- InnoDB, MyIsam 등 여러개의 스토리지 엔진이 존재
- 8.0대 부터는 InnoDB 엔진을 디폴트

#### InnoDB 정리(추가예정)


