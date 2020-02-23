package signalwatcher

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

	}()
}
