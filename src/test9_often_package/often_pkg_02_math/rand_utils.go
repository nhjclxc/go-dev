package main

import (
	"math/rand"
	"sync"
	"time"
)

var allowAllChars []byte

// randPool 随机池
var randPool = sync.Pool{
	New: func() any {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

func init() {
	for i := 48; i <= 57; i++ {
		allowAllChars = append(allowAllChars, byte(i))
	}
	for i := 65; i <= 90; i++ {
		allowAllChars = append(allowAllChars, byte(i))
	}
	for i := 97; i <= 122; i++ {
		allowAllChars = append(allowAllChars, byte(i))
	}
}

// genRandPassword 随机生成密码，密码只允许 ASCII 在[48,57]、[65,90]和[97,122]之间
func genRandPassword(size int) string {
	if size <= 0 {
		return ""
	}

	pwds := make([]byte, 0, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	N := len(allowAllChars)
	for i := 0; i < size; i++ {
		pwds = append(pwds, allowAllChars[r.Intn(N)])
	}
	return string(pwds)
}

// genRandInt 生成[min,max]范围内的随机数
func genRandInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Intn(max-min+1) + min
}
