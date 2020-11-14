import request from "../until/request"

const register = ({name, password, telphone}) => {
    return request.post('auth/register', {name, password, telphone})
}
const login = ({password, telphone}) => {
    return request.post('auth/login', {password, telphone})
}
const info = () => {
    return request.get('auth/info')
}

export default {
    register,
    login,
    info
}