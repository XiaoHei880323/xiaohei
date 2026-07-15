/* editor.js — 轻量富文本编辑器，无外部依赖 */
(function (global) {
    'use strict';

    /* ---- 注入样式（只注入一次）---- */
    if (!document.getElementById('re-style')) {
        var s = document.createElement('style');
        s.id = 're-style';
        s.textContent = [
            '.re-wrap{border:1px solid #2d3f66;border-radius:6px;overflow:hidden}',
            '.re-toolbar{display:flex;flex-wrap:wrap;gap:2px;padding:6px 8px;',
            '  background:#1e2a45;border-bottom:1px solid #2d3f66}',
            '.re-btn{display:inline-flex;align-items:center;justify-content:center;',
            '  min-width:28px;height:26px;padding:0 6px;background:transparent;',
            '  border:1px solid transparent;border-radius:4px;color:#a8b8d8;',
            '  font-size:13px;cursor:pointer;line-height:1;white-space:nowrap}',
            '.re-btn:hover{background:#2a3a5c;border-color:#3a5080;color:#e8eaf0}',
            '.re-btn.active{background:#2a3a5c;border-color:#4a70b0;color:#7eb8f7}',
            '.re-sep{width:1px;background:#2d3f66;margin:2px 4px;align-self:stretch}',
            '.re-editor{min-height:140px;max-height:360px;overflow-y:auto;',
            '  padding:10px 12px;background:#1a2035;color:#e8eaf0;',
            '  font-size:14px;line-height:1.7;outline:none}',
            '.re-editor:empty:before{content:attr(data-ph);color:#4a5a7a;pointer-events:none}',
            '.re-editor ul,.re-editor ol{padding-left:1.5em;margin:.3em 0}',
            '.re-editor a{color:#7eb8f7}'
        ].join('');
        document.head.appendChild(s);
    }

    /* ---- 工具栏配置 ---- */
    var TOOLS = [
        { cmd: 'bold',           label: 'B',   title: '粗体',     style: 'font-weight:700' },
        { cmd: 'italic',         label: 'I',   title: '斜体',     style: 'font-style:italic' },
        { cmd: 'underline',      label: 'U',   title: '下划线',   style: 'text-decoration:underline' },
        { cmd: 'strikeThrough',  label: 'S',   title: '删除线',   style: 'text-decoration:line-through' },
        { sep: true },
        { cmd: 'insertOrderedList',   label: '1≡', title: '有序列表' },
        { cmd: 'insertUnorderedList', label: '•≡', title: '无序列表' },
        { sep: true },
        { cmd: 'justifyLeft',    label: '⬅',  title: '左对齐' },
        { cmd: 'justifyCenter',  label: '≡',  title: '居中' },
        { cmd: 'justifyRight',   label: '➡',  title: '右对齐' },
        { sep: true },
        { cmd: 'createLink',     label: '🔗', title: '插入链接', prompt: '请输入链接地址：' },
        { cmd: 'removeFormat',   label: '✕',  title: '清除格式' }
    ];

    /* ---- 构造函数 ---- */
    function RichEditor(containerId, opts) {
        opts = opts || {};
        var container = document.getElementById(containerId);
        if (!container) { console.error('RichEditor: element #' + containerId + ' not found'); return; }

        var self = this;

        /* 工具栏 */
        var toolbar = document.createElement('div');
        toolbar.className = 're-toolbar';

        TOOLS.forEach(function (t) {
            if (t.sep) {
                var sp = document.createElement('span');
                sp.className = 're-sep';
                toolbar.appendChild(sp);
                return;
            }
            var btn = document.createElement('button');
            btn.type = 'button';
            btn.className = 're-btn';
            btn.title = t.title;
            btn.textContent = t.label;
            if (t.style) btn.style.cssText = t.style;
            btn.addEventListener('mousedown', function (e) {
                e.preventDefault();
                if (t.prompt) {
                    var url = prompt(t.prompt, 'https://');
                    if (url) document.execCommand(t.cmd, false, url);
                } else {
                    document.execCommand(t.cmd, false, null);
                }
                edEl.focus();
                updateActive();
            });
            toolbar.appendChild(btn);
        });

        /* 编辑区 */
        var edEl = document.createElement('div');
        edEl.className = 're-editor';
        edEl.contentEditable = 'true';
        edEl.setAttribute('data-ph', opts.placeholder || '请输入内容（可选）');

        /* 更新工具栏激活状态 */
        function updateActive() {
            toolbar.querySelectorAll('.re-btn').forEach(function (btn, i) {
                var t = TOOLS.filter(function(x){ return !x.sep; })[i];
                if (!t || t.cmd === 'createLink' || t.cmd === 'removeFormat') return;
                btn.classList.toggle('active', document.queryCommandState(t.cmd));
            });
        }
        edEl.addEventListener('keyup', updateActive);
        edEl.addEventListener('mouseup', updateActive);
        edEl.addEventListener('focus', updateActive);

        /* 组装 */
        container.className = 're-wrap';
        container.innerHTML = '';
        container.appendChild(toolbar);
        container.appendChild(edEl);

        /* 公共 API */
        self.getHTML = function () { return edEl.innerHTML; };
        self.setHTML = function (html) { edEl.innerHTML = html || ''; };
        self.focus   = function () { edEl.focus(); };
    }

    global.RichEditor = RichEditor;
}(window));
