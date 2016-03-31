/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: view_show.js
 *    description: 
 *        created: 2016-03-31 19:00:52
 *         author: wystan
 *
 *********************************************************************************
 */

(function(){
    var view = app.namespace("app.modules.team.views.show");
    view.init = function() {
    }
    view.render = function(model) {
        if (model.detail == null || model.members == null) {
            return;
        }
        var html = template("team_detail", model);
        $("#main").html(html);
    }
})();

/************************************* END **************************************/
