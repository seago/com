package com

import (
	"strconv"
	"strings"
)

func IpToInt64(ip string) (int64, error) {
	ips := strings.Split(ip, ".")
	ip1, err := strconv.ParseInt(ips[0], 10, 64)
	if err != nil {
		return 0, err
	}
	ip2, err := strconv.ParseInt(ips[1], 10, 64)
	if err != nil {
		return 0, err
	}
	ip3, err := strconv.ParseInt(ips[2], 10, 64)
	if err != nil {
		return 0, err
	}
	ip4, err := strconv.ParseInt(ips[3], 10, 64)
	if err != nil {
		return 0, err
	}
	return (ip1 << 24) + (ip2 << 16) + (ip3 << 8) + ip4, nil
}

func Int64ToIp(ip int64) string {
	s_ip := strconv.FormatInt(ip>>24, 10) + "."
	s_ip = s_ip + strconv.FormatInt((ip&0x00ffffff)>>16, 10) + "."
	s_ip = s_ip + strconv.FormatInt((ip&0x0000ffff)>>8, 10) + "."
	s_ip = s_ip + strconv.FormatInt(ip&0x000000ff, 10)
	return s_ip
}
