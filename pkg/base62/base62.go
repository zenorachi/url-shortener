package base62

import (
	"math/rand"
	"time"
)

const (
	base62Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	base62Base  = 62
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateCode(length int) string {
	shortCode := make([]byte, length)
	for i := 0; i < length; i++ {
		shortCode[i] = base62Chars[random.Intn(base62Base)]
	}
	return string(shortCode)
}
