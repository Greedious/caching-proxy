package utils

import (
	"caching-proxy/config"
	"net"
	"strconv"
)

func IsPortTaken(port string) bool {
	addr, err := net.Listen(config.DefaultServerProtocol, ":"+port)
	if err != nil {
		return true
	}
	defer addr.Close()
	return false
}

func IsInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
