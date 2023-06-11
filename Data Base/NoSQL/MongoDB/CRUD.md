# CRUD

간단한 CRUD를 공부해보자


## 1. insertOne()

한 개의 데이터 넣기

```javascript
db.employees.insertOne({
    name: "lake",
    age: 21,
    dept: "Database",
    joinDate: new ISODate("2022-10-01"),
    salary: 400000,
    bonus: null
})
```

---

## 2. insertMany()

여러 개의 데이터 넣기

```javascript
db.employees.insertMany([
    {
        name: "ocean",
        age: 45,
        dept: "Network",
        joinDate: new ISODate("1999-11-15"),
        salary: 100000,
        resignationDate: new ISODate("2002-12-23"),
        bonus: null
    },
    {
        name: "river",
        age: 34,
        dept: "DevOps",
        isNegotiating: true
    }
])
```

ocean의 resignationDate는 퇴사 날짜이며, river는 입사전 연봉협상중이라고 가정하자(isNegotiating: true)

---


## 3. find()

데이터가 정상적으로 들어간 것을 볼 수 있다.

```mongodb-json
Atlas atlas-mluprb-shard-0 [primary] test> db.employees.find()
[
  {
    _id: ObjectId("6485cb50a43db35002dfab10"),
    name: 'lake',
    age: 21,
    dept: 'Database',
    joinDate: ISODate("2022-10-01T00:00:00.000Z"),
    salary: 400000,
    bonus: null
  },
  {
    _id: ObjectId("6485cb5ba43db35002dfab11"),
    name: 'ocean',
    age: 45,
    dept: 'Network',
    joinDate: ISODate("1999-11-15T00:00:00.000Z"),
    salary: 100000,
    resignationDate: ISODate("2002-12-23T00:00:00.000Z"),
    bonus: null
  },
  {
    _id: ObjectId("6485cb5ba43db35002dfab12"),
    name: 'river',
    age: 34,
    dept: 'DevOps',
    isNegotiating: true
  }
]
```

---

## 4. insertOne vs insertMany

insertOne과 insertMany의 속도차이를 비교해보자

MongoDB는 JS문법을 사용할 수 있다.


### 4.1. insertOne 300번하기

```javascript
for (i = 0; i < 300; i++){
    db.insertTest.insertOne({a: i})
}
```

### 4.2. insertMany로 한번에 넣기

```javascript
var docs = [];
for (i = 0; i < 300; i++){
    docs.push({a: i})
}
db.insertTest.insertMany(docs);
```

> 해보면 안다 고작 300번 차이지만 엄청난 속도차이가 난다. 따로 초는 구하지 않았따..

---

## 5. updateOne

데이터 한 개 업데이트하기.

```javascript
db.employees.updateOne(
    {name: "river"},
    {
        $set: {
            salary: 350000,
            dept: "Database",
            joinDate: new ISODate("2022-12-31")
        },
        $unset: {
            isNegotiating: ""
        }
    }
)
```

위 에서 연봉협상중인 river가 연봉협상이 끝나고 입사한 상황이다.

첫번째 인자로 name에 river가 들어간다. 이 부분은 필더 부분이다. ```$set```으로 필드를 만들어주고 ```$unset```으로 필드값을 빼주자.

```mongodb-json
# 결과 
{
    _id: ObjectId("6485cb5ba43db35002dfab12"),
    name: 'river',
    age: 34,
    dept: 'Database',
    joinDate: ISODate("2022-12-31T00:00:00.000Z"),
    salary: 350000
}
```

---

## 6. updateMany 

급여를 10퍼센트 올려줄 것이다.

```javascript
db.employees.updateMany(
    { resignationDate: { $exists: false }, joinDate: { $exists: true }},
    { $mul: { salary: Decimal128("1.1")}}
)
```
- ```$exists```: 필드가 있나 없나를 확인한다. 퇴사일 필드는 없고 입사일 필드가 있을경우가 필터 조건이다.
- ```$mul```: Multiply 기능이다. salary에 1.1만큼 곱해준다.
- ```Decimal128```: 소수를 곱하게 되면 소수점 자리가 길게 나와서 Decimal128 타입으로 수정한다.

```mongodb-json
# 결과
[
  {
    _id: ObjectId("6485cb50a43db35002dfab10"),
    name: 'lake',
    age: 21,
    dept: 'Database',
    joinDate: ISODate("2022-10-01T00:00:00.000Z"),
    salary: Decimal128("440000.0"),
    bonus: null
  },
  {
    _id: ObjectId("6485cb5ba43db35002dfab11"),
    name: 'ocean',
    age: 45,
    dept: 'Network',
    joinDate: ISODate("1999-11-15T00:00:00.000Z"),
    salary: 100000,
    resignationDate: ISODate("2002-12-23T00:00:00.000Z"),
    bonus: null
  },
  {
    _id: ObjectId("6485cb5ba43db35002dfab12"),
    name: 'river',
    age: 34,
    dept: 'Database',
    joinDate: ISODate("2022-12-31T00:00:00.000Z"),
    salary: Decimal128("385000.0")
  }
]
```

퇴사자가 아닌 lake와 river의 급여가 올랐다.

---

## 7. updateMany null처리

직원에게 보너스를 줄 예정이다. 기존에 다니고 있던 직원(bonus 필드가 null인 직원)에만 주려고 한다. 

```javascript
db.employees.updateMany(
    { resignationDate: { $exists: false }, bonus: null },
    { $set: { bonus: 100000 }}
)
```

```mongodb-json
# 결과
[
  {
    _id: ObjectId("6485cb50a43db35002dfab10"),
    name: 'lake',
    age: 21,
    dept: 'Database',
    joinDate: ISODate("2022-10-01T00:00:00.000Z"),
    salary: Decimal128("440000.0"),
    bonus: 100000
  },
  {
    _id: ObjectId("6485cb5ba43db35002dfab11"),
    name: 'ocean',
    age: 45,
    dept: 'Network',
    joinDate: ISODate("1999-11-15T00:00:00.000Z"),
    salary: 100000,
    resignationDate: ISODate("2002-12-23T00:00:00.000Z"),
    bonus: null
  },
  {
    _id: ObjectId("6485cb5ba43db35002dfab12"),
    name: 'river',
    age: 34,
    dept: 'Database',
    joinDate: ISODate("2022-12-31T00:00:00.000Z"),
    salary: Decimal128("385000.0"),
    bonus: 100000
  }
]
```

river는 기존에 다니던 직원(bonus 필드가 null인 직원)이 아니였는데 보너스가 지급된걸 볼 수 있다.
이 상황으로 보았을 때 mongodb는 필드가 없을 경우랑 필드가 null인경우를 같게 보는걸 확인할 수 있다.

> #### river의 bonus필드 값을 다시 없에준다.
>
> ```javascript
> db.employees.updateOne(
>     { name: "river" },
>     { $unset: { bonus: "" }}
> )
> ```

#### 위와 같은 경우는 ```$exists```를 이용해 없는 필드를 확인해야한다.

```javascript
db.employees.updateMany(
    { resignationDate: { $exists: false }, bonus: { $exists: true }},
    { $set: { bonus: 200000 }}
)
```

---

## 8. deleteOne

단건 삭제를 해보자.

```javascript
db.employees.deleteOne({name: "river"})
```

---

## 9. deleteMany

많은 데이터 삭제

아레는 모든 데이터를 삭제하는 쿼리이다. 

```javascript
db.employees.deleteMany({})
```