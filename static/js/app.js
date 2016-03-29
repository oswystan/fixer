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

/*
var App = function()
{
    var currentUser = {
        nicky: "wystan",
        id: 1,
        portrait: "/static/images/1.jpg"
    };

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
*/

/*
console.log("show bug list");

this.showProgress();

console.log("start...");
var users = new UserCollection();
users.on("sync", function () {
    console.log("sync users done.");
    _.each(users.models, function (u) {
        console.log(u.toJSON());
    });
});
users.fetch({   
    data: $.param({"q":"j"}),
    reset: true,
});

var user = new UserModel({id: 1});
user.on("sync", function (resp) {
    router.clearMain();
});
user.on("error", function (u, resp) {
    router.showError(resp.responseJSON);
});
user.fetch();

showProgress: function () {
    var html = template("work_in_progress", null);
    $("#main").html(html);
},
showError: function (err) {
    var html = template("error_occurred", err);
    $("#main").html(html);
},
clearMain: function () {
    $("#main").html("");
},
*/

var App = function () {
    return {
        init: function () {
            console.log("app init...");
            _.extend(this.event_bus, Backbone.Events);
            _.each(this.modules, function (m, key, that) {
                m.init(this.event_bus);
            }, this);
            console.log("done.");

            this.event_bus.trigger("current_user", "1", "john");
        },
        modules: {},
        event_bus: {}
    };
};

var app = new App();

/************************************* END **************************************/
