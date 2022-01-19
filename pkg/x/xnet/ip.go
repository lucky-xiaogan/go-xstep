package xnet

import (
	"strconv"
	"strings"
)

/*
将子网掩码和 IP 地址按位计算 AND，就可得到网络号。
*/

func IpToInt(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	l := len(ips)
	rs := 0
	for i := 0; i < l; i++ {
		s, err := strconv.Atoi(ips[i])
		if err != nil {
			return 0, err
		}
		s = s << (8 * (l - i - 1))
		rs |= s
	}
	return rs, nil
}

func IntToIp(ipInt int) (string, error) {
	var ips []string
	for i := 3; i >= 0; i-- {
		pos := 8 * i
		v := ipInt & (255 << pos)
		ips = append(ips, strconv.Itoa(v>>pos))
	}
	return strings.Join(ips, "."), nil
}
