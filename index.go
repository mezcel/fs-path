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

        m3uString += "#EXTINF:0," + trackName + "\n"
        m3uString += "../audio/" + trackPath + "\n"
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

    fmt.Println("\tFile Updated Successfully.",)
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

    fmt.Println("\tFile Updated Successfully.")

}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

// Tty About and instructions splash greeter
func TtyGreeter(ParentDirectory string, Ip string, HostPort string) {

    var (
        hostUrl string = "http://" + Ip + HostPort
        m3uUrl string = hostUrl + "/M3U/playlist.m3u"
        audioDirectory string = ParentDirectory + "/html/audio"
    )

    fmt.Println("## ############################################################################")
    fmt.Println("## fs-path\n##")
    fmt.Println("## About:")
    fmt.Println("##\tHost streaming audio on a Golang file server. ( M3U or HTML5 Audio )\n##")
    fmt.Println("## Instructions:\n##")
    fmt.Println("##    * Play Option 1: (Web Page)", )
    fmt.Println("##       Launch a web browser and enter the following URL.")
    fmt.Println("##        >", hostUrl, "\n##")
    fmt.Println("##    * Play Option 2: (Net Radio)", )
    fmt.Println("##       Launch a media player, like VLC, and run the M3U file.")
    fmt.Println("##        >", m3uUrl, "\n##")
    fmt.Println("##    * (Upload) audio to server playlist", )
    fmt.Println("##       Place individual music into the following server directory.")
    fmt.Println("##        >", audioDirectory)
    fmt.Println("## ############################################################################")
}

// Main()
func main() {
    var (
        Ip               string
        HostPort         string = ":8080"
        JsPlaylistPath   string = "/html/js/jsPlaylist.js"
        M3uPlaylistPath  string = "/html/M3U/playlist.m3u"
    )

    // Get working directory path
    WorkingDirectory, err := os.Getwd()

    if err != nil {
        panic(err)
    }

    Ip = GetLocalIP()
    JsPlaylistPath = WorkingDirectory + JsPlaylistPath
    M3uPlaylistPath = WorkingDirectory + M3uPlaylistPath

    // Display a TTY prompt message regarding file server usage
    TtyGreeter(WorkingDirectory, Ip, HostPort)

    // Array of files in the audio directory
    PopulateFilesArray()

    fmt.Println("\n Generating:", JsPlaylistPath)
    // Write a js script to dynamically add track to the html audio ol list
    WriteJsPlaylist(JsPlaylistPath)

    fmt.Println("\n Generating:", M3uPlaylistPath)
    // Generate a standalone W3M streaming audio player file
    WriteM3UPlaylist(M3uPlaylistPath)

    fmt.Println(" ---\n ( Pres \"Ctrl+c\" to terminate server )\n")

    // File Server
    fs := http.FileServer(http.Dir("./html"))

    // Host streaming audio service
    log.Fatal(http.ListenAndServe(HostPort, fs))
}
