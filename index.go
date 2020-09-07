package main

import (
	"fmt"
	"os"
	"path/filepath"

	"log"
	"net/http"
)

// Multipurpose Struct used in this file server
type FsStruct struct {
	TrackArray   []string
	TrackPlaying string
}

// Global Vars
var (
	textStructs FsStruct
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

// Make an array of the list of items in the audio directory
func PopulateTrackArray() {

	trackDirectory := "html/audio"
	err := filepath.Walk(trackDirectory, func(path string, info os.FileInfo, err error) error {
		textStructs.TrackArray = append(textStructs.TrackArray, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTracks loaded. Track count:", len(textStructs.TrackArray))

}

// TTY prompt message regarding file server usage
func TtyRunPrompt(HostPort string) {

	// Server at port
	fmt.Println("\n\t- Go Server is running the index.go app at,", HostPort)

	fmt.Println("\t- Open a web browser and navigate to \"localhost", HostPort, "\"")
	fmt.Println("\t- Or Open a web browser and navigate to:")
	fmt.Println("\t-\t\"127.0.0.1", HostPort, "\"")
	fmt.Println("\t-\t\"10.42.0.1", HostPort, "\"")
	fmt.Println("\t-\t\"192.168.0.1", HostPort, "\"")

	fmt.Println("\n( From within this prompt,\n\tpress Ctrl-C to terminate server hosting. )\n")

}

// Make place holder file. It will be populated with a JS script which will make a HTML5 Audio playlist
func MakeJsPlaylist(jsPlaylistPath string) {

	// check if file exists
	var _, err = os.Stat(jsPlaylistPath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(jsPlaylistPath)
		if err != nil {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully. [", jsPlaylistPath, "]")
}

// Populate a JS script which will make a HTML5 Audio playlist
func WriteJsPlaylist(jsPlaylistPath string) {

	var (
		javascriptString string
		trackName        string
		trackPath        string
	)

	for i := 1; i < len(textStructs.TrackArray); i++ {
		trackName = filepath.Base(textStructs.TrackArray[i])
		trackPath = "audio/" + trackName
		javascriptString += "AddListItem(\"" + trackName + "\", \"" + trackPath + "\");\n"
	}
	javascriptString += "\naudioPlayer();\n\n"

	// Open file using READ & WRITE permission.

	var file, err = os.OpenFile(jsPlaylistPath, os.O_RDWR, 0644)

	if err != nil {
		return
	}

	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(javascriptString)
	if err != nil {
		return
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		return
	}

	fmt.Println("File Updated Successfully. [", jsPlaylistPath, "]")
}

// Delete the JS HTML5 Audio playlits script. Used prior to writing a new playlist.
func DeleteJsPlaylist(jsPlaylistPath string) {

	// delete file
	var err = os.Remove(jsPlaylistPath)
	if err != nil {
		return
	}

	fmt.Println("File Deleted Successfully. [", jsPlaylistPath, "]")
}

func main() {
	var (
		HostPort       string = ":8080"
		jsPlaylistPath string = "./html/js/jsPlaylist.js"
	)

	// Array of files in the audio directory
	PopulateTrackArray()

	// Delete previous playlist
	DeleteJsPlaylist(jsPlaylistPath)

	// Make a new playlist
	MakeJsPlaylist(jsPlaylistPath)
	WriteJsPlaylist(jsPlaylistPath)

	// Display a TTY prompt message regarding file server usage
	TtyRunPrompt(HostPort)

	// File Server
	fs := http.FileServer(http.Dir("./html"))

	// Host streaming audio service
	log.Fatal(http.ListenAndServe(HostPort, fs))
}
