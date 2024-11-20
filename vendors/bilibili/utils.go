package bilibili

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/synctv-org/vendors/utils"
)

var (
	BVRegex   = regexp.MustCompile(`^(?:https://www\.bilibili\.com/video/)?((?:bv|bV|Bv|BV)\w+)(?:[/\?].*)?$`)
	ARegex    = regexp.MustCompile(`^(?:https://www\.bilibili\.com/video/)?(?:av|aV|Av|AV)(\d+)(?:[/\?].*)?$`)
	SSRegex   = regexp.MustCompile(`^(?:https://www\.bilibili\.com/bangumi/play/)?(?:ss|sS|Ss|SS)(\d+)(?:[/\?].*)?$`)
	EPRegex   = regexp.MustCompile(`^(?:https://www\.bilibili\.com/bangumi/play/)?(?:ep|eP|Ep|EP)(\d+)(?:[/\?].*)?$`)
	B23Regex  = regexp.MustCompile(`^(https://)?b23\.tv/(\w+)$`)
	LIVERegex = regexp.MustCompile(`^(?:https://live\.bilibili\.com/)?(\d+)(?:\?.*)?$`)
)

func Match(url string) (t string, id string, err error) {
	if B23Regex.MatchString(url) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
		if err != nil {
			return "", "", err
		}
		req.Header.Set("User-Agent", utils.UA)
		resp, err := utils.NoRedirectHttpClient().Do(req)
		if err != nil {
			return "", "", err
		}
		resp.Body.Close()
		return Match(resp.Header.Get("Location"))
	}
	if m := BVRegex.FindStringSubmatch(url); m != nil {
		return "bv", m[1], nil
	}
	if m := ARegex.FindStringSubmatch(url); m != nil {
		return "av", m[1], nil
	}
	if m := SSRegex.FindStringSubmatch(url); m != nil {
		return "ss", m[1], nil
	}
	if m := EPRegex.FindStringSubmatch(url); m != nil {
		return "ep", m[1], nil
	}
	if m := LIVERegex.FindStringSubmatch(url); m != nil {
		return "live", m[1], nil
	}
	return "", "", errors.New("match failed")
}
