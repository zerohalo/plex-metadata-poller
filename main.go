package main

import (
	"fmt"
	"github.com/zerohalo/goplexapi"
	"os"
	"time"
)

var (
	plexToken      = os.Getenv("PLEX_TOKEN")
	plexServer     = os.Getenv("PLEX_SERVER")
	plexUserName   = os.Getenv("PLEX_USER_NAME")
	plexClientName = os.Getenv("PLEX_CLIENT_NAME")
	pollInterval   = 5 * time.Second
)

func main() {
	if plexToken == "" {
		fmt.Println("Error: PLEX_TOKEN not set")
		return
	}
	if plexServer == "" {
		fmt.Println("Error: PLEX_SERVER not set")
		return
	}
	plexUrl := fmt.Sprintf("https://%s", plexServer)
	client := goplexapi.NewPlexClient(plexUrl, plexToken)
	var previousTrackInfo *goplexapi.TrackInfo
	for {
		trackInfo, err := client.GetCurrentPlayingSong(plexClientName, plexUserName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if previousTrackInfo == nil || trackInfo.Title != previousTrackInfo.Title {
			fmt.Printf("Currently Playing Song: %s Artist: %s Album: %s\n", trackInfo.Title, trackInfo.Artist, trackInfo.Album)
			//art, err := client.GetAlbumArt(trackInfo.Thumb)
			//if err != nil {
			//	fmt.Println("Error:", err)
			//	return
			//}
			//b64art := base64.StdEncoding.EncodeToString(art)
			//trackInfo.Thumb = b64art
			//mimeType := http.DetectContentType(art)

			writeFile(trackInfo)
			previousTrackInfo = trackInfo
		}
		time.Sleep(pollInterval)
	}
}

func writeFile(trackInfo *goplexapi.TrackInfo) {
	fh, err := os.Create("now_playing.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		err := fh.Close()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}()

	// ArtworkData not currently supported by Icecast, but may be in the future
	//_, err = fh.WriteString(fmt.Sprintf("Title: %s\nArtist: %s\nAlbum: %s\nArtworkData: %s\n", trackInfo.Title, trackInfo.Artist, trackInfo.Album, trackInfo.Thumb))
	_, err = fh.WriteString(fmt.Sprintf("Title: %s\nArtist: %s\nAlbum: %s\n", trackInfo.Title, trackInfo.Artist, trackInfo.Album))

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
