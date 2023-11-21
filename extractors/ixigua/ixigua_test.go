package ixigua

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
			name: "ixigua3",
			args: args{
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
			data, err := got.TransformData(tt.args.url, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("extractor.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(data)
		})
	}
}
