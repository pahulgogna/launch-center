package cmd

import (
	"context"

	"github.com/pahulgogna/launch-center/types"
)

type App struct {
	ctx context.Context
	Store types.ItemStore
}

func NewApp(itemStore types.ItemStore) *App {
	return &App{
		Store: itemStore,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}
