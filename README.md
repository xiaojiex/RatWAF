# RatWAF 项目 README

# RatWAF

一款轻量级、高性能的 Web 应用防火墙（WAF），支持反向代理抓包、攻击日志记录、Web 可视化管理，操作简洁、部署便捷，适用于小型 Web 服务的安全防护。

## 📌 项目简介

RatWAF 核心功能包括反向代理转发、攻击流量抓包、日志入库存储、Web 面板可视化管理，可实时监控攻击者 IP、攻击路径、攻击时间，支持日志详情查看，帮助用户快速识别并追踪异常访问。

项目架构清晰，分为三大模块：反向代理模块（network）、数据库模块（database）、Gin Web 管理面板（web），各模块独立运行、互不冲突，部署简单且可灵活扩展。

## ⚙️ 核心功能

- 反向代理：转发前端请求，拦截异常流量，实现 WAF 网关功能

- 日志记录：自动抓取攻击者 IP、完整请求 URL、攻击时间、请求体，存入 MySQL 数据库

- Web 管理面板：可视化展示攻击日志列表，支持点击查看日志详情

## 📋 环境要求

- Go 1\.18\+（推荐 1\.20\+）

- MySQL 5\.7\+ / MariaDB 10\.0\+

- Git（可选，用于拉取项目）

## 🚀 部署步骤

### 1\. 克隆项目（若有）

```bash
# 克隆项目（若已下载可跳过）
git clone https://github.com/xiaojiex/RatWAF.git
cd RatWAF
go mod init RatWAF
```

### 2\. 配置数据库

1. 启动 MySQL 服务，创建数据库（示例数据库名：waf）

2. 执行数据库表创建语句（创建 waf\_log 表）：

```sql
CREATE DATABASE IF NOT EXISTS waf DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE waf;

CREATE TABLE `waf_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `attacker_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `request` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `createtime` datetime DEFAULT CURRENT_TIMESTAMP,
  `full_url` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'URL',
  `attack_type` text COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

修改 `gorrm/Conn\.go` 中的数据库连接信息（用户名、密码、数据库名）：

```go
dsn := "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
```

### 3\. 安装依赖

```bash
# 安装 Gin 和 GORM 依赖
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
#或
go mod tidy
```

### 4\. 启动项目

```bash
# 直接启动
go run RatWAF

# 后台启动（Linux）
nohup go run RatWAF > ratwaf.log 2>&1 
```

### 5\. 访问面板

启动成功后，访问以下地址即可进入 WAF 管理面板：

Web 管理面板：`http://127\.0\.0\.1:9090`

反向代理入口：`http://127\.0\.0\.1:8080`（转发至本地业务服务 8080 端口）

## 📊 页面说明

### 1\. 日志列表页（首页）

- 左侧侧边栏为功能菜单（攻击总览、日志列表、防御规则等，部分为占位）

- 主内容区展示攻击日志列表，每条日志包含：攻击者 IP、攻击路径、攻击时间

- 点击任意日志卡片，可跳转至详情页查看完整信息

### 2\. 日志详情页

- 展示单条日志的完整信息：日志 ID、攻击者 IP、攻击路径、攻击时间、请求内容

- 支持点击「返回列表」按钮，回到日志列表页

## 🗑️ 清空日志

若需清空数据库中的所有攻击日志，可进入 MySQL 控制台执行以下命令：

```sql
USE waf;
TRUNCATE TABLE waf_log;  -- 清空数据，重置ID（推荐）
-- 或
DELETE FROM waf_log;  -- 仅清空数据，不重置ID
```

## 🔧 常见问题

- Q：启动项目后，Web 面板无法访问？
A：检查端口 9090 是否被占用，或修改 `main\.go` 中 `r\.Run\(\&\#34;:9090\&\#34;\)` 的端口号。

- Q：日志无法正常显示？
A：检查数据库连接是否正确，waf\_log 表字段是否与`gorrm\.WafLog` 结构体对应。

- Q：反向代理无法转发请求？
A：检查 `neet/ReverseProxy\(\)` 中的目标服务端口（默认 8080），确保业务服务正常启动。

## 📌 后续可扩展功能

- 日志筛选（按 IP、时间、URL 关键词筛选）

- 日志分页（解决大量日志加载卡顿问题）

- Web 面板一键清空日志功能

- 防御规则配置（拦截特定 IP、恶意请求）

- 流量统计图表（可视化展示攻击趋势）

## 📄 项目结构

```bash
RatWAF/
├── database/          # 数据库模块（连接、结构体定义）
│   └── Conn.go     # 数据库连接配置
├── network/           # 反向代理模块
│   └── neet.go     # 反向代理逻辑
├── regex/           # 反向代理模块
│   └── rce.go    # 命令执行验证
│   └── sql.go     # sql注入验证
├── web/            # Gin Web 管理面板
│   ├── router.go   # 路由配置（列表页、详情页）
│   ├── handler.go  # 接口逻辑（GetData、GetDetail）
│   └── templates/  # HTML 模板
│       ├── index.html  # 日志列表页
│       └── detail.html # 日志详情页
├── main.go         # 项目入口（启动数据库、代理、Web面板）
└── README.md       # 项目说明文档
```
