var Weasel = {

    parseContent : function() {

        if(typeof content.error != 'undefined') {

            Weasel.RenderError(content.error);
        }

        for (var property in content) {

            if (content.hasOwnProperty(property)) {

                switch (property) {
                    case  "form" :
                        Weasel.parseForm(content[property]);
                        break;
                    case  "grid" :
                        Weasel.parseGrid(content[property]);
                        break;
                    case "page" :
                        //parse page
                        break;
                    case "message" :

                        Weasel.messenger();

                        break;
                    default :
                        // Do nothing
                        break;
                }

            }

        }

    },

    RenderError : function(message) {

        $("#simpleModal").find(".modal-body").html($("#error-template").html()).find(".message-container").html(message);

        $("#simpleModal").modal("toggle");

    },

    parseForm : function(formData) {

        var parsed = ParseForm(formData);

        $("div.body").append(parsed);

    },

    parseGrid : function(gridData) {

        $(document).append('<script type="text/template" id="table-template" src="/static/include/layout/table.html"></script>');

    },

    messenger : function(context) {

        var m = $("#messageModal").find(".alert");

        m.html("");

        m.removeClass("alert-info");
        m.removeClass("alert-danger");

        m.append('<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>');

        m.append(context.label);

        switch (context.message) {

            case "success" :

                m.addClass("alert-info");

                $("#messageModal").modal("toggle");

                break;

            case "fail" :

                m.addClass("alert-danger");

                $("#messageModal").modal("toggle");

                break;

            default :
                // Do nothing
                break;
        }



        setTimeout(function(){$("#messageModal").modal("hide");}, 5000);

    }

};
