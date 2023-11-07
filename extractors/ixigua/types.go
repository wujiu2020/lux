package ixigua

import (
	"github.com/wujiu2020/lux/extractors/proto"
)

type Video struct {
	Title     string `json:"title"`
	Qualities []struct {
		Quality string `json:"quality"`
		Size    int64  `json:"size"`
		URL     string `json:"url"`
		Ext     string `json:"ext"`
	} `json:"qualities"`
}

func (v Video) TransformData(url string, quality string) (*proto.Data, error) {
	streams := make([]proto.Stream, 0)
	for _, quality := range v.Qualities {
		streams = append(streams, proto.Stream{
			Quality: quality.Quality,
			Segs: []proto.Seg{
				{
					URL:  base64Decode(quality.URL),
					Size: quality.Size,
				},
			},
		})
	}
	return &proto.Data{
		Duration: 0,
		Streams:  streams,
		Title:    "",
	}, nil
}
