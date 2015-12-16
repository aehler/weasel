$(document).ready(function() {

    $('.selectpicker').selectpicker();
    //selectpicker doesn't seem to be flexible enough (can't change template), so need to replace span.caret externally
    $('.selectpicker + .bootstrap-select span.caret').replaceWith("<i class='fa fa-caret-down'></i>");
    $('.selectpicker + .bootstrap-select span.pull-left').removeClass("pull-left");

    $(".jsForm").on("click", function(event){

        event.preventDefault();
        event.stopPropagation();

        linkClickHandler($(event.currentTarget).attr("href"));

        return false;
    });

    window.addEventListener("popstate", function() {
        popstate(location.pathname);
    }, false);

//    Weasel.parseContent();

    $("form").on("submit", function(e){

        e.stopPropagation();
        e.preventDefault();

        var responseText = $.ajax({
            type: "POST",
            cache: false,
            async: true,
            url: $(this).attr("action"),
            data: $(this).serializeArray()
        }).always(function () {

        }).fail(function () {
            Weasel.RenderError("<h1>HTTP ERROR</h1>");
        }).success(function(data){

            //Do noting
            if(data == "") {

                return false;
            }

            result = data;

            if (result.hasOwnProperty("redirect") ) {

                window.location.href = result.redirect;

                return

            }

            if (result.hasOwnProperty("loginError") ) {

                Weasel.RenderError("<h1>Ошибка входа в систему</h1><p>"+result.loginError+"</p>");

                return

            }

            if (result.hasOwnProperty("Error") ) {

                $("#formModal").modal("toggle");

                Weasel.RenderError("<h1>ERROR</h1><p>"+result.Error+"</p>");

                return

            }

            //В любой непонятной ситуации - рефреш
            window.location.href = window.location.href;

        });

        return false;

    });

});

function linkClickHandler(url) {

    JSForm.New(url);

}