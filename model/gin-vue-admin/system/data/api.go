package model

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin/system"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

func (a *api) TableName() string {
	var entity model.Api
	return entity.TableName()
}

func (a *api) Initialize() error {
	entities := []model.Api{
		// base
		{Model: global.Model{ID: 1}, Path: "/base/login", Method: "POST", ApiGroup: "base", Description: "用户登录"},
		// user
		{Model: global.Model{ID: 2}, Path: "/user/register", Method: "POST", ApiGroup: "user", Description: "用户注册"},
		{Model: global.Model{ID: 3}, Path: "/user/setUserInfo", Method: "PUT", ApiGroup: "user", Description: "设置用户信息"},
		{Model: global.Model{ID: 4}, Path: "/user/changePassword", Method: "POST", ApiGroup: "user", Description: "修改密码"},
		{Model: global.Model{ID: 5}, Path: "/user/setUserAuthority", Method: "POST", ApiGroup: "user", Description: "修改用户角色"},
		{Model: global.Model{ID: 6}, Path: "/user/deleteUser", Method: "DELETE", ApiGroup: "user", Description: "删除用户"},
		{Model: global.Model{ID: 7}, Path: "/user/getUserList", Method: "POST", ApiGroup: "user", Description: "获取用户列表"},
		// api
		{Model: global.Model{ID: 8}, Path: "/api/createApi", Method: "POST", ApiGroup: "api", Description: "创建api"},
		{Model: global.Model{ID: 9}, Path: "/api/getApiList", Method: "POST", ApiGroup: "api", Description: "获取api列表"},
		{Model: global.Model{ID: 10}, Path: "/api/getApiById", Method: "POST", ApiGroup: "api", Description: "获取api详细信息"},
		{Model: global.Model{ID: 11}, Path: "/api/deleteApi", Method: "POST", ApiGroup: "api", Description: "删除Api"},
		{Model: global.Model{ID: 12}, Path: "/api/deleteApisByIds", Method: "DELETE", ApiGroup: "api", Description: "批量删除api"},
		{Model: global.Model{ID: 13}, Path: "/api/updateApi", Method: "POST", ApiGroup: "api", Description: "更新Api"},
		{Model: global.Model{ID: 14}, Path: "/api/getAllApis", Method: "POST", ApiGroup: "api", Description: "获取所有api"},
		// authority
		{Model: global.Model{ID: 15}, Path: "/authority/createAuthority", Method: "POST", ApiGroup: "authority", Description: "创建角色"},
		{Model: global.Model{ID: 16}, Path: "/authority/updateAuthority", Method: "PUT", ApiGroup: "authority", Description: "更新角色信息"},
		{Model: global.Model{ID: 17}, Path: "/authority/setDataAuthority", Method: "POST", ApiGroup: "authority", Description: "设置角色资源权限"},
		{Model: global.Model{ID: 18}, Path: "/authority/copyAuthority", Method: "POST", ApiGroup: "authority", Description: "拷贝角色"},
		{Model: global.Model{ID: 19}, Path: "/authority/deleteAuthority", Method: "POST", ApiGroup: "authority", Description: "删除角色"},
		{Model: global.Model{ID: 20}, Path: "/authority/getAuthorityList", Method: "POST", ApiGroup: "authority", Description: "获取角色列表"},
		// 菜单
		{Model: global.Model{ID: 21}, Path: "/menu/getMenuList", Method: "POST", ApiGroup: "菜单", Description: "分页获取基础menu列表"},
		{Model: global.Model{ID: 22}, Path: "/menu/addBaseMenu", Method: "POST", ApiGroup: "菜单", Description: "新增菜单"},
		{Model: global.Model{ID: 23}, Path: "/menu/getBaseMenuTree", Method: "POST", ApiGroup: "菜单", Description: "获取用户动态路由"},
		{Model: global.Model{ID: 24}, Path: "/menu/deleteBaseMenu", Method: "POST", ApiGroup: "菜单", Description: "删除菜单"},
		{Model: global.Model{ID: 25}, Path: "/menu/updateBaseMenu", Method: "POST", ApiGroup: "菜单", Description: "更新菜单"},
		{Model: global.Model{ID: 26}, Path: "/menu/getBaseMenuById", Method: "POST", ApiGroup: "菜单", Description: "根据id获取菜单"},
		// 权限菜单
		{Model: global.Model{ID: 27}, Path: "/menu/getMenu", Method: "POST", ApiGroup: "权限菜单", Description: "获取菜单树"},
		{Model: global.Model{ID: 28}, Path: "/menu/addMenuAuthority", Method: "POST", ApiGroup: "权限菜单", Description: "增加menu和角色关联关系"},
		{Model: global.Model{ID: 29}, Path: "/menu/getMenuAuthority", Method: "POST", ApiGroup: "权限菜单", Description: "获取指定角色menu"},
		// casbin
		{Model: global.Model{ID: 30}, Path: "/casbin/updateCasbin", Method: "POST", ApiGroup: "casbin", Description: "更改角色api权限"},
		{Model: global.Model{ID: 31}, Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST", ApiGroup: "casbin", Description: "获取权限列表"},
		// jwt
		{Model: global.Model{ID: 32}, Path: "/jwt/jsonInBlacklist", Method: "POST", ApiGroup: "jwt", Description: "jwt加入黑名单(退出登录)"},
		// 字典
		{Model: global.Model{ID: 33}, Path: "/sysDictionary/createSysDictionary", Method: "POST", ApiGroup: "字典", Description: "新增字典"},
		{Model: global.Model{ID: 34}, Path: "/sysDictionary/findSysDictionary", Method: "GET", ApiGroup: "字典", Description: "根据ID获取字典"},
		{Model: global.Model{ID: 35}, Path: "/sysDictionary/updateSysDictionary", Method: "PUT", ApiGroup: "字典", Description: "更新字典"},
		{Model: global.Model{ID: 36}, Path: "/sysDictionary/deleteSysDictionary", Method: "DELETE", ApiGroup: "字典", Description: "删除字典"},
		{Model: global.Model{ID: 37}, Path: "/sysDictionary/getSysDictionaryList", Method: "GET", ApiGroup: "字典", Description: "获取字典列表"},
		// 字典详情
		{Model: global.Model{ID: 38}, Path: "/sysDictionaryDetail/createSysDictionaryDetail", Method: "POST", ApiGroup: "字典详情", Description: "新增字典内容"},
		{Model: global.Model{ID: 39}, Path: "/sysDictionaryDetail/findSysDictionaryDetail", Method: "GET", ApiGroup: "字典详情", Description: "根据ID获取字典内容"},
		{Model: global.Model{ID: 40}, Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Method: "PUT", ApiGroup: "字典详情", Description: "更新字典内容"},
		{Model: global.Model{ID: 41}, Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Method: "DELETE", ApiGroup: "字典详情", Description: "删除字典内容"},
		{Model: global.Model{ID: 42}, Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Method: "GET", ApiGroup: "字典详情", Description: "获取字典内容列表"},
		// 操作历史
		{Model: global.Model{ID: 43}, Path: "/sysOperationRecord/createSysOperationRecord", Method: "POST", ApiGroup: "操作历史", Description: "新增操作记录"},
		{Model: global.Model{ID: 44}, Path: "/sysOperationRecord/deleteSysOperationRecord", Method: "DELETE", ApiGroup: "操作历史", Description: "删除操作记录"},
		{Model: global.Model{ID: 45}, Path: "/sysOperationRecord/findSysOperationRecord", Method: "GET", ApiGroup: "操作历史", Description: "根据ID获取操作记录"},
		{Model: global.Model{ID: 46}, Path: "/sysOperationRecord/getSysOperationRecordList", Method: "GET", ApiGroup: "操作历史", Description: "获取操作记录列表"},
		{Model: global.Model{ID: 47}, Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Method: "DELETE", ApiGroup: "操作历史", Description: "批量删除操作历史"},
		// 代码生成器
		{Model: global.Model{ID: 48}, Path: "/autoCode/getDB", Method: "GET", ApiGroup: "代码生成器", Description: "获取所有数据库"},
		{Model: global.Model{ID: 49}, Path: "/autoCode/getTables", Method: "GET", ApiGroup: "代码生成器", Description: "获取数据库表"},
		{Model: global.Model{ID: 50}, Path: "/autoCode/getColumn", Method: "GET", ApiGroup: "代码生成器", Description: "获取所选table的所有字段"},
		{Model: global.Model{ID: 51}, Path: "/autoCode/createTemp", Method: "POST", ApiGroup: "代码生成器", Description: "自动化代码"},
		{Model: global.Model{ID: 52}, Path: "/autoCode/preview", Method: "POST", ApiGroup: "代码生成器", Description: "预览自动化代码"},
		{Model: global.Model{ID: 53}, Path: "/autoCode/getSysHistory", Method: "POST", ApiGroup: "代码生成器", Description: "查询回滚记录"},
		{Model: global.Model{ID: 54}, Path: "/autoCode/rollback", Method: "POST", ApiGroup: "代码生成器", Description: "回滚自动生成代码"},
		{Model: global.Model{ID: 55}, Path: "/autoCode/getMeta", Method: "POST", ApiGroup: "代码生成器", Description: "获取meta信息"},
		{Model: global.Model{ID: 56}, Path: "/autoCode/delSysHistory", Method: "POST", ApiGroup: "代码生成器", Description: "删除回滚记录"},
		// 系统
		{Model: global.Model{ID: 57}, Path: "/system/getSystemConfig", Method: "POST", ApiGroup: "系统", Description: "获取配置文件内容"},
		{Model: global.Model{ID: 58}, Path: "/system/setSystemConfig", Method: "POST", ApiGroup: "系统", Description: "设置配置文件内容"},
		{Model: global.Model{ID: 59}, Path: "/system/getServerInfo", Method: "POST", ApiGroup: "系统", Description: "获取服务器信息"},
		// 文件上传下载
		{Model: global.Model{ID: 60}, Path: "/fileUploadAndDownload/upload", Method: "POST", ApiGroup: "文件上传下载", Description: "文件上传示例"},
		{Model: global.Model{ID: 61}, Path: "/fileUploadAndDownload/getFileList", Method: "POST", ApiGroup: "文件上传下载", Description: "获取上传文件列表"},
		{Model: global.Model{ID: 62}, Path: "/fileUploadAndDownload/deleteFile", Method: "POST", ApiGroup: "文件上传下载", Description: "删除文件"},
		// 断点续传
		{Model: global.Model{ID: 63}, Path: "/fileUploadAndDownload/breakpointContinue", Method: "POST", ApiGroup: "断点续传", Description: "断点续传"},
		{Model: global.Model{ID: 64}, Path: "/fileUploadAndDownload/findFile", Method: "POST", ApiGroup: "断点续传", Description: "查询当前文件成功的切片"},
		{Model: global.Model{ID: 65}, Path: "/fileUploadAndDownload/breakpointContinueFinish", Method: "POST", ApiGroup: "断点续传", Description: "断点续传完成"},
		{Model: global.Model{ID: 66}, Path: "/fileUploadAndDownload/removeChunk", Method: "POST", ApiGroup: "断点续传", Description: "移除断点续传区块"},
		// 断点续传(插件版)
		{Model: global.Model{ID: 67}, Path: "/simpleUploader/upload", Method: "POST", ApiGroup: "断点续传(插件版)", Description: "插件版分片上传"},
		{Model: global.Model{ID: 68}, Path: "/simpleUploader/checkFileMd5", Method: "GET", ApiGroup: "断点续传(插件版)", Description: "文件完整度验证"},
		{Model: global.Model{ID: 69}, Path: "/simpleUploader/mergeFileMd5", Method: "GET", ApiGroup: "断点续传(插件版)", Description: "上传完成合并文件"},
		// 邮件
		{Model: global.Model{ID: 70}, Path: "/email/emailTest", Method: "POST", ApiGroup: "邮件", Description: "发送测试邮件"},
		// 客户
		{Model: global.Model{ID: 71}, Path: "/customer/customer", Method: "POST", ApiGroup: "客户", Description: "创建客户"},
		{Model: global.Model{ID: 72}, Path: "/customer/customer", Method: "PUT", ApiGroup: "客户", Description: "更新客户"},
		{Model: global.Model{ID: 73}, Path: "/customer/customer", Method: "DELETE", ApiGroup: "客户", Description: "删除客户"},
		{Model: global.Model{ID: 74}, Path: "/customer/customer", Method: "GET", ApiGroup: "客户", Description: "获取单一客户"},
		{Model: global.Model{ID: 75}, Path: "/customer/customerList", Method: "GET", ApiGroup: "客户", Description: "获取客户列表"},
	}

	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return _errors.Wrap(err, "数据初始化失败!")
	}

	return nil
}

func (a *api) CheckDataExist() bool {
	if _errors.Is(global.Db.Where("id = ?", 74).First(&model.Api{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
