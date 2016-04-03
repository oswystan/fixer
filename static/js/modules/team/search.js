/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: search.js
 *    description: 
 *        created: 2016-04-03 12:14:22
 *         author: wystan
 *
 *********************************************************************************
 */

(function(eb) {
    var search = app.namespace("app.modules.team.search");
    var editorbox = function(eb, input) {
        _.extend(this, Backbone.Events);
        var that = this;
        this.input = input;
        this.eb = eb;
        this.user = null;

        this.done = function(err, data) {
            this.eb.trigger("searchdone:user", err, data);
        };

        $(input).unbind("keyup").keyup(function() {
            var nicky = $.trim(this.value);
            if (nicky.length === 0) {
                return;
            }
            console.log(nicky);
            that.eb.trigger("search:user", nicky, that, that.done);
        });

        this.listenTo(eb, "candinates:user:choose", function(data) {
            if(data !== null) {
                $(input).val(data.nicky);
            }
            this.user = data;
        });
    };

    var candinates = function(eb) {
        _.extend(this, Backbone.Events);
        this.eb = eb;
        this.listenTo(eb, "searchdone:user", function(err, data) {
            this.error = err;
            this.result = data;
            if (err === null) {
                this.render();
            }
        });

        this.render = function() {
            //TODO
            if (this.result && this.result.length > 0) {
                this.eb.trigger("candinates:user:choose", this.result[0]);
            }else {
                this.eb.trigger("candinates:user:choose", null);
            }
        };
    };

    var users = function (eb) {
        _.extend(this, Backbone.Events);
        this.eb = eb;
        var collection = app.namespace("app.modules.user.models.collection");
        this.collection = new collection();
        this.listenTo(eb, "search:user", function (nicky, who, cb) {
            var c = this.collection;
            c.reset();
            c.once("sync", function () {
                var data = [];
                this.each(function (t) {
                    data.push(t.toJSON());
                });
                console.log("sync", data);
                cb.call(who, null, data);
            });
            c.once("error", function () {
                var err = {};
                cb.call(who, err, null);
            })
            c.fetch({data: $.param({q:nicky})});
        });
    };

    var Search = function(eb, input) {
        _.extend(this, Backbone.Events);
        this.editor = new editorbox(eb, input);
        this.candinates = new candinates(eb);
        this.users = new users(eb);
        this.current = function () {
            return this.editor.user;
        };
    };

    search.create = function(eb, input) {
        return new Search(eb, input);
    };

})();



/************************************* END **************************************/