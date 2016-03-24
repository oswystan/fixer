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

    function makeRouter() {
        var Router = Backbone.Router.extend({
            routes: {
                "team"      : "showTeamList",
                "bug"       : "showBugList",
                "stat"      : "showUnsupported",
                "setting"   : "showSetting",
                "logout"    : "doLogout"
            },
            showTeamList : function(){
                console.log("show team list");
            },
            showBugList : function(){
                console.log("show bug list");
            },
            showUnsupported : function(){
                console.log("unsupported");
            },
            showSetting : function(){
                console.log("show setting");
            },
            doLogout : function(){
                console.log("do logout");
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
        initRouter: makeRouter,
        initUserInfo: renderUserInfo
    };
};

/************************************* END **************************************/
