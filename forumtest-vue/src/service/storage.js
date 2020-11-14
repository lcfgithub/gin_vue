import Cookies from 'js-cookie'

//本地缓存服务
const PREFIX = 'forumtest_'

//用户模块
const USER_PREFIX = `${PREFIX}user_`
const USER_INFO = `${USER_PREFIX}info`
const USER_TOKEN = `${USER_PREFIX}token`

//缓存读取
const _get = (name) => {
    // Cookies.get(name)
    localStorage.getItem(name)
}

//缓存设置
const _set = (name, value) => {
    // Cookies.set(name, value)
    localStorage.setItem(name, value)
}

const _cache = (method, name ,value = null) => {
    (method === 'get') && Vue.$cookie.get(name)
    (method === 'set' && value) && Vue.$cookie.set(name, value)
    (method === 'set' && !value) && Vue.$cookie.delete(name)
}

export default {
    USER_INFO,
    USER_TOKEN,
    _cache,
    _set,
    _get,
}
