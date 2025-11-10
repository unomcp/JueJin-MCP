package juejin

import (
	"github.com/go-rod/rod"
)

func (j *JueJin) Login() error {
	page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
	page.MustWaitStable().MustScreenshot("a.png")
	// page := rod.New().MustPages().First().MustNavigate("https://juejin.cn/")
	// page.MustWaitLoad()

	// fmt.Println(page.MustInfo().URL)
	return nil
}
