# Kafjka Local 환경설정

## 실습 환경

- Windows11
- Open JDK 17(Corretto-17.0.7.7.1)
- Kafka 2.8.2

JAVA는 환경설정까지 모두 되어있다는 가정하에 Kafka를 설치해 보겠다.

### [Kafka Download Page](https://kafka.apache.org/downloads)

위 링크에서 자신이 원하는 Kafka 버전을 받으면 된다. 

나는 강의를 보는중이라 강의와 맞게 2.8.2 버전을 다운로드 했다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/5f6bff5f-2e0e-44bb-98de-ba25f1b6227b)

윈도우에서 tar파일로 다운로드 되는데 알집과 같은 애플리케이션을 사용하면 풀 수 있다.

아니면 tar명령어를 이용해 압축을 푸는것도 가능하다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/8ed730e1-d468-469a-8a9f-4b6cfbe14d0b)


압축을 풀게되면 bin폴더가 있는데 bin폴더 안에 windows폴더에 있는 배치피일을 이용해 서버를 실행 할 수 있다.


---

### 실행해보자

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/4648ff60-1e08-4d56-8f88-5674df909f33)

먼저 주키퍼를 실행하자 zookeeper-server-start.bat파일에 주키퍼 설정 파라미터를 넣어 실행하자

> 주키퍼는 기본적으로 2181포트를 사용하며
> 
> 주키퍼는 카프카의 메타정보를 관리하게 된다

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/8c399fb9-f16b-45a4-a523-2394ac9abc58)

이어서 카프카 서버를 올려보자


위와 같은 방법으로 서버를 실행할 수 있다.

근데 나는 뭔가 잘못된거같다.. 우선 강의와 다르게 jdk17을 사용햇는데 주키퍼 올리는 과정중에 문제가 있는것 같다.

다시한번 확인해봐야겠다.


---

## 실행

이것저것 해보다 된거같은데 뭘하다가 된건지 정확히 모르겠다... 젠장,,

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/29cdd24b-ce7e-4542-afe4-395f6ea89064)

- zookeeper

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/ee36b865-5244-467e-be0e-d66f4a3ee70a)

- kafka