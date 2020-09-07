/**
 * Original source code refference (Github):
 *  https://github.com/NelsWebDev/HTML5AudioPlaylist.git
 * My modifications are minor.
 * */

/* Set the next random track on the link list */
function NextTrack() {
    var listLength = $("#playlist li a").length
    var randomTrack = Math.floor(Math.random() * listLength);
    var currentSong = randomTrack

    return currentSong
}

/* Display track name in dom */
function DisplayCurrentTrackName( objectHTMLAudioElement ) {
    var linkText = objectHTMLAudioElement.src
    $("#trackName").text(linkText);
}

/* HTML5 Audio controller */
function audioPlayer(){
    var currentSong = 0;

    $("#audioPlayer")[0].src = $("#playlist li a")[0];
    DisplayCurrentTrackName( $("#audioPlayer")[0] ); // display track url name in dom
    $("#audioPlayer")[0].play();

    $("#playlist li a").click(function(e){
       e.preventDefault();
       $("#audioPlayer")[0].src = this;
       $("#audioPlayer")[0].play();
       $("#playlist li").removeClass("current-song");
        currentSong = $(this).parent().index();
        $(this).parent().addClass("current-song");
    });

    $("#audioPlayer")[0].addEventListener("ended", function(){

        /*currentSong++;
        if(currentSong == $("#playlist li a").length) {
            currentSong = 0;
        }*/

        currentSong = NextTrack(); // randomizer from myScript.js

        $("#playlist li").removeClass("current-song");
        $("#playlist li:eq("+currentSong+")").addClass("current-song");
        $("#audioPlayer")[0].src = $("#playlist li a")[currentSong].href;
        DisplayCurrentTrackName( $("#audioPlayer")[0] ); // display track url name in dom
        $("#audioPlayer")[0].play();
    });
}