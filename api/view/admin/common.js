/* common.js - 后台管理公共工具库 */
'use strict';

/* ========== AUTH ========== */
const Auth = {
    getInfo() {
        try { return JSON.parse(localStorage.getItem('admin_info')); }
        catch(e) { return null; }
    },
    getToken() { return localStorage.getItem('admin_token') || ''; },
    // userId === 1 为超级管理员（与后端 adminRoleService.go 保持一致）
    isSuperAdmin() {
        const i = this.getInfo();
        return !!(i && i.userId === 1);
    },
    check() {
        if (!this.getInfo() || !this.getToken()) {
            window.top.location.href = '/admin/page/login';
            return false;
        }
        return true;
    },
    logout() {
        localStorage.removeItem('admin_token');
        localStorage.removeItem('admin_info');
        window.top.location.href = '/admin/page/login';
    }
};

/* ========== API ========== */
async function api(path, body) {
    try {
        const res = await fetch(path, {
            method: 'POST',
            cache: 'no-store',
            headers: {
                'Content-Type': 'application/json',
                'X-Token': Auth.getToken()
            },
            body: JSON.stringify(body)
        });
        const data = await res.json();
        if (data.code === 300 || data.code === 301) {
            Auth.logout();
            return null;
        }
        return data;
    } catch(e) {
        toast('网络请求失败，请检查服务器连接', 'err');
        return null;
    }
}

/* ========== TOAST ========== */
function toast(msg, type = 'ok') {
    const el = document.createElement('div');
    el.className = `toast ${type}`;
    el.textContent = msg;
    document.body.appendChild(el);
    setTimeout(() => el.remove(), 3000);
}

/* ========== MODAL ========== */
function openModal(id)  { document.getElementById(id).classList.add('on'); }
function closeModal(id) { document.getElementById(id).classList.remove('on'); }

/* ========== PAGINATION ========== */
function renderPager(id, total, page, ps, fn) {
    const el = document.getElementById(id);
    if (!el) return;
    const tp = Math.max(1, Math.ceil(total / ps));
    const h = [`<span class="pager-info">共 ${total} 条</span>`];
    h.push(`<button class="pg-btn" ${page<=1?'disabled':''} onclick="${fn}(${page-1})">上一页</button>`);
    const s = Math.max(1, page-2), e = Math.min(tp, page+2);
    for (let i = s; i <= e; i++) {
        h.push(`<button class="pg-btn ${i===page?'cur':''}" onclick="${fn}(${i})">${i}</button>`);
    }
    h.push(`<button class="pg-btn" ${page>=tp?'disabled':''} onclick="${fn}(${page+1})">下一页</button>`);
    el.innerHTML = h.join('');
}

/* ========== TABLE STATE ROWS ========== */
function trLoading(cols) {
    return `<tr class="st-row"><td colspan="${cols}">加载中...</td></tr>`;
}
function trEmpty(cols, msg) {
    return `<tr class="st-row"><td colspan="${cols}">${msg || '暂无数据'}</td></tr>`;
}

/* ========== IMAGE CELL ========== */
function imgCell(url) {
    if (!url) return '<span class="img-ph">无图</span>';
    return `<img class="img-th" src="${esc(url)}" onerror="this.outerHTML='<span class=\\'img-ph\\'>无图</span>'">`;
}

/* ========== DATE UTILS ========== */
// datetime-local value -> "2024-01-01 12:00:00"
function fmtDT(v) {
    if (!v) return '';
    const s = v.replace('T', ' ');
    return s.split(':').length === 2 ? s + ':00' : s;
}
// "2024-01-01 12:00:00" -> datetime-local input value
function toDTLocal(s) {
    if (!s) return '';
    return s.replace(' ', 'T').slice(0, 16);
}

/* ========== SOURCE LABEL ========== */
function srcLabel(s) {
    return ({1:'注册奖励',2:'活动奖励',3:'管理员增加',4:'管理员扣减',5:'兑换消耗'})[s] || `来源${s}`;
}

/* ========== MISC ========== */
function esc(s) {
    return String(s)
        .replace(/&/g,'&amp;')
        .replace(/</g,'&lt;')
        .replace(/>/g,'&gt;')
        .replace(/"/g,'&quot;');
}

function clearFields(...ids) {
    ids.forEach(id => { const el = document.getElementById(id); if (el) el.value = ''; });
}
