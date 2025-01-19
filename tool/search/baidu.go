package search

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

var ErrorEncodeBaiduResp = errors.New("error encoding Baidu search result")

// NewBaiduSearch creates search for Baidu
func NewBaiduSearch(apiKey string) *Client {
	return NewSearch("baidu",
		apiKey,
		&BingSearchParamsHandler{},
		&BingSearchResultHandler{
			RequiredField: "organic_results",
		})
}

type BaiduSearchParamsHandler struct {
}

func (h *BaiduSearchParamsHandler) Handle(input string, pageIndex, pageSize int) (map[string]string, error) {
	return map[string]string{
		"q":     input,
		"first": strconv.Itoa((pageIndex-1)*pageSize + 1),
		"count": strconv.Itoa(pageSize),
	}, nil
}

type BaiduSearchResultHandler struct {
	RequiredField string
}

func (h *BaiduSearchResultHandler) GetRequiredField() string {
	return h.RequiredField
}

func (h *BaiduSearchResultHandler) Handle(result string) (string, error) {
	var BaiduDatas []BaiduData
	err := json.Unmarshal([]byte(result), &BaiduDatas)
	if err != nil {
		return "", ErrorEncodeBingResp
	}
	var output string
	for _, data := range BaiduDatas {
		output += Title + data.Title + "\n"
		output += Snippet + data.Snippet + "\n"
		output += Link + data.Link + "\n\n"
	}
	return output, nil
}
