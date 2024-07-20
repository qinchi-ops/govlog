import axios from "axios";
import { Message } from '@arco-design/web-vue'


var client = axios.create({
    baseURL:'',
    timeout:5000
})


client.interceptors.response.use(
    (value) => {
        return value.data
    },
    (err) => {
        console.log(err)
        var msg= err.message
        if (err.response && err.response.data ) {
            msg= err.response.data.message
        }

        Message.error(msg)
        return Promise.reject(err)
    }
)

export default client