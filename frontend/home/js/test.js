//小提示框
/*;*/

    //绑定

var vm = avalon.define({
    $id: "test",
    array: [11,22,33],
    items: [

    ],
    information: {
        index:'',
        id:'',
        speakImg:'',
        speakImg1:'',
        speakImg2:'',
        speakImg3:'',
        nickname:'',
        school:'',
        titleTime:'',
        time:'',
        content:'',
        likeNum:'',
        forwardingNum:'',
        reportNum:'',
        reportNumTag:'',
        praiseTag:'',
        forwardingNumTag:'',
        backgroundImg:'',
        headImg:'',
        dynamic:'',
        areLookingAt:'',
        followers:'',
        signature:''
    },
    timetrans: function(date){
        var date = new Date(date*1000);//如果date为10位不需要乘1000
        var Y = date.getFullYear() + '-';
        var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
        var D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate()) + ' ';
        var h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
        var m = (date.getMinutes() <10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
        var s = (date.getSeconds() <10 ? '0' + date.getSeconds() : date.getSeconds());
        return Y+M+D+h+m+s;
    },
    getData: function() {
        var self = this;
        load('dynamic', 'getdynamic', {'tag':1}, function(resultData) {
            var len = resultData['data'].length;
            resultData = resultData['data'];
            if(len >= 1){
                for(var i=0;i<len;i++){
                    self.items.push({
                        id:resultData[i].id,
                        speakImg: resultData[i].pic1,
                        speakImg1:resultData[i].pic2,
                        speakImg3: resultData[i].pic4,
                        headImg:resultData[i].headPortrait,
                        nickname:resultData[i].nickname,
                        backgroundImg:resultData[i].backgroundImage,
                        school:resultData[i].school,
                        signature:resultData[i].motto,
                        dynamic:50,
                        areLookingAt:200,
                        followers:10,
                        time:resultData[i].time,
                        titleTime:self.timetrans(resultData[i].createtime),
                        content:resultData[i].content,
                        likeNum:resultData[i].praise,
                        reportNum:resultData[i].reportNum,
                        forwardingNum:resultData[i].forwardingNum,
                        reportNumTag:resultData[i].reportNumTag,
                        praiseTag:resultData[i].praiseTag,
                        forwardingNumTag:resultData[i].forwardingNumTag
                    });
                    console.log(self.items);
                }
                //我也不知道为什么要放在在这里
                $('*').darkTooltip({
                    gravity:"north",
                });
                $("[data-toggle='tooltip']").tooltip();
                //用于处理图片流
                var $container = $('.masonry-container');
                $container.imagesLoaded( function () {
                    $container.masonry({
                        columnWidth: '.item',
                        itemSelector: '.item'
                    });
                });
                //Reinitialize masonry inside each panel after the relative tab link is clicked -
                $('a[data-toggle=tab]').each(function () {
                    var $this = $(this);
                    $this.on('shown.bs.tab', function () {
                        $container.imagesLoaded( function () {
                            $container.masonry({
                                columnWidth: '.item',
                                itemSelector: '.item'
                            });
                        });

                    }); //end shown
                });  //end each
            }else {
                layer.msg('获取数据失败', {
                    offset: 0,
                    shift: 12
                });
            }
        });
    },
    personalInformation: function (tag,index){
        this.information.index = index;
        this.information.id = tag.id;
        this.information.speakImg = tag.speakImg;
        this.information.speakImg1 = tag.speakImg1;
        this.information.speakImg2 = tag.speakImg2;
        this.information.speakImg3 = tag.speakImg3;
        this.information.nickname = tag.nickname;
        this.information.titleTime = tag.titleTime;
        this.information.backgroundImg = tag.backgroundImg;
        this.information.headImg=tag.headImg;
        this.information.time = tag.time;
        this.information.content = tag.content;
        this.information.likeNum = tag.likeNum;
        this.information.forwardingNum = tag.forwardingNum;
        this.information.reportNum = tag.reportNum;
        this.information.reportNumTag=tag.reportNumTag;
        this.information.school=tag.school;
        this.information.praiseTag=tag.praiseTag;
        this.information.forwardingNumTag=tag.forwardingNumTag;
        this.information.signature = tag.signature;
        this.information.dynamic=50;
        this.information.areLookingAt=200;
        this.information.followers=10;
        $('*').darkTooltip({

        });
    },
    evaluation:function (id,arg,index) {
        var self = this;
        load('evaluation', 'evaluation', {'id':id,'arg':arg}, function(resultData) {
            //@todo 不能只更新数组的某个属性，带观察，性能影响如何
            load('dynamic', 'getdynamic', {'id':id}, function(resultData) {
                var len = resultData['data'].length;
                resultData = resultData['data'];
                if (len >= 1) {
                    for (var i = 0; i < len; i++) {
                        self.items.$set(index, {
                            id: resultData[i].id,
                            speakImg: resultData[i].pic1,
                            speakImg1: resultData[i].pic2,
                            speakImg3: resultData[i].pic4,
                            headImg: resultData[i].headPortrait,
                            nickname: resultData[i].nickname,
                            backgroundImg: resultData[i].backgroundImage,
                            school: resultData[i].school,
                            signature: resultData[i].motto,
                            dynamic: 50,
                            areLookingAt: 200,
                            followers: 10,
                            time: resultData[i].time,
                            titleTime: resultData[i].createtime,
                            content: resultData[i].content,
                            likeNum: resultData[i].praise,
                            reportNum: resultData[i].reportNum,
                            forwardingNum: resultData[i].forwardingNum,
                            reportNumTag: resultData[i].reportNumTag,
                            praiseTag: resultData[i].praiseTag,
                            forwardingNumTag: resultData[i].forwardingNumTag
                        });
                    }
                    self.personalInformation(self.items[index], index);
                } else {
                    layer.msg('获取数据失败', {
                        offset: 0,
                        shift: 12
                    });
                }
            })
        })
    },
})

//
//我也不知道为什么要放在在这里
$('*').darkTooltip({

});
$("[data-toggle='tooltip']").tooltip();
vm.$watch('onReady', function(){
    vm.getData();
})
