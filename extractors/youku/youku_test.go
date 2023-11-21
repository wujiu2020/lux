package youku

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
			name: "youku1",
			args: args{
				url: "https://v.youku.com/v_show/id_XNTIwNTIwMzQyMA==.html",
			},
		},
		{
			name: "youku2",
			args: args{
				url: "https://v.youku.com/v_show/id_XNTIwMjAyODUwMA==.html",
			},
		},
		{
			name: "youku3",
			args: args{
				url: "https://v.youku.com/v_show/id_XNTE5Nzg0ODAyOA==.html",
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
