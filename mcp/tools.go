package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/browser"
	"github.com/unomcp/JueJin-MCP/juejin"
)

// 发布掘金文章
type PublishContentArgs struct {
	Title         string `json:"title" jsonschema:"内容标题（掘金限制：最多20个中文字或英文单词）"`
	Content       string `json:"content" jsonschema:"正文内容 限制为（Markdown 格式）"`
	CategoryIndex int    `json:"categoryIndex" jsonschema:"文章分类索引 0-后端, 1-前端, 2-Android, 3-iOS, 4-人工智能, 5-开发工具, 6-代码人生, 7-阅读"`
	Summary       string `json:"summary" jsonschema:"文章摘要 限制为（100字文本格式）"`
}

// 登录掘金
func loginTool(ctx context.Context, _req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := juejin.Login(p, ctx); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func publishTool(ctx context.Context, _req *goMcp.CallToolRequest, args PublishContentArgs) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := juejin.Publish(p, ctx, juejin.PublishContent{
		Title:         args.Title,
		Content:       args.Content,
		CategoryIndex: args.CategoryIndex,
	}); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
