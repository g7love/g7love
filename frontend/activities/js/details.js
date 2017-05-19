var $table = $('#table'), $remove = $('#remove'), selections = [];
//获取url中的参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg); //匹配目标参数
    if (r != null) return unescape(r[2]); return null; //返回参数值
}
var id= getUrlParam('id');
$(document).ready(function(){
    var columns=[];
    load('activities', 'column', {'id':id}, function(resultData) {
        var len = resultData['data'].length;
        var result = resultData['data'];
        for(var i = 0; i < len; i++){
            columns[i]={field:result[i][0],title:result[i][1], align: 'center'};

        }
        function initTable() {
            $table.bootstrapTable({
                url:"/index.php?r=activities/getdetails&id="+id,
                height: getHeight(),
                columns: columns,
            });
            // sometimes footer render error.
            setTimeout(function () {
                $table.bootstrapTable('resetView');
            }, 200);
        }

        function getHeight() {
            return $(window).height() - $('h1').outerHeight(true)-100;
        }
        var scripts = [
                '/public/js/bootstraptable/js/bootstrap-table.js',
                '/public/js/bootstraptable/js/bootstrap-table-export.js',
                '/public/js/bootstraptable/js/tableExport.js',
            ],
            eachSeries = function (arr, iterator, callback) {
                callback = callback || function () {};
                if (!arr.length) {
                    return callback();
                }
                var completed = 0;
                var iterate = function () {
                    iterator(arr[completed], function (err) {
                        if (err) {
                            callback(err);
                            callback = function () {};
                        }
                        else {
                            completed += 1;
                            if (completed >= arr.length) {
                                callback(null);
                            }
                            else {
                                iterate();
                            }
                        }
                    });
                };
                iterate();
            };
        eachSeries(scripts, getScript, initTable);
    });
});

function getScript(url, callback) {
    var head = document.getElementsByTagName('head')[0];
    var script = document.createElement('script');
    script.src = url;
    var done = false;
    // Attach handlers for all browsers
    script.onload = script.onreadystatechange = function() {
        if (!done && (!this.readyState ||
            this.readyState == 'loaded' || this.readyState == 'complete')) {
            done = true;
            if (callback)
                callback();

            // Handle memory leak in IE
            script.onload = script.onreadystatechange = null;
        }
    };
    head.appendChild(script);
    return undefined;
}