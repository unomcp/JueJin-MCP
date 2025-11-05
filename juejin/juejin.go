package juejin

type JueJin struct{}

func NewJueJin() *JueJin {
	return &JueJin{}
}

func (j *JueJin) GetJueJin() string {
	return "JueJin"
}
