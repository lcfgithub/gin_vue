import storageService from "../../service/storage";
import userService from "../../service/user";

const userModule = {
    namespaced: true,
    state: {
        token: storageService._get(storageService.USER_TOKEN),
        userInfo: storageService._get(storageService.USER_INFO),
    },
    mutations: {
        SET_TOKEN(state, token) {
            //更新本地缓存
            storageService._set(storageService.USER_TOKEN, token)
            //更新state
            state.token = token
        },
        SET_USER_INFO(state, userInfo) {
            //更新本地缓存
            storageService._set(storageService.USER_INFO, userInfo)
            //更新state
            state.userInfo = userInfo
        }
    },
    actions: {
        register(context, {name, password, telphone}) {
            return new Promise((resolve, reject) => {
                userService
                .register({name, password, telphone})
                .then((res) => {
                    if (res.data.code != 200) {
                        return res;
                    }
                  context.commit('SET_TOKEN', res.data.data.token);
                  return userService.info();
                })
                .then((resp) => {
                  context.commit('SET_USER_INFO', resp.data.user);
                  resolve(resp)
                })
                .catch((err) => {
                  reject(err)
                });
            })
        },
        login(context, {password, telphone}) {
            return new Promise((resolve, reject) => {
                userService
                .login({password, telphone})
                .then((res) => {
                    if (res.data.code != 200) {
                        return res;
                    }
                  context.commit('SET_TOKEN', res.data.data.token);
                  return userService.info();
                })
                .then((resp) => {
                  context.commit('SET_USER_INFO', resp.data.user);
                  resolve(resp)
                })
                .catch((err) => {
                  reject(err)
                });
            })
        },
        logout({commit}) {
            commit('SET_USER_INFO', "")
            commit('SET_TOKEN', "")
            storageService._set(storageService.USER_INFO, "")
            storageService._set(storageService.USER_TOKEN)
        }
    },
    modules: {
    }
}

export default userModule