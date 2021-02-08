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

//获取单个菜单
export const getMenuByIdServer = (id) => {
    return service({
        url: "/v1/menus/" + id,
        method: 'get'
    })
}

//删除单个菜单
export const delMenuByIdServer = (id) => {
    return service({
        url: "/v1/menus/" + id,
        method: 'delete'
    })
}
