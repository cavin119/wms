import service from '@/utils/request'

//获取全部菜单
export const getMenusServer = (params) => {
    return service({
        url: "/v1/menus",
        method: 'get',
        params
    })
}
//创建菜单
export const createMenusServer = (data) => {
    return service({
        url: "/v1/menus",
        method: 'post',
        data: data
    })
}