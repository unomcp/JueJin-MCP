package juejin

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func (j *JueJin) Login() error {
	bin, _ := launcher.LookPath()
	u := launcher.New().Bin(bin).Set("headless").Delete("--headless").MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://juejin.cn/")
	page.MustWaitStable().MustScreenshot("a.png")
	return nil
}
