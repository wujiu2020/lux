package ixigua

import (
	"encoding/base64"

	"github.com/wujiu2020/lux/extractors/proto"
)

type VideoInfo struct {
	Status         int     `json:"status"`
	Message        string  `json:"message"`
	EnableSsl      bool    `json:"enable_ssl"`
	AutoDefinition string  `json:"auto_definition"`
	EnableAdaptive bool    `json:"enable_adaptive"`
	VideoID        string  `json:"video_id"`
	VideoDuration  float64 `json:"video_duration"`
	MediaType      string  `json:"media_type"`
	BigThumbs      []struct {
		ImgNum   int      `json:"img_num"`
		URI      string   `json:"uri"`
		ImgURL   string   `json:"img_url"`
		ImgUrls  []string `json:"img_urls"`
		ImgXSize int      `json:"img_x_size"`
		ImgYSize int      `json:"img_y_size"`
		ImgXLen  int      `json:"img_x_len"`
		ImgYLen  int      `json:"img_y_len"`
		Duration float64  `json:"duration"`
		Interval int      `json:"interval"`
		Fext     string   `json:"fext"`
	} `json:"big_thumbs"`
	DynamicVideo struct {
		DynamicType      string `json:"dynamic_type"`
		DynamicVideoList []struct {
			RealBitrate int    `json:"real_bitrate"`
			AvgBitrate  int    `json:"avg_bitrate"`
			Fps         int    `json:"fps"`
			QualityType int    `json:"quality_type"`
			Definition  string `json:"definition"`
			Quality     string `json:"quality"`
			Vtype       string `json:"vtype"`
			Vwidth      int    `json:"vwidth"`
			Vheight     int    `json:"vheight"`
			Bitrate     int    `json:"bitrate"`
			Size        int    `json:"size"`
			CodecType   string `json:"codec_type"`
			FileHash    string `json:"file_hash"`
			FileID      string `json:"file_id"`
			MainURL     string `json:"main_url"`
			BackupURL1  string `json:"backup_url_1"`
			URLExpire   int    `json:"url_expire"`
			InitRange   string `json:"init_range"`
			IndexRange  string `json:"index_range"`
			CheckInfo   string `json:"check_info"`
		} `json:"dynamic_video_list"`
		DynamicAudioList []struct {
			RealBitrate int    `json:"real_bitrate"`
			QualityType int    `json:"quality_type"`
			Quality     string `json:"quality"`
			Vtype       string `json:"vtype"`
			Bitrate     int    `json:"bitrate"`
			CodecType   string `json:"codec_type"`
			FileHash    string `json:"file_hash"`
			MainURL     string `json:"main_url"`
			BackupURL1  string `json:"backup_url_1"`
			URLExpire   int    `json:"url_expire"`
			InitRange   string `json:"init_range"`
			IndexRange  string `json:"index_range"`
			CheckInfo   string `json:"check_info"`
		} `json:"dynamic_audio_list"`
		MainURL    string `json:"main_url"`
		BackupURL1 string `json:"backup_url_1"`
	} `json:"dynamic_video"`
	Volume struct {
		Loudness                 float64 `json:"loudness"`
		Peak                     float64 `json:"peak"`
		MaximumMomentaryLoudness int     `json:"maximum_momentary_loudness"`
		MaximumShortTermLoudness int     `json:"maximum_short_term_loudness"`
		LoudnessRangeStart       int     `json:"loudness_range_start"`
		LoudnessRangeEnd         int     `json:"loudness_range_end"`
		LoudnessRange            int     `json:"loudness_range"`
		Version                  int     `json:"version"`
		VolumeInfoJSON           string  `json:"volume_info_json"`
	} `json:"volume"`
	PopularityLevel     int    `json:"popularity_level"`
	HasEmbeddedSubtitle bool   `json:"has_embedded_subtitle"`
	PosterURL           string `json:"poster_url"`
	ExtraInfos          struct {
		Status                   string `json:"Status"`
		Message                  string `json:"Message"`
		LogoType                 string `json:"LogoType"`
		VideoModelVersion        int    `json:"VideoModelVersion"`
		HelpInfoURL              string `json:"HelpInfoURL"`
		LengthOfVideoList        string `json:"LengthOfVideoList"`
		IsDynamicVideo           bool   `json:"IsDynamicVideo"`
		UserAction               string `json:"UserAction"`
		AccountName              string `json:"AccountName"`
		DeniedVideoModelV1JSON   string `json:"DeniedVideoModelV1JSON"`
		ResTag                   string `json:"ResTag"`
		EncodeUserTag            string `json:"EncodeUserTag"`
		OBSOLETEVideoMeta        string `json:"OBSOLETE_VideoMeta"`
		OBSOLETERemovedVideoMeta string `json:"OBSOLETE_RemovedVideoMeta"`
		VideoMetaList            []struct {
			Definition string `json:"Definition"`
			Fps        string `json:"FPS"`
			FileID     string `json:"FileID"`
		} `json:"VideoMetaList"`
		RemovedVideoMetaList []any `json:"RemovedVideoMetaList"`
		PrivateURL           bool  `json:"PrivateURL"`
	} `json:"extraInfos"`
	RefreshToken  string `json:"refreshToken"`
	InterfaceInfo struct {
		Code       int    `json:"code"`
		Message    string `json:"message"`
		Logid      string `json:"logid"`
		APIStr     string `json:"api_str"`
		Timestamep int64  `json:"timestamep"`
	} `json:"interfaceInfo"`
}

func (v VideoInfo) TransformData(url string, quality string) (*proto.Data, error) {
	if quality == "" {
		quality = "360p"
	}
	stream := proto.Stream{
		Quality: quality,
	}
	for _, item := range v.DynamicVideo.DynamicVideoList {
		if item.Definition == quality {
			sDec1, _ := base64.StdEncoding.DecodeString(item.MainURL)
			sDec2, _ := base64.StdEncoding.DecodeString(item.BackupURL1)
			stream.Segs = append(stream.Segs, proto.Seg{
				Duration:  v.VideoDuration,
				URL:       string(sDec1),
				BackupURL: string(sDec2),
			})
		}
	}
	return &proto.Data{
		Duration: v.VideoDuration,
		Streams: []proto.Stream{
			stream,
		},
		Title: "",
	}, nil
}
