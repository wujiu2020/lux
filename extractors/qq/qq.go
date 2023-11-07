package qq

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
	"github.com/wujiu2020/lux/utils"
)

type extractor struct{}

// New returns a qq extractor.
func New() proto.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	vids := utils.MatchOneOf(url, `vid=(\w+)`, `/(\w+)\.html`)
	if vids == nil || len(vids) < 2 {
		return nil, errors.WithStack(proto.ErrURLParseFailed)
	}
	vid := vids[1]

	if len(vid) != 11 {
		u, err := request.Get(url, url, nil)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vids = utils.MatchOneOf(
			u, `vid=(\w+)`, `vid:\s*["'](\w+)`, `vid\s*=\s*["']\s*(\w+)`,
		)
		if vids == nil || len(vids) < 2 {
			return nil, errors.WithStack(proto.ErrURLParseFailed)
		}
		vid = vids[1]
	}

	vinfo, err := getVinfo(vid, "shd", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// API request error
	if vinfo.Msg != "" {
		return nil, errors.New(vinfo.Msg)
	}
	vinfo.Vid = vid
	return vinfo, nil
}

func getVinfo(vid, defn, refer string) (qqVideoInfo, error) {
	html, err := request.Get(
		fmt.Sprintf(
			"http://vv.video.qq.com/getinfo?otype=json&platform=11&defnpayver=1&appver=%s&defn=%s&vid=%s",
			qqPlayerVersion, defn, vid,
		), refer, nil,
	)
	if err != nil {
		return qqVideoInfo{}, err
	}
	jsonStrings := utils.MatchOneOf(html, `QZOutputJson=(.+);$`)
	if jsonStrings == nil || len(jsonStrings) < 2 {
		return qqVideoInfo{}, errors.WithStack(proto.ErrURLParseFailed)
	}
	jsonString := jsonStrings[1]
	// fmt.Println(jsonString)
	var data qqVideoInfo
	if err = json.Unmarshal([]byte(jsonString), &data); err != nil {
		return qqVideoInfo{}, err
	}
	return data, nil
}
