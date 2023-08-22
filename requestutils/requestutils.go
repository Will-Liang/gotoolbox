package requestutils

import (
	"github.com/imroc/req/v3"
	"strings"
)

// Get请求
// 重试 retry_time 次
func Get(request *req.Request, url, proxy string, retry_time int) (*req.Response, error) {
	// *req.Request,*req.Response From github.com/imroc/req/v3
	check_proxy(request, proxy)
	var resp *req.Response
	var err error
	for i := 0; i < retry_time; i++ {
		resp, err = request.Get(url)
		if !resp.IsSuccessState() && err != nil && resp.Err != nil { // 有错误发生
			continue
		} else {
			break
		}
	}
	return resp, err
}

// 检查代理ip格式
func check_proxy(request *req.Request, proxy string) {
	if proxy != "" && (strings.HasPrefix(proxy, "http://") || strings.HasPrefix(proxy, "https://") || strings.HasPrefix(proxy, "socks")) {
		request.GetClient().SetProxyURL(proxy)
	}
}
