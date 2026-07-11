package middleware

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)



func GetClientIP(c *gin.Context) string {


	// Cloudflare

	if ip:=c.GetHeader("CF-Connecting-IP"); ip!=""{

		return ip

	}



	// Nginx proxy

	if ip:=c.GetHeader("X-Real-IP"); ip!=""{

		return ip

	}



	if ip:=c.GetHeader("X-Forwarded-For"); ip!=""{

		return strings.Split(ip,",")[0]

	}



	return c.ClientIP()

}



func IPVersion(ip string) string {


	parsed:=net.ParseIP(ip)


	if parsed==nil{

		return "unknown"

	}



	if parsed.To4()!=nil{

		return "IPv4"

	}



	return "IPv6"

}
