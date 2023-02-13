package src

import (
	"time"
)

type TOTP struct {
	Time     time.Time
	digits   int
	timeStep int64
}

func NewTOTP() *TOTP {
	return &TOTP{
		Time:     time.Now(),
		digits:   6,
		timeStep: 30,
	}
}

func (t *TOTP) Calc(secret *[]byte) string {
	hotp := NewHOTP(secret, t.digits)
	return hotp.Calc(t.counter())
}

func (t *TOTP) counter() uint64 {
	return uint64(t.Time.Unix() / t.timeStep)
}
