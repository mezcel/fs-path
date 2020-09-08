/* https://github.com/mezcel/fs-path/index.go */

// About: A Golang file server hosting M3U or Html5 streaming audio.
// git repo: https://github.com/mezcel/fs-path.git
package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

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

// Make an M3U playlist
func GenerateM3UPlaylist(m3uPlaylistPath string) {

	var (
		m3uString string
		trackName string
		trackPath string
	)

	// delete file
	var errDel = os.Remove(m3uPlaylistPath)
	if errDel != nil {
		fmt.Println("File does not exist yet. [", m3uPlaylistPath, "]")
		//return
	}

	fmt.Println("File Deleted Successfully. [", m3uPlaylistPath, "]")

	// check if file exists
	var _, errCreate = os.Stat(m3uPlaylistPath)

	// create file if not exists
	if os.IsNotExist(errCreate) {
		var file, errCreate = os.Create(m3uPlaylistPath)
		if errCreate != nil {
			fmt.Println("File does not exist yet. [", m3uPlaylistPath, "]")
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully. [", m3uPlaylistPath, "]")

	m3uString = "#EXTM3U\n"
	//m3uString += "#M3U generated at " + time.Now().String() + "\n"

	// load track paths into a js script
	for i := 1; i < len(textStructs.TrackArray); i++ {
		trackName = filepath.Base(textStructs.TrackArray[i])
		trackPath = trackName

		m3uString += "#EXTINF:" + trackName + " \n"
		m3uString += "../audio/" + trackPath + " \n"
	}

	// Open file using READ & WRITE permission.

	var file, errWrite = os.OpenFile(m3uPlaylistPath, os.O_RDWR, 0644)

	if errWrite != nil {
		return
	}

	defer file.Close()

	// Write some text line-by-line to file.
	_, errWrite = file.WriteString(m3uString)
	if errWrite != nil {
		return
	}

	// Save file changes.
	errWrite = file.Sync()
	if errWrite != nil {
		fmt.Println("File does not exist yet. [", m3uPlaylistPath, "]")
		return
	}

	fmt.Println("File Updated Successfully. [", m3uPlaylistPath, "]")

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

	javascriptString = "/*\n\thttps://github.com/mezcel/fs-path/html/js/jsPlaylist.js\n"
	javascriptString += "\tThis is a Temporary File generated at " + time.Now().String() + "\n"
	javascriptString += "\tUsed as a playlist importer.\n"
	javascriptString += "\n\t*/\n\n"

	// load track paths into a js script
	for i := 1; i < len(textStructs.TrackArray); i++ {
		trackName = filepath.Base(textStructs.TrackArray[i])
		trackPath = "audio/" + trackName
		javascriptString += "AddListItem(\"" + trackName + "\", \"" + trackPath + "\");\n"
	}

	javascriptString += "\n/* Load the html3 audio playlist */\naudioPlayer();\n\n"

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

// TTY prompt message regarding file server usage
func TtyRunPrompt(HostPort string) {

	var ip net.IP

	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		fmt.Println(err)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			fmt.Println(err)
		}

		for _, addr := range addrs {
			//var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
		}
	}

	// Server at port
	fmt.Println("\nReady to serve.")
	fmt.Println("\t- Host Port:\t", HostPort)
	fmt.Println("\t- Web Url:\t", ip, HostPort)
	fmt.Println("\t- M3U Channel:\t", ip, HostPort, "/M3U/playlist.m3u")

	fmt.Println("\n( From within this prompt,\n\tpress Ctrl-C to terminate server hosting. )")
	fmt.Println("")
}

func main() {
	var (
		HostPort        string = ":8080"
		jsPlaylistPath  string = "./html/js/jsPlaylist.js"
		m3uPlaylistPath string = "./html/M3U/playlist.m3u"
	)

	// Array of files in the audio directory
	PopulateTrackArray()

	// Delete previous playlist
	DeleteJsPlaylist(jsPlaylistPath)

	// Make a new playlist
	MakeJsPlaylist(jsPlaylistPath)
	WriteJsPlaylist(jsPlaylistPath)

	GenerateM3UPlaylist(m3uPlaylistPath)

	// Display a TTY prompt message regarding file server usage
	TtyRunPrompt(HostPort)

	// File Server
	fs := http.FileServer(http.Dir("./html"))

	// Host streaming audio service
	log.Fatal(http.ListenAndServe(HostPort, fs))
}
