/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: view_new.js
 *    description:
 *        created: 2016-03-31 13:42:06
 *         author: wystan
 *
 *********************************************************************************
 */

(function() {
    var view = app.namespace("app.modules.team.views.new");

    view.init = function(eb) {};

    view.render = function() {
        this.html = template("team_new");
        $("#main").html(this.html);
        var search = app.namespace("app.modules.user.search");
        view.search = search.create($("#main").find("[name='member-name']")[0]);
    };

    view.validate = function() {
        return true;
    };

    view.getModel = function() {
        return null;
    }
    view.html = null;

    function check_name(n) {
        // body...
        return true;
    }

})();



/************************************* END **************************************/