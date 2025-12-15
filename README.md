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
- [x] 持久化登录
- [x] 文章写入到草稿箱
- [ ] 发布
  - [x] 支持分类
  - [ ] 支持标签
  - [ ] 支持文章封面
  - [ ] 支持专栏
  - [ ] 支持选择话题
  - [x] 支持摘要
- [ ] 支持文章样式 ｜ 代码块样式
- [ ] 智能评论回复与点赞
- [ ] 文章仿写


### 🛠️ 使用示例

目前需要自己启动 `mcp inspector` 来进行手动登录，再运行 `juejin-mcp` 服务。

```bash
# 启动 mcp inspector
npx @modelcontextprotocol/inspector
# 下一步
go run .
```

完成以上操作后，再使用对应的AI工具来做文章生成，通过提示词描述 `使用juejin-mcp将文章发布到掘金` 完成自动化。

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

### ⭐ Star 支持一下！

如果这个工具帮到了你，请点一个 ⭐
对我很重要 🙏

### 📄 协议

MIT License