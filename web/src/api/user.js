import service from '@/utils/request'

export const loginServer = (data) => {
    return service({
        url: "/v1/login",
        method: 'post',
        data: data
    })
}