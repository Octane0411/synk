package controller

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func AddressesController(c *gin.Context) {
	addrs, _ := net.InterfaceAddrs() // 获取当前所有ip
	var result []string
	for _, address := range addrs {
		// check the address type and if it is not a loopback, display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"addresses": result})
}
