function inArray(key, array) {  

    for (var i = 0; i < array.length; i++) {

        if (key == array[i]) {

            return true;
        }
    };

    return false;
}

var JSForm = {

    Meta : {},

    New : function(url) {

        JSForm.Meta.url = url;

        JSForm.Send("GET", {});
    },

    Send : function(method, data) {

        $.ajax({
            url: this.Meta.url,
            cache: false,
            method: method,
            data : data
        })
            .done(function (data) {

                if (data == null || typeof data !== 'object') {

                    Weasel.RenderError(data);

                    return;
                }

                if (typeof data.Error != 'undefined') {

                    Weasel.RenderError(data.Error);

                    return;
                }
                //@todo: #
                if(typeof data.r != 'undefined') {

                    if(data.r == null || data.r == "") {

                        window.location.href = window.location.href;

                    } else if (data.r == "close") {

                        $("#formModal").modal("toggle");

                    } else {

                        window.location.href = data.r;
                    }

                    return;
                }

                var f = ParseForm(data);

                $("#formModal").modal("toggle").find(".modal-body").eq(0).html(f);
                $("#formModal").find("form").eq(0).attr("action", JSForm.Meta.url);

                $('#formModal').find(".date-picker").datepicker({
                    autoclose: true
                });

                if (typeof data.url != "undefined") {

                    JSForm.Meta.url = data.url;

                }

            })
            .fail(
            function (data) {
                Weasel.RenderError(data.responseText)
            })
            .always(function () {
                // On complete actions!
            });
    },

    Post : function(ev) {

        JSForm.Send("POST", $(ev.target).serializeArray());

    }

};

var ParseForm = function(data) {

    var popup = '<fieldset><legend class="section">'+data.Title+'</legend>';

    for (var i = 0; i < data['e'].length; ++i) {

        popup += '<div class="control-group">';

        var error = "";

        if (typeof data['e'][i]['er'] != 'undefined' && data['e'][i]['er']  != '') {
            error = data['e'][i]['er']
        }

        var required = '';

        if (typeof data['e'][i]['r'] != 'undefined' && data['e'][i]['r'] == true) {

            required = ' *';
        }

        var value = '';

        if (data['e'][i]['v'] != null) {

            value = data['e'][i]['v'];
        }



        switch (data['e'][i]['t']) {

            case 'text':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<input class="' + error + ' form-control" type="text" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '" value="' + value + '">';
                popup += '</div></div>';
 
                break;

            case 'uint':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<input class="' + error + ' form-control parsley-validated" type="number" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '" value="' + value + '">';
                popup += '</div></div>';

                break;

            case 'number':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<input class="' + error + ' form-control parsley-validated" type="number" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '" value="' + value + '">';
                popup += '</div></div>';

                break;

            case 'hidden':

                popup += '<input class="' + error + '" type="hidden" name="' + data['e'][i]['n'] + '" value="' + value + '">';

                break;

            case 'login':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<span class="input-group-addon"><i class="fa fa-user"></i></span>';
                popup += '<input class="' + error + ' form-control" type="text" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '" value="' + value + '">';
                popup += '</div></div>';

                break;

            case 'password':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<span class="input-group-addon"><i class="fa fa-lock"></i></span>';
                popup += '<input class="' + error + ' form-control" type="password" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '" value="' + value + '">';
                popup += '</div></div>';
 
                break;

            case 'cpassword': // с подтверждением
                popup += '<h5>Старый пароль*</h5>';
                popup += '<input class="' + error + '" type="password" name="' + data['e'][i]['n'] + 'o">';
                popup += '<h5>Новый пароль*</h5>';
                popup += '<input class="' + error + '" type="password" name="' + data['e'][i]['n'] + 'n">';
                popup += '<h5>Подтвердить *</h5>';
                popup += '<input class="' + error + '" type="password" name="' + data['e'][i]['n'] + 'c">';
                break;

            case 'taglist':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<textarea class="' + error + ' form-control" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '">' + value + '</textarea>';
                popup += '</div></div>';

                break;

            case 'textarea':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<textarea class="' + error + ' form-control" name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '">' + value + '"</textarea>';
                popup += '</div></div>';

                break;

            case 'date':


                popup += '<label for="' + data['e'][i]['n'] + '" class="control-label">' + data['e'][i]['l'] + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<input id="' + data['e'][i]['n'] + '" class="form-control date-picker" name="' + data['e'][i]['n'] + '" value="' + value + '" type="text">';
                popup += '</div></div>';

                break;
            case 'datetimepicker':
                popup += '<h5>' + data['e'][i]['l'] + '</h5>';
                popup += '<input class="datetimepicker ' + error + '" type="text" name="' + data['e'][i]['n'] + '" value="' + value + '">';
 
                break;


            case 'checkbox[]':

                popup += '<h5>' + data['e'][i]['l'] + '</h5>';

                if (typeof data['e'][i]['o'] == 'object') {

                    for (var c = 0; c < data['e'][i]['o'].length; c++) {

                        popup += '<label><input ' + (inArray(data['e'][i]['o'][c].v, value) ? 'checked' : '') + ' type="checkbox" name="' + data['e'][i]['n'] + '[]" value="'+ data['e'][i]['o'][c].v +'"><span>'+ data['e'][i]['o'][c].n +'</span></label>';
                    };
                }
               
                break;

            case 'bool':

                var checked = "";

                console.log(data['e'][i]['v']);

                if (value == "true") {
                    checked = "checked";
                }

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';
                popup += '<input class="' + error + ' form-control" type="checkbox" '+checked+' name="' + data['e'][i]['n'] + '" id="' + data['e'][i]['n'] + '" value="' + value + '">';
                popup += '</div></div>';
                
                break;

            case 'radio':
                var checked = "";
                if (data['e'][i]['checked'] == true) checked = "checked";
                popup += '<label><input ' + checked + ' type="radio" name="' + data['e'][i]['n'] + '"> <span>' + data['e'][i]['l'] + '</span></label>';
                break;

            case 'select':

                popup += '<label class="control-label" for="'+data['e'][i]['n']+'">' + data['e'][i]['l'] + ' ' + required + '</label>';
                popup += '<div class="controls form-group"><div class="input-group col-sm-11"> ';

                //popup += '<h5>' + data['e'][i]['l'] + '</h5>';

                popup += '<select class="selectpicker ' + error + '"  name="' + data['e'][i]['n'] + '">';
                var j;
                for (j = 0; j < data['e'][i]['o'].length; ++j) {
                    var selected = "";

                    if (data['e'][i]['o'][j]['v'] == value) {
                        selected = "selected";
                    }
                    popup += '<option value="' + data['e'][i]['o'][j]['v'] + '" ' + selected + '>' + data['e'][i]['o'][j]['n'] + '</option>';
                }
                popup += '</select>';

                popup += '</div></div>';

                break;

            case 'u_c_list':

                popup += '<h5>' + data['e'][i]['l'] + '</h5>';

                popup += '<select class="selectpicker ' + error + '" name="' + data['e'][i]['n'] + '">';

                var responseText = $.ajax({
                    type: "GET",
                    dataType: 'json',
                    cache: false,
                    async: false,
                    url: "/classifiers/options/" + data['e'][i]['lc']['reference'] + "/" + data['e'][i]['lc']['version_id'] + "/"
                }).responseText;

                var elements = JSON.parse(responseText);
            
                for (var e = 0; e < elements.length; e++) {

                    var selected = '';

                    if (elements[e].v == value) {
                        
                        selected = 'selected';
                    }

                    popup += '<option value="'  + elements[e].v +  '" ' + selected + '>' + elements[e].k + '</option>';
                };

                popup += '</select>';

                break;

            case 'subdivisions':

                popup += '<h5>' + data['e'][i]['l'] + '</h5>';

                popup += '<select class="selectpicker ' + error + '" name="' + data['e'][i]['n'] + '">';

                var responseText = $.ajax({
                    type: "GET",
                    dataType: 'json',
                    cache: false,
                    async: false,
                    url: "/configurator/subdivisions/element-options/"
                }).responseText;

                var elements = JSON.parse(responseText);

                for (var e = 0; e < elements.length; e++) {

                    var selected = '';

                    if (elements[e].v == value) {
                        
                        selected = 'selected';
                    }

                    var level = '';

                    for (var l = 1; l < elements[e].l; l++) {
                        level += '&nbsp;&nbsp;&nbsp;';
                    };

                    popup += '<option value="'  + elements[e].v +  '" ' + selected + '>' + level + elements[e].k + '</option>';
                };

                popup += '</select>';

                break;

            case 'users':

                popup += '<h5>' + data['e'][i]['l'] + '</h5>';

                popup += '<select class="selectpicker ' + error + '" name="' + data['e'][i]['n'] + '">';

                var responseText = $.ajax({
                    type: "GET",
                    dataType: 'json',
                    cache: false,
                    async: false,
                    url: "/configurator/users/element-options/"
                }).responseText;

                var elements = JSON.parse(responseText);

                for (var e = 0; e < elements.length; e++) {

                    var selected = '';

                    if (elements[e].v == value) {
                        
                        selected = 'selected';
                    }

                    popup += '<option value="'  + elements[e].v +  '" ' + selected + '>' + elements[e].k + '</option>';
                };

                popup += '</select>';

                break;

            case 'optgroup':
                popup += '<h5>' + data['e'][i]['l'] + '</h5>';

                var multiple = "";

                if (data['e'][i]['multiple'] == true) {

                    multiple = "multiple";
                }

                popup += '<select class="selectpicker ' + error + '" ' + multiple + ' name="' + data['e'][i]['n'] + '">';
                var j;

                for (j = 0; j < data['e'][i]['og'].length; ++j) {

                    popup += '<optgroup label="' + data['e'][i]['og'][j]['n'] + '">';
                    var k;

                    for (k = 0; k < data['e'][i]['og'][j]['o'].length; ++k) {

                        var selected = "";

                        if (typeof data['e'][i]['multiple'] == "undefined" && data['e'][i]['og'][j]['o'][k]['v'] == value) {

                            selected = "selected";
                        }

                        popup += '<option value="' + data['e'][i]['og'][j]['o'][k]['v'] + '" ' + selected + '>' + data['e'][i]['og'][j]['o'][k]['n'] + '</option>';
                    }
                    popup += '</optgroup>';
                }
                popup += '</select>';

                break;
/*
            case 'p_c': //password constraints

                popup += '<h5>Ограничения</h5>';

                var constraints = ["min", "letters", "number", "upper", "special"];
                var n;


                for (n = 0; n < constraints.length; ++n) {

                    var value = 0;

                    if (typeof data['e'][i]['v'][constraints[n]] != 'undefined') {

                        value = data['e'][i]['v'][constraints[n]];
                    }

                    popup += constraints[n] + ': <input type="text" name="' + data['e'][i]['n'] + '[' + constraints[n] + ']" value="' + value + '">';
                }
                
                break;
*/
            case 'phones':

                var phones = data['e'][i]['v'];
                var name   = data['e'][i]['n'];

                var j;
                
                for (j = 0; j < phones.length; ++j) {

                    if (typeof phones[j]['err'] != 'undefined') {

                        popup += '<i class="error-info">' +  phones[j]['err']  + '</i>';
                    }

                    popup += ' \
                    <input type="hidden" name="' + name + '[][id]" value="' + phones[j]['i'] + '">  \
                    <div class="popup-content popup-content_border_bottom popup-content_pr_60" data-phone-id="' + phones[j]['i'] + '" data-division-id="' + phones[j]['s'] + '"> \
                            <div class="popup-content__left"> \
                                <div class="input-label input-label_inline input-label_mb_5">Название</div> \
                                <div class="input-block input-block_w_100p input-block_mb_7"> \
                                    <input name="' + name + '[' + phones[j]['i'] + '][title]" type="text" class="input input_mb_0" value="' + phones[j]['t'] + '"> \
                                    <i class="placeholder">Введите название телефона...</i> \
                                </div> \
                                <div class="input-label input-label_inline input-label_mb_5">Текстареа</div> \
                                <div class="input-block input-block_w_100p input-block_mb_7"> \
                                    <i class="placeholder-textarea">Описание телефона...</i> \
                                    <textarea class="textarea textarea_h_102" name="' + name + '[' + phones[j]['i'] + '][description]" rows="5">' + phones[j]['d'] + '</textarea> \
                                </div> \
                                <div class="clear"></div> \
                            </div> \
                            <div class="popup-content__right"> \
                                <div class="input-label input-label_inline input-label_mb_5">Код</div> \
                                <div class="input-block input-block_w_100p input-block_mb_7"> \
                                    <input name="' + name + '[' + phones[j]['i'] + '][code]" type="text" class="input input_mb_0" value="' + phones[j]['c'] + '"> \
                                    <i class="placeholder">Введите код города...</i> \
                                </div><div class="input-label input-label_inline input-label_mb_5">Номер</div> \
                                <div class="input-block input-block_w_100p input-block_mb_7"> \
                                    <input name="' + name + '[' + phones[j]['i'] + '][number]" type="text" class="input input_mb_0" value="' + phones[j]['n'] + '"> \
                                    <i class="placeholder">Введите номер телефона...</i> \
                                </div> \
                                <div class="input-label input-label_inline input-label_mb_5">Добавочный номер</div> \
                                <div class="input-block input-block_w_100p input-block_mb_7"> \
                                    <input name="' + name + '[' + phones[j]['i'] + '][extension_number]" type="text" class="input input_mb_0" value="' + phones[j]['e'] + '"> \
                                    <i class="placeholder">Введите добавочный номер...</i> \
                                </div> \
                            </div> \
                            <div class="popup-content__del"> \
                                <button type="button" class="del-btn js-phone-del">Удалить</button> \
                            </div> \
                        </div> \
                    ';

                   
                }

                break;

            case 'addresses':


                var addresses = data['e'][i]['v'];
                var name   = data['e'][i]['n'];

                var j;
                var juridical = false;
                for (j = 0; j < addresses.length; ++j) {

                    var err = '';
                    if (typeof addresses[j]['err'] != 'undefined') {

                        err = '<i class="error-info">' +  addresses[j]['err']  + '</i>';
                    }

                    switch(addresses[j]['e']) {



                        case 'juridical_address' :
                            juridical = true;
                            popup += '<h4 class="title title_mb_5">Юридический адрес</h4>' + err;

                            break;

                        case 'physical_address':

                            popup += '<h4>Фактический адрес</h4>'  + err;

                            if (juridical) {

                                popup += '\
                                    <div class="input-block"> \
                                        <label> \
                                            <input type="checkbox" name="asJuridical" class="js-address" value="1"> \
                                            <span>Совпадает с юридическим адресом</span> \
                                        </label> \
                                    </div> \
                                ';
                            }

                            popup += '<div class="clear"></div>';

                            break;

                        case 'postal_address':
                            
                            popup += '\
                                <h4>Почтовый адрес</h4> ' + err + '\
                                <div class="input-block"> \
                                    <label> \
                                        <input type="checkbox" name="asPhysical" class="js-address" value="1"> \
                                        <span>Совпадает с фактическим адресом</span> \
                                    </label> \
                                </div> \
                                <div class="clear"></div> \
                             ';

                            break;

                    }

                    popup += address(name, data['d']['countries'], addresses[j]);
                }

                break;

            case 'title':

                popup += '<h5>' + data['e'][i]['l'] + '</h5>';
                break;
        }

        if (data['e'][i]['t'] != 'phones' && data['e'][i]['t'] != 'addresses' && error != '') {
            popup += '<i class="error-info">' + error + '</i>';
        }

        popup += '</div>';
    }

    popup += '</fieldset>';

    //popup += '</form>';

    return popup;
};


function address(name, countries, addr) {

    var countriesOptions = '';
    var j;
    for (j = 0; j < countries.length; ++j) {
        
        var selected = '';

        if (addr['r'] == countries[j]['ID']) {

            selected = 'selected'
        }

        countriesOptions += '<option value="' + countries[j]['ID'] + '" ' + selected + ' >' + countries[j]['Name'] + '</option>';
    }

    return '\
    <input type="hidden" name="' + name + '[' + addr['e'] + '][type]" value="' + addr['e'] + '"> \
    <div class="popup-content"> \
        <div class="popup-content__left"> \
            <div class="input-label input-label_inline input-label_mb_5">Страна</div> \
            <div class="input-block input-block_w_100p input-block_mb_7"> \
                <select name="' + name + '[' + addr['e'] + '][country_id]" class="selectpicker select-input_mb_0 "> \
                    ' + countriesOptions + ' \
                </select> \
            </div> \
            <div class="input-label input-label_inline input-label_mb_5">Город</div> \
            <div class="input-block input-block_w_100p input-block_mb_7"> \
                <input name="' + name + '[' + addr['e'] + '][city]" type="text" class="input input_mb_0" value="' + addr['t'] + '"> \
                <i class="placeholder">Введите название города...</i> \
            </div> \
            <div class="input-label input-label_inline input-label_mb_5">Улица</div> \
            <div class="input-block input-block_w_100p input-block_mb_7"> \
                <input name="' + name + '[' + addr['e'] + '][street]" type="text" class="input input_mb_0" value="' + addr['y'] + '">  \
                <i class="placeholder">Введите название улицы...</i> \
            </div> \
        </div> \
        <div class="popup-content__right"> \
            <div class="input-label input-label_inline input-label_mb_5">Дом, строение</div> \
            <div class="input-block input-block_w_100p input-block_mb_7"> \
                <input name="' + name + '[' + addr['e'] + '][house]" type="text" class="input input_mb_0" value="' + addr['u'] + '"> \
                <i class="placeholder">Введите номер дома...</i> \
            </div><div class="input-label input-label_inline input-label_mb_5">Почтовый индекс</div> \
            <div class="input-block input-block_w_100p input-block_mb_7"> \
                <input name="' + name + '[' + addr['e'] + '][postal_code]" type="text" class="input input_mb_0" value="' + addr['i'] + '"> \
                <i class="placeholder">Введите почтовый индекс...</i> \
            </div> \
            <div class="input-label input-label_inline input-label_mb_5">ОКАТО</div> \
            <div class="input-block input-block_w_100p input-block_mb_7"> \
                <input name="' + name + '[' + addr['e'] + '][okato]" type="text" class="input input_mb_0" value="' + addr['o'] + '"> \
                <i class="placeholder">Введите номер ОКАТО...</i> \
            </div> \
        </div> \
    </div> \
    ';
}