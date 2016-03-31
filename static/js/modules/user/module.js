/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: module.js
 *    description: 
 *        created: 2016-03-31 23:16:55
 *         author: wystan
 *
 *********************************************************************************
 */

(function() {
    var user = app.namespace("app.modules.user");

    user.cur_user = {
        id: 0,
        nicky: ""
    };
    user.detail = {
        detail: null,
        team_joined: null,
        team_created: null
    };
    user.team_models = null;

    user.init = function(eb) {
        user.team_models = app.namespace("app.modules.team.models");
        _.extend(this, Backbone.Events);
        _.each(this.views, function(v) {
            v.init(eb);
            console.log(v.name, "initialized");
        });

        this.listenTo(eb, "show:user:edit", show_edit);
        this.listenTo(eb, "show:user:pwd", show_pwd);
        this.listenTo(eb, "show:user:detail", get_detail);
        this.listenTo(eb, "current_user", set_current);
    };

    user.candinates = function(nicky) {
        return [];
    };

    function show_edit() {
        //TODO
    }

    function show_pwd() {
        // body...
    }

    function get_detail(uid) {
        user.detail.detail = user.detail.team_joined = user.detail.team_created = null;
        get_user(uid);
        get_team_joined(uid);
        get_team_created(uid);
    }

    function get_user(uid) {
        var usr = new user.models.user({
            id: uid
        });
        usr.on("sync", function() {
            user.detail.detail = usr.toJSON();
            render_user(user.detail);
        });
        usr.on("error", function() {
            console.log("ERROR: can not get user", uid);
        });

        usr.fetch();
    }

    function get_team_joined(uid) {
        var list = new user.team_models.list();

        list.on("sync", function() {
            var tl = [];
            list.each(function(item) {
                tl.push(item.toJSON());
            });
            user.detail.team_joined = tl;
            render_user(user.detail);
        });

        list.on("error", function() {
            // body...
        });

        list.url = "/users/" + uid + "/teams/joined";
        list.fetch();
    }

    function get_team_created(uid) {
        var list = new user.team_models.list();

        list.on("sync", function() {
            var tl = [];
            list.each(function(item) {
                tl.push(item.toJSON());
            });
            user.detail.team_created = tl;
            render_user(user.detail);
        });

        list.on("error", function() {
            // body...
        });

        list.url = "/users/" + uid + "/teams/created";
        list.fetch();
    }

    function set_current(uid, nicky) {
        user.cur_user.id = uid;
        user.cur_user.nicky = nicky;
    }

    function render_user(detail) {
        if (detail.detail == null || detail.team_joined == null || detail.team_created == null) {
            return;
        }
        if (user.views.show) {
            user.views.show.render(detail);
        }
    }

})();


/************************************* END **************************************/