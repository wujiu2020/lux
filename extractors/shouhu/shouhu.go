package shouhu

import (
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
)

const videoInfoUrl = "https://hot.vrs.sohu.com/vrs_pc_play.action?ver=1&uid=16971209284615788536&ssl=1&pflag=pch5&prod=h5n"

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
	var videoInfo VideoInfo
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Ma cintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		"Origin":     "https://tv.sohu.com",
	}
	html, err := request.Get(url, url, headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	re := regexp.MustCompile(`var vid="(.*?)"`)
	match := re.FindStringSubmatch(html)
	if len(match) < 2 {
		return nil, errors.New("have no mathc url")
	}
	videoInfoBytes, err := request.GetByte(videoInfoUrl+"&vid="+match[1], "", headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := json.Unmarshal(videoInfoBytes, &videoInfo); err != nil {
		return nil, errors.WithStack(err)
	}
	return videoInfo, nil
}
