package iqiyi

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wujiu2020/lux/extractors/proto"
)

func Test_extractor_Extract(t *testing.T) {
	type fields struct {
		siteType SiteType
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*proto.Data
		wantErr bool
	}{
		{
			name: "iqiyi1",
			args: args{
				url: "https://www.iqiyi.com/v_a5tnh7yyr8.html",
			},
		},
		{
			name: "iqiyi2",
			args: args{
				url: "https://www.iqiyi.com/v_1lr0jb5povw.html",
			},
		},
		{
			name: "iqiyi3",
			args: args{
				url: "https://www.iqiyi.com/v_bte31fos94.html",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &extractor{
				siteType: tt.fields.siteType,
			}
			got, err := e.Extract(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractor.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, _ := json.Marshal(got)
			fmt.Println(string(b))
			_, err = got.TransformData(tt.args.url, "")
			if err != nil {
				t.Errorf("got.TransformData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// b, _ := json.Marshal(data)
			// fmt.Println("-------------", string(b), "-------------")
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("extractor.Extract() = %v, want %v", got, tt.want)
			// }
		})
	}
}
