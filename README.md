# channelio-scheduler

채널톡 그룹 내에서 일정을 등록, 확인, 삭제할 수 있습니다.

## 기술 스택

- Server: Fiber (golang)
- DB: PostgreSQL
- Channel Open API
- Channel Webhook

## How to start

### .env 설정

```
TOKEN="(필수) 채널톡 웹훅 토큰"
OPEN_API_ACCESS_KEY="(필수) 채널톡 Open API access key"
OPEN_API_ACCESS_SECRET="(필수) 채널톡 Open API access secret"
PORT="(옵션) 서버 포트, 기본 4000"
```

### 필요 모듈 설치

```bash
go get
```

### 서버 실행

```bash
go run main.go
```

## 기능

### 일정 추가

**형식**

```
-일정 추가 [yyyy/]mm/dd [hh:mm] 일정_이름
```

- 일시와 이름을 기반으로 일정을 생성할 수 있습니다.
- 대괄호 내부 부분은 생략 가능합니다.

![add](https://user-images.githubusercontent.com/34625313/168444888-53684a22-0657-43e8-8e8a-3047c221e2c6.png)

### 일정 조회

**형식**

```
-일정 조회
```

- 현재까지 그룹에 추가된 모든 일정을 확인할 수 있습니다.
- 해당 일정의 시간 순으로 정렬되어 표시됩니다.
- 모든 일정에는 번호가 부여됩니다. 이후 삭제할 때 이 번호가 기준이 됩니다.
- 지난 일정은 표시되지 않습니다.

![get](https://user-images.githubusercontent.com/34625313/168444850-9f65ce57-9e4b-41e8-b5ba-f189fff14b7f.png)

### 일정 삭제

**형식**

```
-일정 삭제 일정_번호
```

- 그룹 내 일정을 번호 기준으로 삭제할 수 있습니다.
- 일정 삭제에 앞서 일정 조회가 필요합니다. 삭제에 사용될 번호가 필요하기 때문입니다.

![delete](https://user-images.githubusercontent.com/34625313/168444856-923b68e1-bca9-41c0-bae0-36a1a42fc7e9.png)


## 모델

2개의 테이블이 존재합니다.

1. schedules

   일정 정보를 관리하는 테이블입니다. Column은 다음과 같습니다:

   - id (PK): uuid
   - created_at, deleted_at: time
   - channel_id: text, 그룹의 고유한 id
   - title: text, 일정의 이름
   - datetime: time 일정의 일시
   - is_time_set: bool 시간까지 지정했는지 여부

2. get_schedule_history

   일정을 삭제할 때 어떤 일정을 삭제할지 확인하기 위한 테이블입니다. Column은 다음과 같습니다:

   - id (PK): uuid
   - created_at, updated_at, deleted_at: time
   - channel_id: text, 그룹의 고유한 id
   - person_id: text, 유저의 고유한 id
   - result: text[], 가장 마지막으로 조회한 일정들의 id 목록

   유저가 번호를 기준으로 삭제를 요청할 경우 이 중 result array를 이용해 일정을 찾고 삭제합니다. 그룹에 추가나 삭제가 일어날 경우 모든 get_schedule_history 내의 row는 삭제됩니다.
