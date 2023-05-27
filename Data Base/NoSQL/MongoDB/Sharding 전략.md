# Sharding 전략

개발자가 Shard를 어떻게 구성을 하는지 전략 선택을 해 주어야 한다.

[Sharded Cluster 보고오기](https://github.com/YoonSeok-Heo/TIL/blob/main/Data%20Base/NoSQL/MongoDB/Sharded%20Cluster.md)

---

# 1. Ranged Sharding

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/6536e1b3-9c34-452e-bb0e-38f0b6db836e)

일반적인 MongoDB의 index를 생성하고, 해당 필드들을 Shard Key로 지정한다. 그리고 Shard Key의 값을 기준으로 청크에 범위를 할당한다.

연속된 숫자들은 같은 Chunk와 같은 Shard에 존재하기 때문에 데이터의 값으로 정확히 어떤 Shard에 존재하는지 알기 때문에 해당 Shard에만 쿼리를 날리는 타겟쿼리가 가능하다.
(효율적이고 빠르다)

> 그림에서 보면 X의 값에 따라 어떤 Shard에 데이터가 존재하는지 명확하기 때문에 타겟 쿼리가 가능한것이다.

### 치명적 단점....

Ranged Sharding은 데이터가 균형있게 분산되지 않을 가능성이 크다.

예를 들어 X의 값이 순차적으로 늘어난다고 하면 항상 마지막 Shard인 C에만 데이터가 쌓이게된다. 
또한 데이터 불균형을 해결하기 위해 다른 Shard로 데이터 **Migration을 내부적으로 하게 되는데 이 작업이 부하가 크기 때문에 치명적 단점으로 사용자 쿼리에도 영향을 주어서 문제가 될 수도 있다.**

> 다른 Sharding 전략을 사용해야 한다. 주로 Hashed Sharding을 사용한다고 한다.

### 그럼 왜 사용하는가?

- Hashed Sharding이 불가능 한 경우가 있는데 그럴경우 Ranged Sharding을 사용한다.
- 버전에 따라 Zone Sharding(?)을 함께 이용하는 경우
- Shard Key값 자체가 잘 분산되어 있을때
- 어떤 Shard에 해당 데이터가 있는지 정확히 알아서 조회를 빠르게 해야하는 경우(타겟 쿼리가 필수로 사용되야 할 경우)

---

# 2. Hashed Sharding

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/a5b05f20-5767-4ea7-8de8-e31366e4c42a)

Shard key의 값에 Hash값을 이용해 청크에 넣어주게 된다. 그렇기 때문에 데이터의 분산이 균등하게 된다.

특정 Shard에 데이터가 몰리는 현상을 막을 수 있다.

> 에시에서도 연속된 값이 들어가도 분산이 잘 되는것을 볼 수 있다.

### 단점..

원본값이 같으면 Hash값도 같게되어 같은 Shard에 들어가게 되므로 데이터가 몰리는 현상을 완전히 막을수는 없다고 한다.

Range Sharding처럼 데이터가 연속되어 있지 않고 **분산되어 있어서 모든 Shard에 모두 요청(브로드 캐스트 쿼리)해서 필터링을 해야한다.**

> 예를 들어 Shard가 10개인데 데이터 10개를 범위 조건으로 검색을 한다면 범위 내의 데이터들이 모든 Shard에 분산되어 있을 확률이 높다. 그렇게 되면 Shard의 갯수가 많을수록 타겟쿼리와의 성능차이는 점점 커질 것이다.

### 단점에도 사용하는 이유!

Sharded Cluster의 목적이 균등하게 데이터를 저장하는 것이기 때문에 주로 Hashed Sharding을 사용하게 된다.

---

# 3. Zone Sharding

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/e662a8da-dcec-4066-abdd-7b3837f90312)

Zones Sharding은 Range, Hashed Sharding이랑 함께 사용한다.

값에 대해 Zone을 생성하고 특정 Shard로 데이터를 위치시키는 것이다.

Shard Alpha는 Zone A만 Beta는 Zone A, B만 그 외에는 Charlie 에 저장되게끔 나누는 방식이다.

### 사용하는곳

#### 글로벌하게 데이터를 분산시켜야 할 경우

- IP를 기준으로 Zone Shard를 결정해 네트워크 레이턴시를 낮추는 방법이 있다.

---

# Summary

- Range와 Hashed Sharding 두 가지 방법이 있다.
- 가능하면 Hshed Sharding을 통해 데이터를 분산한다.

