

<h4 style="text-align:center;">JueJin MCP</h4>

### Directory Tree
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