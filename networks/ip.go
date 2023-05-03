/**
  @author: Zero
  @date: 2023/5/3 11:01:08
  @desc:

**/

package networks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zlx2019/toys/converts"
	"github.com/zlx2019/toys/valida"
	"io"
	"net"
	"net/http"
	"strings"
)

// GetLocalAddressName 获取本机IP所在地区名
func GetLocalAddressName() (string, error) {
	return GetIPAddressName(GetPublicIP())
}

// GetIPAddressName 根据公网IP,获取对应的地区名称
func GetIPAddressName(ip string) (string, error) {
	// 是否是有效的IP地址
	if !valida.IsIp(ip) {
		return "UNKNOWN", errors.New("invalid ip address")
	}
	// 是否是内网IP
	if IsInternalIP(net.IP(ip)) {
		return "0.0.0.0", nil
	}
	// 根据IP获取地区名称网址
	url := fmt.Sprintf("http://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true", ip)
	// 发送GET请求
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	// 读取响应数据
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// 这个网址默认使用BGK格式,中文会产生乱码
	// 将GBK格式转换为UTF-8
	bytes, err = converts.GBKToUTF8(bytes)
	if err != nil {
		return "", err
	}
	addressInfo := make(map[string]string)
	// 将响应映射为map
	err = json.Unmarshal(bytes, &addressInfo)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(addressInfo["addr"]), nil
}

// GetInternalIP 获取本机内网IP
func GetInternalIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err.Error())
	}
	for _, a := range addr {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

// GetPublicIP 获取本机公网IP
func GetPublicIP() string {
	response, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	ipInfo := make(map[string]any)
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return ""
	}
	return ipInfo["query"].(string)
}

// IsPublicIP 是否是公网IP
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// IsInternalIP 是否是内网IP
func IsInternalIP(IP net.IP) bool {
	if IP.IsLoopback() {
		return true
	}
	if ip4 := IP.To4(); ip4 != nil {
		return ip4[0] == 10 ||
			(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
			(ip4[0] == 169 && ip4[1] == 254) ||
			(ip4[0] == 192 && ip4[1] == 168)
	}
	return false
}
