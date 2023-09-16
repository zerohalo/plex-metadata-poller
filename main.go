package main

import (
	"encoding/base64"
	"fmt"
	"github.com/zerohalo/goplexapi"
	"os"
	"time"
)

var (
	plexToken    = os.Getenv("PLEX_TOKEN")
	plexServer   = os.Getenv("PLEX_SERVER")
	pollInterval = 5 * time.Second
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
	for {
		trackInfo, err := client.GetCurrentPlayingSong("Plexamp")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Currently Playing Song: %s by: %s from Album: %s\n", trackInfo.Title, trackInfo.Artist, trackInfo.Album)
		art, err := client.GetAlbumArt(trackInfo.Thumb)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		b64art := base64.StdEncoding.EncodeToString(art)
		trackInfo.Thumb = b64art
		//mimeType := http.DetectContentType(art)

		writeFile(trackInfo)
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

	_, err = fh.WriteString(fmt.Sprintf("Title: %s\nArtist: %s\nAlbum: %s\nArtworkData: %s\n", trackInfo.Title, trackInfo.Artist, trackInfo.Album, trackInfo.Thumb))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	time.Sleep(time.Duration(pollInterval) * time.Second)
}
