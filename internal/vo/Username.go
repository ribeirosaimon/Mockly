package vo

type Username struct {
	value string
}

func NewUserName(v string) Username {
	return Username{
		value: v,
	}
}
