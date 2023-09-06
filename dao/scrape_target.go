package dao

type ScrapeTarget struct {
	ID       string `json:"id,omitempty"`
	Target   string `json:"target,omitempty"`
	Labels   JSON   `json:"labels,omitempty"`
	GroupId  string `json:"groupId,omitempty"`
	Describe string `json:"describe,omitempty"`
}
