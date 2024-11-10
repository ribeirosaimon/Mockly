package vo

import "golang.org/x/crypto/bcrypt"

type Password struct {
	value string
}

func NewPassword(value string) Password {
	return Password{
		value: value,
	}
}

func (p *Password) Encript() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.value = string(hashedPassword)
	return nil
}

func (p *Password) Check(v string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.value), []byte(v))
}
