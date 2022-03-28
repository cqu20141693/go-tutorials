package __18

import (
	"fmt"
	"net/netip"
	"testing"
)

/*
可比较，不可变，可以作为map 键，内存使用少
1. 解析 api
2. 构造api
3. 序列化api
4. port 序列化
5. prefix 序列化，PrefixMasking
6. IP 属性方法
7. ip 比较方法
8. ipv4/v6 方法
*/

func TestIP(t *testing.T) {
	local := netip.AddrFrom4([4]byte{127, 0, 0, 1})
	local.Is6()
	local.Is4()
	local.IsGlobalUnicast()

	// prefix
	private := netip.MustParseAddr("10.22.96.1")
	private.IsPrivate()
	prefix, _ := private.Prefix(24)
	private1 := netip.AddrFrom4([4]byte{10, 22, 96, 10})
	contains(prefix, private1)
	private2 := netip.AddrFrom4([4]byte{10, 22, 97, 1})
	contains(prefix, private2)
}

func contains(prefix netip.Prefix, addr netip.Addr) {
	contains := prefix.Contains(addr)
	if contains {
		fmt.Printf("prefix=%s contains ip=%s \n", prefix, addr)
	} else {
		fmt.Printf("prefix=%s does not contain ip=%s \n", prefix, addr)
	}
}
