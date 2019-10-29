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
)

// App consts
const (
	ProjectID      = "replaceme"
	TokenID        = "replaceme" // nolint: gosec
	DefaultContext = "replaceme"
)

var Contexts []string

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
		signalwire.Log.SetLevel(signalwire.DebugLevelLog)
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
				signalwire.Log.Info("Exit\n")
				os.Exit(0)
			}
		}
	}()

	Contexts = append(Contexts, PContext)
	consumer := new(signalwire.Consumer)
	// setup the Client
	consumer.Setup(PProjectID, PTokenID, Contexts)
	// register callback
	consumer.Ready = func(consumer *signalwire.Consumer) {
		go spinner(100 * time.Millisecond)

		/*prepare the msg first, then send*/
		text := "Hello from Signalwire !"
		context := DefaultContext
		from := "+1XXXXXXXXXX"
		to := "+15XXXXXXXXX"

		message := consumer.Client.Messaging.NewMessage(context, from, to, text)
		message.OnMessageQueued = func(_ *signalwire.SendResult) {
			signalwire.Log.Info("Message Queued.\n")
		}

		message.OnMessageDelivered = func(_ *signalwire.SendResult) {
			signalwire.Log.Info("Message Delivered.\n")
		}

		resultSend1 := consumer.Client.Messaging.SendMsg(message)
		if !resultSend1.GetSuccessful() {
			signalwire.Log.Error("Could not send message\n")
		}

		/* now just send a message using Send() with params */

		resultSend2 := consumer.Client.Messaging.Send(from, to, context, "Hello again from Signalwire !")
		if !resultSend2.GetSuccessful() {
			signalwire.Log.Error("Could not send message\n")
		}

		if err := consumer.Stop(); err != nil {
			signalwire.Log.Error("Error occurred while stopping Signalwire Consumer\n")
		}
	}
	// start
	if err := consumer.Run(); err != nil {
		signalwire.Log.Error("Error occurred while starting Signalwire Consumer\n")
	}
}
