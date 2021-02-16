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
		sig := <-sigchan
		app.Stop(sig)
	}()

	// Start main app
	app.Start()

	app.wg.Wait()
	return
}
