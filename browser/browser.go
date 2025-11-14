package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/unomcp/JueJin-MCP/flags"
)

type Browser struct {
	B     *rod.Browser
	La    *launcher.Launcher
	Flags map[flags.Flag]any `json:"flags"`
}

func New() *Browser {
	bin, _ := launcher.LookPath()

	defaultOptions := map[flags.Flag]any{
		flags.Bin:        bin,
		flags.Headless:   true,
		flags.UserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		flags.ControlURL: "",
		flags.NoSandbox:  true,
		flags.Trace:      true,
	}
	la := launcher.New().
		Bin(defaultOptions[flags.Bin].(string)).
		Headless(defaultOptions[flags.Headless].(bool)).
		NoSandbox(defaultOptions[flags.NoSandbox].(bool)).
		Set("User-Agent", defaultOptions[flags.UserAgent].(string))

	controlURL := la.MustLaunch()
	defaultOptions[flags.ControlURL] = controlURL

	b := rod.New().
		ControlURL(defaultOptions[flags.ControlURL].(string)).
		Trace(defaultOptions[flags.Trace].(bool)).
		MustConnect()

	return &Browser{
		B:     b,
		La:    la,
		Flags: defaultOptions,
	}
}
