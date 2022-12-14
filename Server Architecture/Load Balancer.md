#  Load Balancer

## 로드밸런싱의 사용 이유

사용자가 늘어나게되면 성능이 뛰어난 서버라고해도 모든 트래픽을 감당할 수 없다. 
<br> 기업들은 서버를 추가하고 동일한 데이터를 저장해 트래픽을 효과적으로 분산하게 된다. 
<br> 이러한 과정을 스케일 업이라고 한다. [(Scale-Up Post)](https://github.com/YoonSeok-Heo/TIL/blob/main/Server%20Architecture/Scale-up%20Scale-out.md)

그런데 단순히 서버의 개수를 추가한다고 서버에 부하가 가지 않는건 아니다. 
<br> 트래픽을 각 서버별로 분산해주는 기술이 필요한데 이 기술을 로드밸런싱이라고 한다.

<br>

![image](https://user-images.githubusercontent.com/113662725/209440095-57484586-bb3c-464f-9dac-dd1349dbcd53.png)

<br>


로드 벨런서는 서버에 가해지는 트래픽을 분산해주는 장치 또는 기술을 통칭한다. 
<br> 클라이언트와 서버풀(Server Pool, 분산 네트워크를 구성하는 서버들의 그룹) 사이에 위치하며, 
<br> 한 대의 서버로 부하가 집중되지 않도록 트래픽을 관리해 서버가 최적의 퍼포먼스를 보일 수 있도록 한다.

## 로드밸런싱의 필요한 경우

로드밸런싱은 여러 대의 서버를 두고 서비스를 제공하는 분산 처리 시스템에서 필요한 기술

## 로드밸런싱 알고리즘

### 1. 라운드로빈(Round Robin)

서버에 들어온 요청을 순서대로 서버에 돌아가면서 배정하는 방식. 
<br> 클라이언트의 요청을 순서대로 분배하기 때문에 여러대의 서버가 동일한 스펙을 갖고 있고, 서버와의 연결이 오래 지속되지 않는경우 활용하기 적합.

### 2. 가중 라운드로빈(Round Robin)

각각의 서버마다 가중치를 매기고 가중치가 높은 서버에 클라이언트 요청을 우선적으로 배분.
<br> 주로 서버의 트래픽 처리 능력이 상이한 경우 사용되는 부하 분산 방식.
<br> 예를 들어 A라는 서버가 5라는 가중치를 갖고 B라는 서버가 2라는 가중치를 갖는다면, 로드밸런서는 라운드로빈 방식으로 A서버에 5개 B서버에 2개의 요청을 전달.

### 3. IP 해시(IP Hash)

클라이언트의 IP주소를 특정 서버로 매핑하여 요청을 처리하는 방식.
<br> 사용자의 IP를 해싱해 로드를 분배하기 때문에 사용자가 항상 동일한 서버로 연결되는 것을 보장한다.
<br> (유동 IP의 경우?? 어떻게 되는거지? 고정IP를 사용하는 기업에서만 사용하나??)

### 4. 최소 연결(Least Connection)

요청이 들어온 시점에 가장 적은 연결상태를 보이는 서버에 우선적으로 트래픽을 배분.
<br> 자주 세션이 길어지거나, 서버에 분배된 트래픽들이 일정하지 않은 경우 적합한 방식

### 5. 최소 리스폰타임(Lesat Response Time)

서버의 현재 연결 상태와 응답시간을 모두 고려하여 트래픽을 배분.
<br> 가장 적은 연결 상태와 가장 짧은 응답시간을 보이는 서버에 우선적으로 로드를 배분

[L4, L7 로드밸런서](https://github.com/YoonSeok-Heo/TIL/blob/main/Server%20Architecture/L4%20L7.md)

[참고자료](https://m.post.naver.com/viewer/postView.naver?volumeNo=27046347&memberNo=2521903)
