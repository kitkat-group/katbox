package kontent

// Articles are the full struct for all kontent
type Articles struct {
	GHRepos  []GitHubRepository `json:"ghRepos"`
	Tools    []Tool             `json:"tools"`
	Snippets []Snippet          `json:"snippets"`
	Posts    []Post             `json:"posts"`
}

// GitHubRepository contains an array of repositories that will be useful
type GitHubRepository struct {
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Private  bool     `json:"private"`
	Notes    string   `json:"notes"`
	Keywords []string `json:"keywords"`
}

// Tool contains an array of tools that will be useful
type Tool struct {
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	Notes       string   `json:"notes"`
	Recommended bool     `json:"recommended"`
	Keywords    []string `json:"keywords"`
}

// Snippet contains an array of one-liners, code snippets
type Snippet struct {
	Name     string   `json:"name"`
	OneLiner string   `json:"oneliner"`
	Notes    string   `json:"notes"`
	Keywords []string `json:"keywords"`
}

// Post contain an array of Blog posts, or articles
type Post struct {
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	Recommended bool     `json:"recommended"`
	Internal    bool     `json:"internal"`
	Blog        bool     `json:"blog"`
	Keywords    []string `json:"keywords"`
}
