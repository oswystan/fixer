/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: view_show.js
 *    description: 
 *        created: 2016-03-31 23:35:53
 *         author: wystan
 *
 *********************************************************************************
 */

(function () {
    var view = app.namespace("app.modules.user.views.show");

    view.init = function (eb) {
        // body...
    };

    view.render = function (model) {
        var html = template("user_detail", model);
        $("#main").html(html);
    };
})();




/************************************* END **************************************/
