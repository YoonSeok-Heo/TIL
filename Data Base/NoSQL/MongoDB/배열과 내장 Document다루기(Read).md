# 배열과 내장 Document다루기(Read)

## 환경

MongoDB: Atlas
MongoSh: 1.9.1

샘플데이터를 활용한 학습을 할 것이다.

### [아틀라스 설치 방법](https://github.com/YoonSeok-Heo/TIL/blob/main/Online%20Class/High%20Capacity%20System/Part2.%20MongoDB/Part2-007.%20MongoDB%20Atlas.md)
### [샘플데이터 적재 방법](https://github.com/YoonSeok-Heo/TIL/blob/main/Online%20Class/High%20Capacity%20System/Part2.%20MongoDB/Part2-007.%20MongoDB%20Atlas.md#tip)


## 시작

sample_supplies collection을 사용할 것이다.

```javascript
use sample_supplies
```

## Document를 이용한 조회

customer의 key값으로 데이터를 조회할 것이다.

```mongodb-json
customer: { gender: 'M', age: 51, email: 'worbiduh@vowbu.cg', satisfaction: 5 }
```

```javascript
db.sales.findOne({
    customer: { 
        gender: 'M', 
        age: 51, 
        email: 'worbiduh@vowbu.cg', 
        satisfaction: 5 
    },
})
```

위와 같은 메소드를 이용하면 문제없이 출력되는 것을 확인 할 수 있다. 하지만 customer 내부의 키값의 순서가 달라지면 검색이 안된다.



### 내장도큐먼트에 바로 접근한 검색

매번 이렇게 내장 도큐먼트의 전체를 이용해 검색을 한다면 번거로운 작업이 되니까 내장도큐먼트 내부의 필드 값을 이용해 검색을 할 것이다.

필드로 접근하는 방법은 "."을 이용한다.


```javascript
db.sales.findOne({
    "customer.email": "worbiduh@vowbu.cg"
})
```
- email을 이용한 검색이다.

```javascript
db.sales.findOne({
    "customer.age": {$lt: 20}
})
```
- age를 이용한 검색


---

## 배열 다루기

사용할 데이터를 우선 넣어준다. 몽고디비 홈페이지 튜토리얼에 있는 값에서 마지막줄을 추가해줬다.

```javascript
db.inventory.insertMany([
    { item: "journal", qty: 25, tags: ["blank", "red"], dim_cm: [ 14, 21 ] },
    { item: "notebook", qty: 50, tags: ["red", "blank"], dim_cm: [ 14, 21 ] },
    { item: "paper", qty: 100, tags: ["red", "blank", "plain"], dim_cm: [ 14, 21 ] },
    { item: "planner", qty: 75, tags: ["blank", "red"], dim_cm: [ 22.85, 30 ] },
    { item: "postcard", qty: 45, tags: ["blue"], dim_cm: [ 10, 15.25 ] },
    { item: "postcard", qty: 45, tags: ["blue", "red"], dim_cm: [ 13, 14 ] }
 ]);
```

---

### 배열를 이용한 검색

배열의 순서까지 동일해야 값이 정상적으로 검색된다.

#### 검색
```javascript
db.inventory.find({
    tags: ['red', 'blank']
})
```

#### 결과
```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b956782f"),
    item: 'notebook',
    qty: 50,
    tags: [ 'red', 'blank' ],
    dim_cm: [ 14, 21 ]
  }
]
```

---

### And조건 검색 ($all)

만약 리스트 내부의 값이 순서에 상관없이 모두 가지고 있는걸 검색하고 싶다면 $all을 이용하면 된다.

#### 검색

```javascript
db.inventory.find({
    tags: {$all:  ['red', 'blank'] }
})
```

#### 결과
```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b956782e"),
    item: 'journal',
    qty: 25,
    tags: [ 'blank', 'red' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b956782f"),
    item: 'notebook',
    qty: 50,
    tags: [ 'red', 'blank' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567830"),
    item: 'paper',
    qty: 100,
    tags: [ 'red', 'blank', 'plain' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567831"),
    item: 'planner',
    qty: 75,
    tags: [ 'blank', 'red' ],
    dim_cm: [ 22.85, 30 ]
  }
]
```

$all을 사용할 경우 두개가 모두 들어있는경우 출력해준다.

---

### Or조건 검색 ($in)

위와 다르게 검색 조건에 한 개라도 부합하면 모두 출력하고 싶다면 $in을 사용하자.

#### 검색
```javascript
 db.inventory.find({
    tags: {$in:  ['red', 'blank'] }
 })
```

#### 결과
```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b956782e"),
    item: 'journal',
    qty: 25,
    tags: [ 'blank', 'red' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b956782f"),
    item: 'notebook',
    qty: 50,
    tags: [ 'red', 'blank' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567830"),
    item: 'paper',
    qty: 100,
    tags: [ 'red', 'blank', 'plain' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567831"),
    item: 'planner',
    qty: 75,
    tags: [ 'blank', 'red' ],
    dim_cm: [ 22.85, 30 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567833"),
    item: 'postcard',
    qty: 45,
    tags: [ 'blue', 'red' ],
    dim_cm: [ 13, 14 ]
  }
]
```
---

### 배열의 범위 검색

```javascript
 db.inventory.find({
    dim_cm: {$gt: 15, $lt: 20}
 })
```

위 조건을 보면 15보다 크고 20보다 작은 between에 대한 연산 같지만 둘 중 한개의 조건이라도 만족하면 값이 나오게 된다.

#### 결과
```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b956782e"),
    item: 'journal',
    qty: 25,
    tags: [ 'blank', 'red' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b956782f"),
    item: 'notebook',
    qty: 50,
    tags: [ 'red', 'blank' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567830"),
    item: 'paper',
    qty: 100,
    tags: [ 'red', 'blank', 'plain' ],
    dim_cm: [ 14, 21 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567832"),
    item: 'postcard',
    qty: 45,
    tags: [ 'blue' ],
    dim_cm: [ 10, 15.25 ]
  }
]
```

위 모든 값들이 15보다 큰 값이 들어있고 20보다 작은 값또한 들어있다.

#### 예외의 경우

```mongodb-json
{
    _id: ObjectId("6491d236aa9d3c49b9567831"),
    item: 'planner',
    qty: 75,
    tags: [ 'blank', 'red' ],
    dim_cm: [ 22.85, 30 ]
}
```

15보다는 큰값으로 모두 만족하지만 20보다 작은값에 모두 만족하지 않는다 그렇게 때문에 위 필터에서 걸러졌다.

---

### $elemMatch

15와 20사이의 값이 한 개라도 일치하면 값을 뽑아낸다.

$elemMatch 값에 모두 일치해야 필터링이 된다.

```javascript
 db.inventory.find({
    dim_cm: {$elemMatch: {$gt: 15, $lt: 20}}
 })
```

#### 결과
```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b9567832"),
    item: 'postcard',
    qty: 45,
    tags: [ 'blue' ],
    dim_cm: [ 10, 15.25 ]
  }
]
```
15.25 값이 15와 20사이의 값에 일치해서 위 값만 필터링 된것이다.


---

### 배열의 index번호를 이용한 접근

index번호를 이용해 배열에 접근할 수 있다.

```javascript
 db.inventory.find({
    "dim_cm.1": {$lt: 20}
 })
```

> dim_cm.1에서 볼 수 있듯이 1번째 인덱스 값이 조건에 일치하면 필터링 된다.
> 
> MongoDB도 0번째 인덱스 부터 있다.


#### 결과

15.25와 14는 20보다 작아서 출력되었다.

```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b9567832"),
    item: 'postcard',
    qty: 45,
    tags: [ 'blue' ],
    dim_cm: [ 10, 15.25 ]
  },
  {
    _id: ObjectId("6491d236aa9d3c49b9567833"),
    item: 'postcard',
    qty: 45,
    tags: [ 'blue', 'red' ],
    dim_cm: [ 13, 14 ]
  }
]
```

---

### 배열 크기 확인하기 ($size)

배열의 크기를 가지고 필터링이 가능하다.

```javascript
 db.inventory.find({
    tags: {$size: 3}
 })
```

#### 결과 

tags의 배열 길이가 3이기 때문에 출력된 값이다.

```mongodb-json
[
  {
    _id: ObjectId("6491d236aa9d3c49b9567830"),
    item: 'paper',
    qty: 100,
    tags: [ 'red', 'blank', 'plain' ],
    dim_cm: [ 14, 21 ]
  }
]
```
















