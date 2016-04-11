/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: search.js
 *    description: 
 *        created: 2016-04-03 17:28:34
 *         author: wystan
 *
 *********************************************************************************
 */

(function() {
    var search = app.namespace("app.modules.team.search");

    var editorbox = function(eb, input) {
        _.extend(this, Backbone.Events);
        var that = this;
        this.input = input;
        this.eb = eb;
        this.name = null;

        this.done = function(err, data) {
            this.eb.trigger("searchdone:team", err, data);
        };

        $(input).unbind("change").change(function() {
            var name = $.trim(this.value);
            if (name.length === 0) {
                return;
            }
            that.eb.trigger("search:team", name, that, that.done);
        });

        this.listenTo(eb, "candinates:team:choose", function(data) {
            if (data !== null) {
                var len1 = input.value.length;
                var len2 = data.name.length;
                $(input).val(data.name);
                console.log(len1, len2);
                if (len1 < len2) {
                    input.setSelectionRange(len1, len2);
                }

            }
            this.team = data;
        });
    };

    var candinates = function(eb) {
        _.extend(this, Backbone.Events);
        this.eb = eb;
        this.listenTo(eb, "searchdone:team", function(err, data) {
            this.error = err;
            this.result = data;
            if (err === null) {
                this.render();
            }
        });

        this.render = function() {
            //TODO
            if (this.result && this.result.length > 0) {
                this.eb.trigger("candinates:team:choose", this.result[0]);
            } else {
                this.eb.trigger("candinates:team:choose", null);
            }
        };
    };

    var teams = function(eb) {
        _.extend(this, Backbone.Events);
        this.eb = eb;
        var collection = app.namespace("app.modules.team.models.list");
        this.collection = new collection();
        this.listenTo(eb, "search:team", function(name, who, cb) {
            var c = this.collection;
            c.url = "/teams"
            c.reset();
            c.once("sync", function() {
                var data = [];
                this.each(function(t) {
                    data.push(t.toJSON());
                });
                cb.call(who, null, data);
            });
            c.once("error", function() {
                var err = {};
                cb.call(who, err, null);
            });
            c.fetch({
                data: $.param({
                    q: name
                })
            });
        });
    };

    var Search = function(input) {
        this.eb = {};
        _.extend(this.eb, Backbone.Events);
        this.editor = new editorbox(this.eb, input);
        this.candinates = new candinates(this.eb);
        this.teams = new teams(this.eb);
        this.current = function() {
            return this.editor.team;
        };
    };

    search.create = function(input) {
        return new Search(input);
    };

})();



/************************************* END **************************************/