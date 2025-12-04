package juejin

import (
	"context"
	"testing"

	"github.com/unomcp/JueJin-MCP/browser"
)

func TestPublish(t *testing.T) {
	b := browser.New()
	defer b.Close()

	p := b.MustPage()
	defer p.Close()

	if err := Publish(p, context.Background(), PublishContent{
		Title:   "Test Title",
		Content: "## Test Content",
	}); err != nil {
		t.Fatalf("publish failed: %v", err)
	}
}
