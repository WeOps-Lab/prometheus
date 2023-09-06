package dao

type ScrapeGroup struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	MetricsPath  string `json:"metricsPath,omitempty"`
	Labels       JSON   `json:"labels,omitempty"`
	Scheme       string `json:"scheme,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	Describe     string `json:"describe,omitempty"`
	AuthUser     string `json:"authUser,omitempty"`
	AuthPassword string `json:"authPassword,omitempty"`
	AuthToken    string `json:"authToken,omitempty"`

	ScrapeTarget []ScrapeTarget `json:"scrapeTarget,omitempty"`
}
