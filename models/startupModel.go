package models

type Startup struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Pitch    string `json:"pitch"`
	Image    string `json:"image"`
	Slug     string `json:"slug"`
	Desc     string `json:"desc"`
	User     User   `json:"userId"`
}
