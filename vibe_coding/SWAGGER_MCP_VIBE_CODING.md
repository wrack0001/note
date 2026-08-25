# 在新项目中使用 Swagger MCP：让 API 文档成为 Vibe Coding 的实时上下文

## 摘要

新项目可以从一开始就把 Proto、后端代码或接口注释作为真相源，自动生成
Swagger/OpenAPI，并通过 MCP 直接提供给 AI 编程工具。

开发者描述任务后，AI 可按需查询接口、请求参数和数据结构，无需人工复制整份文档。这样能
减少重复维护、文档漂移和无关 Token 消耗，也让生成结果更容易验证。

## 核心方案

新项目的接口契约生产和使用链路如下：

```text
后端代码、Proto 或接口注释
          ↓
自动生成并部署 Swagger/OpenAPI
          ↓
Swagger Reader MCP（只读检索）
          ↓
Claude Code 等 AI 编程工具按需查询
          ↓
生成或检查客户端代码
```

这里的 MCP 只负责读取接口文档，不调用 Swagger 中描述的真实业务接口，也不需要在后端
服务中增加新的业务 RPC。

它通常由 Claude Code 在开发者本机通过 `stdio` 启动：

1. MCP Server 从指定 URL 下载 Swagger/OpenAPI 文档。
2. MCP Server 解析并缓存接口路径、Operation 和 Schema。
3. Claude Code 根据任务调用检索工具。
4. MCP Server 只返回当前任务需要的局部接口信息。

## 为什么更利于 Vibe Coding

接入 MCP 后，接口上下文不再主要依赖开发者查找和粘贴文档。开发者只需要描述目标，AI 就能
先查询相关接口及 Schema，再根据实际契约编码。例如：

```text
使用 api-docs 查询用户详情接口，根据实际契约实现请求函数和 TypeScript 类型。
不要推测文档中不存在的字段。
```

MCP 按接口路径、Operation ID、Schema 名或关键词返回局部内容，避免把无关接口一起放入
上下文。调用本身仍会消耗 Token，但能减少重复粘贴整份文档造成的浪费。

接口变更后，重新生成并部署 Swagger/OpenAPI，客户端 AI 刷新缓存即可读取新契约。如果
Swagger/OpenAPI 由约定的真相源自动生成，字段不一致的问题就主要收敛在生成链路中，无需
同时检查多份手工副本。

AI 修改代码前可以先列出接口路径、HTTP Method、请求结构和响应 Schema，修改后再用同一份
契约检查客户端实现。项目级 MCP 配置还可以随仓库共享，让团队使用相同的 MCP 名称、文档
地址和 package 版本。

## 建议使用项目级 MCP 配置

对于 Swagger/OpenAPI 这类与具体项目接口契约绑定的 MCP，建议优先使用**项目级配置**，
将 `.mcp.json` 放在客户端项目根目录并提交到 Git。

不建议默认配置成用户级 MCP，主要原因是：

- **接口契约属于项目**：某个服务的 Swagger 地址通常只对对应客户端项目有意义。
- **团队配置一致且可追踪**：MCP 名称、package 版本、Swagger 地址和缓存策略都可接受
  代码审查。
- **新成员开箱即用**：克隆项目后无需根据口头说明重复配置。
- **避免跨项目污染**：只在需要这份接口契约的项目中启用对应能力。

Claude Code 首次加载项目级 MCP 时会要求确认信任。团队成员可先审查配置，再决定是否
启用。

项目级配置中可以提交公开文档地址，但不要提交 Token、API Key 等敏感信息。需要鉴权时，
应在 `.mcp.json` 中引用环境变量，由每位开发者在本机或安全的密钥系统中提供真实值。

## 文档职责应该怎样划分

| 文档载体 | 主要用途 | 适合承载的内容 | 不应重复维护的内容 |
|---|---|---|---|
| Swagger/OpenAPI | 机器可读的接口契约 | Method、路径、参数、请求与响应 Schema、字段类型、必填关系、枚举、格式、示例和标准错误响应 | 需求背景、跨接口业务流程和上线决策 |
| 飞书、Confluence 等知识库 | 供团队理解和讨论业务 | 业务目标、页面交互、调用流程、影响范围、兼容策略、灰度计划、上线时间及尚未固化的决策 | 完整的路径、参数和字段表 |
| 项目内的 `API.md`（可选） | 提供项目内的接口使用指引 | 模块能力概览、复杂调用示例、OpenAPI 难以完整表达的业务约束，以及 Swagger 地址或接口索引 | Swagger/OpenAPI 已经描述的字段级契约 |

字段级契约只维护在能够生成 Swagger/OpenAPI 的真相源中。

## 配置示例：Claude Code + swagger-reader-mcp

下面以开源 package `swagger-reader-mcp` 为例：

- GitHub：[`Abdallahabusnineh/swagger-reader-mcp`](https://github.com/Abdallahabusnineh/swagger-reader-mcp)
- Claude Code MCP 配置说明：[`Connect Claude Code to tools via MCP`](https://code.claude.com/docs/en/mcp)

在需要使用接口文档的客户端项目根目录创建 `.mcp.json`，并将它提交到 Git：

```json
{
  "mcpServers": {
    "api-docs": {
      "type": "stdio",
      "command": "npx",
      "args": [
        "-y",
        "swagger-reader-mcp@0.1.4"
      ],
      "env": {
        "SWAGGER_URL": "https://api.example.com/openapi.json",
        "SWAGGER_CACHE_TTL_MS": "180000"
      }
    }
  }
}
```

需要替换的内容：

- `api-docs`：MCP Server 名称，可以按项目修改。
- `SWAGGER_URL`：替换为团队实际可访问的 Swagger/OpenAPI JSON 地址。
- package 版本：建议锁定明确版本，升级前再进行验证。
- `SWAGGER_CACHE_TTL_MS`：缓存时间，示例中的 `180000` 表示 3 分钟。

直接填写文档 URL 最容易理解。如果开发者使用不同环境，可通过用户级配置或启动环境注入
地址，避免把个人地址提交到共享配置。

### CLI 一条命令安装

在客户端项目根目录执行下面的命令。Claude Code 会创建或更新项目级 `.mcp.json`：

```bash
claude mcp add api-docs --transport stdio --scope project \
  --env SWAGGER_URL=https://api.example.com/openapi.json \
  --env SWAGGER_CACHE_TTL_MS=180000 \
  -- npx -y swagger-reader-mcp@0.1.4
```

执行前将 `SWAGGER_URL` 替换为实际文档地址。`--scope project` 用于生成可供团队共享的
项目级 `.mcp.json`，不要省略。

安装后检查：

```bash
claude mcp get api-docs
```

### 让 AI 一键完成配置

如果不希望手工编辑 JSON，也可以在 Claude Code 或其他能够修改项目文件的 AI 编程工具中，
直接发送下面这段提示词：

```text
请在当前项目中配置项目级 Swagger Reader MCP。

要求：
1. 在项目根目录创建或更新 .mcp.json，不要写入用户级配置。
2. 如果 .mcp.json 已存在，合并到现有 mcpServers，不要覆盖其他 MCP。
3. Server 名称使用 api-docs，transport 使用 stdio。
4. command 使用 npx，args 使用 ["-y", "swagger-reader-mcp@0.1.4"]。
5. SWAGGER_URL 设置为：https://api.example.com/openapi.json
6. SWAGGER_CACHE_TTL_MS 设置为 180000。
7. 不要把 Token、API Key 或其他凭证写进仓库。
8. 修改后校验 JSON 格式，并告诉我新增或修改了哪些内容。
```

使用前替换其中的 Swagger/OpenAPI 地址。修改后仍应检查 `.mcp.json` 差异，并在首次启动
Claude Code 时确认信任提示。

## 首次启用和验证

本机需要安装 Node.js，并确保 `npx` 可用。具体 Node.js 版本以 package 当前要求为准。

在包含 `.mcp.json` 的项目目录启动 Claude Code：

```bash
claude
```

首次读取项目级 MCP 配置时，Claude Code 会请求信任确认。批准后可以检查：

```bash
claude mcp get api-docs
```

在 Claude Code 的 `/mcp` 页面中，确认服务已连接，并能看到搜索接口、读取端点和 Schema、
查看概览及刷新文档等工具。

除确认配置已注册外，至少进行一次真实查询：

```text
使用 api-docs 查询一个已知接口，返回它的路径、Method、请求参数和响应 Schema。
```

## 推荐的日常提示词

### 开发新功能

```text
先使用 api-docs 查询与用户登录有关的接口。
列出实际路径、Method、请求参数和响应 Schema，再实现请求函数和类型。
不要推测接口文档中不存在的字段。
```

### 后端刚更新接口

```text
先刷新 api-docs 的 Swagger 缓存，再重新查询订单详情接口。
检查当前客户端类型和请求代码是否需要同步修改。
```

### 只分析、不修改代码

```text
使用 api-docs 读取目标接口及相关 Schema，对比当前客户端实现，
列出路径、参数、可选字段和枚举值的不一致，暂时不要修改代码。
```

### 开发完成后检查

```text
重新读取本次涉及的接口契约，检查请求路径、Method、参数位置、
TypeScript 类型和错误处理是否与文档一致。
```

检索时优先使用接口路径、Operation ID、Schema 名或单个明确关键词。过多中英文关键词
可能降低部分检索器的命中率。

## 落地与日常协作

新项目先约定接口真相源和 Swagger/OpenAPI 的生成方式，部署稳定可访问的文档地址，再将
项目级 `.mcp.json` 提交到客户端仓库。用一个真实接口完成查询验证后，即可作为团队的标准
配置。

接口更新时，后端只修改真相源并重新生成、部署 Swagger/OpenAPI，不直接修改生成产物。
知识库记录业务背景、影响范围和升级要求，客户端刷新 MCP 缓存后按新契约对接。

CI/CD 负责检查生成产物是否一致、格式和关键路由是否有效，以及部署后的文档地址是否可访问，
避免依赖人工逐项确认。

## 结论

Swagger MCP 的价值不是增加一种文档格式，而是让机器可读的接口契约成为 AI 可主动检索的
开发上下文。

后端负责从真相源生成并部署可靠的 Swagger/OpenAPI；知识库负责业务背景、流程和人工决策；
MCP 负责把当前任务需要的接口契约交给 AI。这样可以减少重复维护和文档传播延迟，降低
接口文档漂移的概率，也让 Vibe Coding 的输入更准确、更精简、更容易验证。
