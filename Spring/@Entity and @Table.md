# @Entity와 @Table

JPA를 사용하는데 매번 Entity만 사용해왔다.
그런데???? Table이란걸 확인했는데.. 이게 뭘까?? 그래서 한번 확인해봤다.

### [참고 자료](https://walkingtechie.blogspot.com/2019/06/difference-between-entity-and-table.html)

## @Entity

### @javax.persistence.Entity

- @Entity는 클래스가 엔티티임을 지정한다.



## @Table

### @javax.persistence.Table

- 엔티티가 기본 테이블임을 지정한다.
- 추가 테이블은 SecondaryTable 또는 SecondaryTables 어노테이션을 사용해 지정할 수 있다.
- 엔티티 클래스에서 Table 어노테이션이 지정되지 않을 경우 기본값으로 적용된다.

---

## @Entity와 Table차이

1. @Entity(name = "someName"): "someName"은 엔티티의 이름을 지정하는 데 사용한다.
2. @Table(name = "otherName"): "otherName"은 DB에서 테이블 이름을 지정하는데  사용한다.

### Entitiy와 Table 사용 시나리오

예시 클래스
```java
// Message.java

package com.walking.techie.entity;
 
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import lombok.Data;
import lombok.NoArgsConstructor;
 
@Data
@NoArgsConstructor
@Entity
public class Message {
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Column(name = "ID")
  private Long id;
 
  @Column(name = "MESSAGE")
  private String message;
 
  public Message(String message) {
    this.message = message;
  }
}
```

---

### Scenario 1: @Entity 어노테이션만 사용

Hibernate 쿼리를 보며 차이를 확인해보자.

#### Scenario 1.1: "name"요소 없이 @Entity 어노테이션만 사용할 경우

```java
// Message.java
@Entity
public class Message {
    // .............
    // .............
}
```

- name 요소를 사용하지 않을경우 테이블 및 엔티티 이름은 기본 값과 같다.
- 기본 값은 엔티티 클래스 이름과 같으며, 지속성 클래스의 이름이 된다.

> Hibernate Query: create table Message (ID bigint not null auto_increment, MESSAGE varchar(255), primary key (ID)) engine=InnoDB

#### Scenario 1.2: "name"요소를 포함한 @Entity 어노테이션만 사용할 경우

```java
// Message.java

@Entity(name = "message")
public class Message {
    // .............
    // .............
}
```

- 이름 요소를 사용할 경우 테이블 및 엔티티 이름은 이름 요소와 값이 동일하다.

> Hibernate Query: create table message (ID bigint not null auto_increment, MESSAGE varchar(255), primary key (ID)) engine=InnoDB
> 
> Scenario 1.1과 테이블 이름이 다르다는걸 확인할 수 있다.




위 경우 테이블과 엔티티의 이름이 같으므로 HQL또는 JPQL을 작성하는 동안 엔티티와 동일한 이름으로 테이블에 엑세스 할 수 있다.

---

### Scenario 2: @Entity, @Table 어노테이션을 모두 사용할 경우

#### Scenario 2.1: 요소 없이 어노테이션을 사용할 경우

```java
// Message.java

@Entity
@Table
public class Message {
   // .............
   // .............
}
```

Scenario 1.1과 같은 결과가 나온다.

> Hibernate Query: create table message (ID bigint not null auto_increment, MESSAGE varchar(255), primary key (ID)) engine=InnoDB

#### Scenario 2.2: Entity 어노테이션의 요소를 사용할 경우

```java
// Message.java

@Entity(name = "MESSAGE")
@Table
public class Message {
    // .............
    // .............
}
```

> Hibernate Query: create table MESSAGE (ID bigint not null auto_increment, MESSAGE varchar(255), primary key (ID)) engine=InnoDB
>
> 테이블 이름만 바뀐것을 확인할 수 있다.


#### Scenario 2.3: Table 어노테이션의 요소를 사용할 경우

```java
// Message.java

@Entity
@Table(name = "text_message")
public class Message {
    // .............
    // .............
}
```

> Hibernate Query: create table text_message (ID bigint not null auto_increment, MESSAGE varchar(255), primary key (ID)) engine=InnoDB
> 
> 테이블 이름이 Table 어노테이션 요소와 같다.

#### Scenario 2.4: Entity, Table 요소를 모두 사용할 경우

```java
// Message.java

@Entity(name = "MESSAGE")
@Table(name = "text_message")
public class Message {
    // .............
    // .............
}
```

> Hibernate Query: create table text_message (ID bigint not null auto_increment, MESSAGE varchar(255), primary key (ID)) engine=InnoDB
> 
> Table의 요소를 기준으로 테이블 이름이 생성된다.

이 경우 쿼리를 작성할 땐 @Entity에 지정된 이름을 사용해야하며, 
@Table에 지정된 이름은 DB의 테이블 이름을 지정하는데 사용하게 된다. 
따라서 JPQL에서 사용하는 엔티티 이름은 DB의 다른 이름을 참조하게 된다.


# 결론 

@Table을 사용하는 경우 테이블이 만들어지는 순간은 Table요소에 맞게 만들어 지지만 사용하면서 Table요소와 Entity요소를 같이 사용하게 될 것이다.
이렇게 되면 대 혼란이 올 것같은데.

처음 시작하는 프로젝트면 그냥 모두 똑같게 통일해서 하는게 맞다고 본다.

부득이하게 서비스 별로 테이블 이름이나 변수명이 다를경우 사용하면 될 것 같다. 나의 느낌상...