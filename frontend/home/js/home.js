var flg = 0;

load('home', 'loginjudge', {}, function(resultData) {
    flg=resultData['data'];
    if(flg){
        $("#postingId").show();
        $("#loginId").hide();
        $("#regis").hide();
    }else {
        $("#postingId").hide();
        $("#loginId").show();
        $("#regis").show();
    }
});

//写入Cookie
function setCookie(name, value, seconds) {
    seconds = seconds || 0;   //seconds有值就直接赋值，没有为0，这个根php不一样。
    var expires = "";
    if (seconds != 0 ) {      //设置cookie生存时间
        var date = new Date();
        date.setTime(date.getTime()+(seconds*1000));
        expires = "; expires="+date.toGMTString();
    }
    document.cookie = name+"="+escape(value)+expires+"; path=/";   //转码并赋值
}
//获取Cookie
function getCookie(name)
{
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
    if(arr=document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}

function home() {
    var href = '/home/index.html';
    setCookie('url',href,500000);
    $('iframe').attr('src', href);
}
function personal() {
    if(flg){
        var href = '/personal/index.html';
        setCookie('url',href,500000);
        $('iframe').attr('src', href);
    }else {
        layer.msg('请登陆后再操作', {
            offset: 0,
            shift: 12
        });
    }
}

function mapchat() {
    if(flg){
        var href = '/mapchat/index.html';
        setCookie('url',href,500000);
        $('iframe').attr('src', href);
    }else {
        layer.msg('请登陆后再操作', {
            offset: 0,
            shift: 12
        });
    }
}

function createactivities() {
    if(flg){
        var href = '/activities/createactivities.html';
        setCookie('url',href,500000);
        $('iframe').attr('src', href);
    }else {
        layer.msg('请登陆后再操作', {
            offset: 0,
            shift: 12
        });
    }
}
function activities() {
    if(flg){
        var href = '/activities/index.html';
        setCookie('url',href,500000);
        $('iframe').attr('src', href);
    }else {
        layer.msg('请登陆后再操作', {
            offset: 0,
            shift: 12
        });
    }
}

function genSearchParams()
{
    var searchParams = $("#frmSearch").serializeArray();
    return searchParams;
}

function loginForm()
{
    var searchParams = $("#login").serializeArray();
    return searchParams;
}

function posting(args) {
    if(args.length == 0 && args['count'] ==''){
        layer.msg('不能发表空白说说', {
            offset: 0,
            shift: 12
        });
    }else {
        load('home', 'posting', {'data':args,'count':args['count']}, function(resultData) {
            $("#post").click();
            personal();
        });
    }
}
$(function () {
    if(getCookie('url') == null){
        var href = '/home/index.html';
        $('iframe').attr('src', href);
    }else {
        var href = getCookie('url');
        $('iframe').attr('src', href);
    }

    //注册
    load('registered', 'provinces', {}, function(resultData) {
        resultData = resultData['data'];
        var temp = '';
        for(var i =0; i<resultData.length; i++){
            temp += '<option value ='+resultData[i].id+'>'+resultData[i].name+'</option>';
        }
        $("#provinces").append(temp);
    });
    $("#provinces").change( function() {
        var provinces = $('#provinces option:selected').val();
        load('registered', 'getschool', {'parentid':provinces}, function(resultData) {
            resultData = resultData['data'];
            var temp = '';
            for(var i =0; i<resultData.length; i++){
                temp += '<option value ='+resultData[i].id+'>'+resultData[i].name+'</option>';
            }
            $("#school").empty();
            $("#school").append(temp);
        });
    });

    var year = '',mouth='',day='';
    for(var i = 1980; i< 2016; i++){
        year += '<option value ='+i+'>'+i+'</option>';
    }
    $("#year").append(year);
    for(var i = 1; i< 13; i++){
        mouth += '<option value ='+i+'>'+i+'</option>';
    }
    $("#mouth").append(mouth);
    for(var i = 1; i< 32; i++){
        day += '<option value ='+i+'>'+i+'</option>';
    }
    $("#day").append(day);
    $("#registered").click(function () {
        var tag = 0;
        var registeredArgs =genSearchParams();
        if(isNaN(registeredArgs[0].value)){
            layer.tips('请选择省份', '#provinces', {
                tips: [4, '#78BA32'],
                tipsMore: true
            });
            tag = 1;
        }
        if(isNaN(registeredArgs[1].value)){
            layer.tips('请选择学校', '#school', {
                tips: [4, '#78BA32'],
                tipsMore: true
            });
            tag = 1;
        }
        var reg = /(^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+$)|(^$)/;
        if(registeredArgs[2].value.length == 0 || !reg.test(registeredArgs[2].value)){
            layer.tips('请填写正确的邮箱（xx@xx.com）', '#inputEmail', {
                tips: [4, '#78BA32'],
                tipsMore: true
            });
            tag = 1;
        }
        if(registeredArgs[3].value.length < 1 || registeredArgs[3].value.length > 10 ){
            layer.tips('昵称字数个数（1-10)', '#nickname', {
                tips: [4, '#78BA32'],
                tipsMore: true
            });
            tag = 1;
        }
        if(!tag){
            load('registered', 'registered', registeredArgs, function(resultData) {
                if(resultData['data'] == 1){
                    layer.msg('注册成功', {
                        offset: 0,
                        shift: 12
                    });
                }else {
                    layer.msg('注册失败', {
                        offset: 0,
                        shift: 12
                    });
                }
            });
        }
    });

    $("#submit").click(function () {
        var tag = 0;
        var login =loginForm();
        if(login[0].value.length < 1 || login[0].value.length > 10 ){
            layer.tips('用户名不能为空', '#username', {
                tips: [4, '#78BA32'],
                tipsMore: true
            });
            tag = 1;
        }
        if(login[1].value.length < 1 || login[1].value.length > 10 ){
            layer.tips('密码不能为空', '#password', {
                tips: [4, '#78BA32'],
                tipsMore: true
            });
            tag = 1;
        }
        if(!tag){
            load('login', 'login', login, function(resultData) {
                if(resultData['data']){
                    layer.msg(resultData['data'], {
                        offset: 0,
                        shift: 12
                    });
                    window.location.reload();
                }else {
                    layer.msg('登陆失败，用户名或密码错误', {
                        offset: 0,
                        shift: 12
                    });
                }
            });
        }
    })
    var img=[], i=0;
    $("#zyupload").zyUpload({
        width            :   "650px",                 // 宽度
        height           :   "400px",                 // 宽度
        itemWidth        :   "140px",                 // 文件项的宽度
        itemHeight       :   "115px",                 // 文件项的高度
        url              :   "/index.php?r=upload/upload",  // 上传文件的路径
        fileType         :   ["jpg","png"],// 上传文件的类型
        fileSize         :   11200000,                // 上传文件的大小
        multiple         :   true,                    // 是否可以多个文件上传
        dragDrop         :   true,                    // 是否可以拖动上传文件
        tailor           :   false,                    // 是否可以裁剪图片
        del              :   true,                    // 是否可以删除文件
        finishDel        :   false,  				  // 是否在上传文件完成后删除预览
        /* 外部获得的回调接口 */
        onSelect: function(selectFiles, allFiles){    // 选择文件的回调方法  selectFile:当前选中的文件  allFiles:还没上传的全部文件
            /*console.info("当前选择了以下文件：");
            console.info(selectFiles);*/
        },
        onDelete: function(file, files){              // 删除一个文件的回调方法 file:当前删除的文件  files:删除之后的文件
            /*console.info("当前删除了此文件：");
            console.info(file.name);*/
        },
        onSuccess: function(file, response){          // 文件上传成功的回调方法
            img[i++]=response;
            /*console.info("此文件上传成功：");
            console.info(file.name);
            console.info("此文件上传到服务器地址：");
            console.info(response);
            $("#uploadInf").append("<p>上传成功，文件地址是：" + response + "</p>");*/
        },
        onFailure: function(file, response){          // 文件上传失败的回调方法
            /*console.info("此文件上传失败：");
            console.info(file.name);*/
        },
        onComplete: function(response){           	  // 上传完成的回调方法
            img['count'] = $("#exampleInput").val();
            posting(img);
            /*console.info("文件上传完成");
            console.info(response);*/
        }
    });


});
