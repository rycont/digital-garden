---
title: PostgreSQL에서 INSERT 할 때만 Computed 되는 Field 구현하기
layout: ../layouts/article.astro
date: 2023-06-06T04:32:10.702Z
---

# PostgreSQL에서 INSERT 할 때만 Computed 되는 Field 구현하기

이런 필드를 원했다

1. count(\*) 함수를 사용할 수 있어야 함

2. UPDATE해도 값이 바뀌면 안됨

정리하자면, INSERT할 때만 Computed되는 필드를 원했다.

#### **GENERATED 필드**

[[PostgreSQL]]에는 필드에 GENERATED라는 속성을 적용할 수 있다. 이는 일반적인 computed field와 동일한 개념으로, 레코드의 다른 필드를 참조해서 새로운 값을 계산할 수 있다. 코드로는 다음과 같이 표현된다.

```javascript
CREATE TABLE Books (
    id SERIAL PRIMARY KEY,
    name TEXT,
    nice_name TEXT GENERATED ALWAYS AS ('NICE ' || name) STORED
);
```

```javascript
INSERT INTO Books(name) VALUES('economy');
UPDATE Books SET name = 'business' WHERE id = 1;

SELECT * FROM Books;
```

```javascript
 id |   name   |   nice_name
----+----------+---------------
  1 | business | NICE business
(1 row)
```

최초에 INSERT를 할 때는 name을 economy로 넣었지만, 이후 update를 통해 name을 business로 바꾸었고 테이블을 출력해보니 `nice_name` 또한 `NICE business`로 바뀌었다. 왜냐하면 GENERATED 필드는 INSERT할 때 뿐만 아니라 UPDATE할 때도 새로 계산하기 때문이다.

또한 Generated는 aggregation 함수를 사용할 수 없다. 그렇기 때문에 계산식에 count(\*)를 포함한다면\
ERROR:  aggregate functions are not allowed in column generation expressions\
라는 오류가 발생한다.

#### **트리거와 함수**

결국 트리거와 함수를 사용하여 해결하였다. 사용한 코드는 다음과 같다.

```javascript
CREATE TABLE Books (
    id SERIAL PRIMARY KEY,
    name TEXT,
    nice_name TEXT NULL
);

CREATE OR REPLACE FUNCTION set_fancy_name()
    RETURNS TRIGGER AS $$
DECLARE BEGIN
    NEW.nice_name := NEW.name || ' is fancy';
    RETURN NEW;
END; $$ LANGUAGE plpgsql;


CREATE OR REPLACE TRIGGER set_nice_name_books
BEFORE INSERT ON Books
FOR EACH ROW EXECUTE FUNCTION set_fancy_name();
```

레코드를 변경하는 함수를 작성하고, 이를 트리거로 테이블에 레코드를 삽입할 때 마다 실행하도록 설정해주었다.

```javascript
INSERT INTO Books(name) VALUES ('economy');
UPDATE Books SET name = 'business' WHERE id = 1;
SELECT * FROM Books;
```

```javascript
 id |   name   |    nice_name
----+----------+------------------
  1 | business | economy is fancy
(1 row)
```

예상한대로 결과가 나오는 것을 알 수 있다.
