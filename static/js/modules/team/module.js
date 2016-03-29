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
    team.init = function (eb) {
        _.extend(this, Backbone.Events);

        // for each function, we are in the team context;
        this.listenTo(eb, "show:team:list", function () {
            console.log("do show team list in team module");
            getList();
        });

        this.listenTo(eb, "show:team:new", function(){
            renderNew();
        });

        this.listenTo(eb, "show:team:detail", function(tid){
            console.log("team module: show detail");
        });

        this.listenTo(eb, "show:team:edit", function(tid){
            console.log("team module: show edit");
        });

        this.listenTo(eb, "current_user", function (id, nicky) {
            this.cur_user.id   = id;
            this.cur_user.nicky = nicky;
            console.log("get user:", this.cur_user);
        });
    };

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

    function getList() {
        if (team.cur_user.id == 0) {
            console.log("ERROR: set current user first");
            return;
        }

        clearList();
        var tl = new TeamList();
        tl.on("sync", function(){
            renderList("teams_joined", this);
        });
        tl.on("error", function(){
            console.log("error occurred");
        });

        tl.url = "/users/" + team.cur_user.id + "/teams/joined";
        tl.fetch();

        var tlCreated = new TeamList();
        tlCreated.on("sync", function(){
            renderList("teams_created", this);
        });
        tlCreated.on("error", function(){
            console.log("error occurred");
        });

        tlCreated.url = "/users/" + team.cur_user.id + "/teams/created";
        tlCreated.fetch();
    }

    function clearList() {
        $("#main").html("");
    }
    function renderList(which, teams) {
        var tl = {list:[]};
        teams.each(function(t) {
            tl.list.push(t.toJSON());
        });

        var old = $("#main").html();
        var html = template(which, tl);
        $("#main").html(old + html);
    }
    function renderNew() {
        var html = template("team_new");
        $("#main").html(html);
    }
})();


/************************************* END **************************************/
