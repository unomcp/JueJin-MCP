package juejin

import (
	"context"

	"github.com/go-rod/rod"
)

const (
	CATEGORY_LIST = `//*[@id="juejin-web-editor"]/div[2]/div/header/div[2]/div[3]/div/div[2]/div[2]`
)

var CategoryMap = map[int]string{
	0: "后端",
	1: "前端",
	2: "Android",
	3: "iOS",
	4: "人工智能",
	5: "开发工具",
	6: "代码人生",
	7: "阅读",
}

func SelectorCategoryItem(page *rod.Page, _ctx context.Context, categoryIndex int) error {
	// 1. 获取分类列表容器
	categoryList := page.MustElementX(CATEGORY_LIST)

	// 2. 获取所有分类项
	items := categoryList.MustElements(".item")

	// 3. 检查索引是否越界
	if categoryIndex < 0 || categoryIndex >= len(items) {
		return nil // 或者返回错误，这里暂时忽略越界
	}

	// 4. 点击对应索引的分类
	items[categoryIndex].MustClick()

	return nil
}
