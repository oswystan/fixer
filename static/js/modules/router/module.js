/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: module.js
 *    description: application routers, it will trigger event bus to notify other
 *                 components of the application, for example: team, bug, user etc.
 *        created: 2016-03-28 23:00:17
 *         author: wystan
 *
 *********************************************************************************
 */


(function () {
    var router = app.namespace("app.modules.router");

    router.init = function (eb) {
        var Router = Backbone.Router.extend({
            routes: {
                // main menu routers
                "team"      : "showTeamList",
                "bug"       : "showBugList",
                "stat"      : "showStat",
                "setting"   : "showSetting",
                "logout"    : "doLogout",

                // team releated routers
                "team/new"          : "showTeamNew",
                "team/:id/show"     : "showTeamDetail",
                "team/:id/modify"   : "showTeamEdit",

                // user related routers
                "user/modify"   : "showUserEdit",
                "user/pwd"      : "showUserPwd",
                "user/:id/show" : "showUserDetail",

                // bugs related routers
                "bug/:tid/show"         : "showBugs",
                "bug/:tid/new"          : "showBugNew",
                "bug/:tid/:bid/show"    : "showBug",
                "bug/:tid/:bid/modify"  : "showBugEdit"
            },

            showTeamList : function(){
                eb.trigger("show:team:list");
            },
            showBugList : function(){
                eb.trigger("show:bug");
            },
            showStat: function(){
                eb.trigger("show:stat");
            },
            showSetting : function(){
                eb.trigger("setting");
            },
            doLogout : function(){
                eb.trigger("logout");
            },

            /**
             * team related operations
             */
            showTeamNew: function() {
                eb.trigger("show:team:new");
            },
            showTeamDetail: function(tid) {
                eb.trigger("show:team:detail", tid);
            },
            showTeamEdit: function(tid) {
                eb.trigger("show:team:edit", tid);
            },

            /**
             * user related operations (TODO)
             */
            showUserDetail: function(uid) {
                eb.trigger("show:user:detail", uid);
            },
            showUserPwd: function () {
                eb.trigger("show:user:pwd");
            },
            showUserEdit: function () {
                eb.trigger("show:user:edit");
            },


            /**
             * bug related operations (TODO)
             */
            showBugs: function(tid){
                eb.trigger("show:bug:list", tid);
            },
            showBug: function (tid, bid) {
                eb.trigger("show:bug:detail", tid, bid);
            },
            showBugEdit: function (tid, bid) {
                eb.trigger("show:bug:edit", tid, bid);
            },
            showBugNew: function (tid) {
                eb.trigger("show:bug:new", tid);
            },
        });
        var router = new Router();
        Backbone.history.start();
    };
})();


/************************************* END **************************************/
