// 获取页面信息并发送消息给后台函数
function sendWebDetailToBackground() {
	chrome.runtime.sendMessage(
        {
            url: document.location.href,
            title: document.title,
            body: document.body.innerHTML
        }, 
        function(response) {
		    console.log('收到来自后台的回复：' + response);
        }
    );
}

// call
sendWebDetailToBackground();