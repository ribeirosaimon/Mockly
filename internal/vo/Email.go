package vo

type Email struct {
	value string
}

func NewEmail(v string) Email {
	return Email{
		value: v,
	}
}
