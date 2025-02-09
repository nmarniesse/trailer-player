# Trailer Player

App to manage trailers playlists and play them.

## Requirements

### VLC media player

Install VLC on the computer

- Disable OSD (On Screen Display) to remove name of playing files: Tools > Preferences > All Settings > Video > Subtitle / OSD > Uncheck "On Screen Display"

Optional? (because we launch the VLC player directly in the app)
- Activate Web Interface: View > Add Interface > Web
- Add a password: Preferences > All Settings > Main interfaces > Lua > Lua HTTP > Password


## Build & Execution

```bash
go build .
./trailer-player
```


## Development

```bash
go run .
```

### Resources

- [go-vlc-ctrl](https://github.com/CedArctic/go-vlc-ctrl) - [Documentation](https://pkg.go.dev/github.com/CedArctic/go-vlc-ctrl)
