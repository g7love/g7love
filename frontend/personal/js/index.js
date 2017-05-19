//小提示框
$(document).ready(function(){
    //绑定数据
    Vue.config.async = false;
    var homedata = new Vue({
        el: '#homedata',
        data: {
            items: [

            ],
            information: {
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
            },
            informationData:{
                headPortrait:'',
                backgroundImage:'',
                motto:'',
                nickname:'',
                birthday:'',
                gender:'',
                school:'',
                self:'',
                createtime:'',
                thumb:'',
            },
        },
        created: function() {

        },
        ready: function() {
            this.getInformation();
            this.getData();
        },
        methods: {
            getData: function() {
                var self = this;
                load('dynamic', 'getdynamic', {}, function(resultData) {
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
                                titleTime:resultData[i].createtime,
                                content:resultData[i].content,
                                likeNum:resultData[i].praise,
                                reportNum:resultData[i].reportNum,
                                forwardingNum:resultData[i].forwardingNum,
                                reportNumTag:resultData[i].reportNumTag,
                                praiseTag:resultData[i].praiseTag,
                                forwardingNumTag:resultData[i].forwardingNumTag
                            });
                        }
                        $('*').darkTooltip({

                        });
                        $("[data-toggle='tooltip']").tooltip();
                    }else {
                        layer.msg('获取数据失败', {
                            offset: 0,
                            shift: 12
                        });
                    }
                });
            },
            //res.speakImg,res.speakImg1,res.speakImg2,res.speakImg3,res.nickname,res.school,res.titleTime,res.time,res.content,res.likeNum,res.forwardingNum,res.reportNum，res.reportNumTag,res.praiseTag,res.forwardingNumTag
            personalInformation: function (tag){
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
                this.information.dynamic=50;
                this.information.areLookingAt=200;
                this.information.followers=10;
                $('*').darkTooltip({

                });
            },
            evaluation:function (id,arg,index) {
                var self = this;
                load('dynamic', 'evaluation', {'id':id,'arg':arg}, function(resultData) {
                    //@todo 不能只更新数组的某个属性，带观察，性能影响如何
                    load('dynamic', 'getdynamic', {'id':id}, function(resultData) {
                        var len = resultData['data'].length;
                        resultData = resultData['data'];
                        if(len >= 1){
                            for(var i=0;i<len;i++){
                                self.items.$set(index,{
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
                                    titleTime:resultData[i].createtime,
                                    content:resultData[i].content,
                                    likeNum:resultData[i].praise,
                                    reportNum:resultData[i].reportNum,
                                    forwardingNum:resultData[i].forwardingNum,
                                    reportNumTag:resultData[i].reportNumTag,
                                    praiseTag:resultData[i].praiseTag,
                                    forwardingNumTag:resultData[i].forwardingNumTag
                                });
                            }
                        }else {
                            layer.msg('获取数据失败', {
                                offset: 0,
                                shift: 12
                            });
                        }
                    });

                });
            },
            getInformation:function () {
                var self = this;
                load('personal', 'getuserinformation', {}, function(resultData) {
                    resultData = resultData['data'];
                    self.informationData.headPortrait=resultData.headPortrait;
                    self.informationData.backgroundImage=resultData.backgroundImage;
                    self.informationData.motto=resultData.motto;
                    self.informationData.nickname=resultData.nickname;
                    self.informationData.birthday=resultData.birthday.substring(0,10);
                    self.informationData.gender=resultData.gender;
                    self.informationData.school=resultData.school;
                    self.informationData.self=resultData.self;
                    self.informationData.createtime=resultData.createtime.substring(0,10);
                    self.informationData.thumb=resultData.thumb;
                })
            },
            thumbUp:function () {
                var self = this;
                load('personal', 'thumbup', {}, function(resultData) {
                    if(resultData.data.thump == 0){
                        layer.msg('亲，每天只有一次机会，明天再来', {
                            offset: 0,
                            shift: 12,
                        });
                    }else {
                        self.getInformation();
                    }
                })
            }

        }
    });

})
