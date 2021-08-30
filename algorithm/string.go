package algorithm

import (
	"fmt"
)

func swap(s []byte, i, j int) {
	s[i], s[j] = s[j], s[i]
}

//一 字符串排列:输出N长度的字符串的全部排列,排列数:N!
func StringRange(str []byte, cur int) {
	if len(str)-1 == cur {
		fmt.Println(string(str))
	} else {
		for i := cur; i < len(str); i++ {
			swap(str, cur, i)
			StringRange(str, cur+1)
			swap(str, i, cur)
		}
	}
}

//二 字符串组合:输入N长度的字符串的所有任意组合,组合数2^N-1
