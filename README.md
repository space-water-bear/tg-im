# 结构
```shell
.
├── consumer
├── go.mod
├── internal
│   └── model
├── job
├── pkg
├── restful
├── script
└── service
```

```html
consumer： 队列消费服务
internal： 工程内部可访问的公共模块
job： cron job 服务
pkg： 工程外部可访问的公共模块
restful：HTTP 服务目录，下存放以服务为维度的微服务
script：脚本服务目录，下存放以脚本为维度的服务
service：gRPC 服务目录，下存放以服务为维度的微服务
```

# 架构图
```mermaid
graph TD
    subgraph Gateway Service
        A[API Gateway]
    end

    subgraph User Service
        B1[User Registration]
        B2[User Authentication]
        B3[User Profile]
    end

    subgraph Message Service
        C1[HTTP Message Send]
        C2[SSE Message Push]
        C3[NSQ Message Queue]
        C4[Message Storage]
    end

    subgraph Friend Service
        D1[Friend Management]
        D2[Friend List]
    end

    subgraph Group Service
        E1[Group Management]
        E2[Group Messaging]
    end

    subgraph Notification Service
        F1[System Notifications]
        F2[Non-Instant Messages]
    end

    subgraph Database
        DB1[MySQL]
        DB2[Redis]
        DB3[NSQ]
    end

    A --> |Register/Login| B1
    A --> |Send Message| C1
    A --> |Receive Message| C2
    A --> |Friend Management| D1
    A --> |Group Management| E1
    A --> |System Notifications| F1

    B1 --> DB1
    B2 --> DB1
    B3 --> DB1

    C1 --> C3
    C3 --> C2
    C3 --> C4
    C2 --> C4
    C4 --> DB1
    C4 --> DB2
    C3 --> DB3

    D1 --> DB1
    D2 --> DB1

    E1 --> DB1
    E2 --> DB1

    F1 --> DB2
    F2 --> DB2
```

```mermaid
graph TD
    Client --> |Register/Login| API_Gateway
    Client --> |Send Message| API_Gateway
    Client --> |Receive Message SSE| API_Gateway
    Client --> |Friend Management| API_Gateway
    Client --> |Group Management| API_Gateway
    Client --> |System Notifications| API_Gateway

    API_Gateway --> |Register/Login| User_Service
    API_Gateway --> |Send Message| Message_Service
    API_Gateway --> |Receive Message SSE| Message_Service
    API_Gateway --> |Friend Management| Friend_Service
    API_Gateway --> |Group Management| Group_Service
    API_Gateway --> |System Notifications| Notification_Service

    subgraph Gateway_Service
        API_Gateway[API Gateway]
    end

    subgraph User_Service
        US[User Service]
    end

    subgraph Message_Service
        MS[Message Service]
    end

    subgraph Friend_Service
        FS[Friend Service]
    end

    subgraph Group_Service
        GS[Group Service]
    end

    subgraph Notification_Service
        NS[Notification Service]
    end

```