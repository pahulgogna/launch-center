package storage

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/pahulgogna/launch-center/config"
	"github.com/pahulgogna/launch-center/types"
)

type Handler struct {
	rootDir string
}

func NewStore() *Handler {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panic("home directory could not be located")
	}

	return &Handler{
		rootDir: path.Join(homeDir, fmt.Sprintf(".%s", config.Env.AppName)),
	}
}

func (h *Handler) GetItems() (*[]types.Item, error) {

	dirs, err := os.ReadDir(h.rootDir)
	if err != nil {
		return nil, err
	}

	var result []types.Item

	for _, file := range dirs {
		if !file.IsDir() {
			
		}
	}

	return &result, nil
}

func (h *Handler) SaveItem(item types.Item) error {

}

func (h *Handler) UpdateItem(item types.Item) error {

}

func (h *Handler) DeleteItem(itemName string) error {

}
