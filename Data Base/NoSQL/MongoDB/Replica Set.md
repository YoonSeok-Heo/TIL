# Replica Set

Replica Set은 H/A에 대한 솔루션이고 멈춤없이 지속적인 운영이 필요한 서비스에서 선택되는 최소 배포형태이다.

H/A를 보장해주기 때문에 내부적으로 약간 복잡해지기도 한다. 이로인해 개발자가 알아햐 하는 것들이 생긴다.

개발자 입장에서 Replication을 이해해야 Replica Set을 이용한 MongoDB를 개발할 수 있을 것이다.

## Replica Set Members

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/436fd79b-9f26-478c-b0e0-0704283fadff)

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/a83626cd-a9ba-4df4-8b04-908781b698ba)

위 그림은 Primary 1개 Secondary 2개 총 3개의 멤버로 이루어져 있는 Replica Set이다. 
각각의 멤버는 **같은 서버에 존재할 수도 있고, 다른 서버에 존재할 수 있다.** 
Replica Set의 주 목적은 H/A기 때문에 **다른 서버에 설치해 사용하는 것이 용도에 맞다고** 볼 수 있다.

> 각 멤버는 상태 값이 존재하는데 여러 상태 값 중에서 정상적으로 데이터를 처리할 수 있는 상태는 **Primary와 Secondary 두개의 상태**이다.

### 1. Primary

Primary는 Replica Set에서 **한 개만 존재**할 수 있다.
Primary는 읽기 요청을 처리할 수 있지만 Repica Set에서 유일하게 **쓰기(삽입/수정/삭제) 요청이 가능**하다.

### 2. Secondary

Secondary는 **읽기요청만 받을 수 있으며**, 
가장 최신상태의 데이터를 가지고 있는 Primary와 동일한 데이터를 가지고 있어야 하기 때문에 **Replication(복제)을 통해 Primary와 동일한 데이터를 유지**한다.

> Replica Set은 쓰기는 Primary로 요청을 하고 읽기에 대한 작업은 Secondary로 요청을 하면서 읽기에 대한 **분산처리**를 하고 있다. 
> 
> 이 작업은 쿼리할때 **Read Preference 옵션을 사용해서 쿼리를 요청해야 가능**하다.
> 
> 아무런 설정없이 Replica Set으로 전달하면 기본적으로 **Primary가 모두 처리하기때문에 읽기에 대한 분산처리가 되지 않는다.**

---

## Replica Set Election(Fail-Over)

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/8da2bab9-1562-4f25-b253-32bceac8e94b)!

Replica Set은 자동 Fail-Over가 가능하기 때문에 H/A가 가능하다고 볼 수 있다.

단순히 '여러개의 서버 중 한개만 내려갔으니 된다.'라는 느낌보다는 좀 더 복잡하다. 

### Primary가 죽을 경우.

Secondary가 죽을 경우 단순히 죽었다고 생각할 수 있다. 하지만 그림처럼 Primary가 죽게 된다면...

쓰기의 요청이 불가능 할 것이다. 읽기만 가능한 상태가 된다는 것이다. 서비스 요구에따라 장애로 판단할 수 있다.


Replica Set은 **Primary가 없다고 판단되면 자동으로 Primary를 선출해서 자동 Fail-Over를 하게된다.**

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/4a55dd49-c108-4c8b-a0db-c39eb2a5acb8)

멤버들은 서로에게 Heartbeat를 날리면서 **살아있는지 수시로 체크**를한다. 
그러다가 Primary에 응답이 없게되면 Primary가 죽은것으로 판단하고 **자동으로 선출 알고리즘 통해 새로운 Primary를 선출**하게 된다.
선출에 대한 기준은 설정값이나, 지연이 없는지, 멤버의 상태 등 여러가지를 확인해서 과반수의 투표를 얻으면 **Secondary에서 Primary로 승격**을 하게 된다.

투표라는 시스템을 가지기 때문에 Replica Set 멤버의 수를 결정할 때 최소 3개 이상 혹은 홀수로 맞추는 등의 개념들이 있다. 

### Arbiter

투표를 할 수 있는 상태가 하나 더있다. Primary Secondary Arbiter이다. 

Arbiter는 데이터를 쓰기 읽기 기능은 하지 않으며 오로지 투표에만 사용된다. 비용 문제로 인하여 서버를 늘리기 힘든경우 리소스가 작은 서버를 이용해 Arbiter가 투표를 대신해서 진행한다고 보면 된다.

> 권장되는 방법은 아니다.

--- 

## OpLog

OpLog를 이용해서 Replica Set의 모든 멤버가 동일한 데이터를 유지하도록 하게 된다.

### 동작 방식

1. 쓰기 요청이 Primary로 들어온다.
2. Primary에 데이터를 저장
3. localdatabase에 있는 Oplog collection에 변경사항을 기록한다.
4. 각 Secondary는 비동기적으로 선반영 되어 있는 멤버의 OpLog를 복사해서 변경을 수행한다.

> 꼭 Primary에서만 Oplog를 복사해 오는 것이 아니라 Secondary가 자신보다 앞서있는 Secondary의 OpLog를 복사할 수 있다.

---

## Summary

- Replica Set은 HA솔루션이다.
- 데이터를 들고 잇는 멤버의 상태는 Primary와 Secondary가 있다.
- Seconday는 선출을 통해 과봔수의 튜표를 얻어서 Primary가 될 수 있다.
- Arbiter는 데이터를 들고 있지 않고 Primary 선출 투표에만 참여하는 멤버이다.
- Replica Set은 loca database의 Oplog Collection을 통해 복제를 수행한다.