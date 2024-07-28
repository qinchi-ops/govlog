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
        var  code = 0
        if (err.response && err.response.data ) {
            msg= err.response.data.message
            // console.log(err.response.data)
            code = err.response.data[0].code
        }
        switch(code){
            case 5000:
                location.assign('/login')
                break
            case 5001:
                location.assign('/login')
                break                
            case 5002:
                location.assign('/login')
                break
            case 5003:
                location.assign('/login')
                break                   
        }
        Message.error(msg)
        return Promise.reject(err)
    }
)

export default client