package juejin

import (
	"context"
	"time"

	"github.com/go-rod/rod"
)

type PublishContent struct {
	Title         string
	Content       string
	CategoryIndex int
	Summary       string
}

var (
	PUBLISH_URL            = "https://juejin.cn/editor/drafts/new?v=2"
	TITLE_INPUT            = `//*[@id="juejin-web-editor"]/div[2]/div/header/input`
	CONTENT_INPUT          = `//*[@id="juejin-web-editor"]/div[2]/div/div/div/div[2]/div[1]/div`
	OPEN_PANEL_BUTTON      = `//*[@id="juejin-web-editor"]/div[2]/div/header/div[2]/div[3]/button`
	CONFIRM_PUBLISH_BUTTON = `//*[@id="juejin-web-editor"]/div[2]/div/header/div[2]/div[3]/div/div[8]/div/button[2]`
)

func Publish(page *rod.Page, ctx context.Context, content PublishContent) error {
	p := page.Context(ctx)
	p.MustNavigate(PUBLISH_URL).MustWaitLoad()

	writeArticle(p, ctx, content)
	PublishPanel(page, ctx)
	SelectorCategoryItem(page, ctx, content.CategoryIndex)
	InputSummary(page, ctx, content.Summary)

	time.Sleep(3 * time.Second)
	return nil
}

func writeArticle(page *rod.Page, _ context.Context, content PublishContent) {
	titleInput := page.MustElementX(TITLE_INPUT)
	titleInput.MustInput(content.Title)

	cm := page.MustElementX(CONTENT_INPUT)
	cm.Eval(`
		(...args) => {
			return this.CodeMirror.setValue(args[0])
		}
	`, content.Content)
}

func PublishPanel(page *rod.Page, _ context.Context) {
	openPanelBtn := page.MustElementX(OPEN_PANEL_BUTTON)
	openPanelBtn.MustClick()
}
