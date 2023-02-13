package src

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

type HOTP struct {
	secret []byte
	digits int
}

func NewHOTP(secret *[]byte, digits int) *HOTP {
	return &HOTP{
		secret: *secret,
		digits: digits,
	}
}

func (h *HOTP) Calc(counter uint64) string {
	hr := h.hmacResult(counter)
	bin := h.dynamicTruncate(hr)
	code := bin % uint64(math.Pow10(h.digits))

	return h.format(code)
}

func (h *HOTP) hmacResult(counter uint64) []byte {
	mac := hmac.New(sha1.New, h.secret)

	byteCounter := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCounter, counter)

	mac.Write(byteCounter)
	return mac.Sum(nil)
}

func (h *HOTP) dynamicTruncate(hr []byte) uint64 {
	offset := hr[len(hr)-1] & 0xf

	return uint64(hr[offset]&0x7f)<<24 |
		uint64(hr[offset+1]&0xff)<<16 |
		uint64(hr[offset+2]&0xff)<<8 |
		uint64(hr[offset+3]&0xff)
}

func (h *HOTP) format(code uint64) string {
	format := fmt.Sprintf("%%0%dd", h.digits)
	return fmt.Sprintf(format, code)
}
