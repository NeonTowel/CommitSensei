// Package main initializes and runs the CommitSensei application using Wails
// [Scope: Application] [Status: Stable]
// Provides a user interface for enhancing commit messages with emojis

package main

import (
	// Import necessary packages for context management, embedding assets, and logging
	"context"
	"embed"
	"log"

	// Import Wails packages for application options and asset server configuration
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// Embed the frontend assets using the embed package
// WHY: Serve static files for the UI without external dependencies
//
//go:embed all:frontend/dist
var assets embed.FS

// main is the entry point of the application
// WHY: Initializes and runs the Wails application with specified options
func main() {
	// Initialize a new App instance
	app := NewApp()

	// Configure and run the Wails application
	err := wails.Run(&options.App{
		Title:  "CommitSensei - Emojis for your commit messages", // Set the application title
		Width:  1024,                                             // Set the application window width
		Height: 768,                                              // Set the application window height
		AssetServer: &assetserver.Options{
			Assets: assets, // Serve embedded assets
		},
		OnStartup:  app.startup,  // Define startup behavior
		OnShutdown: app.shutdown, // Define shutdown behavior
		Bind: []interface{}{
			app, // Bind the App instance to the Wails runtime
		},
	})
	if err != nil {
		log.Fatal(err) // Log any errors that occur during application startup
	}
}

// NewApp creates and returns a new App instance
// WHY: Encapsulates application state and behavior
func NewApp() *App {
	return &App{}
}

// App represents the application state and behavior
// WHY: Manages context and lifecycle events for the application
type App struct {
	ctx context.Context // Context for managing application state
}

// startup is called when the application starts
// WHY: Initializes the application context for use during runtime
func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

// shutdown is called when the application shuts down
// WHY: Performs any necessary cleanup before the application exits
func (b *App) shutdown(ctx context.Context) {}
