# MCP for JueJin

<p align="center">
  <strong>JueJin-MCP是一个能让AI生成的文章自动发布到掘金的工具</strong>
</p>

### 🛠️ 如何使用？

<details>
<summary><b>在 Cursor 使用</b></summary>

Cursor 配置: `Settings` -> `Cursor Settings` -> `MCP` -> `Add new global MCP server`

Pasting the following configuration into your Cursor `~/.cursor/mcp.json` file is the recommended approach. You may also install in a specific project by creating `.cursor/mcp.json` in your project folder. 查看 [Cursor MCP docs](https://docs.cursor.com/context/model-context-protocol) for more info.

> Since Cursor 1.0, you can click the install button below for instant one-click installation.

#### Cursor Local Server Connection

[![Install MCP Server](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en/install-mcp?name=context7&config=eyJjb21tYW5kIjoibnB4IC15IEB1cHN0YXNoL2NvbnRleHQ3LW1jcCJ9)

```json
{
  "mcpServers": {
    "juejin-mcp": {
      "type":"sse",
      "url": "http://localhost:10086/mcp"
    }
  }
}
```

</details>


### 项目结构
```bash
root/
├── main.go               # 🚀 应用入口与引导程序
├── routes.go             # 🌐 HTTP 路由定义
├── browser/              # 🖥️ 浏览器自动化/爬虫模块
│   └── browser.go        #     浏览器实例管理与核心逻辑
├── juejin/               # 💎 掘金业务核心模块
│   ├── juejin.go         #     掘金公共变量
│   └── [FeatureName].go  #     特定功能子模块 (如: login, publish)
├── mcp/                  # 🤖 Model Context Protocol (AI 接口层)
│   ├── mcp.go            #     MCP 服务器实现
│   └── tools.go          #     提供给 AI 调用的具体工具函数
├── middleware/           # 🛡️ 中间件 (鉴权, 日志, 错误处理)
│   └── [Name].go         #     特定中间件实现
├── configs/              # ⚙️ 配置文件与环境管理
│   └── [Name].go         #     不同环境或服务的配置
├── docker/               # 🐳 容器化与部署相关
└── scripts/              # 🛠️ 构建、测试与运维脚本
```