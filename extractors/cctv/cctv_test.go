package cctv

import (
	"testing"

	"github.com/wujiu2020/lux/extractors/proto"
)

func Test_extractor_Extract(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *proto.Data
		wantErr bool
	}{
		{
			name: "cctv",
			args: args{
				url: "https://tv.cctv.com/2022/03/08/VIDEnptauE0cn7I2t7ycCAs8220308.shtml",
			},
		},
		{
			name: "cctv",
			args: args{
				url: "https://tv.cctv.com/2023/03/13/VIDEecCBy1ZBv5L8g16LFOCc230313.shtml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &extractor{}
			_, err := e.Extract(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractor.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
