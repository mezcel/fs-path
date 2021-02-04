/* https://github.com/mezcel/fs-path/html/js/myScript.js */

/* Dynamically adds a list element to the HTML5 Audio Player */
function AddListItem( trackname, trackpath ){
    var a = document.createElement( 'a' );
    var ol = document.getElementById( 'playlist' );
    var li = document.createElement( 'li' );
    var button = document.createElement( 'button' );

    trackname = trackname.replace( /%20/g, " " );
    //trackname = decodeURI( trackname );
    //trackname = unescape( trackname );

    // Delete button
    button.textContent = "delete";
    button.setAttribute( 'style', "width:20%" );
    button.classList.add( "w3-bar-item" );
    button.classList.add( "w3-red" );
    button.classList.add( "w3-hover-pale-red" );
    button.classList.add( "w3-border" );
    button.classList.add( "w3-button" );
    button.classList.add( "w3-round" );
    button.classList.add( "btnDelete" );

    // Track button
    a.textContent = trackname;
    a.setAttribute( 'href', trackpath );
    a.setAttribute( 'style', "width:75%" );
    a.classList.add( "w3-bar-item" );
    a.classList.add( "w3-white" );
    a.classList.add( "w3-hover-blue" );
    a.classList.add( "w3-border" );
    a.classList.add( "w3-button" );
    a.classList.add( "w3-round" );

    // list bar row
    li.classList.add( "w3-bar" );

    // Append nested elements
    li.appendChild( button );    // delete button
    li.appendChild( a );         // track button
    ol.appendChild( li );

    //console.log( newItem.TEXT_NODE );
}

/* Toggle the diplay of the the html audio element */
function ToggleAudio() {
    var audioPlayer = document.getElementById( "audioPlayer" );

    if (  audioPlayer.style.display === "none"  ) {
        audioPlayer.style.display = "block";
        audioPlayer.pause();
    } else {
        audioPlayer.style.display = "none";
        audioPlayer.play();
    }
}

/* Toggle the display of the the playlist */
function TogglePlaylistDisplay() {
    var t = document.getElementById( "deleteModal" );
    var u = document.getElementById( "topLeftContainer" );
    var v = document.getElementById( "topRightContainer" );
    var w = document.getElementById( "mainPlayer" );
    var x = document.getElementById( "playlistModal" );
    var y = document.getElementById( "aboutModal" );
    var z = document.getElementById( "uploadModal" );

    // delete track modal
    if ( t.style.display === "block" ) {
        t.style.display = "none";
    }

    // hide about modal
    if ( y.style.display === "block" ) {
        y.style.display = "none";
    }

    // hide upload modal
    if ( z.style.display === "block" ) {
        z.style.display = "none";
    }

    // view playlist modal
    if ( x.style.display === "none" ) {
        x.style.display = "block";
        u.style.display = "none";
        v.style.display = "none";
        w.style.display = "none";
    } else {
        x.style.display = "none";
        u.style.display = "block";
        v.style.display = "block";
        w.style.display = "block";
    }

}

/* Toggle the display of the the About */
function ToggleAboutDisplay() {
    var t = document.getElementById( "deleteModal" );
    var u = document.getElementById( "topLeftContainer" );
    var v = document.getElementById( "topRightContainer" );
    var w = document.getElementById( "mainPlayer" );
    var x = document.getElementById( "playlistModal" );
    var y = document.getElementById( "aboutModal" );
    var z = document.getElementById( "uploadModal" );

    // delete track modal
    if ( t.style.display === "block" ) {
        t.style.display = "none";
    }

    // hide playlist modal
    if ( x.style.display === "block" ) {
        x.style.display = "none";
    }

    // hide upload modal
    if ( z.style.display === "block" ) {
        z.style.display = "none";
    }

    // view about modal
    if ( y.style.display === "none" ) {
        y.style.display = "block";
        u.style.display = "none";
        v.style.display = "none";
        w.style.display = "none";
    } else {
        y.style.display = "none";
        u.style.display = "block";
        v.style.display = "block";
        w.style.display = "block";
    }

}

/* Toggle upload modal */
function ToggleUploadDisplay() {
    var t = document.getElementById( "deleteModal" );
    var u = document.getElementById( "topLeftContainer" );
    var v = document.getElementById( "topRightContainer" );
    var w = document.getElementById( "mainPlayer" );
    var x = document.getElementById( "playlistModal" );
    var y = document.getElementById( "aboutModal" );
    var z = document.getElementById( "uploadModal" );

    // delete track modal
    if ( t.style.display === "block" ) {
        t.style.display = "none";
    }

    // hide playlist modal
    if ( x.style.display === "block" ) {
        x.style.display = "none";
    }

    // hide about modal
    if ( y.style.display === "block" ) {
        y.style.display = "none";
    }

    // view upload modal
    if ( z.style.display === "none" ) {
        var homeUrl = window.location.href;
        FormActionUrl( "uploadForm", homeUrl + "upload" );

        z.style.display = "block";
        u.style.display = "none";
        v.style.display = "none";
        w.style.display = "none";

    } else {
        z.style.display = "none";
        u.style.display = "block";
        v.style.display = "block";
        w.style.display = "block";
    }

}

/* Set the next random track on the link list */
function NextTrack() {
    var listLength = $( "#playlist li a" ).length - 1
    var randomTrack = Math.floor( Math.random() * listLength );
    var currentSong = randomTrack

    return currentSong
}

/* toggle light/dark mode */
function ToggleDarkmode() {
    var body = document.body;
    body.classList.toggle( "dark-mode" );
}

/* update the form's action's ip url */
function FormActionUrl( formName, eventUrl ) {
    document.getElementById( formName ).action = eventUrl;
}

/* Hide all modals */
function HideModals() {
    var t = document.getElementById( "deleteModal" );
    var u = document.getElementById( "topLeftContainer" );
    var v = document.getElementById( "topRightContainer" );
    var w = document.getElementById( "mainPlayer" );
    var x = document.getElementById( "playlistModal" );
    var y = document.getElementById( "aboutModal" );
    var z = document.getElementById( "uploadModal" );

    // show mainPlayer
    u.style.display = "block";
    v.style.display = "block";
    w.style.display = "block";

    // delete track modal
    if ( t.style.display === "block" ) {
        t.style.display = "none";
    }

    // hide playlist
    if ( x.style.display === "block" ) {
        x.style.display = "none";
    }

    // hide about
    if ( y.style.display === "block" ) {
        y.style.display = "none";
    }

    // hide upload
    if ( z.style.display === "block" ) {
        z.style.display = "none";
    }
}

/* DOM keybindings */
document.onkeyup = function( e ) {

    var keyPress = e.which; // depreciated ???

    switch( keyPress ) {

        case 27: // ESC
            // hide about and playlist if visible
            HideModals();
            break;

        case 81: // q
            document.getElementById( "audioPlayer" ).pause();
            window.close();
            break;

        case 80: // p key
            // Toggle pause and track progress display
            ToggleAudio();
            break;

        case 16: // shift key
            // Toggle Light/Dark Color Theme
            ToggleDarkmode();
            break;

        case 32: // spacebar
            // toggle playlist
            TogglePlaylistDisplay();
            break;

        case 73: // i
            // toggle about/info
            ToggleAboutDisplay();
            break;

        case 85: // u
            // toggle upload
            ToggleUploadDisplay();
            break;
        default:
            console.log( "keyPress:",keyPress );
    }
}

/* Toggle delete track modal Jquery */
$( document ).on( 'click', '.btnDelete', function( even ) {
    // depends on jquery prev() or next()
    var t = document.getElementById( "deleteModal" );
    var u = document.getElementById( "topLeftContainer" );
    var v = document.getElementById( "topRightContainer" );
    var w = document.getElementById( "mainPlayer" );
    var x = document.getElementById( "playlistModal" );
    var y = document.getElementById( "aboutModal" );
    var z = document.getElementById( "uploadModal" );

    // get the element text, next to clicked element
    //var trackText = $( this ).closest( 'button' ).prev( 'a' ).text(); // left element
    var trackText = $( this ).closest( 'button' ).next( 'a' ).text(); // right element

    if ( t.style.display === "none" ) {
        var homeUrl = window.location.href;
        FormActionUrl( "deleteForm", homeUrl + "delete" );

        t.style.display = "block";
        u.style.display = "none";
        v.style.display = "none";
        w.style.display = "none";
        x.style.display = "none";
        y.style.display = "none";
        z.style.display = "none";
    }

    document.getElementById( "inputDelString" ).value = trackText;
    console.log( "trackText:", trackText );
} );