package fenghuang

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/wujiu2020/lux/extractors/proto"
	"github.com/wujiu2020/lux/request"
)

func New() proto.Extractor {
	return &extractor{}
}

type extractor struct{}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string) (proto.TransformData, error) {
	var vedioInfo VedioInfo
	html, err := request.Get(url, "", nil)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`"docData":(.*?)\,"parentColumn"`)
	match := re.FindStringSubmatch(html)
	if len(match) < 2 {
		return nil, errors.New("have no mathc url")
	}
	if err := json.Unmarshal([]byte(match[1]), &vedioInfo); err != nil {
		return nil, errors.New("have no match format")
	}
	return vedioInfo, nil
}
