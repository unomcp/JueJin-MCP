package flags

type Flag string

const (
	// Headless mode. Whether to run browser in headless mode. A mode without visible UI.
	Headless Flag = "headless"

	// User-Agent. The user-agent to use for the browser.
	UserAgent Flag = "User-Agent"

	// ControlURL. The URL to use for the browser's control interface.
	ControlURL Flag = "control-url"

	// Bin is the browser executable file path. If it's empty, launcher will automatically search or download the bin.
	Bin Flag = "juejin-bin"

	// NoSandbox Flag.
	NoSandbox Flag = "no-sandbox"

	// Trace Flag.
	Trace Flag = "trace"
)
