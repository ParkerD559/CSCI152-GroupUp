
function Group(ws) {
    return {
        template: '#tmpl-group',
        mounted: function() {

            ws.addEventListener('message', function (event) {
                    data = event.data;
                    if (!data)
                        return;
                    data = JSON.parse(data);
                    code = data.code;
                    if (!code || code !== "group")
                        return;
                    if (data.groupid === "") {
                        this.showError = true
                    } else {
                        this.showGroup = true
                    }
                }.bind(this));
            
            // Need to send after socket has connected
            if (ws.readyState === 1) {
                ws.send(JSON.stringify({
                    code: "group/join",
                    groupid: this.groupid,
                }));
            } else {
                ws.onopen = function() {
                    ws.send(JSON.stringify({
                        code: "group/join",
                        groupid: this.groupid,
                    }));
                }.bind(this);
            }
        },
        beforeDestroy: function() {
            ws.send(JSON.stringify({
                code: "group/leave",
                groupid: this.groupid,
            }));
        },
        data: function() {
            return {
                drawing: false,
                groupid: this.$route.params.groupid,
                showGroup: false,
                showError: false,
            }
        },
        methods: {
            startDraw: function(e) {
                if (!this.drawing) {
                    this.drawing = true;
                    var x = e.x - e.target.offsetLeft;
                    var y = e.y - e.target.offsetTop;
                    whiteboardCtx.moveTo(x, y);
                    whiteboardCtx.beginPath();
                }
            },
            draw: function(e) {
                if (this.drawing) {
                    console.log(e);
                    var x = e.x - e.target.offsetLeft;
                    var y = e.y - e.target.offsetTop;
                    whiteboardCtx.lineTo(x, y);
                    whiteboardCtx.stroke();
                }
            },
            endDraw: function() {
                if (this.drawing) {
                    this.drawing = false;
                    whiteboardCtx.closePath();
                }
            },
        },
        components: {
            'chat-box': Chatbox(ws),
            'white-board': Whiteboard(ws),
            'user-list': Userlist(ws),
        },
    }
}