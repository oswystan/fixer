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
                console.log(m.name, "initilized.");
            }, this);
            console.log("done.");

            this.event_bus.trigger("current_user", "1", "john");
        },
        namespace: function(name) {
            var parts = name.split(".");
            var tmp = this;
            _.each(parts, function(p) {
                if (p == "app") {
                    return tmp;
                }
                tmp[p] = tmp[p] || {};
                tmp = tmp[p];
            });
            tmp.name=name;
            return tmp;
        },
        event_bus: {}
    };
};

var app = new App();

/************************************* END **************************************/
