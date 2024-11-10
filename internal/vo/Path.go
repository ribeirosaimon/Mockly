package vo

type Path struct {
	value string
}

func NewPath(value string) Path {
	return Path{
		value: value,
	}
}

func (p *Path) GetValue() string {
	return p.value
}
