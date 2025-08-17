import { axios as request } from './request'; 

/*----------------------------------权限中心---------------------------------*/
export function login(parameter){
    return request({
        url: "/login",
        method: 'post',
        data:parameter
    })
}

export function register(parameter){
    return request({
        url: "/user/insert",
        method: 'post',
        data:parameter
    })
}
//权限中心-role
export function roleInsert(parameter){//新增角色
    return request({
        url: "/role/insert",
        method: 'post',
        data: parameter
    })
}

export function roleRemove (parameter) {//删除角色
    return request({
        url: "/role/delete",
        method: 'post',
        data: parameter
    })
}

export function roleEdit(parameter){//修改角色
    return request({
        url: "/role/update",
        method: 'post',
        data: parameter
    })
}

export function roleList(parameter){//查询角色列表
    const queryString = new URLSearchParams(parameter).toString(); 
    return request({
        url: `/role/getList?${queryString}`, 
        method: 'get', 
    });
}

//权限中心-user
export function userList(parameter){//查询用户列表
    const queryString = new URLSearchParams(parameter).toString(); 
    return request({
        url: `/user/getList?${queryString}`, 
        method: 'get', 
    });
}
export function userRemove (parameter) {//删除用户
    return request({
        url: "/user/delete",
        method: 'post',
        data: parameter
    })
}
export function userInsert(parameter){//插入用户
    return request({
        url: "/user/insert",
        method: 'post',
        data: parameter
    })
}

export function userEdit(parameter){//修改用户
    return request({
        url: "/user/update",
        method: 'post',
        data: parameter
    })
}

export function userByRolesList(roleId) {  //查询用户的角色
    return request({
        url: `/user/getUserByRoles?id=${roleId}`,
        method: 'get'
    })
}

//日志中心
export function logByOperationList(parameter){
    const queryString = new URLSearchParams(parameter).toString();    // 将参数转换为查询字符串
    return request({
        url: `/log/operation/getList?${queryString}`,
        method: 'get'
    })
}

