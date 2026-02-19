package main

import (
	"embed"

	"github.com/pahulgogna/launch-center/cmd"
	"github.com/pahulgogna/launch-center/storage"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	itemStore := storage.NewStore()
	app := cmd.NewApp(itemStore)

	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 100},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
		WindowStartState: options.Maximised,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
