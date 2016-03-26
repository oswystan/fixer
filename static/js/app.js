/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: app.js
 *    description: 
 *        created: 2016-03-24 17:02:15
 *         author: wystan
 *
 *********************************************************************************
 */

var App = function()
{
    var currentUser = {
        nicky: "wystan",
        id: 1,
        portrait: "/static/images/1.jpg"
    };

    var team = Team();
    var user = User();
    var bug  = Bug();
    var tmpl = Template();

    function makeRouter() {
        var Router = Backbone.Router.extend({
            routes: {
                // main menu routers
                "team"      : "showTeamList",
                "bug"       : "showBugList",
                "stat"      : "showUnsupported",
                "setting"   : "showSetting",
                "logout"    : "doLogout",

                // team releated routers
                "team/new"          : "showTeamNew",
                "team/:id/show"     : "showTeamDetail",
                "team/:id/modify"   : "showTeamModify",

                // user related routers
                "user/modify"   : "showUserModify",
                "user/pwd"      : "showUserPwd",
                "user/:id/show" : "showUserDetail",

                // bugs related routers
                "bug/:tid/show"         : "showBugs",
                "bug/:tid/new"          : "showBugNew",
                "bug/:tid/:bid/show"    : "showBug",
                "bug/:tid/:bid/modify"  : "showBugModify"
            },
            showTeamList : function(){
                console.log("show team list");
                var mock_data = {
                    list: [
                        {"id": 1, "name": "fox", "leader_name" : "wystan"},
                        {"id": 2, "name": "frog", "leader_name" : "winner"}
                    ]
                };

                var html = template("teams_joined", mock_data);
                $("#main").html(html);
            },
            showBugList : function(){
                console.log("show bug list");
                $("#main").html("");
            },
            showUnsupported : function(){
                console.log("unsupported");
                $("#main").html("");
            },
            showSetting : function(){
                console.log("show setting");
                $("#main").html("");
            },
            doLogout : function(){
                console.log("do logout");
                $("#main").html("");
            }
        });
        var router = new Router();
        Backbone.history.start();
    }

    function renderUserInfo(el) {
        el.find("img").attr("src", currentUser.portrait);
        el.find("a").text(currentUser.nicky);
    }

    return {
        initRouter  : makeRouter,
        initUserInfo: renderUserInfo
    };
};

var Team = function(){
    function show(tid) {
        // get data from server

        // when success render it into #main container.
    }
    function list(uid) {
        // body...
    }
    function modify(tid) {
        // body...
    }
    function create(tid) {
        // body...
    }

    var TView = function() {
        function render(t) {

        }
        return {

        };
    }

    var TLView = function() {
        return {};
    }

    return {};
};

var User = function(){
    return {};
};
var Bug = function(){
    return {};
};
var Template = function(){
    return {};
};
var Stat = function(){
    return {};
};

/************************************* END **************************************/
