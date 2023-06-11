# SQL vs MQL

RDB에서도 SQL이라는 질의 형태를 가지고 있고 MongoDB에서도 MQL(MongoDB Query Language)이라는 쿼리 랭귀지가 갖춰져있다.

### [SQL과 MQL의 비교 내용 확인](https://www.mongodb.com/docs/manual/reference/sql-comparison/)

간단한 쿼리드르을 피교해보자.

## createCollection

RDB에서 database를 생성하는 과정과 같다.

```javascript
db.createCollection("people")
```

---

## Insert

RDB에서는 Table의 형태를 미리 정해주고 Insert작업을 해야하지만 NoSQL인 MongoDB에서는 Insert의 형태 그대로 들어가게 된다.(따로 테이블을 생성할 필요가 없다는 뜻.)

```javascript
db.people.insertOne( {
    user_id: "abc123",
    age: 55,
    status: "A"
 } )
```

insertOne 또는 insertMany를 이용해 데이터를 넣는데 기본 키를 정해주지 않는다면 _id필드에 자동으로 생성된다.

---

## updateMany

RDB에서 alter와 같다 테이블의 구조 수정!

### $set

필드 추가

```javascript
db.people.updateMany(
    { },
    { $set: { join_date: new Date() } }
)
```

### $unset

필드 삭제

```javascript
db.people.updateMany(
    { },
    { $unset: { "join_date": "" } }
)
```

--- 

## ceateIndex

RDB에서 인덱스 작업니다.

### Index 추가

```javascript
db.people.createIndex( { user_id: 1 } )
```

### 여러개 인덱스 추가

1과 -1은 내림차순 오름차순 차이이다.

```javascript
db.people.createIndex( { user_id: 1, age: -1 } )
```

# 결론

위에 링크해준 사이트를 들어가보면 SQL 해당되는 MQL을 찾아볼 수 있다.!!

너무 많아서 다 정리는 불가능 
