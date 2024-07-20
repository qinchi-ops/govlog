import { useStorage } from '@vueuse/core'

const initialState = {
    token: undefined
}

export default useStorage('vblog',initialState,localStorage,{
    mergeDefaults:true
})