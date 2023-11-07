package fenghuang

import "github.com/wujiu2020/lux/extractors/proto"

type VideoInfo struct {
	VSource        string `json:"vSource"`
	CommentURL     string `json:"commentUrl"`
	Base62ID       string `json:"base62Id"`
	SearchPath     string `json:"searchPath"`
	Title          string `json:"title"`
	NewsTime       string `json:"newsTime"`
	URL            string `json:"url"`
	VideoPlayURL   string `json:"videoPlayUrl"`
	GUID           string `json:"guid"`
	Vid            string `json:"vid"`
	Skey           string `json:"skey"`
	BreadCrumbdata []struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"breadCrumbdata"`
	Duration          float64 `json:"duration"`
	PosterURL         string  `json:"posterUrl"`
	ColumnName        string  `json:"columnName"`
	Createdate        string  `json:"createdate"`
	CategoryID        string  `json:"categoryId"`
	Keywords          string  `json:"keywords"`
	WemediaEAccountID string  `json:"wemediaEAccountId"`
	GuideAppSetting   struct {
		Duration int `json:"duration"`
		Enable   int `json:"enable"`
	} `json:"guideAppSetting"`
	LinkSpecialURL string `json:"linkSpecialUrl"`
	HasCopyRight   bool   `json:"hasCopyRight"`
	Subscribe      struct {
		Type               string `json:"type"`
		CateSource         string `json:"cateSource"`
		IsShowSign         int    `json:"isShowSign"`
		Parentid           string `json:"parentid"`
		Parentname         string `json:"parentname"`
		Cateid             string `json:"cateid"`
		Catename           string `json:"catename"`
		Logo               string `json:"logo"`
		Description        string `json:"description"`
		API                string `json:"api"`
		ShowLink           int    `json:"show_link"`
		ShareURL           string `json:"share_url"`
		EAccountID         int    `json:"eAccountId"`
		Status             int    `json:"status"`
		HonorName          string `json:"honorName"`
		HonorImg           string `json:"honorImg"`
		HonorImgNight      string `json:"honorImg_night"`
		ForbidFollow       int    `json:"forbidFollow"`
		ForbidJump         int    `json:"forbidJump"`
		FhtID              string `json:"fhtId"`
		View               int    `json:"view"`
		SourceFrom         string `json:"sourceFrom"`
		Declare            string `json:"declare"`
		OriginalName       string `json:"originalName"`
		RedirectTab        string `json:"redirectTab"`
		AuthorURL          string `json:"authorUrl"`
		NewsTime           string `json:"newsTime"`
		LastArticleAddress string `json:"lastArticleAddress"`
	} `json:"subscribe"`
	SourceAlias  string `json:"sourceAlias"`
	SourceReason string `json:"sourceReason"`
	Summary      string `json:"summary"`
}

func (v VideoInfo) TransformData(url string, quality string) (*proto.Data, error) {
	return &proto.Data{
		Duration: v.Duration,
		Streams: []proto.Stream{
			{
				Segs: []proto.Seg{
					{
						Duration: v.Duration,
						URL:      v.VideoPlayURL,
					},
				},
			},
		},
		Title: v.Title,
	}, nil
}
