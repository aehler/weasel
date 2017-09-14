jsPlumb.ready(function() {

	// setup some defaults for jsPlumb.
	var instance = jsPlumb.getInstance({
		Endpoint : ["Dot", {radius:3}],
		HoverPaintStyle : {strokeStyle:"#1e8151", lineWidth:2 },
		ConnectionOverlays : [
			[ "Arrow", {
				location:1,
				id:"arrow",
                length:14,
                foldback:0.8
			} ],
            [ "Label", { label:"", id:"label", cssClass:"aLabel" }]
		],
		Container:"statemachine-demo"
	});

    window.jsp = instance;

	var windows = jsPlumb.getSelector(".statemachine-demo .w");

    // initialise draggable elements.
	instance.draggable(windows);

    // bind a click listener to each connection; the connection is deleted. you could of course
	// just do this: jsPlumb.bind("click", jsPlumb.detach), but I wanted to make it clear what was
	// happening.
	instance.bind("dblclick", function(c) {
		instance.detach(c);
	});

    windows.bind("dblclick", function(){

        //Не реагировать на первый и последний блоки
        if($(this).attr("id") == 'started-block' || $(this).attr("id") == 'accepted-block') {
            return false;
        }

        var as = approvalStep.NewStep($(this), instance);
        return false;

    });

	// bind a connection listener. note that the parameter passed to this function contains more than
	// just the new connection - see the documentation for a full list of what is included in 'info'.
	// this listener sets the connection's internal
	// id as the label overlay's text.
    instance.bind("connection", function(info) {

		//info.connection.getOverlay("label").setLabel(info.connection.id);
        info.connection.getOverlay("label").setLabel('');

        // Тип связи отключен

        //info.connection.type = "success";
        //
        //var label = '<i class="fa fa-fw fa-lg fa-thumbs-o-up text-'+info.connection.type+'"></i>';
        //
        //if(info.connection.FSMType) {
        //    info.connection.type = info.connection.FSMType;
        //    if(info.connection.type == 'danger') {
        //        label = '<i class="fa fa-fw fa-lg fa-thumbs-o-down text-'+info.connection.type+'"></i>';
        //    }
        //}
        //
        //info.connection.getOverlay("label").setLabel(label);


        //info.connection.bind("click", function(conn) {
        //
        //    var label = '';
        //
        //
        //    switch (conn.type) {
        //        case  "success" :
        //            conn.type = "danger";
        //            label = '<i class="fa fa-fw fa-lg fa-thumbs-o-down text-'+conn.type+'"></i>';
        //            break;
        //        case  "danger" :
        //            conn.type = "success";
        //            label = '<i class="fa fa-fw fa-lg fa-thumbs-o-up text-'+conn.type+'"></i>';
        //            break;
        //        default :
        //            conn.type = "danger";
        //            label = '<i class="fa fa-fw fa-lg fa-thumbs-o-up text-'+conn.type+'"></i>';
        //            break;
        //    }
        //
        //    info.connection.getOverlay("label").setLabel(label);
        //
        //});

    });


	// suspend drawing and initialise.
	instance.doWhileSuspended(function() {
		var isFilterSupported = instance.isDragFilterSupported();
		// make each ".ep" div a source and give it some parameters to work with.  here we tell it
		// to use a Continuous anchor and the StateMachine connectors, and also we give it the
		// connector's paint style.  note that in this demo the strokeStyle is dynamically generated,
		// which prevents us from just setting a jsPlumb.Defaults.PaintStyle.  but that is what i
		// would recommend you do. Note also here that we use the 'filter' option to tell jsPlumb
		// which parts of the element should actually respond to a drag start.
		// here we test the capabilities of the library, to see if we
		// can provide a `filter` (our preference, support by vanilla
		// jsPlumb and the jQuery version), or if that is not supported,
		// a `parent` (YUI and MooTools). I want to make it perfectly
		// clear that `filter` is better. Use filter when you can.
		if (isFilterSupported) {
			instance.makeSource(windows, {
				filter:".ep",
				anchor:"Continuous",
				//connector:[ "StateMachine", { curviness:1 } ],
                connector:[ "Flowchart", { stub:[40, 60], gap:5, cornerRadius:5, alwaysRespectStubs:true } ],
				connectorStyle:{ strokeStyle:"#5c96bc", lineWidth:2, outlineColor:"transparent", outlineWidth:4 },
				maxConnections:2,
				onMaxConnections:function(info, e) {
					alert("Maximum connections (" + info.maxConnections + ") reached");
				}
			});
		}
		else {
			var eps = jsPlumb.getSelector(".ep");
			for (var i = 0; i < eps.length; i++) {
				var e = eps[i], p = e.parentNode;
				instance.makeSource(e, {
					parent:p,
					anchor:"Continuous",
					//connector:[ "StateMachine", { curviness:1 } ],
                    connector:[ "Flowchart", { stub:[40, 60], gap:5, cornerRadius:5, alwaysRespectStubs:true } ],
					connectorStyle:{ strokeStyle:"#5c96bc",lineWidth:2, outlineColor:"transparent", outlineWidth:4 },
					maxConnections:-1,
					onMaxConnections:function(info, e) {
						alert("Maximum connections (" + info.maxConnections + ") reached");
					}
				});
			}
		}
	});

	// initialise all '.w' elements as connection targets.
	instance.makeTarget(windows, {
		dropOptions:{ hoverClass:"dragHover" },
		anchor:"Continuous",
		allowLoopback:false
	});

	// and finally, make a couple of connections, example:

    $("#loadedConnections").find(".loaded-connection").each(function() {

        var conn = instance.connect({ source : $(this).attr("data-out"), target : $(this).attr("data-in")}, {FSMType : $(this).attr("data-type")});

        //conn.type = $(this).attr("data-type");

    });

    instance.bind("connection", function(info) {
        if(info.sourceId == info.targetId) {
            instance.detach(info.connection);
            alert("В маршрутах согласования не допустимы переходы на себя!");
        }
        if(info.targetId == 'started-block') {
            instance.detach(info.connection);
            alert("Переход в начало согласования не допустим!");
        }
        if(info.sourceId == 'accepted-block') {
            instance.detach(info.connection);
            alert("Переход из утвержденного документа на этап согласования не допустим!");
        }
        info.connection.bind("click", function(conn) {
            //alert(conn);
        });
    });

	jsPlumb.fire("jsPlumbDemoLoaded", instance);

    $("#statemachine-save").bind("click", function(){
        if(instance.getConnections() == "") {
            alert("Не указано ни одной связи! Пожалуйста, сформируйте корректный маршрут согласования.");
            return false;
        }
        sendSteps(instance.getConnections());
    });

    $("#statemachine-newelement").bind("click", function(){

        var uuid = guid();

        var nw = $("#statemachine-demo").append('<div class="w" id="'+uuid+'"><span class="container-title">Новый этап согласования*</span><div class="ep" id="ep-'+uuid+'"></div></div>');

        var newWin = jsPlumb.getSelector("#"+uuid);
        var newEp = jsPlumb.getSelector("#ep-"+uuid);

        newWin.css("top", "150px");
        newWin.css("left", "250px");

        var p = newEp.parentNode;

        instance.draggable(newWin);

        if(instance.isDragFilterSupported()) {
            instance.makeSource(newWin, {
                filter: ".ep",
                anchor: "Continuous",
                //connector:[ "StateMachine", { curviness:1 } ],
                connector: ["Flowchart", {stub: [40, 60], gap: 5, cornerRadius: 5, alwaysRespectStubs: true}],
                connectorStyle: {strokeStyle: "#5c96bc", lineWidth: 2, outlineColor: "transparent", outlineWidth: 4},
                maxConnections: 2,
                onMaxConnections: function (info, e) {
                    alert("Maximum connections (" + info.maxConnections + ") reached");
                }
            });
        }
        else {
            instance.makeSource(newEp, {
                parent:p,
                anchor:"Continuous",
                //connector:[ "StateMachine", { curviness:1 } ],
                connector:[ "Flowchart", { stub:[40, 60], gap:5, cornerRadius:5, alwaysRespectStubs:true } ],
                connectorStyle:{ strokeStyle:"#5c96bc",lineWidth:2, outlineColor:"transparent", outlineWidth:4 },
                maxConnections:-1,
                onMaxConnections:function(info, e) {
                    alert("Maximum connections (" + info.maxConnections + ") reached");
                }
            });
        }


        instance.makeTarget(newWin, {
            dropOptions:{ hoverClass:"dragHover" },
            anchor:"Continuous",
            allowLoopback:false
        });

        newWin.bind("dblclick", function() {

            var as = approvalStep.NewStep($(this), instance);

            return false;

        });
    });

});

var RouteId = 0;

var approvalStep = {

    obj  : null,

    inst : null,

    NewStep : function(obj, jPlumbInstance) {

        this.obj = obj;
        this.inst = jPlumbInstance;

        $('#approvalModal').modal();

        this.loadStepSettings();

        return this;
    },

    loadStepSettings : function() {
        //var b = $("#approvalModal .modal-body");
        //b.html("");
        $.ajax({
            url: "/configurator/approval/step_settings/"+this.obj.attr("id")+"/"+RouteId+'/',
            type: "GET",
            processData: false,  // tell jQuery not to process the data
            contentType: false,   // tell jQuery not to set contentType
            success : function(htmlData) {
                //b.html(htmlData);

                Popup.Show('<div class="popup-body">'+htmlData+'</div><div class="popup-actions">'+buttons+'</div> ', 'Редактирование этапа', 'large');

                var sel = $("#formContainer select").bind("change", function(){
                    approvalStep.getApprovalEntityValues($(this).val());
                });

                $('#stepDeleteControl').bind("click", function(){
                        approvalStep.DeleteStep();
                    }
                );
                $('#stepSaveControl').bind("click", function(){
                        approvalStep.SaveStep();
                    }
                );

                //b.find('#entityContainer select[multiple=multiple]').multipleSelect();

            },
            error : function(e) {
                b.html("Error: "+e.status+' '+e.statusText);
            }
        });
    },

    getApprovalEntityValues : function(val) {
      var $entity = $("#entityContainer");

      $entity.html('');
      $.ajax({
        url: "/configurator/approval/entity_values/"+this.obj.attr("id")+"/"+val+"/",
        type: "GET",
        processData: false,  // tell jQuery not to process the data
        contentType: false,   // tell jQuery not to set contentType
        success : function(htmlData) {

          $entity
            .html(htmlData)
            .find('.selectpicker').selectpicker('render');

          //var bb = '<div id="entityContainer">'+htmlData+'</div>';

          //c.append(bb)

        },
        error : function(e) {
          c.append("<div id='entityContainer' class='alert alert-warning'>Error: "+e.status+' '+e.statusText+"</div>");
        }
      });
    },

    DeleteStep : function() {
        if(!confirm("Удалить этап согласования?")) {
            return false;
        }
        var id = this.obj.attr("id");

        $.post(
            "/configurator/approval/step_delete/",
            {id : id},
            function(responce) {
                approvalStep.inst.detachAllConnections(id);
                approvalStep.obj.remove();
                $('#approvalModal').modal('hide');
            }
        );
    },

    SaveStep : function() {

        var m = $(".popup-large");

        var f = [];

        m.find("form").each(function() {

            f.push($(this).serializeArray());
        });

        $.post(
            "/configurator/approval/step_settings/"+this.obj.attr("id")+"/"+RouteId+'/',
            {data : JSON.stringify(f)},
            function(responce) {
                var res = {};
                try {
                  res = JSON.parse(responce);
                  $('.js-popup').hide();
                  $('#overlay').hide();
                }
                catch (ex) {
                    alert(ex);
                    return false;
                }

                approvalStep.obj.find(".container-title").html(res.success.name);

            }
        );

    }

};


$(document).ready(

    function() {

        RouteId = $("#mainRouteId").val();

        $('#stepDeleteControl').bind("click", function(){
                approvalStep.DeleteStep();
            }
        );
        $('#stepSaveControl').bind("click", function(){
                approvalStep.SaveStep();
            }
        );
    }
);

var guid = (function() {
    function s4() {
        return Math.floor((1 + Math.random()) * 0x10000)
            .toString(16)
            .substring(1);
    }
    return function() {
        return s4() + s4() + s4() + '-' + s4() + s4() + s4() + s4() + s4();
    };
})();

var sendSteps = function(params) {

    var url = window.location.href;
    var data = [];

    for(var i=0; i<params.length; i++) {

        data[i] = {
            type : params[i].type,
            out : {id : params[i].sourceId, type : "", position: {top : Math.ceil($(params[i].source).position().top), left: Math.ceil($(params[i].source).position().left)} },
            in  : {id : params[i].targetId, type : "", position: {top : Math.ceil($(params[i].target).position().top), left: Math.ceil($(params[i].target).position().left)} }
        }
    }

    $.post( url, {data : JSON.stringify(data)}, function( data ) {

        if(data == 'true') {
            $( "#result" ).html( '<div class="alert alert-success">Маршрут сохранен</div>');
        }
        else {
            $( "#result" ).html( '<div class="alert alert-danger">Ошибка при сохранении маршрута</div>');
        }

        setTimeout(function(){$( "#result" ).html("")}, 5000);

    });
};