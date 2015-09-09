var dataGrid = function(url) {

    this.URL = url;

    var columns = [];
    var rows = [];

    var ActionsFormatter = _.extend({}, Backgrid.CellFormatter.prototype, {
        fromRaw: function (rawValue, model) {

            return rawValue;
        }
    });

    Backgrid.ActionsCell = Backgrid.Cell.extend({
        className: "actions-cell",
        formatter: ActionsFormatter,
        render: function () {
            this.$el.empty();
            var formattedValue = this.formatter.fromRaw(this.model.get(this.column.get("name")));

            for (var i=0; i<formattedValue.length; i++) {

                for (var property in formattedValue[i]) {
                    if (formattedValue[i].hasOwnProperty(property)) {
                        this.$el.append($("<a>", {
                            href: formattedValue[i][property],
                            title: property,
                            class: "jsForm"
                        }).text(property));
                        //    .on("click",function(event){
                        //
                        //    event.preventDefault();
                        //    event.stopPropagation();
                        //
                        //    linkClickHandler($(event.currentTarget).attr("href"));
                        //
                        //    return false;
                        //});
                    }
                }

            }
            return this;
        }
    });

    this.Model = Backbone.Model.extend({});

    var DataCollection = Backbone.PageableCollection.extend({
        model: this.Model,
        url: this.URL,
        state: {
            pageSize: 25
        },
        mode: "client"
    });

    var data = new DataCollection();

    var grid = {};

    var handle = function(object){

        if(object.Error) {
            Weasel.RenderError(object.Error);
        }

        columns = object.columns;

        for (i=0; i<object.rows.length; i++) {

            var m = new Backbone.Model(object.rows[i]);

            data.push(m);

        }

        grid = new Backgrid.Grid({
            columns: columns,
            collection: data,
            footer: Backgrid.Extension.Paginator.extend({
                template: _.template('<tr><td colspan="<%= colspan %>"><ul class="pagination"><% _.each(handles, function (handle) { %><li <% if (handle.className) { %>class="<%= handle.className %>"<% } %>><a href="#" <% if (handle.title) {%> title="<%= handle.title %>"<% } %>><%= handle.label %></a></li><% }); %></ul></td></tr>')
            }),
            className: 'table table-striped table-editable no-margin'
        });

        $("#table-dynamic").append(grid.render().$el);

//        data.fetch({reset: true});
    };

    this.DataGrid = function() {

        var jxhr = $.ajax({
            url: this.URL,
            dataType: "json"
        }).fail(function( jqXHR, textStatus ) {
            Weasel.RenderError(textStatus);
        }).done(function(msg) {

            handle(msg);

        });

//        data.fetch({reset: true});

    };
/*
    var PageableGrid = Backbone.PageableCollection.extend({
        model: DataGrid,
        url: url,
        state: {
            pageSize: 15
        },
        mode: "client" // page entirely on the client side
    });

    var pageableTerritories = new PageableGrid();

// Set up a grid to use the pageable collection
    var pageableGrid = new Backgrid.Grid({
        columns: [{
            // enable the select-all extension
            name: "",
            cell: "select-row",
            headerCell: "select-all"
        }].concat(columns),
        collection: pageableTerritories
    });

// Render the grid
    var $example2 = $("#example-2-result");
    $example2.append(pageableGrid.render().el)

// Initialize the paginator
    var paginator = new Backgrid.Extension.Paginator({
        collection: pageableTerritories
    });

// Render the paginator
    $example2.after(paginator.render().el);

// Initialize a client-side filter to filter on the client
// mode pageable collection's cache.
    var filter = new Backgrid.Extension.ClientSideFilter({
        collection: pageableTerritories,
        fields: ['name']
    });

// Render the filter
    $example2.before(filter.render().el);

// Add some space to the filter and move it to the right
    $(filter.el).css({float: "right", margin: "20px"});

// Fetch some data
    pageableTerritories.fetch({reset: true});
    */
};