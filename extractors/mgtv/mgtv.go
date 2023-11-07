package mgtv

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
	"github.com/wujiu2020/lux/utils"
)

func init() {
	extractors.Register("mgtv", New())
}

func encodeTk2(str string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(str))
	r1 := regexp.MustCompile(`/\+/g`)
	r2 := regexp.MustCompile(`///g`)
	r3 := regexp.MustCompile(`/=/g`)
	r1.ReplaceAllString(encodeString, "_")
	r2.ReplaceAllString(encodeString, "~")
	r3.ReplaceAllString(encodeString, "-")
	encodeString = utils.Reverse(encodeString)
	return encodeString
}

type extractor struct{}

// New returns a mgtv extractor.
func New() proto.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	html, err := request.Get(url, url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	vid := utils.MatchOneOf(
		url,
		`https?://www.mgtv.com/(?:b|l)/\d+/(\d+).html`,
		`https?://www.mgtv.com/hz/bdpz/\d+/(\d+).html`,
	)
	if vid == nil {
		vid = utils.MatchOneOf(html, `vid: (\d+),`)
	}
	if vid == nil || len(vid) < 2 {
		return nil, errors.WithStack(proto.ErrURLParseFailed)
	}

	// API extract from https://js.mgtv.com/imgotv-miniv6/global/page/play-tv.js
	// getSource and getPlayInfo function
	// Chrome Network JS panel
	headers := map[string]string{
		"Cookie": "PM_CHKID=1",
	}
	clit := fmt.Sprintf("clit=%d", time.Now().Unix()/1000)
	pm2DataString, err := request.Get(
		fmt.Sprintf(
			"https://pcweb.api.mgtv.com/player/video?video_id=%s&tk2=%s",
			vid[1],
			encodeTk2(fmt.Sprintf(
				"did=f11dee65-4e0d-4d25-bfce-719ad9dc991d|pno=1030|ver=5.5.1|%s", clit,
			)),
		),
		url,
		headers,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var pm2 mgtvPm2Data
	if err = json.Unmarshal([]byte(pm2DataString), &pm2); err != nil {
		return nil, errors.WithStack(err)
	}

	dataString, err := request.Get(
		fmt.Sprintf(
			"https://pcweb.api.mgtv.com/player/getSource?video_id=%s&tk2=%s&pm2=%s",
			vid[1], encodeTk2(clit), pm2.Data.Atc.Pm2,
		),
		url,
		headers,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var mgtvData mgtv
	if err = json.Unmarshal([]byte(dataString), &mgtvData); err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Println(dataString)
	return mgtvData, nil
}
