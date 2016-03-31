/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: model.js
 *    description: 
 *        created: 2016-03-29 16:38:59
 *         author: wystan
 *
 *********************************************************************************
 */


(function() {
    var models = app.namespace("app.modules.team.models");

    models.list_item = Backbone.Model.extend({
        defaults: {
            id: 0,
            name: "",
            leader_id: 0,
            leader_name: "",
            created_date: "",
        },
    });

    models.list = Backbone.Collection.extend({
        url: "",
        model: models.list_item,
    });

    models.team = Backbone.Model.extend({
        defaults: {
            id: 0,
            name: "",
            leader_id: 0,
            leader_name: "",
            goal: "",
            created_date: "",
            bug_table: "",
            bug_table_status: "",
            status: "",
            logo: ""
        },
        urlRoot: "/teams"
    });

    models.member = Backbone.Model.extend({
        defaults: {
            id: 0,
            nicky: "",
            portrait: "",
            email: "",
            last_login_time: "",
            register_date: ""
        },
    });

    models.members = Backbone.Collection.extend({
        url: "",
        model: models.member
    });
})();



/************************************* END **************************************/