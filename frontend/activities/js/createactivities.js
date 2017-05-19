var row_tsf_index = 0;
/**删除行对象*/
function delRow(row_id){
    $(row_id).remove(); // 删除匹配row_id的元素
}

function addNewTsf(){
    var row_id = "tsf" + row_tsf_index;// 所插入行的id
    var namea = "fielda"+row_tsf_index;
    var nameb = "fieldb"+row_tsf_index;
    var row_obj = "<section class='form-group'  id='"+row_id+"' style='margin-left: 20%;'><label class='col-sm-2 control-label' style='width: auto;margin-top: 7px; font-size: medium;'>字段:</label> <input type='text' class='form-control' placeholder='' name= '"+namea+"' style='width: auto;float: left;height: 30px;margin-top: 5px;'><label class='col-sm-2 control-label' style='width: auto;margin-top: 7px; font-size: medium;'>描述:</label><input type='text' class='form-control'  name= '"+nameb+"' placeholder='' style='width: auto;float: left;height: 30px;margin-top: 5px;'><img src='/activities/img/sub.png' style='cursor: pointer; margin-top: 7px;' onclick='delRow("+row_id+");'><br>";
    $("#tsfto").after(row_obj); // 插入行
    row_tsf_index = row_tsf_index + 1;
}

function genSearchParams()
{
    var searchParams = $("#frmSearch").serializeArray();
    return searchParams;
}

$(document).ready(function() {
    $("#create").click(function () {
        var tag = 0;
        var registeredArgs =genSearchParams();
        var i=0;
        $.each(registeredArgs, function(n, value) {
            if(i % 2 == 0){
                if(value.value == ''){
                    tag = 1;
                }
                console.log(value.name);
            }
            i++;
        });
        if(!tag){
            load('activities', 'createactivities', registeredArgs, function(resultData) {
                layer.msg('创建成功', {
                    offset: 0,
                    shift: 12
                });
            });
        }else {
            layer.msg('所有的字段名称必填，或移除不需要的字段', {
                offset: 0,
                shift: 12
            });
        }
    });
})
