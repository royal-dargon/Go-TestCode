package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	p := []string{"熊吉祥", "赵潇帆", "陈为婷", "段元元", "雷杰睿", "王美龄", "夏智佳", "李劲哲", "王静蕾", "杨卿怡", "钟鸣", "曾瑶", "刘鑫", "张迎媂", "冯旭冉"}
	n, _ := rand.Int(rand.Reader, big.NewInt(2))
	fmt.Println(p[n.Int64()])
}
	