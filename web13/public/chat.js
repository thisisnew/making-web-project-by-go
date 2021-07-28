document.addEventListener("DOMContentLoaded", function(){
    if(!window.EventSource) {
        alert("No EventSource");
        return;
    }

    if (!String.prototype.trim) {
        String.prototype.trim = function () {
            return this.replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g, '');
        };
    };

    const chatLog = document.querySelector("#chat-log");
    const chatMsg = document.querySelector("#chat-msg");

    let inputName;
    while (isBlank(inputName)) {
        inputName = prompt("What is your name?")
        if(!isBlank(inputName)){
            const userName = document.querySelector("#user-name");
            userName.innerHTML = '<b>' + inputName + '</b>'
        }
    }

    const btn = document.querySelector("#btn")
    btn.addEventListener("click", () => {
        post('/messages', {
            msg: chatMsg.value,
            name : inputName
        });
        chatMsg.value = '';
        chatMsg.focus();
    });
});

function isBlank(string) {
    return string == null || string.trim() === "";
};

function post(path, params, method='post') {
    const form = document.createElement('form');
    form.method = method;
    form.action = path;
    for (const key in params) {
        if (params.hasOwnProperty(key)) {
            const hiddenField = document.createElement('input');
            hiddenField.type = 'hidden';
            hiddenField.name = key;
            hiddenField.value = params[key];
            form.appendChild(hiddenField);
        }
    }
    document.body.appendChild(form);
    form.submit();
}