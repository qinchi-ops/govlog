import client  from "./client";

export var LOGIN = (data) => {
    return client.post('/vblog/api/v1/tokens',data)
}

export var LOGOUT = () => {
    return client.delete('/vblog/api/v1/tokens',{})
}