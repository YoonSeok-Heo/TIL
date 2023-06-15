# Intro

기존의 웹으로 관리하던 배치 솔루션을 사내 배치 관리 솔루션으로 옮겨야한다. 그런데 배치 관리 솔루션은 쉘로 만들어져있다. 쉘을 이용해야한다.

MSA형태로 만들어진 솔루션에 각 모듈마다 REST API를 재공한다. 그렇다는건? 그냥 API호출로만 해결 할 수 있다고 생각한다. 그래서 API를 호출하기 위해 curl을 알아보았다.

## cURL
서버와 통신할 수 있는 명령어 툴이다.

다양한 프로토콜을 지원한다.(DICT, FILE, FTP, FTPS, Gopher, HTTP, HTTPS, IMAP, IMAPS, LDAP, LDAPS, POP3, POP3S, RTMP, RTSP, SCP, SFTP, SMB, SMBS, SMTP, SMTPS, Telnet, TFTP)
그냥 url을 이용한다면 모든 다 할수 있다고 보면된다. SSL인증도 지원한다.

나는 REST API를 사용하기 위해 cURL을 정리할 것이다.

### option

|short|long|설명|비고|
|---|----|---|---|
|-k|--insecure|https 사이트를 SSL certificate 검증없이 연결한다.|wget 의 --no-check-certificate 과 비슷한 역할 수행|
|-l|--head|HTTP header 만 보여주고 content 는 표시하지 않는다||
|-D|--dump-header [file]|[file] 에 HTTP header 를 기록한다.|
|-L|--location|서버에서 HTTP 301이나 HTTP 302 응답이 왔을 경우 redirection URL 로 따라간다.--max-redirs 뒤에 숫자로 redirection 을 몇 번 따라갈지 지정할 수 있다. 기본 값은 50이다.|curl -v daum.net 을 실행하면 결과값으로 다음과 같이 HTTP 302 가 리턴된다.< HTTP/1.1 302 Object Moved< Location: http://www.daum.net/-L 옵션을 추가하면 www.daum.net 으로 재접속하여 결과를 받아오게 된다.|
|-d|--data|HTTP Post data|FORM 을 POST 하는 HTTP나 JSON 으로 데이타를 주고받는 REST 기반의 웹서비스 디버깅시 유용한 옵션이다|
|-v|--verbose|동작하면서 자세한 옵션을 출력한다.
|-J|--remote-header-name|어떤 웹서비스는 파일 다운로드시 Content-Disposition Header 를 파싱해야 정확한 파일이름을 알 수 있을 경우가 있다. -J 옵션을 주면 헤더에 있는 파일 이름으로 저장한다.|curl 7.20 이상부터 추가된 옵션|
|-o|--output FILE|	curl 은 remote 에서 받아온 데이타를 기본적으로는 콘솔에 출력한다. -o 옵션 뒤에 FILE 을 적어주면 해당 FILE 로 저장한다. (download 시 유용)||
|-O|--remote-name|	file 저장시 remote 의 file 이름으로 저장한다. -o 옵션보다 편리하다.||
|-s|--silent|	정숙 모드. 진행 내역이나 메시지등을 출력하지 않는다. -o 옵션으로 remote data 도 /dev/null 로 보내면 결과물도 출력되지 않는다|HTTP response code 만 가져오거나 할 경우 유리|
|-X|--request|Request 시 사용할 method 종류(GET, POST, PUT, PATCH, DELETE) 를 기술한다.||
|-i|--include|응답에 Content 만 출력하지 않고 서버의 Reponse 도 포함해서 출력한다. (디버깅에 유용)||

<br><br>
**내가 자주 사용하게 될 것은 -s -S -X 이정도 일것같다.**

**RestAPI에서 JSON형식으로 응답을 주면 jq를 사용해서 데이터를 다루려고한다.**