package main

import (
	"fmt"
	"time"
	"os/exec"
	"log"

	vlcctrl "github.com/CedArctic/go-vlc-ctrl"
)

func main() {
	fmt.Println("Hello!")

	fmt.Println("Launch VLC")
	cmd := exec.Command("vlc", "--intf", "http", "--extraintf", "qt", "--http-password", "the-password")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second) // Be sure VLC is ready

	// Declare a local VLC instance on port 8080 with password "the-password"
	myVLC, err := vlcctrl.NewVLC("127.0.0.1", 8080, "the-password")
	if (err != nil) {
		fmt.Println(err)
		return
	}

	/**
	 * Add items to playlist.
	 * We add a black in order to configure VLC, when ready we launch the playlist.
	 */
	fmt.Println("Add videos...")
	myVLC.Add("file:///home/Nicolas-Marniesse/workspace/trailer-cinema/samples/black-20s.mp4")
	myVLC.Add("file:///home/Nicolas-Marniesse/workspace/trailer-cinema/samples/5538137-hd_1920_1080_25fps.mp4")
	myVLC.Add("file:///home/Nicolas-Marniesse/workspace/trailer-cinema/samples/7233646-hd_1080_1920_25fps.mp4")
	myVLC.Add("file:///home/Nicolas-Marniesse/workspace/trailer-cinema/samples/minimalistic-design-ambient-background-201445.mp3")
	myVLC.Add("file:///home/Nicolas-Marniesse/workspace/trailer-cinema/samples/black-20s.mp4")
	

	// Play first item and wait for 10 seconds
	fmt.Println("Play black and configure VLC")
	myVLC.Play()
	time.Sleep(1 * time.Second) // Be sure video starts. Otherwise fullscreen doesn't work
	status, _ := myVLC.GetStatus()
	if (!status.Fullscreen) {
		myVLC.ToggleFullscreen()
	}
	myVLC.Volume("200")
	time.Sleep(2 * time.Second)

	fmt.Println("Let's start!")
	myVLC.Next()

	fmt.Println("End")
}
