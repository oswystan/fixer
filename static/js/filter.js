/*
 ************************************************
 *
 *
 *
 ************************************************
 */

var baseUrl = "/filter";

var TeamListFilter = {
    creator: "",
    member: "",
    creatorId: -1,
    memberId: -1,
    success: 0,
    result: "",

    go: function () {
        var url = baseUrl + "/teamlist.html";
        if (this.creator != "") {
            url = url + "?creator=" + this.creator;
        } else if (this.member != "") {
            url = url + "?member=" + this.member;
        } else if (this.creatorId != -1) {
            url = url + "?creatorid=" + this.creatorId;
        } else if (this.memberId != -1) {
            url = url + "?memberid=" + this.memberId;
        }

        var ret = $.ajax({url: url, async: false, cache: false});
        if (ret.status == 200) {
            this.success = 1;
            this.result = ret.responseText;
        } else {
            this.success = 0;
            this.result = "";
        }
    },

    json: function () {
        return eval("(" + this.result + ")");
    },

    reset: function () {
        with (this) {
            creator = "";
            member = "";
            creatorId = -1;
            memberId = -1;
            success = 0;
            result = "";
        }
    },
};

var TeamFilter = {
    name: "",
    id: -1,
    success: 0,
    result: "",

    go: function () {
        var url = baseUrl + "/teamlist.html";
        if (this.name != "") {
            url = url + "?name=" + this.name;
        } else if (this.id != -1) {
            url = url + "?id=" + this.id;
        }

        var ret = $.ajax({url: url, async: false, cache: false});
        if (ret.status == 200) {
            this.success = 1;
            this.result = ret.responseText;
        } else {
            this.success = 0;
            this.result = "";
        }
    },

    json: function () {
        return eval("(" + this.result + ")");
    },

    reset: function () {
        with (this) {
            name = "";
            id = -1;
            success = 0;
            result = "";
        }
    },
};

/*
 ************************************************
 *
 *
 *
 ************************************************
 */
var MemberListFilter = {
    team_id: -1,
    success: 0,
    result: "",

    go: function () {
        var url = baseUrl + "/memberlist.html";
        if (this.team_id != -1) {
            url = url + "?team_id=" + this.team_id;
        }

        var ret = $.ajax({url: url, async: false, cache: false});
        if (ret.status == 200) {
            this.success = 1;
            this.result = ret.responseText;
        } else {
            this.success = 0;
            this.result = "";
        }
    },

    json: function () {
        return eval("(" + this.result + ")");
    },

    reset: function () {
        with (this) {
            team_id = -1;
            success = 0;
            result = "";
        }
    },
};


/*
 ************************************************
 *
 *
 *
 ************************************************
 */

var UserFilter = {
    nicky: "",
    id: -1,
    success: 0,
    result: "",

    go: function () {
        var url = baseUrl + "/user-detail.html";
        if (this.nicky != "") {
            url = url + "?nicky=" + this.nicky;
        } else if (this.id != -1) {
            url = url + "?id=" + this.id;
        }

        var ret = $.ajax({url: url, async: false, cache: false});
        if (ret.status == 200) {
            this.success = 1;
            this.result = ret.responseText;
        } else {
            this.success = 0;
            this.result = "";
        }
    },

    json: function () {
        return eval("(" + this.result + ")");
    },

    reset: function () {
        with (this) {
            nicky = "";
            id = -1;
            success = 0;
            result = "";
        }
    },
};

/*
 ************************************************
 *
 *
 *
 ************************************************
 */

var BugListFilter = {
    team_id: -1,
    priority: "",
    handler: "",
    created_by: "",
    status: "",
    date_from: "",
    date_to: "",
    success: 0,
    result: "",

    go: function () {
        var url = baseUrl + "/buglist.html";
        url = url + (this.team_id != -1 ? "?team_id=" + this.team_id : "");
        url = url + (this.priority != "" ? "?priority=" + this.priority : "");
        url = url + (this.handler != "" ? "?handler=" + this.handler : "");
        url = url + (this.created_by != "" ? "?created_by=" + this.created_by : "");
        url = url + (this.status != "" ? "?status=" + this.status : "");
        url = url + (this.date_from != "" ? "?date_from=" + this.date_from : "");
        url = url + (this.date_to != "" ? "?date_to=" + this.date_to : "");


        var ret = $.ajax({url: url, async: false, cache: false});
        if (ret.status == 200) {
            this.success = 1;
            this.result = ret.responseText;
        } else {
            this.success = 0;
            this.result = "";
        }
    },

    json: function () {
        return eval("(" + this.result + ")");
    },

    reset: function () {
        with (this) {
            team_id = -1;
            priority = "";
            handler = "";
            created_by = "";
            status = "";
            date_from = "";
            date_to = "";
            success = 0;
            result = "";
        }
    },
};


/*
 ************************************************
 *
 *
 *
 ************************************************
 */
Date.prototype.Format = function(fmt) {

    str = fmt;
    str = str.replace(/yyyy|YYYY/, this.getFullYear());

    return str;
}


/*
 ************************************************
 *
 *
 *
 ************************************************
 */


/*
 ************************************************
 *
 *
 *
 ************************************************
 */


/*
 ************************************************
 *
 *
 *
 ************************************************
 */


/*
 ************************************************
 *
 *
 *
 ************************************************
 */

