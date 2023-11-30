package youku

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	netURL "net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
	"github.com/wujiu2020/lux/utils"
)

// https://g.alicdn.com/player/ykplayer/0.5.61/youku-player.min.js
// {"0505":"interior","050F":"interior","0501":"interior","0502":"interior","0503":"interior","0510":"adshow","0512":"BDskin","0590":"BDskin"}

// var ccodes = []string{"0510", "0502", "0507", "0508", "0512", "0513", "0514", "0503", "0590"}

func youkuUps(vid string) (*youkuData, error) {
	var (
		url   string
		utid  string
		utids []string
		data  youkuData
	)
	cookie := "isI18n=false; cna=nAqpHaJwSWsCATwRD9TpnDDQ; __ysuid=1697123838217CG9; __ayft=1697123838217; __arycid=dv-3-00; __arcms=dv-3-00; __aysid=1700545456431H3F; __ayscnt=2; _m_h5_tk=9ca417e8d7a3a764b40285245d329c4d_1700550136926; _m_h5_tk_enc=8a78fcdaad0363effd199ff376fe82df; xlly_s=1; __ayvstp=30; __aysvstp=3; tfstk=dpJMRq2ZCwa5wEQMNVBsQo8AMehdBP6f4EeAktQqTw7QMsPOB-uD2wpOWhBageK27NUOQCIDJC9Lhhw9cMHB4pqAf5K1dpSdJEevXEE6OUKzBdh1H-X1htu-yYH8fh6f3EIjH0UACYSA47k-eht_XGoJYYnsvwi1UO3n1I3MpJ6droNn7838gwbybN-9vpJbyaxN--yVtGXZrWPPRIyfYmpUGS1NAMbJdp6MK; l=fBrdg41HPaEwBFLMKOfaourza77THIRXmuPzaNbMi9fP9aBB55cNW1eKK4t6CnGVEsB2535htC1BByYglyUg7xv9-eZbte9sndLnwpzHU; isg=BIKCaM-JUwErf0-p8fWfaker047kU4Ztylfc-8yb1PWgHyOZsODWfZ-bzxtje_4F; __arpvid=1700545465766N7YUZP-1700545465773; __aypstp=10; __ayspstp=2"
	ccode := []string{"0502"}
	youkuKey := "140#LdQojCgmzzWpnQo2+bJu4pN8s9xJB7C8vEOU+TtrPIy1p2yXFKEjpMtF13IrQPZqKQ62lp1zzq1QTXM3ozzxbxqK9ph/zzrb22U3lp1xz7fIDbiltFrz2PzvL6hqzFxwD9pyONdOHaU+PGDcH3N9Udbz5opCjH+Je4nZ3WGkqLJxoDHUeTikBemD72dD2b2y/Inkz5X6adLm3dCNpaV9cV/WAUuhArf/y2eWfkFvHPrY9x+B0zGoSzL7f4SeQDUBUcbXkJLHTC6IgqiK5ZAV0en2h6Ni3Bs9YS7b4JTYa0IXsv67P/Xjqat0igjoZtw8e+uIgSQvHMT3xsHAHgiNbjQiHQbGCfw8OQEheh6nyWGlgL5Qvvs1iJKAo5lLtX736oT2M1dvxShYDbE4KxK9IofJjM1lpM03WvUomgoVIsqOLJlGc7M+vB1pHXTn8FYX127089zVXnOuS9w0ab2TxNfboja9jeoO18Rhdvx6pVaEnf1ZSNC3XSuopbm10zDZ9+ch4KwR0eMj9OnmHQp1RkrMztOz8UeE8zUIrKJ/cktU7+QiET/9fLIcvSbKcZv98wBJHxLFBzHuVXTzV23cVaGlqm9USmCZO8TB941KN9f0yuSwl7UXUVq8JbSIAfOJdncuXmHbw3LqwDhDc4lFSMbgORt4XizUVloAZQ1Oi7y48pylfJsvdD9iuDClR+yzrQSuppOfSQ4i1QZJjc9si4PmvKiIrU5WyJpSy5oIa7tYCtHdS/pryhe83KJRPbmlosKsuH8Z3gKXL3Pu3WcnSfztz5v6NrA9fvVZdOONzJBi0bMbTYqB914pOmkw0/pbR7q69VcWmwW7QqoCrx9+zNkmo5ttlpsazQ=="
	if strings.Contains(cookie, "cna") {
		utids = utils.MatchOneOf(cookie, `cna=(.+?);`, `cna\s+(.+?)\s`, `cna\s+(.+?)$`)
	} else {
		headers, err := request.Headers("http://log.mmstat.com/eg.js", youkuReferer)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		setCookie := headers.Get("Set-Cookie")
		utids = utils.MatchOneOf(setCookie, `cna=(.+?);`)
	}
	if utids == nil || len(utids) < 2 {
		return nil, errors.WithStack(proto.ErrURLParseFailed)
	}
	utid = utids[1]

	// https://g.alicdn.com/player/ykplayer/0.5.61/youku-player.min.js
	// grep -oE '"[0-9a-zA-Z+/=]{256}"' youku-player.min.js
	for _, ccode := range ccode {
		if ccode == "0103010102" {
			utid = generateUtdid()
		}
		url = fmt.Sprintf(
			"https://ups.youku.com/ups/get.json?vid=%s&ccode=%s&client_ip=192.168.1.1&client_ts=%d&utid=%s&ckey=%s",
			vid, ccode, time.Now().Unix()/1000, netURL.QueryEscape(utid), netURL.QueryEscape(youkuKey),
		)
		html, err := request.GetByte(url, youkuReferer, nil)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// data must be emptied before reassignment, otherwise it will contain the previous value(the 'error' data)
		data = youkuData{}
		if err = json.Unmarshal(html, &data); err != nil {
			return nil, errors.WithStack(err)
		}
		if data.Data.Error == (errorData{}) {
			return &data, nil
		}
	}
	return &data, nil
}

func getBytes(val int32) []byte {
	var buff bytes.Buffer
	binary.Write(&buff, binary.BigEndian, val) // nolint
	return buff.Bytes()
}

func hashCode(s string) int32 {
	var result int32
	for _, c := range s {
		result = result*0x1f + c
	}
	return result
}

func hmacSha1(key []byte, msg []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(msg) // nolint
	return mac.Sum(nil)
}

func generateUtdid() string {
	timestamp := int32(time.Now().Unix())
	var buffer bytes.Buffer
	buffer.Write(getBytes(timestamp - 60*60*8))
	buffer.Write(getBytes(rand.Int31()))
	buffer.WriteByte(0x03)
	buffer.WriteByte(0x00)
	imei := fmt.Sprintf("%d", rand.Int31())
	buffer.Write(getBytes(hashCode(imei)))
	data := hmacSha1([]byte("d6fc3a4a06adbde89223bvefedc24fecde188aaa9161"), buffer.Bytes())
	buffer.Write(getBytes(hashCode(base64.StdEncoding.EncodeToString(data))))
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

type extractor struct{}

// New returns a youku extractor.
func New() proto.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	vids := utils.MatchOneOf(
		url, `id_(.+?)\.html`, `id_(.+)`,
	)
	if vids == nil || len(vids) < 2 {
		return nil, errors.WithStack(proto.ErrURLParseFailed)
	}
	vid := vids[1]

	youkuData, err := youkuUps(vid)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if youkuData.Data.Error.Code != 0 {
		return nil, errors.New(youkuData.Data.Error.Note)
	}
	return youkuData, nil
}
