$(document).ready(function() {

    //$("a").on("click", function(event){
    //
    //    event.preventDefault();
    //    event.stopPropagation();
    //
    //    linkClickHandler($(event.currentTarget).attr("href"));
    //
    //    return false;
    //});
    //
    //window.addEventListener("popstate", function() {
    //    popstate(location.pathname);
    //}, false);

    Weasel.parseContent();

});

function linkClickHandler(url) {

    var responseText = $.ajax({
        type: "GET",
        dataType: 'json',
        cache: false,
        async: false,
        url: url
    }).always(function () {
        history.pushState(null,null, url);
    }).fail(function () {
        $("#block-content").html("<h1>HTTP ERROR</h1>");
    }).responseText;

    var result;

    try {
        result = JSON.parse(responseText);
    }
    catch (ex) {
        $("#block-content").html("<h1>ERROR</h1><p>"+responseText+"</p>");

        return;
    }

}

function popstate(url) {

    var responseText = $.ajax({
        type: "GET",
        dataType: 'json',
        cache: false,
        async: false,
        url: url
    }).always(function () {

    }).fail(function () {
        $("#block-content").html("<h1>HTTP ERROR</h1>");
    }).responseText;

    var result;

    try {
        result = JSON.parse(responseText);
    }
    catch (ex) {
        $("#block-content").html("<h1>ERROR</h1><p>"+responseText+"</p>");

        return;
    }

}