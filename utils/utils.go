package utils

import "net/http"

const (
	UA = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.69`
)

var (
	noRedirectHttpClient = NoRedirectClient(&http.Client{})
)

func NoRedirectHttpClient() *http.Client {
	return noRedirectHttpClient
}

func NoRedirectClient(client *http.Client) *http.Client {
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return client
}

func MapToHttpCookies(m map[string]string) []*http.Cookie {
	var cookies = make([]*http.Cookie, 0, len(m))
	for k, v := range m {
		cookies = append(cookies, &http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	return cookies
}
