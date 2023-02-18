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

- 데이터베이스 구조를 정의하고 변경하는 기능을 제공하는 언어
- Create: 새로운 데이터베이스 오브젝트들을 생성(schema, table, view 등)
- Alter: 존재하는 오브젝트의 정의를 변경
- Drop: 존재하는 오브젝트를 데이터베이스에서 삭제

### 3.2.1. Create Table syntax
```
CREATE TABLE <table name> { <column definitions> | <like_table_clause> } 
  <column definitions> := (<data column definition>, ...)
  <data column definition> := <column name> <data type> [<default value>] [ <integrity constraint>]
  <default value> := DEFAULT <value>
  <integrity constraint> := { UNIQUE | <primary key> | <reference contraint> | <check constraint>}
```

### 3.2.2. Example
```sql
CREATE TABLE student (
  ID varchar(8) primary key,
  name varchar(20) not null,
  dept_name varchar(20),
  grade smallint,
  foreign key (dep_name) reference department
);
```

### 3.2.3. 데이터타입 on Sql

|데이터 타입|설명|
|---|---|
|CHAR(n)|길이가 n으로 고정인 문자열|
|VARCHAR(n)|최대 길이가 n인 가변 길이의 문자열|
|BIGINT|8Bytes 정수|
|INT or INTEGER|4Bytes 정수|
|SAMLLINT|2Bytes 정수|
|NUMERIC(p, s) or DECIMAL (p, s)|고정 소수점 실수 <br>p: 소수점을 제외한 전체 길이, s : 소수점 이하 숫자의 길이|
|FLOAT(n)|길이가 n인 부동소수점 실수|
|REAL|부동 소수점 실수|
|DATE|연, 월, 일로 표현되는 시간|
|TIME|시, 분, 초로 표현되는 시간|
|TIMESTAMP|DATE + TIME|
|BLOB|Binary large objects|
|CLOB|Charater large objects|

#### 3.2.3.1 Decimal example
![image](https://user-images.githubusercontent.com/113662725/212526697-6dc7fc21-9a8f-46b7-8abe-5d16067d9536.png)


#### 3.2.3.2. default clause
**SQL분류\<default value\> := DEFAULT \<value\>**
- INSERT문에 컬럼에 대한 값이 지정되지 않았을 때 null대신에 <value> 사용
- credit INT DEFAULT 0
- contact char(20) DEFAULT 'james'

#### 3.2.3.3. contraint clause
**\<integrity constraint\> := {UNIQUE | \<primary key\> | \<reference contraint\> | \<check constraint\>}**
- UNIQUE : 대체 키를 지정하고 유일성을 가지는 컬럼임을 지정
- PRIMARY KEY : ROW를 식별할 수 있는 기본 키 ( unique + not null )
- <reference constraint> : 참조하는 테이블을 지정하는 외래키 설정 (다음 장에서 설명) 
- <check constraint> : column 값의 가능 domain 지정 (다음 장에서 설명

#### 3.2.3.4. CREATE TABLE LIKE
**CREATE TABLE \<table name\> \<like_table_clause\> \<like_table_clause\> <br>:= LIKE \<table_name\> \[WITH DATA\]**
<br>Ex) CREATE TABLE customer_backup LIKE customer WITH DATA;

#### 3.2.3.5 ALTER TABLE
-새로운 컬럼 추가
  - ALTER TBLE <table_name> ADD <column_name> <column_type> [not null] [<default_value>
  - Ex) ALTER TABLE customer ADD reg_date DATE;
- 기존 컬럼 삭제
  - ATLER TABLE <table_name> DROP COLUMN <column_name>
  - Ex) ALTER TABLE customer drop age;
- 새로운 제약조건 추가
  - ALTER TABLE <table_name> ADD CONSTRAINT <constraint_name> <contraints>
  - ATLER TBLE customer ADD CONSTRAINT set_pri_key PRIMARY KEY (id)
- 제약조건 삭제
  - ALTER TBLE <tble_name> DROP CONTRAINT set_pri_key

#### 3.2.3.6 DROP TABLE
- 테이블 데이터 및 catalog 삭제
- DROP TABLE <table_name>

## 3.3. SELECT: 데이터 조회
- 테이블에서 원하는 데이터를 검색하는 SQL

### 3.3.1. SELECT
- SELECT clause 는 산술식 (arithmetic expression) 을 포함 가능
  - +, -, * and /
  - SELECT ID, name, salary/12 from employee

#### 3.3.1.1. Syntax
- SELECT \[DISTINCT\] \{* | \<column_name\>, ...\} FROM \<table list\>;

#### 3.3.1.2. Example
- SELECT DISTINCT dept_name from instructor;
  - DISTINCT : 결과 테이블에서 중복된 레코드를 제거하는 키워드

### 3.3.2. SELECT with WHERE
- 테이블에서 조건에 맞는 데이터만 검색하기 위해 WHERE clause 이용

#### 3.3.2.1. Syntax
- SELECT \[DISTINCT\] \{* | \<column_name\>, ...\} FROM \<table list\> \[ WHERE condition \];
- Condition은 비교 연산자 (=, <>, <, >, <= , =>) 와 논리연산자 (AND, OR, NOT)

#### 3.3.2.2. Example
```sql
SELECT name from employee WHERE dept_name = ‘SALES’ and salary > 100000;
```

#### 3.3.2.3. LIKE keyword
- 문자열 컬럼인 경우 부분적인 매칭으로 조건 검색 가능
- % : 0개 이상의 문자 , _ : 1개의 문자
```sql
SELECT * from employee where name like ‘이%’;
```
```sql
SELECT * from employee where name like ‘이정_’
```

### 3.3.3. NULL Value
- null은 empty가 아닌 unknown
- null값과 산술 연한 및 비교 연산 결과는 null
  - 5 + null returns null
  - null > 5 returns null
  - null = null returns null
  

- null과 논리 연산 결과
  - OR : (null OR true) = true, (null or false) = null, (null or null) = null
  - AND : (true AND null) = null, (false AND null) = false, (null and null) = null
  - NOT : (NOT null) = null
  

- null 값을 가진 row을 찾기 위한 조건식 – IS NULL
```sql
SELECT id, name FROM customer WHERE aga IS NULL;
```
  
### 3.3.4. ORDER BY
- SELECT의 검색 결과를 정렬하기 위해 ORDER BY 키워드 사용

#### 3.3.4.1. Syntax
```sql
SELECT [DISTINCT] {* | <column_name>, ...} 
FROM <table list> [ WHERE condition ]
[ ORDER BY <column_name>, ... [ASC | DESC] ];
```
- ASC: 오름 차순 정렬 (default), DESC: 내림 차순 정렬

#### 3.3.4.2. Example
```sql
SELECT id, name, age, address FROM customer ORDER BY age DESC;
```
```sql
SELECT id, name, dept_name, salary FROM employee ORDER BY salary;
```
```sql
SELECT id, product, quantity, order_date FROM sales_orders
ORDER BY order_date ASC, quantity DESC;
```

### 3.3.5. Aggregation Function

#### 3.3.5.1. 특정 컬럼의 값을 통계적으로 계산한 결과를 보여주는 SQL 함
- count: 컬럼 값의 개수
- max: 컬럼 값의 최대값
- min: 컬럼 값의 최소값
- sum: 컬럼 값의 합계
- avg: 컬럼 값의 평균

#### 3.3.5.2. COUNT, MAX, MIN은 LOB 타입을 제외한 모든 타입에서 사용 가능 SUM, AVG는 숫자 데이터만 가능
```sql
SELECT sum (salary) FROM employee WHERE dept_name = ‘SALES’;
```
```sql
SELECT count (distinct id) FROM sales_orders WHERE product = ‘computer’;
```

### 3.3.6. Null value in aggregation function

#### 3.3.6.1. SELECT sum (salary) FROM employee;
- 만약 컬럼 값이 null이면 합을 계산할 때 무시
- 만약 컬럼 값이 모두 null이면 null이 출력된다.

#### 3.3.6.2. MAX, MIN, SUM, AVG 경우 null이 아닌 값으로만 계산
- 모든 null인 경우 null을 리턴

#### 3.3.6.3. COUNT : null이 아닌 값의 개수를 출력
- 컬럼이 모두 null인 경우 0리턴

#### 3.3.6.4. 테이블에 레코드가 없는 경우
- sum, avg, min, max: return null
- count: return 0

### 3.3.7. GROUP BY

#### 3.3.7.1 GROUP BY
- 테이블에서 특정 컬럼의 값이 같은 rows를 모아 그룹을 만들고, 그룹별로 검색하기 위해 사용되는 keyword
- 그룹에 대한 조건을 추가하려면 HAVING 키워드를 이용
  - Where 조건 : 레코드를 grouping 하기 전에 조건을 검색
  - Having 조건 : 레코드들을 grouping 후에 그룹에 대한 조건을 검색
  
#### 3.3.7.2. Syntax
```sql
SELECT [DISTINCT] {* | <column_name>, ...} 
FROM <table list> [ WHERE condition ]
  [ GROUP BY <column_name list> [ HAVING condition ] ]
  [ ORDER BY <column_name_list> [ASC | DESC] ]
```
#### 3.3.7.3. Aggregation function을 제외한 SELECT문의 컬럼은 GROUP BY 리스트에 포함된 컬럼들만 가능
```sql
SELECT dept_name, ID, avg (salary) from employee group by dept_name;
```
- Return error : SELECT list is not in GROUP BY clause and contains nonaggregated column

--- 
#### 3.3.7.4. Example
```sql
SELECT dept_name, avg (salary) FROM employee GROUP BY dept_name;
```
```sql
SELECT dept_name, avg (salary) FROM employee GROUP BY dept_name
HAVING avg(salary) > 50000000;
```
```sql
SELECT company_name, 
      count(product_name) as number_of_product, 
      max(price) as highest_price
FROM product GROUP BY company_name;
```
```sql
SELECT company_name, 
      count(product_name) as number_of_product, 
      max(price) as highest_price
FROM product GROUP BY company_name
HAVING COUNT(product_name) >= 3;
```

### 3.3.8. Set

#### 3.3.8.1. Union : r∪s
```sql
select a, b from r
union
select a, b, from s;
```

**r**
|A|B|
|---|---|
|a|1|
|a|2|
|b|1|

**s**
|A|B|
|---|---|
|a|2|
|b|3|

**r∪s**
|A|B|
|---|---|
|a|1|
|a|2|
|b|1|
|b|3|

### 3.3.8.2. Difference (except) : r - s
```sql
select a, b from r
except
select a, b, from s;
```

**r**
|A|B|
|---|---|
|a|1|
|a|2|
|b|1|

**s**
|A|B|
|---|---|
|a|2|
|b|3|

**r-s**
|A|B|
|---|---|
|a|1|
|b|1|

#### 3.3.8.3. intersection : r ∩ s
```sql
select a, b from r
intersect
select a, b, from s;
```
**r**
|A|B|
|---|---|
|a|1|
|a|2|
|b|1|

**s**
|A|B|
|---|---|
|a|2|
|b|3|

**r-s**
|A|B|
|---|---|
|a|2|


### 3.3.9. JOINs

#### 3.3.9.1. Cartesian Product : r × s
```sql
SELECT * FROM r, s
```
**r**
|A|B|
|---|---|
|a|1|
|b|5|

**s**
|C|D|E|
|---|---|---|
|c|100|5|
|d|200|6|
|e|300|7|


**r X s**
|A|B|C|D|E|
|---|---|---|---|---|
|a|1|c|100|5|
|a|1|d|200|6|
|a|1|e|300|7|
|b|5|c|100|5|
|b|5|c|200|6|
|a|5|c|300|7|

#### 3.3.9.2. Natural join : r ⋈ s
```sql
SELECT * from r natural join s;
```
- 많은 DBMS가 지원을 안함

**r**
|A|B|C|D|
|---|---|---|---|
|a|1|a|a|
|b|2|y|a|
|y|4|b|b|
|a|1|y|a|
|b|2|b|b|
|b|5|

**s**
|B|D|E|
|---|---|---|
|1|a|a|
|3|a|b|
|1|a|y|
|2|b|b|
|3|b|e|


**r ⋈ s**
|r.A|r.B|r.C|r.D|s.E|
|---|---|---|---|--|
|a|1|a|a|a|
|a|1|a|a|y|
|a|1|y|a|a|
|a|1|y|a|y|
|b|2|b|b|b|


#### 3.3.9.3. Theta join : r ⋈θ s
```sql
select r.a, r.b, r.c, r.d, s.e
from r join s on r.b = s.b and r.d = s.d;
```
- 많은 DBMS에서 사용되는 inner join 실행문


**r**
|A|B|C|D|
|---|---|---|---|
|a|1|a|a|
|b|2|y|a|
|y|4|b|b|
|a|1|y|a|
|b|2|b|b|
|b|5|

**s**
|B|D|E|
|---|---|---|
|1|a|a|
|3|a|b|
|1|a|y|
|2|b|b|
|3|b|e|


**r ⋈(r.a=a ∧ r.b=s.b) s**
|r.A|r.B|r.C|r.D|s.E|
|---|---|---|---|--|
|a|1|a|a|a|
|a|1|a|a|y|
|a|1|y|a|a|
|a|1|y|a|y|
|b|2|b|b|b|







