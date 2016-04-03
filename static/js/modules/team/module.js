/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: module.js
 *    description: team module controller
 *        created: 2016-03-29 11:35:18
 *         author: wystan
 *
 *********************************************************************************
 */

(function() {
    var team = app.namespace("app.modules.team");
    team.cur_user = {
        id: 0,
        nicky: ""
    };
    team.list = {
        joined: null,
        created: null
    };
    team.detail = {
        detail: null,
        members: null
    };

    team.init = function(eb) {
        _.extend(this, Backbone.Events);
        _.each(this.views, function (v) {
            v.init(eb);
            console.log(v.name, "initialized");
        });

        // for each function, we are in the team context;
        this.listenTo(eb, "show:team:list", function() {
            getList();
        });

        this.listenTo(eb, "show:team:new", function() {
            if(this.views.new) {
                this.views.new.render();
            }
        });

        this.listenTo(eb, "show:team:detail", function(tid) {
            team.detail.detail = team.detail.members = null;
            getTeamDetail(tid);
            getTeamMembers(tid);
        });

        this.listenTo(eb, "show:team:edit", function(tid) {
            if (team.detail.detail == null || team == null) {
                getTeamDetail(tid);
                getTeamMembers(tid);
            } else {
                renderEdit();
            }
        });

        this.listenTo(eb, "current_user", function(id, nicky) {
            this.cur_user.id = id;
            this.cur_user.nicky = nicky;
            console.log("get user:", id, nicky);
        });
    };

    function getList() {
        if (team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        clear();
        team.list.joined = team.list.created = null;

        // get the joined team list
        var tlJoined = new team.models.list();
        tlJoined.on("sync", function() {
            team.list.joined = this;
            renderList();
        });
        tlJoined.on("error", function() {
            console.log("error occurred");
            renderError();
        });
        tlJoined.url = "/users/" + team.cur_user.id + "/teams/joined";
        tlJoined.fetch();

        //get the created team list
        var tlCreated = new team.models.list();
        tlCreated.on("sync", function() {
            team.list.created = this;
            renderList();
        });
        tlCreated.on("error", function() {
            console.log("error occurred");
        });

        tlCreated.url = "/users/" + team.cur_user.id + "/teams/created";
        tlCreated.fetch();
    }

    function getTeamDetail(tid) {
        if (team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        var teamDetail = new team.models.team({
            id: tid
        });
        teamDetail.on("sync", function() {
            team.detail.detail = this.toJSON();
            renderTeam();
        });
        teamDetail.on("error", function() {
            renderError();
        });
        teamDetail.fetch();
    }

    function getTeamMembers(tid) {
        if (team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        var members = new team.models.members();
        members.on("sync", function() {
            team.detail.members = [];
            this.each(function(t) {
                team.detail.members.push(t.toJSON());
            });
            renderTeam();
        });
        members.on("error", function() {
            renderError();
        });
        members.url = "/teams/" + tid + "/users";
        members.fetch();
    }

    function clear() {
        render("");
    }

    function renderList() {
        if (team.list.joined == null || team.list.created == null) {
            return;
        }
        var joined = {
            list: []
        };
        team.list.joined.each(function(t) {
            joined.list.push(t.toJSON());
        });

        var created = {
            list: []
        };
        team.list.created.each(function(t) {
            created.list.push(t.toJSON());
        });

        if (team.views.list) {
            team.views.list.render(joined, created);
        }
    }

    function renderTeam() {
        if(team.views.show) {
            team.views.show.render(team.detail);
        }
    }

    function renderEdit() {
        if(team.views.edit) {
            team.views.edit.render(team.detail);
        }
    }

    function renderError() {}

    function render(html) {
        $("#main").html(html);
    }
})();


/************************************* END **************************************/