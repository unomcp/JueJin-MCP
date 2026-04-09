package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/unomcp/JueJin-MCP/configs"
)

func New(headless ...bool) *rod.Browser {
	h := configs.BrowserHeadless
	if len(headless) > 0 {
		h = headless[0]
	}

	bin, _ := launcher.LookPath()

	la := launcher.New().UserDataDir(".juejin-mcp")
	la.Bin(bin)
	la.Headless(h)
	la.NoSandbox(configs.BrowserNoSandbox)
	la.Set("User-Agent", configs.BrowserUserAgent)

	controlURL := la.MustLaunch()
	b := rod.New().
		ControlURL(controlURL).
		Trace(configs.BrowserTrace).
		MustConnect()

	return b
}
