### 1. "Register a new user"

1. route definition

- Url: /api/v1/api/v1/register
- Method: POST
- Request: `RegisterRequest`
- Response: `RegisterResponse`

2. request definition



```golang
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
```


3. response definition



```golang
type RegisterResponse struct {
	Success bool `json:"success"`
}
```

### 2. "Login"

1. route definition

- Url: /api/v1/login
- Method: POST
- Request: `LoginRequest`
- Response: `LoginResponse`

2. request definition



```golang
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResponse struct {
	UserId int64 `json:"userId"`
	Token string `json:"token"`
}
```

### 3. "Get user list"

1. route definition

- Url: /api/v1/user
- Method: GET
- Request: `-`
- Response: `UserListResponse`

2. request definition



3. response definition



```golang
type UserListResponse struct {
	Users []UserInfoResponse `json:"users"`
}
```

### 4. "Add a new user"

1. route definition

- Url: /api/v1/user
- Method: POST
- Request: `AddUserRequest`
- Response: `AddUserResponse`

2. request definition



```golang
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
```


3. response definition



```golang
type AddUserResponse struct {
	UserId int64 `json:"userId"`
}
```

### 5. "Update user info"

1. route definition

- Url: /api/v1/user
- Method: PUT
- Request: `UpdateUserInfoRequest`
- Response: `UpdateUserInfoResponse`

2. request definition



```golang
type UpdateUserInfoRequest struct {
	UserId int64 `json:"userId"`
	Username string `json:"username,omitempty"`
	Email string `json:"email,omitempty"`
}
```


3. response definition



```golang
type UpdateUserInfoResponse struct {
	Success bool `json:"success"`
}
```

### 6. "Update user info"

1. route definition

- Url: /api/v1/user
- Method: PATCH
- Request: `UpdateUserInfoRequest`
- Response: `UpdateUserInfoResponse`

2. request definition



```golang
type UpdateUserInfoRequest struct {
	UserId int64 `json:"userId"`
	Username string `json:"username,omitempty"`
	Email string `json:"email,omitempty"`
}
```


3. response definition



```golang
type UpdateUserInfoResponse struct {
	Success bool `json:"success"`
}
```

### 7. "Delete user"

1. route definition

- Url: /api/v1/user
- Method: DELETE
- Request: `DeleteUserRequest`
- Response: `DeleteUserResponse`

2. request definition



```golang
type DeleteUserRequest struct {
	UserId int64 `json:"userId"`
}
```


3. response definition



```golang
type DeleteUserResponse struct {
	Success bool `json:"success"`
}
```

### 8. "Get user info"

1. route definition

- Url: /api/v1/user/:id
- Method: GET
- Request: `-`
- Response: `UserInfoResponse`

2. request definition



3. response definition



```golang
type UserInfoResponse struct {
	UserId int64 `json:"userId"`
	Username string `json:"username"`
	NickName string `json:"nickName"`
}
```

