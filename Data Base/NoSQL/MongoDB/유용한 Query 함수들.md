# 유용한 Query 함수

## db.collection.bulkWrite

**bulkWrite는 ```mongosh```에서 사용하는 메서드이다. Node.js와 같은 언어에서 사용하는 메서드가 아니다.**

[문서 보기](https://www.mongodb.com/docs/manual/reference/method/db.collection.bulkWrite/#mongodb-method-db.collection.bulkWrite)

bulkWrite는 여러 변경사항을 한번에 진행하는 배치성 작업에서 유용하게 사용하는 메서드이다.

여러 종류의 write관련(insert, update, delete) 작업을 하나의 함수에서 실행하면서 실행하면서 Network Round Trip을 최소화하고 애플리케이션의 성능을 높일수 있다.

```javascript
db.collection.bulkWrite(
   [ <operation 1>, <operation 2>, ... ],
   {
      writeConcern : <document>,
      ordered : <boolean>
   }
)
```

첫번째 인자 ```[ <operation 1>, <operation 2>, ... ]```에 작업들을 리스트 형식으로 넣어준다.

다음에 옵션을 넣는 옵션을 넣는다.

> 옵션들은 docs를 확인해보자!

### 예제

```javascript
db.bulk.bulkWrite(
    [
        {insertOne: {doc: 1, order: 1}},
        {insertOne: {doc: 2, order: 2}},
        {insertOne: {doc: 3, order: 3}},
        {insertOne: {doc: 4, order: 4}},
        {insertOne: {doc: 4, order: 5}},
        {insertOne: {doc: 5, order: 6}},
        {
            deleteOne: {
                filter: {doc: 3}
            }
        },
        {
            updateOne: {
                filter: {doc: 2},
                update: {
                    $set: {doc: 12}
                }
            }
        }
    ]
)
```

6개의 데이터를 모두 넣은 뒤 doc이 3인 데이터를 지우고 doc이 2인 데이터에 doc값을 12로 수정하는 과정이다.

### 결과

```mongodb-json
[
  { _id: ObjectId("64906f4ae40fe2a3c5262160"), doc: 1, order: 1 },
  { _id: ObjectId("64906f4ae40fe2a3c5262161"), doc: 12, order: 2 },
  { _id: ObjectId("64906f4ae40fe2a3c5262163"), doc: 4, order: 4 },
  { _id: ObjectId("64906f4ae40fe2a3c5262164"), doc: 4, order: 5 },
  { _id: ObjectId("64906f4ae40fe2a3c5262165"), doc: 5, order: 6 }
]
```


### ordered 옵션

ordered의 default값은 true인데 true일경우 순서대로 bulkWrite()가 진행된다.

false일 경우 list내부의 명령어들이 순서에 상관없이 속도가 최적화 된 상태로 실행된다. 즉 위에있다고 먼저 실행되는것이 아니다.

---

## db.collection.countDocuments()

**countDocuments는 ```mongosh```에서 사용하는 메서드이다. Node.js와 같은 언어에서 사용하는 메서드가 아니다.**

```countDocuments```는 말그대로 Document의 갯수를 가져온다. ```countDocuments```는 Document를 모두 불러서 개수를 확인하는 과정을 거친다. 
그렇기 때문에 Document가 많아지면 값을 불러오는데 시간이 오래 걸리게 된다.

> 예전 버전에서는 count를 사용했는데 deplecated 됐다.

---

## db.collection.estimatedDocumentCount()

```countDocuments```와 같은 역할을 하는데 실제로 실행하지않고 예상되는 Document개수를 반환한다. 그러므로 빠르다.

---

## db.collection.distinct()

고유한 값들을 리스트로 뽑아낸다.

```javascript
db.bulk.distinct("doc")
```

```mongodb-json
[ 1, 4, 5, 12 ]
```