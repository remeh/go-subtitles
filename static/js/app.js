// go-subtitles webapp.
// 
// Easily find the best subtitles for your video file.
//
// Copyright © 2015 - Rémy MATHIEU

// On DOM ready.
document.addEventListener("DOMContentLoaded", function(){
    "use strict";
    prepareDragSurface();
});

// prepareDragSurface is called when the DOM is ready
// to prepare the drag surface.
function prepareDragSurface() {
    var dragSurface =  document.querySelector("#drag_surface");
    dragSurface.addEventListener("dragover", handleDragOver, false);
    dragSurface.addEventListener("dragenter", handleDragEnter, false);
    dragSurface.addEventListener("drop", handleDrop, false);
}

// We must cancel this event.
function handleDragOver(event) {
    "use strict";
    event.preventDefault();
    return false;
}
// We must cancel this event.
function handleDragEnter(event) {
    "use strict";
    event.preventDefault();
    return false;
}

// handleDrop is called when files are dropped onto
// the drop surface.
function handleDrop(event) {
    event.preventDefault(); // We don't want the browser to load the file.

    var dt    = event.dataTransfer;
    var files = dt.files;

    // TODO Many files
    if (files.length == 0) {
        return;
    }

    // Launch the interesting part.
    findBestSubtitle(files[0]);
}

// findBestSubtitle is calling the API to
// retrieve the best subtitle for the given filename,
// the response of the API is handled in readApiResponse.
function findBestSubtitle(file) {
    "use strict";
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/1.0/get?f="+file.name, true);
    xhr.send();
    xhr.onreadystatechange = function (event) { readApiResponse(event, xhr); }
}



function readApiResponse(event, xhr) {
    "use strict";
    if (xhr.readyState == 4) {
        // TODO
    }
}
