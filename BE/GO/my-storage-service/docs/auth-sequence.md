# Auth Sequence Diagrams

เอกสารนี้สรุปลำดับการทำงานของระบบ Auth แบบ REST ใน my-storage-service ด้วย Mermaid

## 1) Register

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant H as AuthHandler
    participant U as UserUsecase
    participant P as PasswordService
    participant DB as UserRepository(memory)
    participant AJ as AccessJWT
    participant RJ as RefreshJWT

    C->>R: POST /api/v1/auth/register
    R->>H: Register(req)
    H->>U: Register(email,password)
    U->>DB: GetByEmail(email)
    DB-->>U: not found
    U->>P: Hash(password)
    P-->>U: passwordHash
    U->>DB: Create(user)
    DB-->>U: createdUser
    U->>AJ: Issue(userID,email)
    AJ-->>U: accessToken
    U->>RJ: Issue(userID,email)
    RJ-->>U: refreshToken
    U-->>H: AuthResult
    H-->>C: 201 + Set-Cookie(access/refresh)
```

## 2) Login

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant H as AuthHandler
    participant U as UserUsecase
    participant P as PasswordService
    participant DB as UserRepository(memory)
    participant AJ as AccessJWT
    participant RJ as RefreshJWT

    C->>R: POST /api/v1/auth/login
    R->>H: Login(req)
    H->>U: Login(email,password)
    U->>DB: GetByEmail(email)
    DB-->>U: user + passwordHash
    U->>P: Compare(passwordHash,password)
    P-->>U: ok
    U->>AJ: Issue(userID,email)
    AJ-->>U: accessToken
    U->>RJ: Issue(userID,email)
    RJ-->>U: refreshToken
    U-->>H: AuthResult
    H-->>C: 200 + Set-Cookie(access/refresh)
```

## 3) Refresh Token

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant H as AuthHandler
    participant U as UserUsecase
    participant DB as UserRepository(memory)
    participant RJ as RefreshJWT
    participant AJ as AccessJWT

    C->>R: POST /api/v1/auth/refresh (cookie refresh_token)
    R->>H: Refresh()
    H->>U: Refresh(refresh_token)
    U->>RJ: Parse(refresh_token)
    RJ-->>U: claims(email,userID)
    U->>DB: GetByEmail(email)
    DB-->>U: user
    U->>AJ: Issue(userID,email)
    AJ-->>U: newAccessToken
    U->>RJ: Issue(userID,email)
    RJ-->>U: newRefreshToken
    U-->>H: AuthResult
    H-->>C: 200 + Set-Cookie(access/refresh)
```

## 4) Me (Protected Endpoint)

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant M as JWTAuthMiddleware
    participant AJ as AccessJWT
    participant H as AuthHandler
    participant U as UserUsecase
    participant DB as UserRepository(memory)

    C->>R: GET /api/v1/auth/me (cookie access_token)
    R->>M: JWTAuth()
    M->>AJ: Parse(access_token)
    AJ-->>M: claims(userID,email)
    M->>H: next(auth_user_email in context)
    H->>U: Me(email)
    U->>DB: GetByEmail(email)
    DB-->>U: user
    U-->>H: UserInfo
    H-->>C: 200 user profile
```

## 5) Logout

```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant H as AuthHandler

    C->>R: POST /api/v1/auth/logout
    R->>H: Logout()
    H-->>C: 200 + Clear-Cookie(access/refresh)
```
