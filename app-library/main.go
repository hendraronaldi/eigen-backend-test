package main

import (
	"app-library/internal/servers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := servers.Init() // initialize server
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh,
		syscall.SIGTERM, // terminate: stopped by `kill -9 PID`
		syscall.SIGINT,  // interrupt: stopped by Ctrl + C
	)

	go func() {
		defer func() {
			exitCh <- syscall.SIGTERM // send terminate signal when
			// application stop naturally
		}()
		s.Run() // run server / start the application
	}()

	<-exitCh  // blocking until receive exit signal
	s.Close() // close server
}
