package iqiyi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
	"github.com/wujiu2020/lux/utils"
)

func getMacID() string {
	var macID string
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n", "m", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	size := len(chars)
	for i := 0; i < 32; i++ {
		macID += chars[rand.Intn(size)]
	}
	return macID
}

func getVF(params string) string {
	var suffix string
	for j := 0; j < 8; j++ {
		for k := 0; k < 4; k++ {
			var v8 int
			v4 := 13 * (66*k + 27*j) % 35
			if v4 >= 10 {
				v8 = v4 + 88
			} else {
				v8 = v4 + 49
			}
			suffix += string(rune(v8)) // string(97) -> "a"
		}
	}
	params += suffix

	return utils.Md5(params)
}

func getVPS(tvid, vid, refer string) (*iqiyi, error) {
	t := time.Now().Unix() * 1000
	host := "http://cache.video.qiyi.com"
	params := fmt.Sprintf(
		"/vps?tvid=%s&vid=%s&v=0&qypid=%s_12&src=01012001010000000000&t=%d&k_tag=1&k_uid=%s&rs=1",
		tvid, vid, tvid, t, getMacID(),
	)
	vf := getVF(params)
	fmt.Println(vf)
	apiURL := fmt.Sprintf("%s%s&vf=%s", host, params, vf)
	info, err := request.Get(apiURL, refer, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data := new(iqiyi)
	if err := json.Unmarshal([]byte(info), data); err != nil {
		return nil, errors.WithStack(err)
	}
	return data, nil
}

type extractor struct {
	siteType SiteType
}

// New returns a iqiyi extractor.
func New(siteType SiteType) proto.Extractor {
	return &extractor{
		siteType: siteType,
	}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	refer := iqiyiReferer
	// headers := make(map[string]string)
	// if e.siteType == SiteTypeIQ {
	// 	headers = map[string]string{
	// 		"Accept-Language": "zh-TW",
	// 	}
	// 	refer = iqReferer
	// }
	// html, err := request.Get(url, refer, headers)
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }
	// tvid := utils.MatchOneOf(
	// 	url,
	// 	`#curid=(.+)_`,
	// 	`tvid=([^&]+)`,
	// )
	// if tvid == nil {
	// 	tvid = utils.MatchOneOf(
	// 		html,
	// 		`data-player-tvid="([^"]+)"`,
	// 		`param\['tvid'\]\s*=\s*"(.+?)"`,
	// 		`"tvid":"(\d+)"`,
	// 		`"tvId":(\d+)`,
	// 	)
	// }
	// if tvid == nil || len(tvid) < 2 {
	// 	return nil, errors.WithStack(proto.ErrURLParseFailed)
	// }

	// vid := utils.MatchOneOf(
	// 	url,
	// 	`#curid=.+_(.*)$`,
	// 	`vid=([^&]+)`,
	// )
	// if vid == nil {
	// 	vid = utils.MatchOneOf(
	// 		html,
	// 		`data-player-videoid="([^"]+)"`,
	// 		`param\['vid'\]\s*=\s*"(.+?)"`,
	// 		`"vid":"(\w+)"`,
	// 	)
	// }
	// if vid == nil || len(vid) < 2 {
	// 	return nil, errors.WithStack(proto.ErrURLParseFailed)
	// }
	// videoDatas, err := getVPS(tvid[1], vid[1], refer)
	videoDatas, err := getVPS("5808984258104400", "196d2372d6ceea1980843c558907764d", refer)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if videoDatas.Code != "A00000" {
		return nil, errors.New("can't play this video")
	}
	return *videoDatas, nil
}
