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

    // replace the content of the dropdown surface.
    document.querySelector("#drag_text").innerHTML = 'Looking for the best subtitles...';

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
    xhr.open("GET", "/api/1.0/search?f="+file.name, true);
    xhr.send();
    xhr.onreadystatechange = function (event) { readApiResponse(event, xhr); }
}

// Reads the API response of a search request.
function readApiResponse(event, xhr) {
    "use strict";
    if (xhr.readyState == 4) {
        try {
            var resp = JSON.parse(xhr.responseText)
            if (resp.subtitles != undefined) {
                renderResponse(resp.subtitles, resp.metadata);
            }
        } catch (exception) {
            // TODO
            console.error("Parsing error:", exception);
        }
    }
}

// Display the given subtitles.
function renderResponse(subtitles, metadata) {
    // First, hide the dropdown surface.
    document.querySelector("#drag_surface").style.display = 'none'; 

    var container = document.querySelector("#results_container");
    var metadataContainer = document.querySelector("#metadata_container");

    // Empty the container
    container.innerHTML = '';
    metadataContainer.innerHTML = '';

    // No result
    if (subtitles == null || subtitles.length == 0) {
        container.innerHTML = '<pre>No results found.</pre>';
    }
    
    // gets the template
    var template = document.querySelector("#subtitle_template").innerHTML;

    var content = '';
    // Applies the template on each result.
    for (var i = 0; i < subtitles.length; i++) {
        var compiled_html = _.template(template)({
            url: subtitles[i].download_link,
            name: subtitles[i].filename,
            filename_score: subtitles[i].filename_score
        });
        content += compiled_html; 
    }
    container.innerHTML = content;

    // metadata
    if (metadata != null) {
        var mTemplate = document.querySelector("#metadata_template").innerHTML;
        console.log(metadata);
        var metadataContent = _.template(mTemplate)({
            image: metadata.image
        });
        metadataContainer.innerHTML = metadataContent;
    }
}
