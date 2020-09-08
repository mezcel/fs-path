/**
 * Original source code refference (Github):
 *  https://github.com/NelsWebDev/HTML5AudioPlaylist.git
 *
 * I added some decorative functions and passed some strings arround.
 * */

/* Set the next random track on the link list */
function NextTrack() {
    /*currentSong++;
    if(currentSong == $("#playlist li a").length) {
        currentSong = 0;
    }*/

    var listLength = $("#playlist li a").length
    var randomTrack = Math.floor(Math.random() * listLength);
    var currentSong = randomTrack

    return currentSong
}

/* Display track name in dom */
function DisplayCurrentTrackName( objectHTMLAudioElement ) {
    // file path
    var linkUrl = objectHTMLAudioElement.src
    var dirUrl = linkUrl.split("/audio");
        dirUrl = dirUrl[0] + "/audio";

    $("#serverUrl").text(dirUrl);

    // basae name
    var linkText = linkUrl.substring(linkUrl.lastIndexOf('/')+1);


    linkText = linkText.replace(/%20/g, " ");
    //trackname = decodeURI(trackname);
    //trackname = unescape(trackname);

    if ( linkText == "f-s-path0-fx.mp3") { linkText = "(splash sound) 🌊"; } // decorate splash intro track
    $("#trackName").text(linkText);
}

/* HTML5 Audio controller */
function audioPlayer(){
    var currentSong = 0;

    $("#audioPlayer")[0].src = $("#playlist li a")[0];
    $("#audioPlayer")[0].play();
    DisplayCurrentTrackName( $("#audioPlayer")[0] ); // display track url name in dom

    $("#playlist li a").click(function(e){
       e.preventDefault();
       $("#audioPlayer")[0].src = this;
       $("#audioPlayer")[0].play();
       DisplayCurrentTrackName( $("#audioPlayer")[0] ); // display track url name in dom
       $("#playlist li").removeClass("current-song");
        currentSong = $(this).parent().index();
        $(this).parent().addClass("current-song");
    });

    $("#audioPlayer")[0].addEventListener("ended", function(){

        currentSong = NextTrack(); // randomizer from myScript.js

        $("#playlist li").removeClass("current-song");
        $("#playlist li:eq("+currentSong+")").addClass("current-song");
        $("#audioPlayer")[0].src = $("#playlist li a")[currentSong].href;
        $("#audioPlayer")[0].play();
        DisplayCurrentTrackName( $("#audioPlayer")[0] ); // display track url name in dom
    });

    $("#audioPlayer")[0].onplaying = function() {
        var isPlaying = true;
    };

    $("#audioPlayer")[0].onpause = function() {
        var isPlaying = false;
    };
}