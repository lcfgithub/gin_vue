
import axios from 'axios'
import storageService from "../service/storage";

const VUE_APP_BASE_API = 'http://localhost:1010/api/'

const httpGet = axios.create({
    timeout: 5000,
    baseURL: VUE_APP_BASE_API,// process.env.VUE_APP_BASE_API, //设置默认的URL前缀
    withCredentials: true, // send cookies when cross-domain requests
    // headers: {'Authorization': 'Bearer ' + localStorage.getItem(storageService.USER_TOKEN)}
    // headers: {'Authorization': 'Bearer ' + storageService._get(storageService.USER_TOKEN)}
  });
// 添加请求拦截器
httpGet.interceptors.request.use(function (config) {
  // 在发送请求之前做些什么
  Object.assign(config.headers, {'Authorization': 'Bearer ' + localStorage.getItem(storageService.USER_TOKEN)})
  return config;
}, function (error) {
  // 对请求错误做些什么
  return Promise.reject(error);
});

// 添加响应拦截器
httpGet.interceptors.response.use(function (response) {
  // 对响应数据做点什么
  return response;
}, function (error) {
  // 对响应错误做点什么
  return Promise.reject(error);
});
  export default httpGet