package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/browser"
	"github.com/unomcp/JueJin-MCP/juejin"
)

// 发布掘金文章
type PublishContentArgs struct {
	Title         string   `json:"title" jsonschema:"内容标题（掘金限制：最多20个中文字或英文单词）"`
	Content       string   `json:"content" jsonschema:"正文内容 限制为（Markdown 格式）"`
	Theme         string   `json:"theme" jsonschema:"Markdown 主题，可选值：juejin, github, smartblue, cyanosis, channing-cyan, fancy, hydrogen, condensed-night-purple, greenwillow, v-green, vue-pro, healer-readable, mk-cute, jzman, geek-black, awesome-green, qklhk-chocolate, orange, scrolls-light, simplicity-green, arknights, vuepress, Chinese-red, nico, devui-blue, serene-rose, z-blue, minimalism, koi, yu, lilsnake, keepnice"`
	Highlight     string   `json:"highlight" jsonschema:"代码块主题，可选值：a11y-dark, a11y-light, agate, an-old-hope, androidstudio, arduino-light, arta, ascetic, atom-one-dark, atom-one-light, darcula, dark, default, dracula, github, github-gist, googlecode, gruvbox-dark, gruvbox-light, hybrid, idea, monokai, monokai-sublime, night-owl, nord, obsidian, ocean, solarized-dark, solarized-light, vs, vs2015, xcode, zenburn 等"`
	CategoryIndex int      `json:"categoryIndex" jsonschema:"文章分类索引 0-后端, 1-前端, 2-Android, 3-iOS, 4-人工智能, 5-开发工具, 6-代码人生, 7-阅读"`
	Summary       string   `json:"summary" jsonschema:"文章摘要 限制为（100字文本格式）"`
	Tags          []string `json:"tags" jsonschema:"文章标签列表，只能使用允许的标签（参见 juejin.AllowedTags），最多3个标签"`
}

// 登录掘金
func loginTool(ctx context.Context, _req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New(false)
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := juejin.Login(p, ctx); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

// 草稿掘金文章
type DraftContentArgs struct {
	Title     string `json:"title" jsonschema:"内容标题（掘金限制：最多20个中文字或英文单词）"`
	Content   string `json:"content" jsonschema:"正文内容 限制为（Markdown 格式）"`
	Theme     string `json:"theme" jsonschema:"Markdown 主题，可选值：juejin, github, smartblue, cyanosis, channing-cyan, fancy, hydrogen, condensed-night-purple, greenwillow, v-green, vue-pro, healer-readable, mk-cute, jzman, geek-black, awesome-green, qklhk-chocolate, orange, scrolls-light, simplicity-green, arknights, vuepress, Chinese-red, nico, devui-blue, serene-rose, z-blue, minimalism, koi, yu, lilsnake, keepnice"`
	Highlight string `json:"highlight" jsonschema:"代码块主题，可选值：a11y-dark, a11y-light, agate, an-old-hope, androidstudio, arduino-light, arta, ascetic, atom-one-dark, atom-one-light, darcula, dark, default, dracula, github, github-gist, googlecode, gruvbox-dark, gruvbox-light, hybrid, idea, monokai, monokai-sublime, night-owl, nord, obsidian, ocean, solarized-dark, solarized-light, vs, vs2015, xcode, zenburn 等"`
}

func draftTool(ctx context.Context, _req *goMcp.CallToolRequest, args DraftContentArgs) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := juejin.Draft(p, ctx, juejin.PublishContent{
		Title:     args.Title,
		Content:   args.Content,
		Theme:     args.Theme,
		Highlight: args.Highlight,
	}); err != nil {
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
		Theme:         args.Theme,
		Highlight:     args.Highlight,
		CategoryIndex: args.CategoryIndex,
		Summary:       args.Summary,
		Tags:          args.Tags,
	}); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
