package signal

import (
	"os"
	"os/signal"
)

func Wait(signals ...os.Signal) os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	s := <-c
	return s
}
