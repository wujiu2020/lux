package kuaishou

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
)

func init() {
	extractors.Register("kuaishou", New())
}

type extractor struct{}

// New returns a kuaishou extractor.
func New() proto.Extractor {
	return &extractor{}
}

// fetch url and get the cookie that write by server
func fetchCookies(url string, headers map[string]string) (string, error) {
	res, err := request.Request(http.MethodGet, url, nil, headers)
	if err != nil {
		return "", err
	}

	defer res.Body.Close() // nolint

	cookiesArr := make([]string, 0)
	cookies := res.Cookies()

	for _, c := range cookies {
		cookiesArr = append(cookiesArr, c.Name+"="+c.Value)
	}

	return strings.Join(cookiesArr, "; "), nil
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	}

	cookies, err := fetchCookies(url, headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	headers["Cookie"] = cookies

	html, err := request.Get(url, url, headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return kuaishou(html), nil
}
