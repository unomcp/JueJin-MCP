# MCP for JueJin

<div align="center">
  <img src="./assets/mcp.svg" alt="MCP" height="100" />
  <img src="./assets/plus.svg" alt="+" height="100" />
  <img src="./assets/juejin.svg" alt="JueJin Logo" height="100" />
</div>

<br/>

<p align="center">
  <strong>JueJin-MCP 是一个能让AI生成的文章自动发布到掘金的工具</strong>
</p>

项目愿景是解放你的双手，让ai帮你的文章质量兜底，让每个人都能有一个掘金Lv8的账号

> [!IMPORTANT]
> 项目目前还在开发中，尚未正式发布。你可以直接构建源代码进行已有功能的使用。

### 🌟 功能特性
- [x] 登录
- [x] 持久化登录
- [x] 文章写入到草稿箱
- [x] 发布
  - [x] 支持分类
  - [x] 支持标签
  - [ ] 支持文章封面
  - [ ] 支持专栏
  - [ ] 支持选择话题
  - [x] 支持摘要
- [ ] 支持文章样式 ｜ 代码块样式


### 🛠️ MCP配置

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

### 🤖 如何使用

第一步启动 `mcp/inspector`。(为了让用户无感发布，所以使用的无头浏览器。代表需要进行手动登录）
```bash
npx @modelcontextprotocol/inspector
```
![这是图片](/assets/inspector.png)

第二步启动mcp
```bash
git clone https://github.com/unomcp/JueJin-MCP

cd JueJin-MCP

go run .
```

第三部使用，在你写完提示词后，输入 `使用 juejin-mcp 将文章发布到掘金` 即可！

### ⭐ Star 支持一下！

如果这个工具帮到了你，请点一个 ⭐
对我很重要 🙏

### 📄 协议

MIT License
