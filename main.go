package main

import (
	"os"
	"os/signal"
)

func main() {
	app := NewUpMaster()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan)

	// Graceful Shutdown
	go func() {
		_ = <-sigchan
		app.Stop()
	}()

	// Start main app
	app.Start()
	return
}
