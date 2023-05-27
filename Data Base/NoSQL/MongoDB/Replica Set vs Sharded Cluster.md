# Replica Set vs Sharded Cluster

개발자는 배포전에 크게 두 가지를 고려하면 된다.

## 1. 서비스 요구 사항에 맞게 배포

- 서비스가 하루에 쌓이는 데이터가 얼마나 되는지
- 쓰기가 많은 서비스인지 읽기가 많은 서비스인지

## 2. 배포 환경에 따른 선택

- 내가 배포하려고 하는 서버의 환경

> ### 예 1) 용량에 대한 고려 -> 환경에 따라 다르다.
> 
> 매일 1GB씩 데이터가 증가하고 3년간 보관한다.
> 
> ![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/256a6cf0-97aa-42ae-b90d-df14f8976da5)
> 
> 위 상황으로 보았을 때 안전 수치 35% 정도 남겨둔다고 할 경우 필요한 서버의 스팩은 5TB정도의 서버가 필요하다.
> 
> 하지만 준비할 수 있는 서버가 **2TB의 저장공간을 가진 서버뿐**이라면, 서버의 수를 늘리더라도 **Sharded Cluster를 이용해 분산처리**를 해주어야한다.
> 
> 반면에 10TB의 서버를 구할 수 있으므로 Replica Set을 이용해 배포하면 된다.


> ### 예 2) Write 요청이 압도적으로 많은 서비스 -> Sharded Cluster
>
> ![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/256a6cf0-97aa-42ae-b90d-df14f8976da5)
> 
> 예를들어 증권 어플이다. 수시로 주가가 변경되고 거래가 성립되니 읽기보다 쓰기요청이 압도적으로 많다.
> 
> - **Replica Set은 쓰기에 대한 분산이 불가**능 하다. Replica Set의 **Primary멤버가 부하를 감당하지 못하는 경우가 발생**할 수 있다.
> - **Sharded Cluster는 읽기의 분산이 가능**하므로 Sharded Cluster를 사용해야한다.


> ### 예 3) 논리적인 데이터베이스가 많은 경우 -> 여러 Replica Set으로 분리
> 
> 데이터의 영구적인 보관 (3년 보관 X)
> 
> 대용량의 데이터를 논리적으로 분리해서 여러 데이터베이스를 하나의 MongoDB Cluster에서 사용하게끔 하는 요구사항
> 
>> #### 문제가 발생 할 수 있는 점.
>>
>> - MongoDB의 Cluster를 한 개만 사용하다보면 특정 부분에서만의 장애가 서비스 전체가 마비되는 치명적인 단점이 있을 수 있다.
> 
> ![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/47b250a7-55d7-4b39-a4d7-9dc981050fc6)
> 
> 논리적 데이터베이스에서 한 개의 Replica Set에 배포하기 보다는 여러 Replica Set

---

## Replica Set vs Sharded Cluster

X축은 실행시간 Y축은 쿼리 방식이다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/151b8097-36a3-4f3a-8f36-5d73cc48be55)

### 성능

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/4fc0f6d5-3894-4776-9e32-b13b107aec77)

쿼리를 실행 했을 때 실행 시간만 확인한 경우

---

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/3428a416-edab-4b65-8a26-4040e545efeb)

클라이언트 트라이버를 이용한 쿼리. 네트워크 레이턴시와 같은 변수가 포함되어 있다.

---

### AGGREGATION (집계 프래임 워크)

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/1aa8472c-c5b3-46d9-b5a7-eba7846be489)

일반적이 쿼리 이외에 집계를 위한 쿼리문들

---
### 단일 건 INSERT

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/1a939e55-edad-4cf3-b61b-78c964bd7f6e)

큰 차이가 없음을 알 수 있다.

---

### INSERT Many

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/bc77d138-28c1-4265-b105-b765da409be4)

여러 개의 데이터를 넣을 땐 차이가 많이 나는 것을 확인할 수 있다. DML을 수행 할 때마다 Chunk split이나 migration에 대한 여부를 확인하고 내부에서 split, migration이 진행중이라면 더 오래 걸릴 수 있다.

---

큰 차이가 나지 않으나. 대체적으로 Replica set이 빠르며 Hash shard에서 Range Query가 느린이유는 모든 shard에 쿼리요청을 하기 때문에 느려진다.

Shard의 갯수가 많아지면 더 큰 차이가 발생할 수 있다고한다.


# Summary
- 가능하면 Replica Set으로 배포를 한다.
- 서비스의 요구사항이 Replica Set으로 충족하지 못할 경우 Sharded Cluster를 사용한다.