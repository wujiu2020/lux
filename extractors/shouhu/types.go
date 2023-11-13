package shouhu

import (
	"encoding/json"

	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
)

type VideoInfo struct {
	Isp2P             int    `json:"isp2p"`
	Iswebp2P          int    `json:"iswebp2p"`
	Isai              int    `json:"isai"`
	ID                int    `json:"id"`
	Prot              int    `json:"prot"`
	Status            int    `json:"status"`
	Play              int    `json:"play"`
	TvPlayType        int    `json:"tvPlayType"`
	URL               string `json:"url"`
	Fee               int    `json:"fee"`
	Caid              int    `json:"caid"`
	Caname            string `json:"caname"`
	Pid               int    `json:"pid"`
	Age               string `json:"age"`
	Year              int    `json:"year"`
	AreaID            int    `json:"areaId"`
	IsMemberPlay      int    `json:"is_member_play"`
	TrySeeTime        int    `json:"try_see_time"`
	MainActorID       string `json:"mainActorId"`
	Act               string `json:"act"`
	TvApplicationTime string `json:"tv_application_time"`
	Catcode           string `json:"catcode"`
	Pvpic             struct {
		Big   string `json:"big"`
		Small string `json:"small"`
	} `json:"pvpic"`
	Tvid                    int    `json:"tvid"`
	Syst                    int64  `json:"syst"`
	Isdrm                   int    `json:"isdrm"`
	Dm                      int    `json:"dm"`
	Eye                     int    `json:"eye"`
	Early                   int    `json:"early"`
	Tvfee                   int    `json:"tvfee"`
	Vr                      int    `json:"vr"`
	LabelFirstIds           string `json:"labelFirstIds"`
	TvPromotePic            string `json:"tvPromotePic"`
	IsAIStarRealtime        int    `json:"isAIStarRealtime"`
	AiStarDetectResultExist string `json:"aiStarDetectResultExist"`
	SupportType             int    `json:"supportType"`
	Data                    struct {
		TvName           string    `json:"tvName"`
		SubName          string    `json:"subName"`
		Ch               string    `json:"ch"`
		IPLimit          int       `json:"ipLimit"`
		Width            int       `json:"width"`
		Version          int       `json:"version"`
		ClipsBytes       []int64   `json:"clipsBytes"`
		AudioCodec       string    `json:"audioCodec"`
		VideoCodec       string    `json:"videoCodec"`
		CoverImg         string    `json:"coverImg"`
		Height           int       `json:"height"`
		TotalDuration    float64   `json:"totalDuration"`
		TotalBytes       int64     `json:"totalBytes"`
		ClipsDuration    []float64 `json:"clipsDuration"`
		RelativeID       int       `json:"relativeId"`
		NorVid           int       `json:"norVid"`
		HighVid          int       `json:"highVid"`
		SuperVid         int       `json:"superVid"`
		OriVid           int       `json:"oriVid"`
		H2644KVid        int       `json:"h2644kVid"`
		H265NorVid       int       `json:"h265norVid"`
		H265HighVid      int       `json:"h265highVid"`
		H265SuperVid     int       `json:"h265superVid"`
		H265OriVid       int       `json:"h265oriVid"`
		H2654MVid        int       `json:"h2654mVid"`
		H2654KVid        int       `json:"h2654kVid"`
		NorVidNs         int       `json:"norVid_ns"`
		HighVidNs        int       `json:"highVid_ns"`
		SuperVidNs       int       `json:"superVid_ns"`
		OriVidNs         int       `json:"oriVid_ns"`
		P1080HdrVid      int       `json:"p1080HdrVid"`
		P1080Hdr265Vid   int       `json:"p1080Hdr265Vid"`
		P1080HdrVidNs    int       `json:"p1080HdrVid_ns"`
		P1080Hdr265VidNs int       `json:"p1080Hdr265Vid_ns"`
		TvVer35Vid       int       `json:"tvVer35_vid"`
		TvVer36Vid       int       `json:"tvVer36_vid"`
		TvVer34Vid       int       `json:"tvVer34_vid"`
		TvVer284Vid      int       `json:"tvVer284_vid"`
		TvVer285Vid      int       `json:"tvVer285_vid"`
		TvVer260Vid      int       `json:"tvVer260_vid"`
		TvVer262Vid      int       `json:"tvVer262_vid"`
		TvVer264Vid      int       `json:"tvVer264_vid"`
		TvVer266Vid      int       `json:"tvVer266_vid"`
		TvVer301Vid      int       `json:"tvVer301_vid"`
		TvVer302Vid      int       `json:"tvVer302_vid"`
		TvVer303Vid      int       `json:"tvVer303_vid"`
		TvVer304Vid      int       `json:"tvVer304_vid"`
		TvVer306Vid      int       `json:"tvVer306_vid"`
		TvVer307Vid      int       `json:"tvVer307_vid"`
		TvVer321Vid      int       `json:"tvVer321_vid"`
		TvVer322Vid      int       `json:"tvVer322_vid"`
		TvVer323Vid      int       `json:"tvVer323_vid"`
		TvVer324Vid      int       `json:"tvVer324_vid"`
		TvVer326Vid      int       `json:"tvVer326_vid"`
		TvVer327Vid      int       `json:"tvVer327_vid"`
		TotalBlocks      int       `json:"totalBlocks"`
		Hc               []string  `json:"hc"`
		Su               any       `json:"su"`
		EP               []struct {
			K   int    `json:"k"`
			V   string `json:"v"`
			URL string `json:"url"`
			Pt  int    `json:"pt"`
		} `json:"eP"`
		Adpo       any      `json:"adpo"`
		Mp4PlayURL []string `json:"mp4PlayUrl"`
		Num        int      `json:"num"`
		ST         int      `json:"sT"`
		ET         int      `json:"eT"`
	} `json:"data"`
	Fnor    int    `json:"fnor"`
	Keyword string `json:"keyword"`
	Company string `json:"company"`
	Crid    int    `json:"crid"`
	Plcatid int    `json:"plcatid"`
	Cmscat  string `json:"cmscat"`
	Hcap    int    `json:"hcap"`
	Vt      int    `json:"vt"`
	Isdl    int    `json:"isdl"`
	Systype int    `json:"systype"`
	Pl      int    `json:"pl"`
	Us      int    `json:"us"`
}

type VideoSegInfo struct {
	Servers []struct {
		Nid   int    `json:"nid"`
		Isp2P int    `json:"isp2p"`
		URL   string `json:"url"`
	} `json:"servers"`
}

func (v VideoInfo) TransformData(url string, quality string) (*proto.Data, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Ma cintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		"Origin":     "https://tv.sohu.com",
	}
	var segs []proto.Seg
	var videoSegInfo VideoSegInfo
	for i := 0; i < len(v.Data.Mp4PlayURL); i++ {
		videoSegInfoBytes, err := request.GetByte("http:"+v.Data.Mp4PlayURL[i], "", headers)
		if err != nil {
			continue
		}
		if err := json.Unmarshal(videoSegInfoBytes, &videoSegInfo); err != nil {
			continue
		}
		segUrl := ""
		segBackUrl := ""
		if len(videoSegInfo.Servers) == 1 {
			segUrl = videoSegInfo.Servers[0].URL
		}
		if len(videoSegInfo.Servers) == 2 {
			segUrl = videoSegInfo.Servers[0].URL
			segBackUrl = videoSegInfo.Servers[1].URL
		}
		segs = append(segs, proto.Seg{
			Duration:  v.Data.ClipsDuration[i],
			URL:       segUrl,
			BackupURL: segBackUrl,
			Size:      v.Data.ClipsBytes[i],
		})
	}
	return &proto.Data{
		Title:    v.Data.TvName,
		Duration: float64(v.Data.TotalDuration),
		Streams: []proto.Stream{{
			Segs:    segs,
			Quality: "default",
		}},
	}, nil
}
