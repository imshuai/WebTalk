function get_timestamp(dateObj) {
    var year = '' + dateObj.getFullYear();
    var month = '' + (dateObj.getMonth() + 1);
    var day = '' + dateObj.getDate();
    var hour = '' + dateObj.getHours();
    var minute = '' + dateObj.getMinutes();
    var second = '' + dateObj.getSeconds();
    return year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + second;
}

function reciveMsg(msg) {
    var $content = $('.chat-content');
    var $stamp = $('<p>');
    $stamp.addClass('rev-timestamp');
    $stamp.text(msg.timestamp);
    var $icon = $('<div>');
    $icon.addClass('chat-rev-icon');
    var $send = $('<p>');
    $send.addClass('chat-msg-rev pull-left');
    $send.text(msg.content);
    var $box = $('<div>');
    $box.addClass('chat-msg-rev-box');
    $box.append($stamp);
    $box.append($icon);
    $box.append($send);
    $content.append($box);
    $content.animate({ scrollTop: this.innerHeight }, 800);
}

function sendMsg(msg) {
    c.send(JSON.stringify(msg));
    var $content = $('.chat-content');
    var $stamp = $('<p>');
    $stamp.addClass('send-timestamp');
    $stamp.text(msg.timestamp);
    var $icon = $('<div>');
    $icon.addClass('chat-send-icon');
    var $send = $('<p>');
    $send.addClass('chat-msg-send pull-right');
    $send.text(msg.content);
    var $box = $('<div>');
    $box.addClass('chat-msg-send-box');
    $box.append($stamp);
    $box.append($icon);
    $box.append($send);
    $content.append($box);
    $content.animate({ scrollTop: this.innerHeight }, 800);
}
$(function() {
    $('#talk-list').height(window.innerHeight);
    $('.chat-content').height(window.innerHeight - $('.chat-editor').height() - $('.chat-title').height() - 12);
    $('#connect').click(function() {
        url = 'ws://127.0.0.1:8080/user/login/' + $('input[name="nickname"]').val();
        c = new WebSocket(url);

        c.onmessage = function(msg) {
            if (msg.data == "control:ping") {
                c.send("control:pong")
            } else if (msg.data == "control:pong") {} else {
                reciveMsg(JSON.parse(msg.data));
            }
        }

        c.onopen = function() {
            setInterval(
                function() {
                    c.send("control:ping")
                }, 5000)
        }

    });
    $('.chat-editor>textarea').keydown(function(e) {
        if (e.ctrlKey && e.keyCode == 13) {
            var msg = new Object;
            msg.author = $('input[name="nickname"]').val();
            msg.sendto = $('input[name="sendto"]').val();
            msg.timestamp = get_timestamp(new Date());
            msg.content = this.value;
            this.value = '';
            sendMsg(msg);
        }
    });
    $('#chat-close').on('click.close.talk', function(event) {
        $('#talk-list').addClass('list-hidden');
        $('#talk-button').css('visibility', 'visible');
    });
    $('#talk-button').on('click.button.talk', function(event) {
        $('#talk-list').removeClass('list-hidden');
        $('#talk-button').css('visibility', 'hidden');
    });
});