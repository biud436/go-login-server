# Go Login Server

Golang을 사용한 간단한 JWT 기반 로그인 서버입니다.

## 기술 스택

- **Golang**
- **Gin**: 웹 프레임워크
- **GORM**: ORM 라이브러리
- **PostgreSQL**
- **JWT**: 인증 토큰
- **Docker Compose**

## 설치 및 실행

1. **Docker Compose로 PostgreSQL 생성 및 실행**

   ```bash
   docker-compose up -d
   ```

2. **프로젝트 의존성 설치**

   ```bash
   go mod tidy
   ```

3. **애플리케이션 실행**

   ```bash
   go run main.go
   ```

   서버는 `localhost:8080`에서 실행됩니다.

## API 엔드포인트

- **POST /register**

  회원가입을 위한 엔드포인트.

  **Request Body:**

  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```

- **POST /login**

  로그인을 위한 엔드포인트.

  **Request Body:**

  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```

  **Response:**

  ```json
  {
    "token": "your_jwt_token"
  }
  ```

- **GET /protected/dashboard**

  인증이 필요한 보호된 엔드포인트.

  **Headers:**

  ```
  Authorization: Bearer your_jwt_token
  ```

  **Response:**

  ```json
  {
    "message": "Welcome your_username"
  }
  ```

## 환경 변수

프로젝트의 환경 변수는 `.env` 파일에 정의됩니다.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=login_db
JWT_SECRET=your_secret_key
```

## 실행

서버를 실행하려면 다음 명령을 실행하십시오.

```bash
go run main.go
```

스웨거 문서는 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)에서 확인할 수 있습니다.

## 스웨거 문서 생성

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

```bash
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
```

`$GOPATH/bin` 디렉토리가 PATH에 포함되어 있는지 확인합니다. 없으면 다음 명령어로 `PATH`에 추가합니다.

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

swag init 명령을 실행하여 스웨거 문서를 생성합니다.

```bash
swag init
```

만일 서버가 제대로 실행되지 않는다면, 다음 명령을 실행하여 swag를 업데이트합니다.

```bash
go get -u github.com/swaggo/swag
```

## 버전 관리 및 업데이트

```bash
go get -u
go mod tidy
```
