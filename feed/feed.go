package feed

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/zorth/zorth-rss/model"
)

func UseVpnUrlToRSSFeedDMHY(urlStr string) (model.RSSFeedDMHY, error) {
	proxyStr := "http://127.0.0.1:10809" // 替换为你的 V2RayN 代理地址和端口
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}

	httpClient := http.Client{
		Timeout: 30 * time.Second, // 增加超时时间
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	// 添加 User-Agent 头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	var resp *http.Response
	for i := 0; i < 3; i++ { // 尝试重试3次
		resp, err = httpClient.Do(req)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second) // 等待2秒后重试
	}
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	var feed model.RSSFeedDMHY
	err = xml.Unmarshal(dat, &feed) // 修正 Unmarshal 调用
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	return feed, nil
}

func UrlToRSSFeedDMHY(urlStr string) (model.RSSFeedDMHY, error) {
	httpClient := http.Client{
		Timeout: 30 * time.Second, // 增加超时时间
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	// 添加 User-Agent 头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	var resp *http.Response
	for i := 0; i < 3; i++ { // 尝试重试3次
		resp, err = httpClient.Do(req)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second) // 等待2秒后重试
	}
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	var feed model.RSSFeedDMHY
	err = xml.Unmarshal(dat, &feed) // 修正 Unmarshal 调用
	if err != nil {
		return model.RSSFeedDMHY{}, err
	}
	return feed, nil
}
