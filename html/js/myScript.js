/*var a = document.createElement("a");
var ulist = document.getElementById("playlist");
var newItem = document.createElement("li");

a.textContent = "ghost";
a.setAttribute('href', "audio/ghost.mp3");
newItem.appendChild(a);
ulist.appendChild(newItem);*/

// loads the audio player
//audioPlayer();


/* Dynamically adds a list element to the HTML5 Audio Player */
function AddListItem(trackname, trackpath){
    var a = document.createElement('a');
    var ulist = document.getElementById('playlist');
    var newItem = document.createElement('li');

    a.textContent = trackname;
    a.setAttribute('href', trackpath);
    a.classList.add("w3-btn");
    a.classList.add("w3-round");
    newItem.appendChild(a);
    ulist.appendChild(newItem);

    console.log(newItem.TEXT_NODE);
}

/* Toggle the diplay of the the playlist */
function TogglePlaylistDisplay() {
    var x = document.getElementById("playlist");
    if (x.style.display === "none") {
        x.style.display = "block";
    } else {
        x.style.display = "none";
    }
}

/* Set the next random track on the link list */
function NextTrack() {
    var listLength = $("#playlist li a").length - 1
    var randomTrack = Math.floor(Math.random() * listLength);
    var currentSong = randomTrack

    return currentSong
}