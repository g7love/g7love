layui.use('layim', function(layim){
    var autoReplay = [
        '您好，我现在有事不在，一会再和您联系。',
        '你没发错吧？face[微笑] ',
        '洗澡中，请勿打扰，偷窥请购票，个体四十，团体八折，订票电话：一般人我不告诉他！face[哈哈] ',
        '你好，我是主人的美女秘书，有什么事就跟我说吧，等他回来我会转告他的。face[心] face[心] face[心] ',
        'face[威武] face[威武] face[威武] face[威武] ',
        '<（@￣︶￣@）>',
        '你要和我说话？你真的要和我说话？你确定自己想说吗？你一定非说不可吗？那你说吧，这是自动回复。',
        'face[黑线]  你慢慢说，别急……',
        '(*^__^*) face[嘻嘻] ，是贤心吗？'
    ];

    var layim = layui.layim;
    layim.config({
        //获取主面板列表信息
        init: {
            url: 'http://socialnetworking.com/mapchat/index.json' //接口地址（返回的数据格式见下文）
            ,type: 'get' //默认get，一般可不填
            ,data: {} //额外参数
        }
        //获取群员接口
        ,members: {
            url: '' //接口地址（返回的数据格式见下文）
            ,type: 'get' //默认get，一般可不填
            ,data: {} //额外参数
        }
        //上传图片接口（返回的数据格式见下文）
        ,uploadImage: {
            url: '' //接口地址（返回的数据格式见下文）
            ,type: 'post' //默认post
        }
        //上传文件接口（返回的数据格式见下文）
        ,uploadFile: {
            url: '' //接口地址（返回的数据格式见下文）
            ,type: 'post' //默认post
        }
        //增加皮肤选择，如果不想增加，可以剔除该项
        ,skin: [
            'http://xxx.com/skin.jpg',

        ]
        ,brief: false //是否简约模式（默认false，如果只用到在线客服，且不想显示主面板，可以设置 true）
        ,title: '我的LayIM' //主面板最小化后显示的名称
        ,min: false //用于设定主面板是否在页面打开时，始终最小化展现。默认false，即记录上次展开状态。
        ,minRight: null //【默认不开启】用户控制聊天面板最小化时、及新消息提示层的相对right的px坐标，如：minRight: '200px'
        ,maxLength: 3000 //最长发送的字符长度，默认3000
        ,isfriend: true //是否开启好友（默认true，即开启）
        ,isgroup: true //是否开启群组（默认true，即开启）
        ,right: '0px' //默认0px，用于设定主面板右偏移量。该参数可避免遮盖你页面右下角已经的bar。
        ,chatLog: 'http://socialnetworking.com/mapchat/index.html' //聊天记录地址（如果未填则不显示）
        ,find: '/find/' //查找好友/群的地址（如果未填则不显示）
        ,copyright: true //是否授权，如果通过官网捐赠获得LayIM，此处可填true
    });

    var socket = new WebSocket('ws://127.0.0.1:9502');



    //监听发送消息
    layim.on('sendMessage', function(data){
        var To = data.to;

        /*//连接成功时触发
        socket.onopen = function(){
            socket.send('XXX连接成功');
            socket.send('Hi Server, I am LayIM!');
        };*/

        socket.send(JSON.stringify({
            username: "纸飞机" //消息来源用户名
            ,avatar: "http://tp1.sinaimg.cn/1571889140/180/40030060651/1" //消息来源用户头像
            ,id: "100000" //聊天窗口来源ID（如果是私聊，则是用户id，如果是群聊，则是群组id）
            ,type: "friend" //聊天窗口来源类型，从发送消息传递的to里面获取
            ,content: "嗨，你好！本消息系离线消息。" //消息内容
            ,mine: false //是否我发送的消息，如果为true，则会显示在右方
            ,timestamp: 1467475443306 //服务端动态时间戳
        }));

        //连接成功时触发
        /*socket.onopen = function(){

        };*/

        //监听收到的消息
        /*socket.onmessage = function(res){
            //res为接受到的值，如 {"emit": "messageName", "data": {}}
            //emit即为发出的事件名，用于区分不同的消息
            console.log(res)
        };*/

        //监听收到的聊天消息，假设你服务端emit的事件名为：chatMessage
        socket.onmessage = function(res){
            if(res.emit === 'chatMessage'){
                layim.getMessage(res.data); //res.data即你发送消息传递的数据（阅读：监听发送的消息）
            }
        };

        console.log(data);
        //演示自动回复
/*
        setTimeout(function(){
            var obj = {};
            if(To.type === 'group'){
                obj = {
                    username: '模拟群员'+(Math.random()*100|0)
                    ,avatar: layui.cache.dir + 'images/face/'+ (Math.random()*72|0) + '.gif'
                    ,id: To.id
                    ,type: To.type
                    ,content: autoReplay[Math.random()*9|0]
                }
            } else {
                obj = {
                    username: To.name
                    ,avatar: To.avatar
                    ,id: To.id
                    ,type: To.type
                    ,content: autoReplay[Math.random()*9|0]
                    ,timestamp: new Date().getTime()
                }
            }
            layim.getMessage(obj);
        }, 1000);
*/
    });
    //监听在线状态的切换事件
    layim.on('online', function(data){
        console.log(data);
    });
    //监听查看群员
    layim.on('members', function(data){
        console.log(data);
    });

    //监听聊天窗口的切换
    layim.on('chatChange', function(data){
        console.log(data);
    });
    //面板外的操作
    var $ = layui.jquery, active = {
        chat: function(){
            layim.chat({
                name: '小闲'
                ,type: 'friend'
                ,avatar: 'http://tva3.sinaimg.cn/crop.0.0.180.180.180/7f5f6861jw1e8qgp5bmzyj2050050aa8.jpg'
                ,id: 1008612
            });
            layer.msg('也就是说，此人可以不在好友面板里');
        }
        ,message: function(){
            layim.getMessage({
                username: "贤心"
                ,avatar: "http://tp1.sinaimg.cn/1571889140/180/40030060651/1"
                ,id: "100001"
                ,type: "friend"
                ,content: "嗨，你好！欢迎体验LayIM。演示标记："+ new Date().getTime()
                ,timestamp: new Date().getTime()
            });
        }
        ,messageTemp: function(){
            layim.getMessage({
                username: "小酱"
                ,avatar: "http://tva1.sinaimg.cn/crop.7.0.736.736.50/bd986d61jw8f5x8bqtp00j20ku0kgabx.jpg"
                ,id: "198909151014"
                ,type: "friend"
                ,content: "临时："+ new Date().getTime()
            });
        }
        ,addFriend: function(){
            layer.msg('已成功把[冲田杏梨]添加到好友【网红】组里', {
                icon: 1
            });
            layim.addList({
                type: 'friend'
                ,avatar: "http://tp2.sinaimg.cn/2386568184/180/40050524279/0"
                ,username: '冲田杏梨'
                ,groupid: 2
                ,id: "1233333312121212"
                ,remark: "本人冲田杏梨将结束AV女优的工作"
            });
        }
        ,addGroup: function(){
            layer.msg('已成功把[Angular开发]添加到群组里', {
                icon: 1
            });
            layim.addList({
                type: 'group'
                ,avatar: "http://tva3.sinaimg.cn/crop.64.106.361.361.50/7181dbb3jw8evfbtem8edj20ci0dpq3a.jpg"
                ,groupname: 'Angular开发'
                ,id: "12333333"
                ,members: 0
            });
        }
        ,removeFriend: function(){
            layer.msg('已成功删除[凤姐]', {
                icon: 1
            });
            layim.removeList({
                id: 121286
                ,type: 'friend'
            });
        }
        ,removeGroup: function(){
            layer.msg('已成功删除[前端群]', {
                icon: 1
            });
            layim.removeList({
                id: 101
                ,type: 'group'
            });
        }
    };
    $('.site-demo-layim').on('click', function(){
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
    });

});

var map = new AMap.Map('container', {
    resizeEnable: true,
    center: [110.408075, 34.950187],
    zoom: 5
});
function refresh(e) {
    map.setMapStyle(e);
}

var marker = new AMap.Marker({
    position: map.getCenter(),
    draggable: true,
    cursor: 'move'
});
marker.setMap(map);
// 设置点标记的动画效果，此处为弹跳效果
marker.setAnimation('AMAP_ANIMATION_BOUNCE');