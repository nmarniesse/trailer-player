package main

func main() {
	logInfo("Start")

	httpPort := 8088

	tracks := []Track{
		{fullpath: "/home/Nicolas-Marniesse/workspace/trailer-cinema/samples/5538137-hd_1920_1080_25fps.mp4"},
		{fullpath: "/home/Nicolas-Marniesse/workspace/trailer-cinema/samples/7233646-hd_1080_1920_25fps.mp4"},
		{fullpath: "/home/Nicolas-Marniesse/workspace/trailer-cinema/samples/minimalistic-design-ambient-background-201445.mp3"},
	}

	player, _ := CreatePlayer("127.0.0.1", httpPort, tracks)

	player.Start()

	logInfo("End")
}
