/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: view_list.js
 *    description: 
 *        created: 2016-03-31 19:00:43
 *         author: wystan
 *
 *********************************************************************************
 */

(function(){
    var view = app.namespace("app.modules.team.views.list");
    view.init = function (eb) {
        // body...
    };

    view.render = function (joined, created) {
        var h1 = template("teams_joined", joined);
        var h2 = template("teams_created", created);
        $("#main").html(h1+h2);
    };

})();




/************************************* END **************************************/
