package iqiyi

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
)

type SiteType int

const (
	// SiteTypeIQ indicates the site is iq.com
	SiteTypeIQ SiteType = iota
	// SiteTypeIqiyi indicates the site is iqiyi.com
	SiteTypeIqiyi
	iqReferer    = "https://www.iq.com"
	iqiyiReferer = "https://www.iqiyi.com"
)

type iqiyi struct {
	Code string `json:"code"`
	Data struct {
		VP struct {
			Du  string `json:"du"`
			Tkl []struct {
				Vs []struct {
					Bid   int    `json:"bid"`
					Scrsz string `json:"scrsz"`
					Vsize int64  `json:"vsize"`
					Fs    []struct {
						L string `json:"l"`
						B int64  `json:"b"`
					} `json:"fs"`
					Duration float64 `json:"duration"`
				} `json:"vs"`
			} `json:"tkl"`
		} `json:"vp"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type iqiyiURL struct {
	L string `json:"l"`
}

func (v iqiyi) TransformData(url string, quality string) (*proto.Data, error) {
	streams := make([]proto.Stream, 0)
	urlPrefix := v.Data.VP.Du
	for _, video := range v.Data.VP.Tkl[0].Vs {
		urls := make([]proto.Seg, 0)
		for _, v := range video.Fs {
			realURLData, err := request.Get(urlPrefix+v.L, url, map[string]string{"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"})
			if err != nil {
				return nil, err
			}
			var realURL iqiyiURL
			if err = json.Unmarshal([]byte(realURLData), &realURL); err != nil {
				return nil, errors.WithStack(err)
			}
			urls = append(urls, proto.Seg{
				URL:      realURL.L,
				Size:     v.B,
				Duration: video.Duration,
			})
		}
		streams = append(streams, proto.Stream{
			Referer: url,
			Segs:    urls,
			Quality: video.Scrsz,
		})
	}

	return &proto.Data{
		Duration: 0,
		Streams:  streams,
		Title:    "",
	}, nil
}
