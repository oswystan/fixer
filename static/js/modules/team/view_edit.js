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

(function() {
    var view = app.namespace("app.modules.team.views.edit");

    view.cur_team = null;

    view.init = function(eb) {
        view.search = null;
        view.eb = eb;
    };

    view.render = function(model) {
        if (model.detail == null || model.members == null) {
            return;
        }
        view.cur_team = model;
        var main = $("#main");
        var html = template("team_edit", model);
        main.html(html);
        main.find("[name='team-name']").unbind('blur').blur(change_name);
        main.find("[name='team-goal']").unbind('blur').blur(change_goal);
        main.find("[name='check-active']").unbind('click').click(change_status);
        main.find("[name='member-name']").unbind('keyup;').keyup(check_user);
        main.find("[name='op-submit']").unbind('click').click(do_submit);
        main.find("[name='op-add']").unbind('click').click(do_add_member);
        main.find("[name='op-delete']").unbind('click').click(do_del_member);

        var search = app.namespace("app.modules.team.search");
        view.search = search.create(view.eb, main.find("[name='member-name']")[0]);
    };

    view.validate = function() {
        return true;
    };

    // team detail changed handlers
    function change_name() {
        view.cur_team.detail.name = this.value;
    }
    function change_goal() {
        view.cur_team.detail.goal = this.value;
    }
    function change_status() {
        view.cur_team.detail.status = this.checked ? 1 : 0;
    }
    function check_user() {
        var nicky = $.trim(this.value);
        if(nicky.length == 0) {
            return;
        }
        console.log(nicky);
    }

    // operations
    function do_add_member() {
        var user = view.search.current();
        if(user !== null) {
            console.log("add", user.nicky, user.id, user.email);
        }
    }
    function do_del_member() {
        $("#main").find("[name='team-members']").find(":checked").parentsUntil("div").remove();
    }
    function do_submit() {
        if (view.validate() != true) {
            return;
        }
        console.log(view.cur_team.detail);
    }

})();



/************************************* END **************************************/