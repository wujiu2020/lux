package mgtv

import (
	"testing"
)

func Test_extractor_Extract(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		e       *extractor
		args    args
		wantErr bool
	}{
		{
			name: "mgtv",
			args: args{
				url: "https://www.mgtv.com/b/354045/10948467.html",
			},
		},
		{
			name: "mgtv",
			args: args{
				url: "https://www.mgtv.com/b/388252/15472835.html",
			},
		},
		{
			name: "mgtv",
			args: args{
				url: "https://www.mgtv.com/b/401876/15491442.html",
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
			if _, err := got.TransformData(tt.args.url, ""); err != nil {
				t.Errorf("got.TransformData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
