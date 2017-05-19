
function genSearchParams()
{
    var searchParams = $("#frmSearch").serializeArray();
    return searchParams;
}
$(document).ready(function() {
    load('registered', 'provinces', {}, function(resultData) {
        resultData=resultData['data']
        var temp = '';
        for(var i =0; i<resultData.length; i++){
            temp += '<option value ='+resultData[i].id+'>'+resultData[i].name+'</option>';
        }
        $("#provinces").append(temp);
    });
    $("#provinces").change( function() {
        var provinces = $('#provinces option:selected').val();
        load('registered', 'getschool', {'parentid':provinces}, function(resultData) {
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
                if(resultData == 1){
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
})