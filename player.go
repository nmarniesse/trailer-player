package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/google/uuid"

	vlcctrl "github.com/CedArctic/go-vlc-ctrl"
)

const blackFile = "/samples/black-20s.mp4"

type Track struct {
	fullpath string
}

type Player struct {
	vlc *vlcctrl.VLC
}

func CreatePlayer(host string, port int, tracks []Track) (*Player, error) {
	logInfo("Create VLC player")

	vlcPassword := uuid.New().String()

	/**
	 * To launch on a specific screen:
	 * 	--directx-device=\\.\\DISPLAY2
	 * 	or --qt-fullscreen-screennumber=1 (0 is the first display, 1 the second, ...)
	 */
	cmd := exec.Command("vlc", "--fullscreen", "--intf", "http", "--extraintf", "qt", "--http-port", strconv.Itoa(port), "--http-password", vlcPassword, "--no-video-title-show", "--no-embedded-video", "--no-qt-fs-controller")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second) // Be sure VLC is ready

	vlc, err := vlcctrl.NewVLC(host, port, vlcPassword)
	if err != nil {
		logError(err)

		return nil, err
	}

	player := Player{vlc: &vlc}

	logInfo("Add playlist")

	currentDirectory, err := os.Getwd()
	if err != nil {
		logError(err)

		return nil, err
	}

	err = player.vlc.Add("file://" + currentDirectory + blackFile)
	if err != nil {
		logError(err)
	}
	for _, track := range tracks {
		logInfo("Add file [" + track.fullpath + "] to playlist")
		err = player.vlc.Add(track.fullpath)
		if err != nil {
			logError(err)
		}
	}
	err = player.vlc.Add("file://" + currentDirectory + blackFile)
	if err != nil {
		logError(err)
	}

	player.vlc.Play()
	time.Sleep(1 * time.Second) // Be sure video starts. Otherwise fullscreen doesn't work

	player.vlc.Volume("200")
	status, _ := player.vlc.GetStatus()
	if !status.Fullscreen {
		player.vlc.ToggleFullscreen()
	}

	player.vlc.Pause()
	player.vlc.Next()

	return &player, nil
}

func (player *Player) Start() {
	logInfo("Player Start")
	player.vlc.Play()
}
