package kuaishou

import (
	"encoding/json"
	"regexp"

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

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	var videoInfo []VideoInfo
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Ma cintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	}
	headers["Cookie"] = "did=web_78dc0711d2b146819feaa030598b1350; didv=1701046998000; kpf=PC_WEB; clientid=3; kpn=KUAISHOU_VISION"
	html, err := request.Get(url, url, headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	re := regexp.MustCompile(`"adaptationSet":(.*?)\,"playInfo"`)
	match := re.FindStringSubmatch(html)
	if len(match) < 2 {
		return nil, errors.New("have no mathc url")
	}
	if err := json.Unmarshal([]byte(match[1]), &videoInfo); err != nil {
		return nil, errors.New("have no mathc format")
	}
	if len(videoInfo) < 1 {
		return nil, errors.New("have no mathc video")
	}
	return videoInfo[0], nil
}
