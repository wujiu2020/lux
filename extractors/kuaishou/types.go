package kuaishou

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
	"github.com/wujiu2020/lux/utils"
)

type kuaishou string

func (v kuaishou) TransformData(url string, quality string) (*proto.Data, error) {
	titles := utils.MatchOneOf(string(v), `<title>([^<]+)</title>`)
	if titles == nil || len(titles) < 2 {
		return nil, errors.New("can not found title")
	}

	title := regexp.MustCompile(`\n+`).ReplaceAllString(strings.TrimSpace(titles[1]), " ")

	qualityRegMap := map[string]*regexp.Regexp{
		"sd": regexp.MustCompile(`"photoUrl":\s*"([^"]+)"`),
	}

	streams := make([]proto.Stream, 0)
	for quality, qualityReg := range qualityRegMap {
		matcher := qualityReg.FindStringSubmatch(string(v))
		if len(matcher) != 2 {
			return nil, errors.WithStack(proto.ErrURLParseFailed)
		}

		u := strings.ReplaceAll(matcher[1], `\u002F`, "/")

		size, err := request.Size(u, url)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		urlData := proto.Seg{
			URL:  u,
			Size: size,
		}
		streams = append(streams, proto.Stream{
			Segs:    []proto.Seg{urlData},
			Quality: quality,
		})
	}

	return &proto.Data{
		Duration: 10,
		Streams:  streams,
		Title:    title,
	}, nil
}
