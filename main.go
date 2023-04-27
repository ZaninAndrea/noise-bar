package main

import (
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/getlantern/systray"
)

func getStreamer(sound string) (beep.StreamSeekCloser, beep.Format, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	f, err := os.Open(filepath.Join(exPath, sound))
	if err != nil {
		return nil, beep.Format{}, err
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}
	return streamer, format, nil
}

func main() {
	err := speaker.Init(48000, 4800)
	if err != nil {
		panic(err)
	}

	app := &App{}
	systray.Run(app.OnTrayReady, app.OnTrayQuit)
}
