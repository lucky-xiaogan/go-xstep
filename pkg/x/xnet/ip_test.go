package xnet

import "testing"

func TestIpToInt(t *testing.T) {
	ip := "192.168.0.1"
	i, err := IpToInt(ip)
	t.Log(i, err)
}

func TestIntToIp(t *testing.T) {
	num := 3232235521
	i, err := IntToIp(num)
	t.Log(i, err)
}
