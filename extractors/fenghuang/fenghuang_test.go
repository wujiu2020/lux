package fenghuang

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
		e       *extractor
		args    args
		want    proto.TransformData
		wantErr bool
	}{
		{
			name: "fenghuang1",
			args: args{
				url: "https://v.ifeng.com/c/8O1MxOEuQ4k",
			},
		},
		{
			name: "fenghuang2",
			args: args{
				url: "https://v.ifeng.com/c/8Mhn0o5wXPi",
			},
		},
		{
			name: "fenghuang3",
			args: args{
				url: "https://v.ifeng.com/c/8LYJO5wUt7z",
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
			data, err := got.TransformData(tt.args.url, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("fenghuang.TransformData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(data)
		})
	}
}
