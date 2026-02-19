package types

type ItemStore interface {
	GetItems() (*[]Item, error)
	SaveItem(item Item) error
	UpdateItem(item Item) error
	DeleteItem(itemName string) error
}

type Item struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Url               string `json:"url"`
	ThumbnailLocation string `json:"thumbnail"`
}
