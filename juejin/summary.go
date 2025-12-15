package juejin

import (
	"context"

	"github.com/go-rod/rod"
)

var (
	SUMMARY_INPUT = `//*[@id="juejin-web-editor"]/div[2]/div/header/div[2]/div[3]/div/div[7]/div[2]/div/textarea`
)

func InputSummary(page *rod.Page, _ctx context.Context, summary string) error {
	summaryInput := page.MustElementX(SUMMARY_INPUT)
	summaryInput.MustInput(summary)
	return nil
}
