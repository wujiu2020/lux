package ixigua

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
)

func init() {
	extractors.Register("ixigua", New())
	extractors.Register("toutiao", New())
}

type extractor struct{}

// New returns a ixigua extractor.
func New() proto.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Ma cintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	}
	var videoInfo VideoInfo
	html, err := request.Get(url+"?wid_try=1", "", headers)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`"dash_120fps":(.*?)\,"normal_6min"`)
	match := re.FindStringSubmatch(html)
	if len(match) < 2 {
		return nil, errors.New("have no mathc url")
	}
	if err := json.Unmarshal([]byte(strings.ReplaceAll(match[1], "undefined", "\"\"")), &videoInfo); err != nil {
		return nil, errors.New("have no match format")
	}
	return videoInfo, nil
}
