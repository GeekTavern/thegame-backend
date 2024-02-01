// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Image struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	URL         *string `json:"url,omitempty"`
}

type Inventory struct {
	ID    string  `json:"id"`
	Items []*Item `json:"items,omitempty"`
}

type Item struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mutation struct {
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Query struct {
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type StoryPoint struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Image *Image `json:"image,omitempty"`
}

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Storypoint *StoryPoint `json:"storypoint"`
	Inventory  *Inventory  `json:"inventory"`
}
