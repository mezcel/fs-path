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

// I did not need to make structs, but I felt this could come in handy if I need to scale the project later
// Multipurpose Struct used in this file server
type FsStruct struct {
	TrackArray   []string
	TrackPlaying string
}

// Global Struct Vars
var (
	fsStructs FsStruct
)

// Make an array of the list of items in the audio directory
func PopulateFilesArray() {

	trackDirectory := "html/audio"
	err := filepath.Walk(trackDirectory, func(path string, info os.FileInfo, err error) error {
		fsStructs.TrackArray = append(fsStructs.TrackArray, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTracks loaded. Track count:", len(fsStructs.TrackArray))

}

// Make place holder file. It will be populated with a JS script which will make a HTML5 Audio playlist
func MakeTextFile(textfilePath string) {

	// check if file exists
	var _, err = os.Stat(textfilePath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(textfilePath)
		if err != nil {
			return
		}
		defer file.Close()
	}

	fmt.Println("\tFile Created Successfully. [", textfilePath, "]")
}

// Delete the JS HTML5 Audio playlist script. Used prior to writing a new playlist.
func DeleteTextFile(textfilePath string) {

	// delete file
	var err = os.Remove(textfilePath)
	if err != nil {
		return
	}

	fmt.Println("\tFile Deleted Successfully. [", textfilePath, "]")
}

// Generate JS script which will make a HTML5 Audio playlist
func GenerateJSScript() string {
	var (
		javascriptString string
		trackName        string
		trackPath        string

		// fsStructs.TrackArray was a global struct var
	)

	javascriptString = "/*\n\thttps://github.com/mezcel/fs-path/html/js/jsPlaylist.js\n"
	javascriptString += "\tThis is a Temporary File generated at " + time.Now().String() + "\n"
	javascriptString += "\tUsed as a playlist importer.\n"
	javascriptString += "\n\t*/\n\n"

	// load track paths into a js script
	for i := 1; i < len(fsStructs.TrackArray); i++ {
		trackName = filepath.Base(fsStructs.TrackArray[i])
		trackPath = "audio/" + trackName
		javascriptString += "AddListItem(\"" + trackName + "\", \"" + trackPath + "\");\n"
	}

	javascriptString += "\n/* Load the html3 audio playlist */\naudioPlayer();\n\n"

	return javascriptString
}

// Generate M3U script listing file paths in the audio directory
func GenerateM3UScript() string {

	var (
		m3uString string
		trackName string
		trackPath string

		// fsStructs.TrackArray was a global struct var
	)

	m3uString = "#EXTM3U\n"
	m3uString += "#M3U generated at " + time.Now().String() + "\n"

	// load track paths into a js script
	for i := 1; i < len(fsStructs.TrackArray); i++ {
		trackName = filepath.Base(fsStructs.TrackArray[i])
		trackPath = trackName

		m3uString += "#EXTINF:0," + trackName + " \n"
		m3uString += "../audio/" + trackPath + " \n"
	}

	return m3uString
}

// Make a JS script file
func WriteJsPlaylist(jsPlaylistPath string) {

	var javascriptString string

	DeleteTextFile(jsPlaylistPath)
	MakeTextFile(jsPlaylistPath)

	javascriptString = GenerateJSScript()

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

	fmt.Println("\tFile Updated Successfully. [", jsPlaylistPath, "]")
}

// Make an M3U playlist file
func WriteM3UPlaylist(m3uPlaylistPath string) {

	var m3uString string = GenerateM3UScript()

	DeleteTextFile(m3uPlaylistPath)
	MakeTextFile(m3uPlaylistPath)
	m3uString = GenerateM3UScript()

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
		fmt.Println("\tFile does not exist yet. [", m3uPlaylistPath, "]")
		return
	}

	fmt.Println("\tFile Updated Successfully. [", m3uPlaylistPath, "]")

}

// TTY prompt message regarding file server usage
func TtyRunPrompt(HostPort string) {

	var ip net.IP

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
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
	PopulateFilesArray()

	fmt.Println("\nGenerating:", jsPlaylistPath)
	// Write a js script to dynamically add track to the html audio ol list
	WriteJsPlaylist(jsPlaylistPath)

	fmt.Println("\nGenerating:", m3uPlaylistPath)
	// Generate a standalone W3M streaming audio player file
	WriteM3UPlaylist(m3uPlaylistPath)

	// Display a TTY prompt message regarding file server usage
	TtyRunPrompt(HostPort)

	// File Server
	fs := http.FileServer(http.Dir("./html"))

	// Host streaming audio service
	log.Fatal(http.ListenAndServe(HostPort, fs))
}
