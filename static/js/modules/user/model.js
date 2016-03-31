/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: model.js
 *    description: 
 *        created: 2016-03-28 13:42:40
 *         author: wystan
 *
 *********************************************************************************
 */

(function() {
    var models = app.namespace("app.modules.user.models");

    models.user = Backbone.Model.extend({
        defaults: {
            id: 0,
            nicky: '',
            email: '',
            portrait: '',
            register_date: '',
            lastlogin_time: ''
        },
        urlRoot: '/users'
    });

    models.collection = Backbone.Collection.extend({
        model: models.user,
        url: '/users'
    });

})();


/************************************* END **************************************/