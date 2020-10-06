// Package tools tools
// 参考: https://blog.csdn.net/w616589292/article/details/51078787
package tools

import "fmt"

// Print print key val
func Print(key string, val ...interface{}) {
	fmt.Printf("\n\n%c[%d;%d;%dm[key]%c[0m %c[%d;%d;%dm %s %c[0m \n",
		0x1B, 1, 40, 31, 0x1B,
		0x1B, 1, 40, 33, key, 0x1B,
	)
	fmt.Printf("%c[%d;%d;%dm[val]%c[0m %c[%d;%d;%dm %+v %c[0m \n\n",
		0x1B, 1, 40, 31, 0x1B,
		0x1B, 1, 40, 33, val, 0x1B,
	)
}
