package utils

import (
	"caching-proxy/config"
	"net"
	"strconv"
)

func IsPortTaken(port string) bool {
	addr, err := net.Listen(config.DefaultServerProtocol, ":"+port)
	if err != nil {
		return true // Port is taken
	}
	defer addr.Close() // Now it's safe, because Listen() succeeded
	return false       // Port is free
}

func IsInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
