package main

import (
	"embed"

	"github.com/pahulgogna/launch-center/cmd"
	"github.com/pahulgogna/launch-center/storage"
	"github.com/pahulgogna/launch-center/types"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {

	var itemStore types.ItemStore = storage.NewStore()
	app := cmd.NewApp(itemStore)

	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			itemStore,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
