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


var UserModel = Backbone.Model.extend({
    defaults: {
        id              : 0,
        nicky           : '',
        email           : '',
        portrait        : '',
        register_date   : '',
        lastlogin_time  : ''
    },
    urlRoot: '/users'
});

var UserCollection = Backbone.Collection.extend({
    model: UserModel,
    url: '/users'
});

/************************************* END **************************************/
