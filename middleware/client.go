package middleware

import (
	"strings"
)



func DetectClient(userAgent string) string {


	ua := strings.ToLower(userAgent)



	switch {


	case strings.Contains(ua, "shadowrocket"):

		return "Shadowrocket"



	case strings.Contains(ua, "clash"):

		return "Clash"



	case strings.Contains(ua, "sing-box"):

		return "sing-box"



	case strings.Contains(ua, "surge"):

		return "Surge"



	case strings.Contains(ua, "v2rayn"):

		return "v2rayN"



	case strings.Contains(ua, "nekobox"):

		return "NekoBox"



	default:

		return "Unknown"

	}

}
