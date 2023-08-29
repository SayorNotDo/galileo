package pagination

type State struct {
	Offset        int         `json:"offset"`
	SortBy        string      `json:"sort_by"`
	FilterFields  []string    `json:"filter_field"`
	FilterValues  []string    `json:"filter_value"`
	SearchKeyword string      `json:"search_keyword"`
	Additional    interface{} `json:"additional"`
}
