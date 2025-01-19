package search

const (
	Title   = "title: "
	Link    = "link: "
	Snippet = "snippet: "
)

// SerpSearchResult holds response
type SerpSearchResult map[string]interface{}

// SerpSearchParamsHandler is a handler for search params
type SerpSearchParamsHandler interface {
	Handle(input string, pageIndex, pageSize int) (map[string]string, error)
}

// SerpSearchResultHandler is a handler for search result
type SerpSearchResultHandler interface {
	GetRequiredField() string
	Handle(result string) (string, error)
}

// Factory is a factory function for creating a search engine client
type Factory func(string) *Client

// GoogleData Google Response struct
// refer to https://serpapi.com/search-api
type GoogleData struct {
	Position      int    `json:"position"`
	Title         string `json:"title"`
	Link          string `json:"link"`
	RedirectLink  string `json:"redirect_link"`
	DisplayedLink string `json:"displayed_link"`
	Snippet       string `json:"snippet"`
	Sitelinks     struct {
		Inline []struct {
			Title string `json:"title"`
			Link  string `json:"link"`
		} `json:"inline"`
	} `json:"sitelinks"`
	RichSnippet struct {
		Bottom struct {
			Extensions         []string `json:"extensions"`
			DetectedExtensions struct {
				IntroducedThCentury int `json:"introduced_th_century"`
			} `json:"detected_extensions"`
		} `json:"bottom"`
	} `json:"rich_snippet"`
	AboutThisResult struct {
		Source struct {
			Description    string `json:"description"`
			SourceInfoLink string `json:"source_info_link"`
			Security       string `json:"security"`
			Icon           string `json:"icon"`
		} `json:"source"`
		Keywords  []string `json:"keywords"`
		Languages []string `json:"languages"`
		Regions   []string `json:"regions"`
	} `json:"about_this_result"`
	AboutPageLink        string `json:"about_page_link"`
	AboutPageSerpapiLink string `json:"about_page_serpapi_link"`
	CachedPageLink       string `json:"cached_page_link"`
	RelatedPagesLink     string `json:"related_pages_link"`
}

// BingData Bing Response struct
// refer to https://serpapi.com/bing-search-api
type BingData struct {
	Position       int    `json:"position"`
	TrackingLink   string `json:"tracking_link"`
	Link           string `json:"link"`
	Title          string `json:"title"`
	DisplayedLink  string `json:"displayed_link"`
	Thumbnail      string `json:"thumbnail"`
	Snippet        string `json:"snippet"`
	CachedPageLink string `json:"cached_page_link"`
}

// BaiduData Baidu Response struct
// refer to https://serpapi.com/baidu-search-api
type BaiduData struct {
	Position       int    `json:"position"`
	Title          string `json:"title"`
	Link           string `json:"link"`
	Thumbnail      string `json:"thumbnail"`
	DisplayedBrand string `json:"displayed_brand"`
	RichSnippet    []struct {
		Extensions         []string `json:"extensions"`
		DetectedExtensions struct {
			简介 string `json:"简介"`
		} `json:"detected_extensions"`
	} `json:"rich_snippet"`
	Sitelinks []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"sitelinks"`
	Snippet string `json:"snippet"`
}
