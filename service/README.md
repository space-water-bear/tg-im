# 服务器关系

```mermaid
graph TD
    Client[客户端]
    Gateway[网关]
    UserService[用户服务]
    FriendService[好友服务]
    MessageService[消息服务]

    Client -->|常规HTTP请求| Gateway
    Client <--[消息推送](链接地址)  Gateway
    Gateway -->|用户操作| UserService
    Gateway -->|好友操作| FriendService
    Gateway -->|消息操作| MessageService
```