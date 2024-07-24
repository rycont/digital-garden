---
title: "리눅스에서 .mdb(Microsoft Access) 파일을 CSV로 변환하기"
date: 2024-06-06T01:14:15.957Z
---

# 리눅스에서 .mdb([[Microsoft]] Access) 파일을 [[CSV]]로 변환하기

mdbtools라는 패키지를 사용하면 된다.

```sh
# For Debian / Ubuntu
apt install mdbtools
```

다음 명령어로 테이블 목록을 확인한다

```sh
mdb-tables file.mdb
```

다음 명령어로 테이블을 변환한다

```sh
mdb-export file.mdb <table_name>
```
