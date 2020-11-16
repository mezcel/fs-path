# fs-path

## 1.0 ABOUT

Host streaming audio on a Golang file server. ( [M3U](https://wiki.videolan.org/M3U/) or Web Browser ) <img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" height="16"> <img src="https://www.w3.org/html/logo/badge/html5-badge-h-css3-device-multimedia.png" height="16"> <img src="https://images.videolan.org/images/VLC-IconSmall.png" height="16">

---

## 2.0 PORTABLE SERVER

* Download: [golang.org](https://golang.org/dl/)
* Install: ```git clone https://github.com/mezcel/fs-path.git ~/github/mezcel/fs-path.git```
* Run:
    ```sh
    ## launch web server
    cd ~/github/mezcel/fs-path.git
    go run index.go

    ## Build options:
    ## step 1:   # go build index.go     # go build      # go build -o <desired executable destination>/fs-path
    ## step 2:   # ./index.exe           # ./fs-path     # ./fs-path
    ```
* Play:
    | Host Server IP (example) | Streaming Radio Client, like VLC | Web Browser, like Edge |
    | :---: | :---: | :---: |
    | ```10.42.0.1``` | ```10.42.0.1:8080/M3U/playlist.m3u``` | ```10.42.0.1:8080``` |

### 2.1 SERVER LOCATION

* Go server Port 8080 was set within the Golang script.

### 2.2 REMOTE SERVER ( Optional yet preferred )

* The use case for this project is targeted toward playing remotely hosted audio.
* SSH port 22 is set by the server's system admin.
    * Typically it is set as available by default on non-configured systems.

| host (example) | url (example) | ssh (example) |
| --- | --- | --- |
| ```ssh -p 22 mezcel@10.42.0.1``` | ```10.42.0.1:8080``` | ```mezcel@10.42.0.1:~/github/mezcel/fs-path.git``` |

---

## 3.0 AUDIO LIBRARY:

* Remotely update the audio files on the host server

### 3.1 GIT

* git push the files you want the server to have

```sh
## Clone server repo into client repo
## In my case, mezcel@10.42.0.1 is a user on a linux ad-hoc server set at 10.42.0.1

git clone mezcel@10.42.0.1:~/github/mezcel/fs-path.git ~/github/mezcel/fs-path.git
```

### 3.2 AUDIO FILES

#### 3.2.1 Uploading resource:

1. Make a new git branch, on client side, which contains music.
2. Then git push to the server and then git checkout on the server machine.
* audio directory: ```~/github/mezcel/fs-path.git/html/audio```


#### 3.2.2 ```.gitignore```

* Review this git repo's ```.gitignore```
* I set it to ignore all file in the audio directory to prevent audio uploads to Github.
    * remove the ignore script in production use
    * Playlist and Audio files will not be uploaded back to Github.com.

### 3.3 SCP (ssh)

* Manually upload audio

```sh
## Dir
#scp -r <local-dir> mezcel@10.42.0.1:~/github/mezcel/fs-path.git/html/audio

## File
scp <my-local-file> mezcel@10.42.0.1:~/github/mezcel/fs-path.git/html/audio
```

---

## 4.0 Splash Audio

* Demo audio and intro sounds were generated using the ( [Watson Text to Speech](https://text-to-speech-demo.ng.bluemix.net/?_ga=2.149277174.1746788865.1577973300-883782623.1576869895&cm_mc_uid=15278110739115689857415&cm_mc_sid_50200000=20950731577973297095&cm_mc_sid_52640000=33641591577973297117) / [IBM](https://www.ibm.com/cloud/watson-text-to-speech?p1=Search&p4=43700051010023756&p5=b&cm_mmc=Search_Google-_-1S_1S-_-WW_NA-_-%2Btext%20%2Bto%20%2Bspeech_b&cm_mmca7=71700000062156796&cm_mmca8=aud-309367918490:kwd-18391235536&cm_mmca9=EAIaIQobChMIvLr8y_rW6wIVAtvACh1XXwtwEAAYASAAEgIQFPD_BwE&cm_mmca10=412803414889&cm_mmca11=b&gclid=EAIaIQobChMIvLr8y_rW6wIVAtvACh1XXwtwEAAYASAAEgIQFPD_BwE&gclsrc=aw.ds) ).
* Plan on adding in additional effects just for jollies and personalization. Aiming to make the transitions sound more seamless.

### 4.1 Rendering Scripts

* American English (en-US): AllisonV3 (female, enhanced dnn)
    ```xml
    f<p>s<s>path</s></p>
    ```
* American English (en-US): AllisonV3 (female, expressive, transformable)
    ```xml
    <speak>ef es, <express-as type="GoodNews">path </express-as></speak>
    ```
* American English (en-US): AllisonV3 (female, expressive, dnn)
    ```xml
    <p><s>eff ess,<prosody rate="-15%"> path.</prosody></s></p>
