/*
 *********************************************************************************
 *                     Copyright (C) 2016 wystan
 *
 *       filename: helper.js
 *    description: 
 *        created: 2016-03-29 19:29:41
 *         author: wystan
 *
 *********************************************************************************
 */

template.helper('dateFormat', function(d, fmt) {
    var date = new Date(d);
    str = fmt;
    str = str.replace(/yyyy|YYYY/, date.getFullYear());
    str = str.replace(/mm/, date.getMonth() >= 9 ? date.getMonth() + 1 : '0' + (date.getMonth() + 1));
    str = str.replace(/dd/, date.getDate() > 9 ? date.getDate() : '0' + date.getDate());
    str = str.replace(/HH/, date.getHours() > 9 ? date.getHours() : '0' + date.getHours());
    str = str.replace(/MM/, date.getMinutes() > 9 ? date.getMinutes() : '0' + date.getMinutes());
    str = str.replace(/SS/, date.getSeconds() > 9 ? date.getSeconds() : '0' + date.getSeconds());
    return str;
});



/************************************* END **************************************/