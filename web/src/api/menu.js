import service from '@/utils/request'

export const getMenusServer = () => {
    return service({
        url: "/v1/menus",
        method: 'get'
    })
}