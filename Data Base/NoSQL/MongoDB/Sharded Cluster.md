# Sharded Cluster


> 참고. 모든 Shard는 Replica Set으로 구성되어 있다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/d18fb34d-f8f2-4e2a-94ce-a7d46f5aa395)

좌측에 이미지는 A-Z까지의 Collection을 가지고있는 Replica Set이 있다.

그런데 Replica Set이 감당할 수 없을 정도로 많은 데이터가 유입이 되면서 오른쪽과 같은 Sharded Cluster 형태로 배포하기로 한다.

우측의 sharded cluster는 알파벳 컬렉션이 알파벳이라는 필드를 Shard key로 Sharding되어서 총 3개의 Shard로 분산된 것을 확인할 수 있다.

> 각각의 Shard는 Replica Set으로 관리가 된다.
> 그러므로 H/A 또한 보장이 된다.
> 
> 복제는 H/A를 위한 솔루션이고 Sharding은 분산처리를 위한 솔루션이라고 보면된다.

## 간단한 용어정리.

Sharding: 하나의 큰 데이터를 여러 장비에 걸쳐서 분할하는 과정
Shard: 분할된 데이터 셋의 모음
Shard key: 분산이 되는 기준. Collection필드에 생긴 인덱스를 이용해 활성화한다.

---

## 장단점

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/d62b5ace-c3c8-4223-9bc9-537c182e9e21)

### 장점

1. 용량의 한계를 극복할 수 있다: Scale out의 개념이므로 용량을 언제든지 늘릴 수 있다.
2. 데이터 규모와 부하가 크더라도 처리량이 좋다: 서버가 처리할 수 있는 용량 만큼만 라우팅을 해주기 때문에 대용량의 경우 성능이 더 좋아진다.
3. 고가용성 보장(H/A): Shard가 Replica Set으로 이루어져 있으므로 고가용성이 보장된다.
4. 하드웨어에 대한 제약을 해결할 수 있다: Scale Up으로 처리하지 못하는 것을 Scale Out으로 처리하기 때문에 하드웨어 제약을 해결할 수 있다.

### 단점

1. 관리가 비교적 복잡하다: 구성 자체가 복잡해 지기 때문에 관리 또한 복잡하진다. 운영 입장에서는 가능하면 Replica Set으로만 가려고한다.(관리가 편하기 때문에)
2. Replica Set과 비교해서 쿼리가 느리다: 각각 shard가 다른 데이터를 가지고 있기 때문에 요청이 들어올 경우 어느 Shard에 데이터가 있는지 확인하는 Route시스템이 필요하기 때문이다. 또한 범위가 큰 데이터가 들어올 경우 여러 Shard에서 데이터를 찾고 머지해주는 과정까지 시간이 소요된다.


---

## 비슷하면서 햇갈리는 개념들 정리

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/bbe9294c-8dc9-4cf7-977d-518afda37e2d)

---

## Architecture

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/41acc0ce-5140-405f-8b7d-b62a612897dd)

Sharded Cluster는 router, Shard 그리고 Config Server로 구성되어 있다. 

1. Shard: Replica Set으로 구성되어 있으면 데이터를 직접 들고있다.
2. Router(mongos): 쿼리를 Shard로 전달하고 결과를 머지해서 사용자에게 반환한다.
3. Config Server: Sharded Cluster의 메타데이터를 가지고 있는 컴포넌트이다. 어떤 Shard가 어떤 데이터를 가지고있나, Cluster의 상태 등 중요한 정보를 가지고있다.(Config Server도 H/A를 보장하는 Replica Set으로 구성되어 있따.)

### 흐름

1. 사용자가 Driver를 통해 router에 요청을 한다.
2. router는 Config Server에 메타데이터 정보를 확인해서 찾으려고 하는 데이터가 어떤 Shard에 있는지 확인한다.
3. router가 Shard에 직접 접근해서 필요한 데이터를 가져와 머지하고 사용자에게 반환한다.

---

## Sharding Collectin

Sharding은 Collection단위로 하게되는데 다음 이미지는 Collection이 Sharding된 모습이다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/133293e5-862c-41d4-8cc0-ca1c99dda418)

Collection 1: Sharding이 된 상태로 데이터가 나누어진 상태이다.

Collection 2: Sharding이 되지 않고 하나의 Shard에만 들어있다.

### 왜?? Collection 2는 Shading을 하지 않았을까?

상황에 따라서 일부 Colection은 Sharding을 하지 않는다. **분산을 시키지 않는게 유리하다!** 라고 판단될 경우이다.

Sharded Cluster가 Replica Set보다 느리다고 했는데, 
그 이유가 데이터가 분산되어 있기 때문에 데이터를 **각 Shard에서 조회하고
merge하는데 시간이 소요**되기 때문이다.

> 예로 메타 데이터가 있다.데이터의 크기는 작고 조회수는 잦은 데이터의 경우 한개의 Shard에서 관리하고 한 개의 Shard만 조회한다면 각 Shard에 데이터를 조회하고 merge하는 시간이 소요되지 않기 때문이다.

--- 

## Chunks

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/133293e5-862c-41d4-8cc0-ca1c99dda418)

MongoDB는 sharding된 데이터를 분할하고 그것을 청크라고 한다. 각 분할된 조각은 여러 개의 Shard에 분산되어 저장되는데 이 분산 조각을 Chunks라고 한다.

Chunks는 각 Shard서버에 균등하게 저장되어야 좋은 성능을 낼 수 있는데, 균등하게 저장하기 위에서는 Chunks를 Split하고 Migration 하는 과정을 거치게 된다.

> 위 이미지에서 사각형 한 개를 청크라고 할 수 있다. 









