package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pahulgogna/launch-center/config"
	"github.com/pahulgogna/launch-center/types"
)

type Handler struct {
	rootAppDir string
	itemsDir   string
}

func NewStore() *Handler {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panic("home directory could not be located")
	}

	rootAppDir := path.Join(homeDir, fmt.Sprintf(".%s", config.Env.AppName))
	itemsDir := path.Join(rootAppDir, "collection")

	if notExists(rootAppDir) {
		os.Mkdir(rootAppDir, 0770)
	}
	if notExists(itemsDir) {
		os.Mkdir(itemsDir, 0770)
	}

	return &Handler{
		rootAppDir: rootAppDir,
		itemsDir:   itemsDir,
	}
}

func (h *Handler) GetItems() ([]types.Item, error) {

	dirs, err := os.ReadDir(h.itemsDir)
	if err != nil {
		return nil, err
	}

	var result []types.Item

	for _, file := range dirs {
		if !file.IsDir() && strings.Contains(file.Name(), ".json") {
			data, err := os.ReadFile(path.Join(h.itemsDir, file.Name()))
			if err != nil {
				continue
			}

			var item types.Item

			err = json.Unmarshal(data, &item)
			if err != nil {
				continue
			}
			if item.Name == "" || item.Url == "" {
				continue
			}

			result = append(result, item)
		}
	}

	return result, nil
}

func (h *Handler) SaveItem(item types.Item) error {

	if item.Name == "" {
		return fmt.Errorf("name is required")
	}
	if (item.Url == "") {
		return  fmt.Errorf("url is required")
	}

	filePath := path.Join(h.itemsDir, item.Name)
	if !notExists(filePath) {
		return fmt.Errorf("name already exists, please choose another name")
	}

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(h.itemsDir, fmt.Sprintf("%s.json", item.Name)), data, 0660)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateItem(item types.Item) error {
	return nil
}

func (h *Handler) DeleteItem(itemName string) error {
	return nil
}

func notExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
