import axios from "axios";

axios.interceptors.request.use((config) => {
    return config
}, (err) => {
    return Promise.reject(err)
})

axios.interceptors.response.use((response) => {
    if (response.status === 200) {
        return response.data
    }
}, (err) => {
    return Promise.reject(err)
})
export default axios;