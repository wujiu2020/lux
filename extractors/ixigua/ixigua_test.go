package ixigua

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wujiu2020/lux/extractors/proto"
)

func Test_extractor_Extract(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		e       *extractor
		args    args
		want    proto.TransformData
		wantErr bool
	}{
		{
			name: "ixigua1",
			args: args{
				url: "https://www.ixigua.com/pseries/6762421329418256907_6760625597132571147",
			},
		},
		{
			name: "ixigua2",
			args: args{
				url: "https://www.ixigua.com/6726810241054278158",
			},
		},
		{
			name: "ixigua2",
			args: args{
				// https://v3-xg-web-pc.ixigua.com/b810b1ce6c85f2d980cc0a3eae276446/652df50a/video/tos/cn/tos-cn-vd-0026/81e01389e31f447ea70fe44888767bb0/media-video-hvc1/?a=1768&ch=0&cr=7&dr=0&er=0&cd=0%7C0%7C0%7C1&cv=1&br=248&bt=248&cs=4&ds=2&mime_type=video_mp4&qs=0&rc=NDNnOjY1ZjY3aTpkaDo3N0BpMzpxajVkd2p4cTMzNDczM0BeNDIvXi42XzYxLzZhLzQyYSM0LWg0ZTUzZDZfLS0xLS9zcw%3D%3D&btag=e00030000&dy_q=1697506597&l=20231017093637DA2D649CCEE8C914FCE1
				url: "https://www.ixigua.com/7067856027621949964",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &extractor{}
			got, err := e.Extract(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractor.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, _ := json.Marshal(got)
			fmt.Println(string(b))
		})
	}
}
