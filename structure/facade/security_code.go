package facade

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("code is error")
	}
	fmt.Println("code verified")
	return nil
}
