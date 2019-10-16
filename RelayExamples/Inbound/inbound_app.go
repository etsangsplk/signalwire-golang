package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/signalwire/signalwire-golang/signalwire"
	log "github.com/sirupsen/logrus"
)

// App consts
const (
	ProjectID      = "replaceme"
	TokenID        = "replaceme" // nolint: gosec
	DefaultContext = "replaceme"
)

// Contexts needed for inbound calls
var Contexts = []string{}

// PProjectID passed from command-line
var PProjectID string

// PTokenID passed from command-line
var PTokenID string

// PContext passed from command line (just one being passed, although we support many)
var PContext string

/*gopl.io spinner*/
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// MyOnIncomingCall - gets executed when we receive an incoming call
func MyOnIncomingCall(consumer *signalwire.Consumer, call *signalwire.CallObj) {
	fmt.Printf("got incoming call.\n")

	resultAnswer := call.Answer()
	if !resultAnswer.Successful {
		if err := consumer.Stop(); err != nil {
			log.Errorf("Error occurred while trying to stop Consumer")
		}

		return
	}

	log.Info("Playing audio on call..")

	go spinner(100 * time.Millisecond)

	if _, err := call.PlayAudio("https://cdn.signalwire.com/default-music/welcome.mp3"); err != nil {
		log.Errorf("Error occurred while trying to play audio")
	}

	if err := call.Hangup(); err != nil {
		log.Errorf("Error occurred while trying to hangup call")
	}

	if err := consumer.Stop(); err != nil {
		log.Errorf("Error occurred while trying to stop Consumer")
	}
}

func main() {
	var printVersion bool

	var verbose bool

	flag.BoolVar(&printVersion, "v", false, " Show version ")
	flag.StringVar(&PProjectID, "p", ProjectID, " ProjectID ")
	flag.StringVar(&PTokenID, "t", TokenID, " TokenID ")
	flag.StringVar(&PContext, "c", DefaultContext, " Context ")
	flag.BoolVar(&verbose, "d", false, " Enable debug mode ")
	flag.Parse()

	if printVersion {
		fmt.Printf("%s\n", signalwire.SDKVersion)
		fmt.Printf("Blade version: %d.%d.%d\n", signalwire.BladeVersionMajor, signalwire.BladeVersionMinor, signalwire.BladeRevision)
		fmt.Printf("App built with GO Lang version: " + fmt.Sprintf("%s\n", runtime.Version()))

		os.Exit(0)
	}

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
		for {
			s := <-interrupt
			switch s {
			case syscall.SIGHUP:
				fallthrough
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGINT:
				log.Printf("Exit")
				os.Exit(0)
			}
		}
	}()

	Contexts = append(Contexts, PContext)
	consumer := new(signalwire.Consumer)
	// setup the Client
	consumer.Setup(PProjectID, PTokenID, Contexts)
	// register callback
	consumer.OnIncomingCall = MyOnIncomingCall

	log.Info("Wait incoming call..")

	// start
	if err := consumer.Run(); err != nil {
		log.Errorf("Error occurred while starting Signalwire Consumer")
	}
}