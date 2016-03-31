/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: view_edit.js
 *    description: 
 *        created: 2016-03-31 19:01:03
 *         author: wystan
 *
 *********************************************************************************
 */

(function(){
    var view = app.namespace("app.modules.team.views.edit");

    view.init = function(eb) {

    };
    view.render = function(model) {
        if(model.detail == null || model.members == null) {
            return;
        }
        var html = template("team_edit", model);
        $("#main").html(html);
    };

})();




/************************************* END **************************************/
