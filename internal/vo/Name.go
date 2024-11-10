package vo

type Name struct {
	value string
}

func NewName(v string) Name {
	return Name{value: v}
}

func (n Name) GetValue() string {
	return n.value
}
