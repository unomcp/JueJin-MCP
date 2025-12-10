package juejin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

var (
	LOGIN_URL    = "https://juejin.cn/"
	LOGIN_BUTTON = `//*[@id="juejin"]/div[1]/div/header/div/nav/ul/ul/li[2]/div/button`
	USERAVATAR   = `//*[@id="juejin"]/div[1]/div/header/div/nav/ul/ul/li[3]/div/div/img`
)

func Login(page *rod.Page, ctx context.Context) error {
	p := page.Context(ctx)
	p.MustNavigate(LOGIN_URL).MustWaitLoad()

	// 检查是否已经登录（登录按钮不存在说明已登录）
	if exist, _, _ := p.HasX(LOGIN_BUTTON); exist {
		// 登录按钮存在，需要点击登录
		p.MustElementX(LOGIN_BUTTON).MustClick()
		time.Sleep(10 * time.Second)
	}

	// 检查是否登录成功（通过用户头像判断）
	if exist, el, _ := p.HasX(USERAVATAR); exist {
		fmt.Println(el.Describe(0, true))
		return nil
	}

	return errors.New("login failed")
}
