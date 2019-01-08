package utils

import (
	"fmt"
	utils2 "github.com/zcong1993/libgo/utils"
	"math/rand"
	"strings"
	"time"
)

const TOKEN_LEN = 32

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateToken() string {
	tk, err := utils2.GenerateRandomStringURLSafe(TOKEN_LEN)
	if err != nil {
		panic(err)
	}
	return tk
}

func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func RandomColor() string {
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)

	return getHex(r) + getHex(g) + getHex(b)
}

func GenerateAvatar(name string) string {
	return fmt.Sprintf("https://via.placeholder.com/125/%s/%s?text=%s", RandomColor(), RandomColor(), strings.Replace(name, " ", "%20", -1))
}
