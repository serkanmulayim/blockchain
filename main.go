package main

import (
	"os"
	"os/signal"
	"serkanmulayim/blockchain/p2p/peer"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	setLogLevel()
	peer.Start()

	<-stopCh
	log.Info("Detected shutdown signal. Shutting down gracefully.")
	peer.Stop()
}

func setLogLevel() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")

	if !ok { //TODO change to info
		lvl = "debug"
	}

	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}

	log.SetLevel(ll)
	if lvl == "debug" {
		//log.SetReportCaller(true)
	}

}
