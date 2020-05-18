package main

import (
	"flag"
	"fmt"
)

// 注意
// 1. p g 必须为质数
// 2. e 必须是小于 φ(N) 而与 φ(N) 互质的自然数，否则无法找到对应的模反元素
// 3. 加密信息必须小于 N，否则解密后无法得到原文
var (
	p       = flag.Int64("p", 13, "p")
	q       = flag.Int64("q", 19, "q")
	e       = flag.Int64("e", 11, "e")
	message = flag.Int64("message", 12, "message")
)

type rsaKey struct {
	n int64
	s int64
}

func main() {
	flag.Parse()

	// 选取互不相等的两个质数 p q 相乘得出 N
	N := *p * *q
	// 计算 N 的欧拉函数 φ(N)
	FaiN := (*p - 1) * (*q - 1)

	publicKey := &rsaKey{n: N, s: *e}
	fmt.Println("公钥:", publicKey)

	// 找出一个模反元素 d, ed ≡ 1 (mod φ(N))
	d := findD(FaiN, *e)

	privateKey := &rsaKey{n: N, s: d}
	fmt.Println("私钥:", privateKey)

	// m^e mod N = c
	c := encrypt(*message, publicKey)
	fmt.Println("密文:", c)

	// c^d mod N = m
	m := decrypt(c, privateKey)
	fmt.Println("明文:", m)

}

// 计算模反元素
func findD(FaiN int64, e int64) int64 {
	var d int64
	d = 1
	for {
		if e*d%FaiN == 1 {
			return d
		}
		d++
	}
}

// 加密
func encrypt(data int64, k *rsaKey) int64 {
	var i int64
	var c int64
	c = 1
	for i = 0; i < k.s; i++ {
		c = c * data % k.n
	}
	return c
}

// 解密
func decrypt(data int64, k *rsaKey) int64 {
	var i int64
	var c int64
	c = 1
	for i = 0; i < k.s; i++ {
		c = c * data % k.n
	}
	return c
}
