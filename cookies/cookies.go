package cookies

// CookieStore 定义 cookies 操作接口
type CookieStore interface {
	// Save 保存 cookies 到文件
	Save(cookies []byte) error
	// Load 从文件加载 cookies
	Load() ([]byte, error)
	// GetPath 获取 cookies 文件路径
	GetPath() string
	// Clear 清除 cookies 文件
	Clear() error
	// Delete 删除 cookies 文件
	Delete() error
}

// Cookies 实现 CookieStore 接口
type Cookies struct {
	path string
}

// NewCookies 创建一个新的 Cookies 实例
func NewCookies(path string) *Cookies {
	return &Cookies{
		path: path,
	}
}

// Save 保存 cookies 到文件
func (c *Cookies) Save(cookies []byte) error {
	// TODO: 实现保存逻辑
	return nil
}

// Load 从文件加载 cookies
func (c *Cookies) Load() ([]byte, error) {
	// TODO: 实现加载逻辑
	return nil, nil
}

// GetPath 获取 cookies 文件路径
func (c *Cookies) GetPath() string {
	return c.path
}

// Clear 清除 cookies 文件
func (c *Cookies) Clear() error {
	// TODO: 实现清除逻辑
	return nil
}
