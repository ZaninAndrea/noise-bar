package main

import (
	_ "embed"
	"os"
	"os/signal"
	"syscall"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/getlantern/systray"
)

type Sound struct {
	Name     string
	Path     string
	MenuItem *systray.MenuItem
}

func AddSound(name string, path string, icon []byte) Sound {
	item := systray.AddMenuItem(name, "Play "+name)
	item.SetIcon(icon)

	return Sound{
		Name:     name,
		Path:     path,
		MenuItem: item,
	}
}

type App struct {
	streamer *beep.StreamSeekCloser
	mQuit    *systray.MenuItem
	mStop    *systray.MenuItem
	sounds   []Sound
}

func (a *App) OnTrayReady() {
	systray.SetIcon(trayIcon)
	systray.SetTitle("")
	systray.SetTooltip("Background noise machine")
	a.sounds = append(a.sounds, AddSound("Airplane", "./sounds/air-plane.ogg", airPlane))
	a.sounds = append(a.sounds, AddSound("Birds", "./sounds/birds-tree.ogg", birdsTree))
	// a.sounds = append(a.sounds, AddSound("Brown noise", "./sounds/brown-noise.ogg", brownNoise))
	// a.sounds = append(a.sounds, AddSound("Brown noise 2", "./sounds/brown-noise2.ogg", brownNoise2))
	a.sounds = append(a.sounds, AddSound("Brown noise", "./sounds/brown-noise3.ogg", brownNoise3))
	a.sounds = append(a.sounds, AddSound("Cave", "./sounds/cave-drops.ogg", cave))
	a.sounds = append(a.sounds, AddSound("Coffee", "./sounds/coffee.ogg", coffee))
	a.sounds = append(a.sounds, AddSound("Drops", "./sounds/drops.ogg", drops))
	a.sounds = append(a.sounds, AddSound("Fire", "./sounds/fire.ogg", fire))
	a.sounds = append(a.sounds, AddSound("Leaves", "./sounds/leaves.ogg", leaves))
	a.sounds = append(a.sounds, AddSound("Night", "./sounds/night.ogg", night))
	a.sounds = append(a.sounds, AddSound("Rain", "./sounds/rain.ogg", rain))
	a.sounds = append(a.sounds, AddSound("Storm", "./sounds/storm.ogg", storm))
	a.sounds = append(a.sounds, AddSound("Stream water", "./sounds/stream-water.ogg", streamWater))
	a.sounds = append(a.sounds, AddSound("Train", "./sounds/train.ogg", train))
	a.sounds = append(a.sounds, AddSound("Underwater", "./sounds/underwater.ogg", underwater))
	a.sounds = append(a.sounds, AddSound("Washing machine", "./sounds/washing-machine.ogg", washingMachine))
	a.sounds = append(a.sounds, AddSound("Waterfall", "./sounds/waterfall.ogg", waterfall))
	a.sounds = append(a.sounds, AddSound("Waves", "./sounds/waves.ogg", waves))
	a.sounds = append(a.sounds, AddSound("Wind", "./sounds/wind.ogg", wind))

	systray.AddSeparator()
	a.mStop = systray.AddMenuItem("Stop", "Stop the noise")
	a.mQuit = systray.AddMenuItem("Quit", "Quit NoiseBar")

	for _, sound := range a.sounds {
		go a.HandleSoundButton(sound)
	}

	go a.HandleShutdownSignals()
	go a.HandleStopSound()
}

func (a *App) OnTrayQuit() {
	if a.streamer != nil {
		(*a.streamer).Close()
	}
	speaker.Close()

	os.Exit(0)
}

func (a *App) HandleShutdownSignals() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-shutdownSignal:
	case <-a.mQuit.ClickedCh:
	}

	systray.Quit()
}

func (a *App) HandleStopSound() {
	for {
		<-a.mStop.ClickedCh

		speaker.Clear()
		if a.streamer != nil {
			err := (*a.streamer).Close()
			if err != nil {
				panic(err)
			}
			a.streamer = nil
		}
	}
}

func (a *App) HandleSoundButton(sound Sound) {
	for {
		select {
		case <-sound.MenuItem.ClickedCh:
		}

		speaker.Clear()
		if a.streamer != nil {
			err := (*a.streamer).Close()
			if err != nil {
				panic(err)
			}
			a.streamer = nil
		}

		streamer, _, err := getStreamer(sound.Path)
		if err != nil {
			panic(err)
		}
		a.streamer = &streamer

		volume := &effects.Volume{
			Streamer: beep.Loop(-1, streamer),
			Base:     2,
			Volume:   0,
			Silent:   false,
		}
		speaker.Play(volume)
	}
}
