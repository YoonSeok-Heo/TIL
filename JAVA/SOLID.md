# SOLID

객체 지향 프로그래밍 및 설계에서 다섯 가지 기본 원칙을 마이클 페더스가 소개한 것이다.
프로그래머가 시간이 지나도 유지보수와 확장이 쉬운 시스템을 만들고자 할 때 이 원칙들을 적용할 수 있다.

SOLID원칙들은 소프트웨어 작업에서 프로그래머가 코드를 읽기 쉽고 확장하기 쉽게 될 때까지 소프트웨어 소스코드를 리팩터링하여 코드 냄새를 제거하기 위해 적용할 수 있느 지침이다.

<table class="wikitable"
       style="width: auto; font-size: 95%; table-layout: fixed; line-height:1.25; margin-left: auto; margin-right: auto;">

  <tbody>
  <tr>
    <th>두문자</th>
    <th>약어</th>
    <th>개념
    </th>
  </tr>
  <tr>
    <th>S
    </th>
    <td><a href="https://ko.wikipedia.org/wiki/%EB%8B%A8%EC%9D%BC_%EC%B1%85%EC%9E%84_%EC%9B%90%EC%B9%99"
           title="단일 책임 원칙">SRP</a>
    </td>
    <td>
      <dl>
        <dt><a href="https://ko.wikipedia.org/wiki/%EB%8B%A8%EC%9D%BC_%EC%B1%85%EC%9E%84_%EC%9B%90%EC%B9%99"
               title="단일 책임 원칙">단일 책임 원칙 (Single responsibility principle)</a></dt>
        <dd>한 <a
          href="https://ko.wikipedia.org/wiki/%ED%81%B4%EB%9E%98%EC%8A%A4_(%EC%BB%B4%ED%93%A8%ED%84%B0_%EA%B3%BC%ED%95%99)"
          class="mw-redirect" title="클래스 (컴퓨터 과학)">클래스</a>는 하나의 책임만 가져야 한다.
        </dd>
      </dl>
    </td>
  </tr>
  <tr>
    <th>O
    </th>
    <td><a href="https://ko.wikipedia.org/wiki/%EA%B0%9C%EB%B0%A9-%ED%8F%90%EC%87%84_%EC%9B%90%EC%B9%99"
           title="개방-폐쇄 원칙">OCP</a>
    </td>
    <td>
      <dl>
        <dt><a href="https://ko.wikipedia.org/wiki/%EA%B0%9C%EB%B0%A9-%ED%8F%90%EC%87%84_%EC%9B%90%EC%B9%99"
               title="개방-폐쇄 원칙">개방-폐쇄 원칙 (Open/closed principle)</a></dt>
        <dd>“소프트웨어 요소는 확장에는 열려 있으나 변경에는 닫혀 있어야 한다.”</dd>
      </dl>
    </td>
  </tr>
  <tr>
    <th>L
    </th>
    <td><a
      href="https://ko.wikipedia.org/wiki/%EB%A6%AC%EC%8A%A4%EC%BD%94%ED%94%84_%EC%B9%98%ED%99%98_%EC%9B%90%EC%B9%99"
      title="리스코프 치환 원칙">LSP</a>
    </td>
    <td>
      <dl>
        <dt><a
          href="https://ko.wikipedia.org/wiki/%EB%A6%AC%EC%8A%A4%EC%BD%94%ED%94%84_%EC%B9%98%ED%99%98_%EC%9B%90%EC%B9%99"
          title="리스코프 치환 원칙">리스코프 치환 원칙 (Liskov substitution principle)</a></dt>
        <dd>“프로그램의 <a href="https://ko.wikipedia.org/wiki/%EA%B0%9D%EC%B2%B4" class="mw-redirect" title="객체">객체</a>는
          프로그램의 정확성을 깨뜨리지 않으면서 하위 타입의 인스턴스로 바꿀 수 있어야 한다.” <a
            href="/w/index.php?title=%EA%B3%84%EC%95%BD%EC%97%90_%EC%9D%98%ED%95%9C_%EC%84%A4%EA%B3%84&amp;action=edit&amp;redlink=1"
            class="new" title="계약에 의한 설계 (없는 문서)">계약에 의한 설계</a>를 참고하라.
        </dd>
      </dl>
    </td>
  </tr>
  <tr>
    <th>I
    </th>
    <td><a
      href="https://ko.wikipedia.org/wiki/%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4_%EB%B6%84%EB%A6%AC_%EC%9B%90%EC%B9%99"
      title="인터페이스 분리 원칙">ISP</a>
    </td>
    <td>
      <dl>
        <dt><a
          href="https://ko.wikipedia.org/wiki/%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4_%EB%B6%84%EB%A6%AC_%EC%9B%90%EC%B9%99"
          title="인터페이스 분리 원칙">인터페이스 분리 원칙 (Interface segregation principle)</a></dt>
        <dd>“특정 클라이언트를 위한 인터페이스 여러 개가 범용 인터페이스 하나보다 낫다.”<sup id="cite_ref-martin-design-principles_4-0"
                                                             class="reference"><a
          href="#cite_note-martin-design-principles-4">&#91;4&#93;</a></sup></dd>
      </dl>
    </td>
  </tr>
  <tr>
    <th>D
    </th>
    <td><a
      href="https://ko.wikipedia.org/wiki/%EC%9D%98%EC%A1%B4%EA%B4%80%EA%B3%84_%EC%97%AD%EC%A0%84_%EC%9B%90%EC%B9%99"
      title="의존관계 역전 원칙">DIP</a>
    </td>
    <td>
      <dl>
        <dt><a
          href="https://ko.wikipedia.org/wiki/%EC%9D%98%EC%A1%B4%EA%B4%80%EA%B3%84_%EC%97%AD%EC%A0%84_%EC%9B%90%EC%B9%99"
          title="의존관계 역전 원칙">의존관계 역전 원칙 (Dependency inversion principle)</a></dt>
        <dd>프로그래머는 “추상화에 의존해야지, 구체화에 의존하면 안된다.”<sup id="cite_ref-martin-design-principles_4-1" class="reference"><a
          href="#cite_note-martin-design-principles-4">&#91;4&#93;</a></sup> <a
          href="/wiki/%EC%9D%98%EC%A1%B4%EC%84%B1_%EC%A3%BC%EC%9E%85" title="의존성 주입">의존성 주입</a>은 이 원칙을 따르는 방법 중 하나다.
        </dd>
      </dl>
    </td>
  </tr>
  </tbody>
</table>