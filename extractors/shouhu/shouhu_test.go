package shouhu

import (
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
			name: "shouhu1",
			args: args{
				url: "https://tv.sohu.com/v/MjAyMjA2MTYvbjYwMTE5MDc4NC5zaHRtbA==.html",
			},
		},
		{
			name: "shouhu2",
			args: args{
				url: "https://tv.sohu.com/v/MjAyMDA3MDMvbjYwMDg3NzE0OS5zaHRtbA==.html",
			},
		},
		{
			name: "shouhu3",
			args: args{
				url: "https://tv.sohu.com/v/MjAyMTA4MDMvbjYwMTAzMTc3NC5zaHRtbA==.html",
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
			fmt.Println(data)
		})
	}
}
