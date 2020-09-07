# fs-path

## 1.0 ABOUT

A Golang file server hosting html5 streaming audio. <img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" height="16"> <img src="https://www.w3.org/html/logo/badge/html5-badge-h-css3-device-multimedia.png" height="16">

---

## 2.0 PERSONAL SERVER

* Download: [golang.org](https://golang.org/dl/)
* Install:
    * ```git clone https://github.com/mezcel/fs-path.git ~/github/mezcel/fs-path.git```
* Run:
    ```sh
    ## launch web server

    cd ~/github/mezcel/fs-path.git

    go run index.go
    ```

### 2.1 SERVER LOCATION

* Go server Port 8080 was set in the golang script
* SSH port 22 is set by the server's system admin.
    * Typically it is set as available by default on non-configured systems.

| host | url | ssh |
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

### 3.3 SCP (ssh)

* Manually upload audio

```sh
## Dir
#scp -r <local-dir> mezcel@10.42.0.1:~/github/mezcel/fs-path.git/html/audio

## File
scp <my-local-file> mezcel@10.42.0.1:~/github/mezcel/fs-path.git/html/audio
```
