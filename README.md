# MCP for JueJin

<div align="center">
  <img src="./assets/mcp.svg" alt="MCP" height="100" />
  <img src="./assets/plus.svg" alt="+" height="100" />
  <img src="./assets/juejin.svg" alt="JueJin Logo" height="100" />
</div>

<br/>

<p align="center">
  <strong>JueJin-MCP是一个能让AI生成的文章自动发布到掘金的工具</strong>
</p>

> [!IMPORTANT]
> 项目目前还在开发中，尚未正式发布。你可以直接构建源代码进行已有功能的使用。

### 🌟 功能特性
- [x] 登录
- [x] 文章写入到草稿箱
- [ ] 发布
  - [ ] 支持分类
  - [ ] 支持标签
  - [ ] 支持文章封面
  - [ ] 支持专栏
  - [ ] 支持选择话题
  - [ ] 支持摘要
- [ ] 智能评论回复与点赞
- [ ] 文章仿写


### 🛠️ 使用示例

<details>
<summary><b>在 Cursor 使用</b></summary>

Cursor 配置: `Settings` -> `Cursor Settings` -> `MCP` -> `Add new global MCP server`

将以下配置粘贴到您的 Cursor `~/.cursor/mcp.json` 文件中是推荐的方法。您也可以通过在项目文件夹中创建来在特定项目中安装。. 查看 [Cursor MCP docs](https://docs.cursor.com/context/model-context-protocol) 获取更多.

> 自 Cursor 1.0 以来，您可以点击下面的安装按钮进行一键快速安装。

#### Cursor 本地服务连接

[![下载 MCP Server](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en-US/install-mcp?name=juejin-mcp&config=eyJ0eXBlIjoic3NlIiwidXJsIjoiaHR0cDovL2xvY2FsaG9zdDoxMDA4Ni9tY3AifQ%3D%3D)

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

<details>
<summary><b>在 Trae 使用</b></summary>

使用“手动添加”功能，并填写该 MCP 服务器的 JSON 配置信息。更多详情，请访问 Trae 文档 。

#### Trae 本地服务连接

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


### 🌲 项目结构

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

### ⭐ Star 支持一下！

如果这个工具帮到了你，请点一个 ⭐
对我很重要 🙏

### 📄 协议

MIT License