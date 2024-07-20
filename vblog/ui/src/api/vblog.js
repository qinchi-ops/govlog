import client  from "./client";

export var LOGIN = (data) => {
    return client.post('/vblog/api/v1/tokens',data)
}

export var LOGOUT = (refreshToken) => {
    return client.delete('/vblog/api/v1/tokens',{
        headers:{
            'X-REFRESH-TOKEN': refreshToken
        }
    })
}