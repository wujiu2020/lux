package youku

import (
	"fmt"

	"github.com/wujiu2020/lux/extractors/proto"
)

type errorData struct {
	Note string `json:"note"`
	Code int    `json:"code"`
}

type segs struct {
	TotalMilliSecondsVideo float64 `json:"total_milliseconds_video"`
	Size                   int64   `json:"size"`
	URL                    string  `json:"cdn_url"`
}

type stream struct {
	Size      int64  `json:"size"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Segs      []segs `json:"segs"`
	Type      string `json:"stream_type"`
	AudioLang string `json:"audio_lang"`
}

type youkuVideo struct {
	Title string `json:"title"`
}

type youkuShow struct {
	Title string `json:"title"`
}

type data struct {
	Error  errorData  `json:"error"`
	Stream []stream   `json:"stream"`
	Video  youkuVideo `json:"video"`
	Show   youkuShow  `json:"show"`
}

type youkuData struct {
	Data data `json:"data"`
}

const youkuReferer = "https://v.youku.com"

func getAudioLang(lang string) string {
	var youkuAudioLang = map[string]string{
		"guoyu": "国语",
		"ja":    "日语",
		"yue":   "粤语",
	}
	translate, ok := youkuAudioLang[lang]
	if !ok {
		return lang
	}
	return translate
}

func (v youkuData) TransformData(url, quality string) (*proto.Data, error) {
	streams := make([]proto.Stream, len(v.Data.Stream))
	for index, stream := range v.Data.Stream {
		var streamQuality string
		if stream.AudioLang == "default" {
			streamQuality = fmt.Sprintf(
				"%s %dx%d", stream.Type, stream.Width, stream.Height,
			)
		} else {
			streamQuality = fmt.Sprintf(
				"%s %dx%d %s", stream.Type, stream.Width, stream.Height,
				getAudioLang(stream.AudioLang),
			)
		}
		urls := make([]proto.Seg, len(stream.Segs))
		for index, data := range stream.Segs {
			urls[index] = proto.Seg{
				URL:      data.URL,
				Size:     data.Size,
				Duration: data.TotalMilliSecondsVideo / 1000,
			}
		}
		streams[index] = proto.Stream{
			Segs:    urls,
			Quality: streamQuality,
		}
	}
	return &proto.Data{
		Streams: streams,
	}, nil
}
