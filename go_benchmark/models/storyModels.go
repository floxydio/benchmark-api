package models

type StoryModels struct {
	StoryId     uint   `json:"story_id" db:"story_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Detail      string `json:"detail" db:"detail"`
	MetaTag     string `json:"meta_tag" db:"meta_tag"`
	Kategori    string `json:"kategori" db:"kategori"`
}
