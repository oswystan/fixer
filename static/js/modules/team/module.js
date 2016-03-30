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

(function () {
    var team = app.modules.team = {};
    team.name = "team";
    team.cur_user = {
        id: 0,
        nicky: ""
    };
    team.list = {
        joined : null,
        created : null
    };
    team.detail = {
        detail : null,
        members : null
    };
    team.init = function (eb) {
        _.extend(this, Backbone.Events);

        // for each function, we are in the team context;
        this.listenTo(eb, "show:team:list", function () {
            getList();
        });

        this.listenTo(eb, "show:team:new", function(){
            renderNew();
        });

        this.listenTo(eb, "show:team:detail", function(tid){
            team.detail.detail = team.detail.members = null;
            getTeamDetail(tid);
            getTeamMembers(tid);
        });

        this.listenTo(eb, "show:team:edit", function(tid){
            console.log("team module: show edit");
        });

        this.listenTo(eb, "current_user", function (id, nicky) {
            this.cur_user.id    = id;
            this.cur_user.nicky = nicky;
            console.log("get user:", id, nicky);
        });
    };

    var TeamModel = Backbone.Model.extend({
        defaults: {
            id              : 0,
            name            : "",
            leader_id       : 0,
            leader_name     : "",
            goal            : "",
            created_date    : "",
            bug_tab         : "",
            bug_tab_status  : "",
            status          : "",
            logo            :""
        },
        urlRoot:"/teams"
    });
    var TeamMember = Backbone.Model.extend({
        defaults: {
            id              : 0,
            nicky           : "",
            portrait        : "",
            email           : "",
            last_login_time : "",
            register_date   : ""
        },
    });

    var TeamItem = Backbone.Model.extend({
        defaults:{
            id              : 0,
            name            : "",
            leader_id       : 0,
            leader_name     : "",
            created_date    : "",
        },
    });

    var TeamList = Backbone.Collection.extend({
        url: "",
        model: TeamItem,
    });

    var TeamMembers = Backbone.Collection.extend({
        url: "",
        model: TeamMember
    });

    function getList() {
        if (team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        clear();
        team.list.joined = team.list.created = null;

        // get the joined team list
        var tlJoined = new TeamList();
        tlJoined.on("sync", function(){
            team.list.joined = this;
            renderList();
        });
        tlJoined.on("error", function(){
            console.log("error occurred");
            renderError();
        });
        tlJoined.url = "/users/" + team.cur_user.id + "/teams/joined";
        tlJoined.fetch();

        //get the created team list
        var tlCreated = new TeamList();
        tlCreated.on("sync", function(){
            team.list.created = this;
            renderList();
        });
        tlCreated.on("error", function(){
            console.log("error occurred");
        });

        tlCreated.url = "/users/" + team.cur_user.id + "/teams/created";
        tlCreated.fetch();
    }

    function getTeamDetail(tid) {
        if(team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        var teamDetail = new TeamModel({id: tid});
        teamDetail.on("sync", function () {
            team.detail.detail = this.toJSON();
            renderTeam();
        });
        teamDetail.on("error", function(){
            renderError();
        });
        teamDetail.fetch();
    }
    function getTeamMembers(tid) {
        if(team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        var members = new TeamMembers();
        members.on("sync", function(){
            team.detail.members = [];
            this.each(function(t){
                team.detail.members.push(t.toJSON());
            });
            renderTeam();
        });
        members.on("error", function(){
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
        var joined = {list:[]};
        team.list.joined.each(function(t) {
            joined.list.push(t.toJSON());
        });
        h1 = template("teams_joined", joined);
        var created = {list:[]};
        team.list.created.each(function(t) {
            created.list.push(t.toJSON());
        });
        h2 = template("teams_created", created);
        render(h1 + h2);
    }
    function renderTeam() {
        if (team.detail.detail == null || team.detail.members == null) {
            return;
        }
        console.log(team.detail);
        var html = template("team_detail", team.detail);
        render(html);
    }
    function renderNew() {
        var html = template("team_new");
        render(html);
    }
    function renderEdit() {
    }
    function renderError() {
    }
    function render(html) {
        $("#main").html(html);
    }
})();


/************************************* END **************************************/