$.fn.extend({
    wTree : function(options) {

        options.Target = $(this);
        options.node = 0;

        if (typeof options.selected != "function") {

            options._selectedParam = options.selected;

            options.selected = WTree._selected;
        }

        WTree.opts = options;

        WTree.Load();
    }
});

WTree = {

    opts : {},

    data : {},

    _selected : function(data) {

        return WTree.opts._selectedParam == data.Id;

    },

    Load : function() {

        $.getJSON(WTree.opts.url, {"node" : WTree.opts.node}, function(res) {

            if(typeof res.error != 'undefined') {

                alert(res.error);

                return;
            }

            WTree.populate(WTree.opts.node, res);

            $(".wtree-toggle").bind("click", function(){

                if ($(this).hasClass("open")) {

                    return false;
                }

                WTree.opts.node = $(this).attr("data-node");

                WTree.Load();

                $(this).toggleClass("open");
            });
        });

    },

    populate : function(id, d) {

        for(var i=0; i < d.length; i++) {

            if(id != d[i].Parent) {
                continue;
            }

            var container;

            if(id === 0) {
                container = WTree.opts.Target;
            } else {
                container = WTree.opts.Target.find(".subcontainer-id-"+id);
            }

            var txt = "<div><p>";

            if (!WTree.opts.selected(d[i])) {

                txt += "<a href='"+d[i].A.href+"'>"+d[i].Text+"</a>";

            }
            else {
                txt += '<span class="wtree-selected">'+d[i].Text+'</span>';
            }

            txt += "</p>";

            if(d[i].HasChildren) {
                txt += '<div class="wtree-toggle" data-node="'+d[i].Id+'"></div>';
            }
            txt += "<div class='offset-wtree-1 subcontainer-id-"+d[i].Id+"'></div></div>";

            container.append(txt);

            if(d[i].Open == true && d[i].HasChildren) {

                container.find(".wtree-toggle[data-node="+d[i].Id+"]").toggleClass("open");
            }



            WTree.populate(d[i].Id, d);

        }

    }


};