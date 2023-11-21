package qq

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
		want    []*proto.Data
		wantErr bool
	}{
		{
			name: "qq1",
			args: args{
				url: "https://v.qq.com/x/cover/mzc00200my8s5sr/w0044vllwhj.html",
			},
		},
		{
			name: "qq2",
			args: args{
				url: "https://v.qq.com/x/cover/mzc00200c2ku08a/d0041bwpuu7.html",
			},
		},
		{
			name: "qq3",
			args: args{
				url: "https://v.qq.com/x/cover/mzc00200wl6js55/o0042uh3rfa.html",
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
