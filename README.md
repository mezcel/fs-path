# fs-path

## 1.0 ABOUT

A Golang file server hosting html5 streaming audio. <img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" height="16"> <img src="https://www.w3.org/html/logo/badge/html5-badge-h-css3-device-multimedia.png" height="16">

---

## 2.0 PERSONAL SERVER

* Download: [golang.org](https://golang.org/dl/)
* Install:
    * ```git clone https://github.com/mezcel/fs-path.git ~/github/mezcel/fs-path.git```

### 2.1 SERVER LOCATION

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

git clone mezcel@10.42.0.1:~/github/mezcel/fs-path.git ~/github/mezcel/fs-path.git
```

* make a new branch with music, and git git checkout on the server machine \
* audio directory: ```~/github/mezcel/fs-path.git/html/audio```

### 3.2 SCP (ssh)

* Manually upload audio

```sh
## Dir
#scp -r <local-dir> mezcel@10.42.0.1:~/github/mezcel/fs-path.git/html/audio

## File
scp <my-local-file> mezcel@10.42.0.1:~/github/mezcel/fs-path.git/html/audio
```
