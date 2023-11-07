package mgtv

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
	"github.com/wujiu2020/lux/utils"
)

func mgtvM3u8(url string) ([]mgtvURLInfo, int64, error) {
	var data []mgtvURLInfo
	var temp mgtvURLInfo
	var size, totalSize int64
	urls, err := utils.M3u8URLs(url)
	if err != nil {
		return nil, 0, err
	}
	m3u8String, err := request.Get(url, url, nil)
	if err != nil {
		return nil, 0, err
	}
	sizes := utils.MatchAll(m3u8String, `#EXT-MGTV-File-SIZE:(\d+)`)
	// sizes: [[#EXT-MGTV-File-SIZE:1893724, 1893724]]
	for index, u := range urls {
		size, err = strconv.ParseInt(sizes[index][1], 10, 64)
		if err != nil {
			return nil, 0, err
		}
		totalSize += size
		temp = mgtvURLInfo{
			URL:  u,
			Size: size,
		}
		data = append(data, temp)
	}
	return data, totalSize, nil
}

type mgtvVideoStream struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Def  string `json:"def"`
}

type mgtvVideoInfo struct {
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Duration string `json:"duration"`
}

type mgtvVideoData struct {
	Stream       []mgtvVideoStream `json:"stream"`
	StreamDomain []string          `json:"stream_domain"`
	Info         mgtvVideoInfo     `json:"info"`
}

type mgtv struct {
	Data mgtvVideoData `json:"data"`
}

type mgtvVideoAddr struct {
	Info string `json:"info"`
}

type mgtvURLInfo struct {
	URL  string
	Size int64
}

type mgtvPm2Data struct {
	Data struct {
		Atc struct {
			Pm2 string `json:"pm2"`
		} `json:"atc"`
		Info mgtvVideoInfo `json:"info"`
	} `json:"data"`
}

func (v mgtv) TransformData(url string, quality string) (*proto.Data, error) {
	headers := map[string]string{
		"Cookie": "PM_CHKID=1",
	}
	title := strings.TrimSpace(
		v.Data.Info.Title + " " + v.Data.Info.Desc,
	)
	mgtvStreams := v.Data.Stream
	var addr mgtvVideoAddr
	streams := make([]proto.Stream, len(mgtvStreams))
	for _, stream := range mgtvStreams {
		if stream.URL == "" {
			continue
		}
		// real download address
		addr = mgtvVideoAddr{}
		addrInfo, err := request.GetByte(v.Data.StreamDomain[0]+stream.URL, url, headers)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if err = json.Unmarshal(addrInfo, &addr); err != nil {
			return nil, errors.WithStack(err)
		}

		m3u8URLs, _, err := mgtvM3u8(addr.Info)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		urls := make([]proto.Seg, len(m3u8URLs))
		for _, u := range m3u8URLs {
			urls = append(urls, proto.Seg{
				Size: u.Size,
				URL:  u.URL,
			})
		}
		streams = append(streams, proto.Stream{
			Segs: urls,
		})
	}

	return &proto.Data{
		Title:    title,
		Duration: 10,
		Streams:  streams,
	}, nil
}
