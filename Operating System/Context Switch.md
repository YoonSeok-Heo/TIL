# Context Switch

## Context Switching이란? 

멀티프로세스 환경의 CPU가 한 개의 Task(Process/Thread)를 실행하고 있는 상태에서 Interrupt요청에 의해 
다른 Task로 실행이 전환되어야 할 때 기존의 Task의 상태 및 레지스터 값들에 대한 정보(Context)를 저장하고 
CPU가 다음 프로세스를 수행하도록 새로운 Task의 상태 및 레지스터 값(Context)을 교체하는 작업을 말한다.

> **※ 위에서 Context란, CPU가 다루는 Task(Process/Thread)에 대한 정보로 대부분의 정보는 Register에 저장되며 PCB(Process Control Block)으로 관리된다.**
>> 여기서 Process와 Thread를 처리하는 Context Switching은 조금 다른데, PCB는 OS에 의해 스케줄링 되는 Process Control Block이고,
>> Thread의 경우 Process내의 TCB(Task Control Block)라는 내부 구조를 통해 관리된다.

### PCB의 저장 정보

- 프로세스 상태: 생성, 준비, 수행, 대기, 중지
- 프로그램 카운터: 프로세스가 다음에 실행할 명령어 주소
- 레지스터: 누산기, 스택, 색인 레지스터
- 프로세스 번호

> ※ Context Switching때 해당 CPU는 작업을 하지 못한다. 따라서 컨텍스트 스위칭이 잦아지면 오버헤드가 발생해 성능이 떨어진다.
> 
> ※ 일반적으로 멀티 프로세스를 통해 PCB를 Context Switching하는 것보다 멀티 쓰레드를 통해 TCB를 Context Switching하는 비용이 더 적다고 알려져 있다. 


## Interrupt

인터럽트는 CPU가 프로그램을 실행하고 있을 때 실행중인 프로그램 밖에서 예외 상황이 발생하여 
처리가 필요한 경우에 CPU에게 알려 예외 상황을 처리할 수 있도록 하는 것을 말한다.

### 인터럽트의 종류

1. I/O request(입출력 요청)
2. time slice expired(CPU 사용시간이 만료 되었을 때)
3. fork a child(자식 프로세스를 만들 때)
4. wait for an interrupt(인터럽트 처리를 기다릴 때)
5. 등...

