# Sub Audit Gateway

订阅审计网关

一个部署在订阅服务前端的安全代理网关，用于隐藏真实订阅地址，并提供访问审计、客户端识别、访问风控和异常报警功能。

---

## Features

✅ Subscription Reverse Proxy  
隐藏真实订阅源地址


✅ Access Audit  
记录订阅访问行为：

- IP地址
- IPv4 / IPv6
- User-Agent
- 客户端类型
- 请求路径
- 状态码
- 请求时间


✅ Client Detection

自动识别：

- Shadowrocket
- Clash
- sing-box
- Surge
- v2rayN
- NekoBox


✅ Rate Limit

访问频率控制：

- 单IP请求限制
- 防止订阅扫描
- 防止恶意刷请求


✅ Telegram Alert

支持报警：

- 上游订阅异常
- IP访问超限
- 风控事件


✅ Docker Deployment

支持：

- Docker
- Docker Compose
- VPS快速部署


---

# Architecture


```
用户客户端

Shadowrocket
Clash
sing-box
Surge

        ↓

Sub Audit Gateway

        ↓

真实订阅服务器

        ↓

节点服务
```


---

# Project Structure


```
sub-audit-gateway

├── admin
│   └── status.go

├── config
│   ├── config.go
│   └── config.yaml

├── database
│   ├── database.go
│   └── model.go

├── handler
│   └── subscribe.go

├── middleware
│   ├── audit.go
│   ├── ip.go
│   ├── client.go
│   ├── rate.go
│   └── save.go

├── notify
│   ├── telegram.go
│   └── alert.go

├── Dockerfile

├── docker-compose.yml

└── README.md
```


---

# Deploy


## 1. Clone


```bash
git clone https://github.com/username/sub-audit-gateway.git

cd sub-audit-gateway
```


---

## 2. Configure


复制环境变量：

```bash
cp .env.example .env
```


编辑：

```bash
nano .env
```


配置：

```env
TELEGRAM_ENABLE=true

TELEGRAM_BOT_TOKEN=

TELEGRAM_CHAT_ID=
```


---

## 3. Start


Docker启动：

```bash
docker compose up -d
```


查看运行状态：

```bash
docker ps
```


查看日志：

```bash
docker logs -f sub-audit-gateway
```


---

# Configuration


配置文件：

```
config/config.yaml
```


示例：

```yaml
server:

  port: 8080


subscription:

  upstream: https://your-subscription-server.com


telegram:

  enable: true

  bot_token: ""

  chat_id: ""
```


---

# API


## Health Check


```
GET /status
```


返回：

```json
{
  "service": "sub-audit-gateway",
  "status": "running",
  "version": "v1.0"
}
```


---

## Subscription Proxy


```
GET /sub/{token}
```


示例：

```
https://gateway.example.com/sub/xxxxxx
```


请求流程：

```
用户

↓

网关

↓

真实订阅服务器

↓

返回订阅内容
```


---

# Audit Database


系统使用 SQLite 保存审计记录。


记录内容：

- Token
- IP
- IP版本
- User-Agent
- 客户端类型
- 请求路径
- HTTP状态
- 请求时间


数据库：

```
/app/data/audit.db
```


---

# Security


建议：

- 使用 HTTPS
- 配合 Cloudflare
- 配合 Nginx/Caddy
- 定期备份数据库
- 不公开数据库文件


---

# Monitoring


支持接入：

- Uptime Kuma
- 哪吒监控
- Telegram Bot


状态检测：

```
GET /status
```


---

# Roadmap


计划：

- Web管理后台
- IP国家查询
- ASN识别
- 黑名单系统
- 多节点管理
- PostgreSQL支持
- 数据统计面板


---

# License


MIT License
