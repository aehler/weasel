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
                        //parse grid
                        break;
                    case "page" :
                        //parse page
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

    }

};
