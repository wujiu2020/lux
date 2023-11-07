package kuaishou

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
			name: "kuaishou1",
			args: args{
				url: "https://v.kuaishou.com/diUCxC",
			},
		},
		{
			name: "kuaishou2",
			args: args{
				url: "https://v.kuaishou.com/hiZKok",
			},
		},
		{
			name: "kuaishou3",
			args: args{
				url: "https://v.kuaishou.com/gJ6yQL",
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
			res, _ := got.TransformData(tt.args.url, "")
			b, _ := json.Marshal(res)
			fmt.Println(string(b))
		})
	}
}
