package types

type JSResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

type Paginated struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type TreeNode struct {
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url"`
	Children    []TreeNode `json:"children,omitempty"`
}
