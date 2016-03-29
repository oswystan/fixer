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

var App = function () {
    return {
        init: function () {
            console.log("app init...");
            _.extend(this.event_bus, Backbone.Events);
            _.each(this.modules, function (m, key, that) {
                m.init(this.event_bus);
                console.log("module " + key + " is initilized.")
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
