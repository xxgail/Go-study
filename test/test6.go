package main

/**

通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
*/

import (
	"fmt"
	"strconv"
)

type IPAddr []byte

func (ip IPAddr) String() string {
	var res string
	for i, ch := range ip {
		res += strconv.Itoa(int(ch))
		if i < 3 {
			res += string(".")
		}
	}
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
