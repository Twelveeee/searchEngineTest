package model

type Article struct {
	Rid           string
	Tags          string
	Author        string
	Author_avatar string
	Name          string
	Full_name     string
	Title         string
	Description   string
	Summary       string
	Lang_color    string
	Primary_lang  string
	Stars_str     string
	Publish_at    string
	Stars         int
	Has_chinese   bool
	Is_active     bool
}
