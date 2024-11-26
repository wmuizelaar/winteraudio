package easteregg

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

//go:embed winter.mp3
var winterAudio []byte

func PlayWinter() {
	readCloser := io.NopCloser(bytes.NewReader(winterAudio))

	streamer, format, err := mp3.Decode(readCloser)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
