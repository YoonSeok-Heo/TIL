# Rollback이란?
#DB #Rollback

## Rollback
- 작업 중 문제가 발생되어 트랜잭션 처리중 발생한 변경사항을 취소하는 명령
- Rollback역시 하나의 트랜잭션 과정을 종료
- 트랜잭션으로 인한 하나의 묶음 처리가 시작되기 이전 상태로 되돌림
- 이전 Commit지점까지 복구


### 한줄요약 
트랜잭션으로 묶여있는 단위의 변경내용이 비정상적으로 종료되면 트랜잭션이 시작하기 전 상태로 되돌린다.!
