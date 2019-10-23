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
	ProjectID = "replaceme"
	TokenID   = "replaceme" // nolint: gosec
)

// Contexts not needed for only outbound calls
var Contexts = []string{"replaceme"}

// PProjectID passed from command-line
var PProjectID string

// PTokenID passed from command-line
var PTokenID string

// CallThisNumber get the callee phone number from command line
var CallThisNumber string

/*gopl.io spinner*/
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// MyReady - gets executed when Blade is successfully setup (after signalwire.receive)
func MyReady(consumer *signalwire.Consumer) {
	signalwire.Log.Info("calling out...\n")

	fromNumber := "+13XXXXXXXXX"

	var toNumber = "+16XXXXXXXXX"

	if len(CallThisNumber) > 0 {
		toNumber = CallThisNumber
	}

	resultDial := consumer.Client.Calling.DialPhone(fromNumber, toNumber)
	if !resultDial.Successful {
		if err := consumer.Stop(); err != nil {
			signalwire.Log.Error("Error occurred while trying to stop Consumer\n")
		}

		return
	}

	resultDial.Call.OnPrompt = func(promptaction *signalwire.PromptAction) {
		// we could do somethin here and this gets called when the Prompt Action finishes.
	}

	playAudioParams := signalwire.PlayAudioParams{
		URL: "https://www.voiptroubleshooter.com/open_speech/american/OSR_us_000_0010_8k.wav",
	}

	playTTSParams := signalwire.PlayTTSParams{
		Text: "Hello from Signalwire!",
	}

	playRingtoneParams := signalwire.PlayRingtoneParams{
		Duration: 5,
		Name:     "us",
	}

	play := []signalwire.PlayStruct{{
		Type:   "audio",
		Params: playAudioParams,
	}, {
		Type:   "tts",
		Params: playTTSParams,
	}, {
		Type:   "ringtone",
		Params: playRingtoneParams,
	}}

	collectDigits := new(signalwire.CollectDigits)
	collectDigits.Max = 2

	collectSpeech := new(signalwire.CollectSpeech)
	collectSpeech.EndSilenceTimeout = 5
	collectSpeech.SpeechTimeout = 10
	collect := signalwire.CollectStruct{
		Speech: collectSpeech,
		Digits: collectDigits,
	}

	promptAction, err := resultDial.Call.PromptAsync(&play, &collect)

	if err != nil {
		signalwire.Log.Error("Error occurred while trying to start Prompt Action\n")

		if err := consumer.Stop(); err != nil {
			signalwire.Log.Error("Error occurred while trying to stop Consumer. Err: %v\n", err)
		}

		return
	}

	// do something here
	go spinner(100 * time.Millisecond)
	time.Sleep(10 * time.Second)

	promptAction.Stop()

	if _, err := resultDial.Call.Hangup(); err != nil {
		signalwire.Log.Error("Error occurred while trying to hangup call. Err: %v\n", err)
	}

	if err := consumer.Stop(); err != nil {
		signalwire.Log.Error("Error occurred while trying to stop Consumer. Err: %v\n", err)
	}
}

func main() {
	var printVersion bool

	var verbose bool

	flag.BoolVar(&printVersion, "v", false, " Show version ")
	flag.StringVar(&CallThisNumber, "n", "", " Number to call ")
	flag.StringVar(&PProjectID, "p", ProjectID, " ProjectID ")
	flag.StringVar(&PTokenID, "t", TokenID, " TokenID ")
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

	consumer := new(signalwire.Consumer)
	// setup the Client
	consumer.Setup(PProjectID, PTokenID, Contexts)
	// register callback
	consumer.Ready = MyReady
	// start
	if err := consumer.Run(); err != nil {
		signalwire.Log.Error("Error occurred while starting Signalwire Consumer\n")
	}
}