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

    view.init = function(eb) {
        view.search = null;
        view.cur_team = null;
        var rs = app.namespace("app.modules.team.models.relationship");
        view.rs = new rs();
    };

    view.render = function(model) {
        if (arguments.length === 0) {
            model = view.cur_team;
        } else {
            //deep clone the team information.
            view.cur_team = $.extend(true, {}, model);
        }
        if (model.detail == null || model.members == null) {
            return;
        }

        var main = $("#main");
        var html = template("team_edit", model);
        console.log(model.members);
        main.html(html);
        main.find("[name='team-name']").unbind('blur').blur(change_name);
        main.find("[name='team-goal']").unbind('blur').blur(change_goal);
        main.find("[name='check-active']").unbind('click').click(change_status);
        main.find("[name='op-submit']").unbind('click').click(do_submit);
        main.find("[name='op-add']").unbind('click').click(do_add_member);
        main.find("[name='op-delete']").unbind('click').click(do_del_member);

        var search = app.namespace("app.modules.user.search");
        view.search = search.create(main.find("[name='member-name']")[0]);
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

    // operations
    function do_add_member() {
        var user = view.search.current();
        if (user !== null) {
            console.log("begin to add", user.id, user.nicky, user.email);
        }

        var result = _.find(view.cur_team.members, function(m) {
            return m.id === user.id;
        });
        if (result !== undefined) {
            console.log(user.nicky, "already in the team");
            return;
        }

        view.rs.set_team(view.cur_team.detail.id);
        view.rs.add_rs(user.id, function() {
            view.cur_team.members.push(user);
            view.render();
        }, function() {
            console.log("relationship ", user.id, "=>", view.cur_team.detail.id, "failed");
            //TODO need to show error on the web;
        });
    }

    function do_del_member() {
        var selected = $("#main").find("[name='team-members']").find(":checked");
        if (selected.length === 0) {
            return;
        }

        var ml = [];
        $.each(selected, function(i, obj) {
            ml.push($(obj).attr("data-id"));
        });
        process_del_member(ml);
    }

    function process_del_member(ml) {
        if (ml.length === 0) {
            return;
        }

        function on_success() {
            var uid = view.rs.get('id');
            $.each(view.cur_team.members, function(j, m) {
                if (uid == m.id) {
                    view.cur_team.members.splice(j, 1);
                    console.log("remove ", m.id, m.nicky);
                    return false;
                }
            });
            view.render();
            process_del_member(ml);
        }
        function on_error() {
            console.log("fail to del relationship");
            process_del_member(ml);
        }
        var uid = ml.pop();
        view.rs.set_team(view.cur_team.detail.id);
        view.rs.del_rs(uid, on_success, on_error);
    }

    function do_submit() {
        if (view.validate() != true) {
            return;
        }
        console.log(view.cur_team.detail);
    }

})();



/************************************* END **************************************/