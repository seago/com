package com

import "testing"

func TestIpToInt64(t *testing.T) {
	_, err := IpToInt64("220.220.220.220")
	if err != nil {
		t.Error(err)
	}
}

func TestInt64ToIp(t *testing.T) {
	ip := Int64ToIp(3705461980)
	if ip != "220.220.220.220" {
		t.Error("int64 to ip error")
	}
}
