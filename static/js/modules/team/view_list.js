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
        view.s1 = null;
        view.s2 = null;
    };

    view.render = function (joined, created) {
        var h1 = template("teams_joined", joined);
        var h2 = template("teams_created", created);
        $("#main").html(h1+h2);

        var input1 = $("#teamlist-joined").find("[name='op-name']")[0];
        var input2 = $("#teamlist-created").find("[name='op-name']")[0];

        var search  = app.namespace("app.modules.team.search");
        view.s1 = search.create(input1);
        view.s2 = search.create(input2)
    };

})();




/************************************* END **************************************/
