package kuaishou

import (
	"github.com/wujiu2020/lux/extractors/proto"
)

type VideoInfo struct {
	ID             int     `json:"id"`
	Duration       float64 `json:"duration"`
	Representation []struct {
		ID         int      `json:"id"`
		URL        string   `json:"url"`
		BackupURL  []string `json:"backupUrl"`
		MaxBitrate int      `json:"maxBitrate"`
		AvgBitrate int      `json:"avgBitrate"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		FrameRate  int      `json:"frameRate"`
		Quality    float64  `json:"quality"`
		KvqScore   struct {
			Fr     float64 `json:"FR"`
			Nr     float64 `json:"NR"`
			FRPost int     `json:"FRPost"`
			NRPost int     `json:"NRPost"`
		} `json:"kvqScore"`
		QualityType     string `json:"qualityType"`
		QualityLabel    string `json:"qualityLabel"`
		FeatureP2Sp     bool   `json:"featureP2sp"`
		Hidden          bool   `json:"hidden"`
		DisableAdaptive bool   `json:"disableAdaptive"`
		DefaultSelect   bool   `json:"defaultSelect"`
		Comment         string `json:"comment"`
		HdrType         int    `json:"hdrType"`
		FileSize        int    `json:"fileSize"`
	} `json:"representation"`
}

func (v VideoInfo) TransformData(url string, quality string) (*proto.Data, error) {
	var segs []proto.Seg
	for _, item := range v.Representation {
		segs = append(segs, proto.Seg{
			Duration: v.Duration,
			URL:      item.URL,
		})
	}
	return &proto.Data{
		Duration: v.Duration,
		Streams: []proto.Stream{
			{
				Segs: segs,
			},
		},
		Title: "",
	}, nil
}
