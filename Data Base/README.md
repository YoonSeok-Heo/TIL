# Data Base

Data Base의 종류는 여러가지가 있으며, 그 중에 가장 기본적으로 사용하는 데이터 베이스는 관계형 데이터 베이스(Relational Data Base)이다. 
<br> [**(DB Ranking 바로가기)**](https://db-engines.com/en/ranking) DB 랭킹사이트를 확인해보면 1위부터 4위까지 RDB가 차지하고 있음을 알 수 있다.(2022년 12월 기준)
<br> 범위를 넓혀 20위까지 확인해보아도 대부분의 DB는 RDB임을 알 수 있다. 그러므로 RDB부터 다시 공부할 것이다.!

---
# 1. 데이터베이스 기본개념
## 1-1. 데이터베이스 정의

### 1-1-1. 데이터베이스
- 특정 조직의 여러 사용자가 공유하여 사용할 수 있도록 통합해서 저장한 운영 데이터의 집합

### 1-1-2. 테이터베이스 예제
- 은행 : 계좌정보, 입출금 내역 등
- 항공사 : 예약정보, 비행기 스케쥴
- 대학교 : 학생정보, 수강신청
- 온라인 쇼핑몰 : 고객기록, 주문내역
- 제조업 : 제품목록, 주문, 재고, 공급망
- 회사 인사시스템 : 직원정보, 연봉

### 1-1-3. 데이터베이스 특징
![image](https://user-images.githubusercontent.com/113662725/210138766-6ab8483e-e2c9-48c0-8326-6c3a4fec926a.png)


### 1-1-4. 데이터의 유형
- 정형 데이터(structured data)
  - 엑셀 스프레드 시트
  - 관계형데이터베이스 테이블

- 반정형 데이터(semi-structured data)
  - self-describing data : HTML, XML, JSON

- 비정형 데이터(unstructured data)
  - 정해진 구조가 없이 저장된 데이터
  - text, 멀티미디어 데이터

### 1-1-5. 파일시스템을 사용했을 때의 문제점
- 데이터 중복성 문제 : 공간 낭비
- 업데이트 및 데이터 일관성(Data consistency)유지에 어려움
- 데이터 무결성(Data integrity constraints)유지 어려움
  - 응용프로그램이 모두 체크해야함(예제. 나이 > 0)
  - 엑셀의 스프레드시트, 관계데이터베이스의 테이블
- 데이터 종속성
  - 응용프로그램이 파일 데이터 구조에 종속적
  - 파일구조가 바뀔 때마다 응용프로그램 교체 필요
**위와 같은 이유료 DBMS를 사용해야한다!**

## 1-2. 데이터베이스 관리시스템(DBMS)

### 1-2-1. 데이터베이스 관리시스템
- 파일시스템의 데이터중복과 데이터 종속 문제를 해결하기 위해 제시된 소프트웨어
- 데이터베이스의 생성과 관리를 담당
- 모든 응용프로그램은 데이터베이스 공유 가능, DBMS 통해 데이터 삽입, 수정, 검색

#### 1-2-1-1. DBMS 종류
- Oracle
- MS SQL
- MySQL 

### 1-2-2. 데이터베이스 관리시스템의 주요 기능
![image](https://user-images.githubusercontent.com/113662725/210139597-6a2d1b3f-5c62-4349-abba-a9ad96fe3ab9.png)

### 1-2-3. 데이터베이스 관리시스템의 장단점

#### 1-2-3-1. DBMS 장점
- 파일시스템의 데이터 중복 문제 해결
- 데이터 독립성 확보(데이터 구조가 바뀌어도 응용프로그램을 바꿀 필요 없음)
- 데이터 동시 공유(Concurrency control)
- 데이터 보안 향상 - 사용자 별 접근 가능한 데이터베이스 영역 제한 가능
- 데이터 무결성 유지
- 표준화 방식으로 데이터에 접근(with 데이터베이스 언어)
- 장애 발생 후 회복 시 데이터 일관성과 무결성 유지 - 트랜잭션 가능
- 응용 프로그램 개발 비용이 줄어 듬


#### 1-2-3-2. DBMS 단점
- DBMS 구매비용
- 응용프로그램이 DBMS를 통해 데이터에 접근하기 때문에 추가적인 오버헤드 발생
- 응용프로그래머가 DBMS가 어떻게 동작하는 지와 표준 데이터 언어에 대한 지식 필요
- DBMS 장애가 발생할 때 모든 응용프로그램 장애가 발생함

## 1-3. 데이터베이스 시스템 구조

### 1-3-1. 데이터베이스 시스템

#### 1-3-1-1. 데이터베이스 시스템?
- 데이터베이스에 데이터를 저장하고, 저장된 데이터를 관리하여 조직에 필요한 정보를 생성해주는 시스템
- 데이터베이스와 이를 관리하는 소프트웨어인 DBMS와 응용프로그램 모두를 칭하는 용어
- 통상적으로 DBMS와 구분없이 사용

### 1-3-2. 데이터베이스의 중요 개념

#### 1-3-2-1. 스키마(schema)
- 데이터베이스에 저장되는 데이터의 논리적구조와 제약조건을 정의한 
<br><br>
![image](https://user-images.githubusercontent.com/113662725/210140637-990cff16-ca28-488c-8898-728f229dad6e.png)

#### 1-3-2-1. 인스턴스(instance)
- 정의된 스키마에 따라 데이터베이스에 실제로 저장된 값

### 1-3-3. 3단계 데이터베이스 구조 
- 미국의 기관인 ANSI/SPARC에서 데이터베이스의 내부 구조를 감추고 일반 사용자가 데이터베이스를 쉽게 이용하고 사용할 수 있게 하는 구조 제안 

#### 1-3-3-1. Physical Level - Internal Schema(내부 스키마)
- 실제 데이터가 데이터베이스에 어떻게 저장되는지 기술
- 파일에 저장되는 레코드의 구조(사이즈), 압축방법, 암호화 방법 등

#### 1-3-3-2. Conceptual Level - Conceptual Schema(개념 스키마)
- 일반적인 스키마
- 데이터베이스를 개념적으로 표현 그리고 데이터 테이블 간의 관계를 기술
- 물리적 데이터 독립성
- 
#### 1-3-3-3. External Level - External Schema(외부 스키마)
- 사용자가 보는 View Level
- 사용자에게 view의 형태로 필요한 정보만 보여주는 단계
- 각 사용자들이 개별 요구 사항에 따라 다른 view를 정의해서 데이터를 볼 수 있다.
- 데이터베이스 하나에 여러개의 외부 스키마 가능
- 개념스키마와 논리적 데이터 독립성

## 1-4. 데이터베이스 시스템 언어

### 1-4-1. 데이터베이스 사용자

#### 1-4-1-1. 데이터베이스 사용자 (DBA)
- 데이터베이스 시스템을 운영하고 관리
- 데이터베이스를 설계 및 구축, 제어
- DBMS 자체를 물론 데이터베이스 구축, 관리에 해박한 지식과 많은 경험이 요구됨

#### 1-4-1-2. 응용프로그래머
- 데이터베이스 언어를 이용하여 응용프로그램을 작성

#### 1-4-1-3. 최종 사용자 (End user)
- 데이터베이스에 접근하여 데이터를 조작
- casual end user : 데이터베이스 언어를 이용해 데이터베이스 접근 및 조작
- native end user : 데이터베이스 언어를 쓰지 않고 응용프로그램을 통해 데이터베이스 접근


### 1-4-2. 데이터 언어
- 사용자가 데이터베이스를 구축하고 이에 접근하기 위해 데이터베이스 관리 시스템과 통신하는 수단

### 1-4-3. 데이터 언어 종류

#### 1-4-3-1. 데이터 정의어(DDL)
- 스키마 구조와 제약조건 등을 정의, 삭제, 수정

```sql
create table instructor ( 
  ID char(5), 
  name varchar(20),
  deptname varchar(20), 
  salary numeric(8,2), primary key(ID)
);
```

#### 1-4-3-2. 데이터 조작어(DML)
- 데이터 삽입, 삭제, 수정, 검색

#### 1-4-3-3. 데이터 제어어(DCL)
- 내부적으로 필요한 규칙(보안, 무결성 체크) 등을 정의하는 데이터 언어
- 데이터베이스 백업하거나 백업이미지로 회복(recover)하는 명령어

### 1-4-4. 질의 처리기

#### 1-4-4-1. DDL Interpreter
- DDL로 작성된 스키마의 정의를 해석, 데이터 딕셔너리(Catalog)에 저장

#### 1-4-4-2. DML Compiler
- DML로 작성된 데이터 처리 요구를 processing engine이 이해할 수 있는 코드로 해석하여 plan을 작성

#### 1-4-4-3. DML(Query) Processing Engine
- 컴파일된 plan을 데이터베이스에 실제로 실행



-----
# 2. 관계형 데이터 모델

## 2.1. 데이터 

### 2.1.1. 데이터 모델링(data modeling)
- 현실 세계에 존재하는 데이터를 컴퓨터 세계의 데이터베이스로 옮기는 과정
- 데이터베이스 설계의 핵심 과정

### 2.1.2. 데이터 모델링 3단계
![image](https://user-images.githubusercontent.com/113662725/210243632-63dfab78-a592-4274-bf90-b31dd1b06040.png)

#### 2.1.2.1. 개념적 데이터 모델링
- 현실세계를 추상화하여 중요 데이터를 개념 세계로 추출해 가는 과정
- 결과물로 개념적 데이터 모델 (객체 - 관계 (E - R)모델)

#### 2.1.2.2. 논리적 데이터 모델링
- 개념 세계의 데이터를 데이터베이스가 저장할 구조로 변환하는 과정
- 결과물로 관계 데이터 모델

#### 2.1.2.3. 물리적 데이터 모델링
- 논리 데이터 모델이 실제 데이터베이스 저장소에 저장되는 저장 구조(테이블, 컬럼)로 변경

### 2.1.3. 데이터 모델링 예제
![image](https://user-images.githubusercontent.com/113662725/210244807-954c3bcc-147b-4add-b745-492748457fd6.png)


## 2.2. 관계형 데이터 모델

### 2.2.1. 관계 데이터 모델
![image](https://user-images.githubusercontent.com/113662725/210245075-589c4914-f09a-4337-8c03-3af8cb95bfc8.png)

### 2.2.2. 릴레이션의 특징
- 튜플의 유일성 : 동일한 튜플이 존재할 수 없다.
- 튜플의 무순서 : 튜플 사이의 순서는 무의미
- 속성(애트리뷰트)의 무순서 : 속성 사이의 순서는 무의미
- 속성의 원자성(Atomic) : 애트리뷰트 값으로 하나(나누어지지 않는)의 값만 가짐

### 2.2.3. Key의 종류

#### 2.2.3.1. Key?
- 릴레이션에 튜플을 구별하는 역할을 하는 속성 또는 속성의 집합

#### 2.2.3.2. Super Key 
- 튜플을 구별하기 위해 유일성을 제공할 수 있는 속성 또는 속성의 집합
- ex) {ID}, {ID, name}

#### 2.2.3.3. Candidate Key
- Super Key 중에서 개수가 가장 작은 키
- ex) {ID}

#### 2.2.3.4. Primary Key
- Candidate key 중에서 디자인을 고려하여 선택된 키

#### 2.2.3.5. Foreign Key
- 다른 릴레이션의 primary key를 참조하는 속성 또는 속성의 집합

#### 2.2.3.5 PK와 FK의 관계
![image](https://user-images.githubusercontent.com/113662725/210246652-d493288a-2b37-4386-9fa7-e4ba5043f631.png)


### 2.2.4. 관계데이터 모델의 제약조건

#### 2.2.4.1. 도메인 무결성 제약조건 (domain integrity constraint)
- 릴레이션 내의 튜플들이 각 속성의 도메인에 지정된 값만 가져야 한다.

#### 2.2.4.2. 개체 무결성 제약조건 (entity integrity constraint, = primary key constraint)
- 기본키를 구성하는 모든 속성은 null를 가질 수 없다.

#### 2.2.4.3. 참조 무결성 제약조건 (referential integrity constraint, =foreign key constraint)
- Foreign key는 참조하는 릴레이션의 primary key 속성 값 중 하나여야 한다 (null 가능)



### 2.3. 관계대수 - 1

### 2.3.1. 관계데이터 연산

모든 DBMS는 데이터 처리를 위해 하나 이상의 데이터 언어를 제공

#### 2.3.1.1. Formal query language
- 수학기호(notation)을 사용하여 데이터 처리를 기술한 언어
- 새로운 언어의 개념과 유용성을 검증하는 기준
- 관계 대수(Relation algebra)

#### 2.3.1.2. Commercial language
- 수학적인 원리를 기반으로 사용하기 쉽게 만들어진 단어
- 관계 대수로 만들어진 모든 질의가 표현 가능 - Relationally complete
- SQL

#### 2.3.1.3. 관계 대수 연산자 (Relational Algebra Operations)
- 피연산자로 하나 또는 두 개의 릴레이션(Unary and binary operaions)
- 각 연산자의 연산 결과는 새로운 릴레이션-연산의 합성(compose)&체이닝(chaining)가능

#### 2.3.1.4. 연산자 종류
- select
- project
- union
- difference
- intersection
- cartesian product
- natural join
- theta join
- outer join

### 2.3.2. select operaion

![image](https://user-images.githubusercontent.com/113662725/210572202-149a4de8-f45e-4601-8a99-f73c29e2be9c.png)

### 2.3.3. project operaion

![image](https://user-images.githubusercontent.com/113662725/210572258-fbdce1d8-7dcf-460d-b90c-0e2c9d59e172.png)

### 2.3.4. union operation(합집합)

![image](https://user-images.githubusercontent.com/113662725/210572317-ef8cd9a3-3b3a-4bdf-b4b3-9b240067d194.png)

### 2.3.5. difference operation(차집합)

![image](https://user-images.githubusercontent.com/113662725/210572424-04fcff65-a8cc-45b4-a96c-6608324f7010.png)

### 2.3.6 intersection operation(교집합)

![image](https://user-images.githubusercontent.com/113662725/210572565-dbfd24e0-ba0e-48bb-bbcd-c8875864aea1.png)

## 2.4. 관계 대수 - 2

### 2.4.1 Cartesioan product operation

![image](https://user-images.githubusercontent.com/113662725/210573645-665a4082-a029-4ea8-b2e9-ecacbfd95f84.png)

### 2.4.2 Natural join operation

![image](https://user-images.githubusercontent.com/113662725/210576419-2a7c696e-038f-46a8-98df-83e67720d950.png)

### 2.4.3 Theta join operation

![image](https://user-images.githubusercontent.com/113662725/210576822-600e4a40-dc5e-458b-a4d3-9767ae0a78b4.png)

### 2.4.4 Outer join operation

![image](https://user-images.githubusercontent.com/113662725/210577305-e8a0cb14-d28d-4891-a395-eac4c93e7f0e.png)

---

# 3. SQL 데이터베이스 언어

## 3.1. SQL 소개

### 3.1.1. SEQUEL(Structured English QUEry Language)
- 1974년 IBM San Jose Research Lab에서 연구용 DBMS 인 SYSTEM R 를 위한 언어로 개발된 언어 연산자로 하나 또는 두 개의 릴레이션 (Unary and binary operations)
**SQL (Structured Query Language) : SEQUEL에서 이름이 바뀜**

### 3.1.2. SQL?
- 대부분의 DBMS 는 SQL-92 표준의 대부분을 지원하고 추후의 표준에 지정된 기능을 추가하여 지원하고 고유의 기능도 제공함
- 최근 SQL-2011 에 추가된 기능 예제
  - Temporal table (system versioned temporal table)
  - 데이터 변경 내용의 전체 기록을 유지해 간편한 지정 시간 분석을 허용하도록 설계된 사용자 테이블의 종류
### 3.1.3. SQL 분류

#### 3.1.3.1. DDL (Data Definition Language)
- 테이블을 생성하고 변경, 제거하는 기능을 제공

#### 3.1.3.2. DML (Data Manipulation Language)
- 테이블에 새 데이터를 삽입하거나, 테이블에 저장된 데이터를 수정, 삭제하는 기능을 제공

#### 3.1.3.3. SELECT
- 테이블 데이터를 조회하는 기능을 제공

#### 3.1.3.4. DCL (Data Control Language)
- 보안을 위해 데이터에 대한 접근 및 사용권한을 조절하는 기능을 제


![image](https://user-images.githubusercontent.com/113662725/210800621-954a7ad5-8e60-462c-ab43-47cca716f85e.png)


## 3.2. DDL : 데이터 정의

