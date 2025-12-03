package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/unomcp/JueJin-MCP/configs"
)

func New() *rod.Browser {
	bin, _ := launcher.LookPath()

	la := launcher.New().UserDataDir(".juejin-mcp")
	la.Bin(bin)
	la.Headless(configs.BrowserHeadless)
	la.NoSandbox(configs.BrowserNoSandbox)
	la.Set("User-Agent", configs.BrowserUserAgent)

	controlURL := la.MustLaunch()
	b := rod.New().
		ControlURL(controlURL).
		Trace(configs.BrowserTrace).
		MustConnect()

	return b
}
