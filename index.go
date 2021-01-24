/* https://github.com/mezcel/fs-path/index.go */

// About: A Golang file server hosting M3U or Html5 streaming audio.
// git repo: https://github.com/mezcel/fs-path.git
package main

import (
	"fmt"
	"io/ioutil"
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
	// Array for the track names in the /html/audio directory
	TrackArray       []string
	Ip               string
	HostPort         string
	WorkingDirectory string
	JsPlaylistPath   string
	M3uPlaylistPath  string
}

/* Global variables */
var (
	fsStructs FsStruct
)

// Make an array of the list of items in the audio directory
func PopulateFilesArray() {
	// Nill any existing track array values
	fsStructs.TrackArray = nil

	trackDirectory := "html/audio"
	err := filepath.Walk(trackDirectory, func(path string, info os.FileInfo, err error) error {
		fsStructs.TrackArray = append(fsStructs.TrackArray, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\n Tracks loaded. Track count:", len(fsStructs.TrackArray))
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

	//fmt.Println("\tFile Created Successfully.")
}

// Delete the JS HTML5 Audio playlist script. Used prior to writing a new playlist.
func DeleteTextFile(textfilePath string) {
	// delete file
	var err = os.Remove(textfilePath)
	if err != nil {
		return
	}

	//fmt.Println("\tFile Deleted Successfully.")
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
		trackPath = "html/audio/" + trackName
		javascriptString += "AddListItem(\"" + trackName + "\", \"" + trackPath + "\");\n"
	}

	javascriptString += "\n/* Load the html3 audio playlist */\naudioPlayer();\n\n"

	return javascriptString
}

// Generate M3U script listing file paths in the audio directory
func GenerateM3UScript(Ip string, HostPort string) string {

	var (
		m3uString string
		trackName string
		trackPath string

		// fsStructs.TrackArray was a global struct var
	)

	m3uString = "#EXTM3U\n"

	// load track paths into a js script
	for i := 1; i < len(fsStructs.TrackArray); i++ {
		trackName = filepath.Base(fsStructs.TrackArray[i])
		trackPath = trackName

		m3uString += "#EXTINF:0," + trackName + "\n"
		m3uString += "http://" + Ip + HostPort + "/html/audio/" + trackPath + "\n"
	}

	m3uString += "\n#M3U generated at: " + time.Now().String()

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

	fmt.Println("\tJS File Updated Successfully.")
}

// Make an M3U playlist file
func WriteM3UPlaylist(m3uPlaylistPath string, Ip string, HostPort string) {

	var m3uString string = GenerateM3UScript(Ip, HostPort)

	DeleteTextFile(m3uPlaylistPath)
	MakeTextFile(m3uPlaylistPath)
	//m3uString = GenerateM3UScript(Ip)

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

	fmt.Println("\tM3U File Updated Successfully.")

}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	var returnIp string = "127.0.0.1" // localhost

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				returnIp = ipnet.IP.String()
				return returnIp
			}
		}
	}

	return returnIp
}

// Tty About and instructions splash greeter
func TtyGreeter(WorkingDirectory string, Ip string, HostPort string) {

	var (
		hostUrl        string = "http://" + Ip + HostPort
		m3uUrl         string = hostUrl + "/M3U/playlist.m3u"
		audioDirectory string = WorkingDirectory + "/html/audio"
	)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("## ############################################################################")
	fmt.Println("## fs-path\n##")
	fmt.Println("## Host Server:", hostname)
	fmt.Println("## ")
	fmt.Println("## About:")
	fmt.Println("##\tHost streaming audio on a Golang file server. ( M3U or HTML5 Audio )\n##")
	fmt.Println("## Instructions:\n##")
	fmt.Println("##    * Play Option 1: (Web Page)")
	fmt.Println("##       Launch a web browser and enter the following URL.")
	fmt.Println("##        >", hostUrl, "\n##")
	fmt.Println("##    * Play Option 2: (Net Radio)")
	fmt.Println("##       Launch a media player, like VLC, and run the M3U file.")
	fmt.Println("##        >", m3uUrl, "\n##")
	fmt.Println("##    * (Upload) audio to server playlist")
	fmt.Println("##       Place individual music into the following server directory.")
	fmt.Println("##        >", audioDirectory)
	fmt.Println("## ############################################################################")
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	var (
		tempFilename string
		jsString     string
	)

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Fprintf(w, "Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf(" Uploaded File:\t%+v\n", handler.Filename)
	fmt.Printf(" File Size:\t\t%+v\n", handler.Size)
	//fmt.Printf("MIME Header:\t%+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	tempFilename = handler.Filename + "_"
	tempFile, err := ioutil.TempFile("html/audio", tempFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// return that we have successfully uploaded our file!
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<center><h4>Successfully uploaded file.</h4>Go back and refresh.</center>")

	// Reload DOM and reset navigation to home view
	jsString = "<script> var linkUrl = window.location.href; var dirUrl = linkUrl.split(\"html/audio\"); window.location.href = dirUrl[0]; </script> "
	fmt.Fprintf(w, jsString)

	fmt.Printf("Refreshing server track list ...")
	PopulateFilesArray()
	WriteJsPlaylist(fsStructs.JsPlaylistPath)
	WriteM3UPlaylist(fsStructs.M3uPlaylistPath, fsStructs.Ip, fsStructs.HostPort)
	fmt.Println("Updated M3U and JS Playlist\n")
}

/* Serve local file path in goserver */
func ServeFiles(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	if p == "./" {
		p = "./html/index.html"
	}

	http.ServeFile(w, r, p)
}

// Initialize the server resource path strings
func InitializeServerPath() {
	// Get working directory path
	WorkingDirectory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	// Update file paths
	fsStructs.HostPort = ":8080"
	fsStructs.Ip = GetLocalIP()
	fsStructs.WorkingDirectory = WorkingDirectory
	fsStructs.JsPlaylistPath = WorkingDirectory + "/html/js/jsPlaylist.js"
	fsStructs.M3uPlaylistPath = WorkingDirectory + "/html/M3U/playlist.m3u"
}

// Main()
func main() {
	// Initialize the server resource path strings
	InitializeServerPath()

	// Display a TTY prompt message regarding file server usage
	TtyGreeter(fsStructs.WorkingDirectory, fsStructs.Ip, fsStructs.HostPort)

	// Array of files in the audio directory
	PopulateFilesArray()

	fmt.Println("\n Generating:", fsStructs.JsPlaylistPath)
	// Write a js script to dynamically add track to the html audio ol list
	WriteJsPlaylist(fsStructs.JsPlaylistPath)

	fmt.Println("\n Generating:", fsStructs.M3uPlaylistPath)
	// Generate a standalone W3M streaming audio player file
	WriteM3UPlaylist(fsStructs.M3uPlaylistPath, fsStructs.Ip, fsStructs.HostPort)

	fmt.Println(" ---\n ( Pres \"Ctrl+c\" to terminate server )\n")

	http.HandleFunc("/html/audio", UploadFile)
	http.HandleFunc("/", ServeFiles)

	// File Server
	//fs := http.FileServer(http.Dir("html"))

	// Host streaming audio service
	//log.Fatal(http.ListenAndServe(HostPort, fs))
	log.Fatal(http.ListenAndServe(fsStructs.HostPort, nil))
}
