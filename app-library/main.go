package main

import (
	"app-library/internal/servers"
	"os"
	"os/signal"
	"syscall"

	_ "app-library/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Library server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000

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
