package extractors

import (
	"net/url"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors/cctv"
	"github.com/wujiu2020/lux/extractors/iqiyi"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/extractors/qq"
	"github.com/wujiu2020/lux/utils"
)

func init() {
	Register("cctv", cctv.New())
	Register("iqiyi", iqiyi.New(iqiyi.SiteTypeIqiyi))
	Register("iq", iqiyi.New(iqiyi.SiteTypeIQ))
	Register("qq", qq.New())
}

var lock sync.RWMutex
var extractorMap = make(map[string]proto.Extractor)

// Register registers an Extractor.
func Register(domain string, e proto.Extractor) {
	lock.Lock()
	extractorMap[domain] = e
	lock.Unlock()
}

// Extract is the main function to extract the data.
func Extract(u, quality string) (*proto.Data, error) {
	u = strings.TrimSpace(u)
	var domain string

	uri, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}
	domain = utils.Domain(uri.Host)
	extractor := extractorMap[domain]
	if extractor == nil {
		return nil, errors.New("have not extractor")
	}
	videos, err := extractor.Extract(u)
	if err != nil {
		return nil, err
	}
	return videos.TransformData(u, quality)
}
