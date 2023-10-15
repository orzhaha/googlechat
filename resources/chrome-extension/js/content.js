document.addEventListener("DOMContentLoaded", function() {
    // 使用事件代理捕獲DOM元素插入事件
    document.addEventListener('DOMNodeInserted', function(event) {
        var target = event.target;

        if (target instanceof Element && target.matches('c-wiz:not(.wiz-added)')) {
            target.classList.add('wiz-added');

            if (target.querySelector("[data-group-synthetic-id]+[role=button]:not(.msg-url-btn-added)")) {
                target.querySelectorAll("[data-group-synthetic-id]+[role=button]:not(.msg-url-btn-added)").forEach((btn) => {
                    addUrlBtn(btn)
                })

                target.addEventListener('DOMNodeInserted', function(event) {
                    var wizTarget = event.target;

                    if (wizTarget instanceof Element) {
                        var btn = wizTarget.querySelector("[data-group-synthetic-id]+[role=button]:not(.msg-url-btn-added)");

                        if (btn) {
                            addUrlBtn(btn)
                        }
                    }
                })
            }
        }
    });

    // 這裡放置你的代碼，確保 DOM 準備好後運行
    var styleTag = document.createElement("style");
    styleTag.type = "text/css";
    
    var cssRules = `
        .btn-msg-url > .icon-link {
            margin: 13px 0px 0px 9px;
        }

        .icon-link {
            position: absolute;
            display: inline-block;
            width: 8px;
            border-width: 1px;
            border-style: solid;
            border-color: initial;
            border-image: initial;
        }

        .btn-msg-url {
            position: relative;
            width: 28px;
            height: 28px;
            margin-left: 5px;
            color: var(--icon-color,#1f1f1f);
            cursor: pointer;
            border-radius: 50%;
        }

        .btn-msg-url:active {
            transform: scale(0.96);
            box-shadow: 0 0 1px 2px rgba(158, 158, 158, 0.4), 0 0 1px 4px rgba(209, 209, 209, 0.1), 0 0 1px 6px rgba(252, 252, 252, 0.1);
        }

        .btn-thread-url {
            position: sticky;
            top: 5px;
            z-index: 10;
        }


        .btn-thread-url .icon-link {
            margin: 15px 0px 0px 11px;
        }


        .btn-thread-url > div{
            color: var(--icon-color,#1f1f1f);
            cursor: pointer;
            border-radius: 50%;
        }

        .icon-link::before, .icon-link::after {
            position: absolute;
            content: "";
            width: 6px;
            height: 6px;
            border-width: 2px;
            border-style: solid;
            border-color: initial;
            border-image: initial;
        }

        .icon-link::before {
            border-top-left-radius: 5px;
            border-bottom-left-radius: 5px;
            margin: -5px 0px 0px -6px;
            border-right: 0px;
        }

        .icon-link::after {
            border-top-right-radius: 5px;
            border-bottom-right-radius: 5px;
            margin: -5px 0px 0px 6px;
            border-left: 0px;
        }

        .btn-thread-url:hover > div, .btn-msg-url:hover {
            background-color: rgb(238, 238, 238);
        }

        .btn-thread-url > div {
            position: absolute;
            width: 32px;
            height: 32px;
            top: 9px;
            right: 9px;
        }

        .btn-thread-url:hover {
            display: block;
        }
        `;
    

    styleTag.innerHTML = cssRules;
    
    // 確保 document.head 不為 null 再附加
    if (document.head) {
        document.head.appendChild(styleTag);
    }
});


function addUrlBtn(btn) {
    btn.classList.add('msg-url-btn-added');
    
    var btnelement = btn;
    var element = document.createElement('div');
    
    element.classList.add('btn-msg-url');
    element.setAttribute('title', 'Copy the Message URL');
    element.addEventListener('click', function() {
        var input = document.createElement('input');
    
        input.classList.add('input-for-copy');
        input.setAttribute('type', 'text');
        input.value = 'https://chat.google.com/' + document.querySelector('[data-group-id]').getAttribute('data-group-id').replace('space', 'room') + '/' + btnelement.parentElement.closest('[data-topic-id]').getAttribute('data-topic-id') + '/' + btnelement.parentElement.closest('[data-id]').getAttribute('data-id');
        document.body.appendChild(input);
        input.focus();
        input.select();
        document.execCommand('copy');
        input.remove();
    });
    
    var elementIconLink = document.createElement('div');
    elementIconLink.classList.add('icon-link');
    element.appendChild(elementIconLink);
    
    var b = btnelement.parentElement.querySelector('[aria-haspopup="true"]');
    
    if (b) {
        b.parentNode.insertBefore(element, b);
    } else {
        btnelement.parentElement.appendChild(element);
    }
}