# Query Filter와 Operator

MongoDB CRUD작업에서 Query Filter와 Operator가 필수이다.!

### [참고 링크](https://www.mongodb.com/docs/manual/crud/)

## Create Operations

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/3eae3052-5bb5-46f2-89d7-8f56810192bf)

```javascript
# sample
db.users.insertOnd(
  {
    name: "sue",
    age: 26,
    status: "pending"
  }
)
```

데이터를 삽인하는 쿼리이다. 컬랙션이 존재하지 않을경우 삽입 작업이 진행되면서 컬렉션도 만들어진다.

몽고디비에서 삽입 메소드는 총 두개가 있으며 단일 컬렉션에만 삽입이 가능하다.

- db.collection.insertOne() 
- db.collection.insertMany()

> 삽입은 별도의 설명이 없어도 쉽게 이해할 것이라고 본다!

---

## Read Operations

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/985466d8-2700-49f5-9f2f-7187b9c184a1)

```javascript
db.users.find(
    { age: { $gt: 18 } },
    { name: 1, address: 1 }
).limit(5)
```

user collection에서 find()를 가지고 해당하는 데이터를 찾는 방식이다.

age가 18보다 큰($gt: greater than)데이터의 name필드와 address필드를 출력하고 마지막으로 몇개의 데이터를 출력하는지 표기한 것이다.

---

## Update Operations

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/b7870c0a-3d38-4577-aa54-75d0691ff6a2)

```javascript
db.users.updateMany(
    { age: { $lt: 18 } },
    { $set: { status: "reject" } }
)
```

updateMany()를 이용해서 업데이트를 진행하는데 첫번째 인자에는 age의 초건을 두 번째 인자는 업데이트 하려는 값을 넣어주었따.

---

## Delete Operations

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/caf6ff57-e953-4ab8-820a-1f1e78efa3f7)

```javascript
db.users.deleteMany(
    { status: "reject" }
)
```

deleteMany()를 이용해 필터만 넣어주면 필터에 맞는 데이터들이 지워진다.

---

## query filer와 operator 확인해보기 

Create를 제외하고 나머지 3개의 함수의 공통점은 첫번째 인자에 작업을 수행하려는 데이터의 조건을 넣어주었다. 
첫번째 인자를 ```query filter```라고 한다.

find()나 update()를 확인해보면 ```$gt```, ```$lt```를 사용한 것을 볼 수 있는데 이런것들을 ```operator```라고 한다.
각종 연산을 시행하기 위한 몽고디비에서 제공하는 도구이며 ```query filter```와 함께 사용해서 다양한 조건을 걸 수 있다.

> operator는 모두 외운다는건 힘들다. 그러니 Document를 확인해보는걸 추천한다.
> 
> ### [Query Operator 링크](https://www.mongodb.com/docs/manual/reference/operator/query/)
> ### [Update Operator 링크](https://www.mongodb.com/docs/manual/reference/operator/update/)