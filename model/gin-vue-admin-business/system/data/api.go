package data

import (
	"errors"
	"github.com/flipped-aurora/gva/library/fmt"
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin-business/system"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

var Api = new(api)

type api struct{}

// DataInit system_apis 表数据初始化
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *api) DataInit() error {
	entities := []model.Api{
		// base
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/login", Method: "POST", ApiGroup: "base", Description: "用户登录"},
		// user
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/register", Method: "POST", ApiGroup: "user", Description: "用户注册"},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserInfo", Method: "PUT", ApiGroup: "user", Description: "设置用户信息"},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/changePassword", Method: "POST", ApiGroup: "user", Description: "修改密码"},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthority", Method: "POST", ApiGroup: "user", Description: "修改用户角色"},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUser", Method: "DELETE", ApiGroup: "user", Description: "删除用户"},
		{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserList", Method: "POST", ApiGroup: "user", Description: "获取用户列表"},
		// api
		{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/createApi", Method: "POST", ApiGroup: "api", Description: "创建api"},
		{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiList", Method: "POST", ApiGroup: "api", Description: "获取api列表"},
		{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiById", Method: "POST", ApiGroup: "api", Description: "获取api详细信息"},
		{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApi", Method: "POST", ApiGroup: "api", Description: "删除Api"},
		{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApisByIds", Method: "DELETE", ApiGroup: "api", Description: "批量删除api"},
		{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/updateApi", Method: "POST", ApiGroup: "api", Description: "更新Api"},
		{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getAllApis", Method: "POST", ApiGroup: "api", Description: "获取所有api"},
		// authority
		{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/createAuthority", Method: "POST", ApiGroup: "authority", Description: "创建角色"},
		{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/updateAuthority", Method: "PUT", ApiGroup: "authority", Description: "更新角色信息"},
		{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/setDataAuthority", Method: "POST", ApiGroup: "authority", Description: "设置角色资源权限"},
		{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/copyAuthority", Method: "POST", ApiGroup: "authority", Description: "拷贝角色"},
		{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/deleteAuthority", Method: "POST", ApiGroup: "authority", Description: "删除角色"},
		{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/getAuthorityList", Method: "POST", ApiGroup: "authority", Description: "获取角色列表"},
		// 菜单
		{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuList", Method: "POST", ApiGroup: "菜单", Description: "分页获取基础menu列表"},
		{Model: global.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addBaseMenu", Method: "POST", ApiGroup: "菜单", Description: "新增菜单"},
		{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuTree", Method: "POST", ApiGroup: "菜单", Description: "获取用户动态路由"},
		{Model: global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/deleteBaseMenu", Method: "POST", ApiGroup: "菜单", Description: "删除菜单"},
		{Model: global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/updateBaseMenu", Method: "POST", ApiGroup: "菜单", Description: "更新菜单"},
		{Model: global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuById", Method: "POST", ApiGroup: "菜单", Description: "根据id获取菜单"},
		// 权限菜单
		{Model: global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenu", Method: "POST", ApiGroup: "权限菜单", Description: "获取菜单树"},
		{Model: global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addMenuAuthority", Method: "POST", ApiGroup: "权限菜单", Description: "增加menu和角色关联关系"},
		{Model: global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuAuthority", Method: "POST", ApiGroup: "权限菜单", Description: "获取指定角色menu"},
		// casbin
		{Model: global.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/updateCasbin", Method: "POST", ApiGroup: "casbin", Description: "更改角色api权限"},
		{Model: global.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST", ApiGroup: "casbin", Description: "获取权限列表"},
		// jwt
		{Model: global.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/jwt/jsonInBlacklist", Method: "POST", ApiGroup: "jwt", Description: "jwt加入黑名单(退出登录)"},
		// 字典
		{Model: global.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/createSysDictionary", Method: "POST", ApiGroup: "字典", Description: "新增字典"},
		{Model: global.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/findSysDictionary", Method: "GET", ApiGroup: "字典", Description: "根据ID获取字典"},
		{Model: global.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/updateSysDictionary", Method: "PUT", ApiGroup: "字典", Description: "更新字典"},
		{Model: global.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/deleteSysDictionary", Method: "DELETE", ApiGroup: "字典", Description: "删除字典"},
		{Model: global.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/getSysDictionaryList", Method: "GET", ApiGroup: "字典", Description: "获取字典列表"},
		// 字典详情
		{Model: global.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/createSysDictionaryDetail", Method: "POST", ApiGroup: "字典详情", Description: "新增字典内容"},
		{Model: global.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/findSysDictionaryDetail", Method: "GET", ApiGroup: "字典详情", Description: "根据ID获取字典内容"},
		{Model: global.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Method: "PUT", ApiGroup: "字典详情", Description: "更新字典内容"},
		{Model: global.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Method: "DELETE", ApiGroup: "字典详情", Description: "删除字典内容"},
		{Model: global.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Method: "GET", ApiGroup: "字典详情", Description: "获取字典内容列表"},
		// 操作历史
		{Model: global.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/createSysOperationRecord", Method: "POST", ApiGroup: "操作历史", Description: "新增操作记录"},
		{Model: global.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecord", Method: "DELETE", ApiGroup: "操作历史", Description: "删除操作记录"},
		{Model: global.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/findSysOperationRecord", Method: "GET", ApiGroup: "操作历史", Description: "根据ID获取操作记录"},
		{Model: global.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/getSysOperationRecordList", Method: "GET", ApiGroup: "操作历史", Description: "获取操作记录列表"},
		{Model: global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Method: "DELETE", ApiGroup: "操作历史", Description: "批量删除操作历史"},
		// 代码生成器
		{Model: global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getDB", Method: "GET", ApiGroup: "代码生成器", Description: "获取所有数据库"},
		{Model: global.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getTables", Method: "GET", ApiGroup: "代码生成器", Description: "获取数据库表"},
		{Model: global.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getColumn", Method: "GET", ApiGroup: "代码生成器", Description: "获取所选table的所有字段"},
		{Model: global.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/createTemp", Method: "POST", ApiGroup: "代码生成器", Description: "自动化代码"},
		{Model: global.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/preview", Method: "POST", ApiGroup: "代码生成器", Description: "预览自动化代码"},
		{Model: global.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getSysHistory", Method: "POST", ApiGroup: "代码生成器", Description: "查询回滚记录"},
		{Model: global.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/rollback", Method: "POST", ApiGroup: "代码生成器", Description: "回滚自动生成代码"},
		{Model: global.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getMeta", Method: "POST", ApiGroup: "代码生成器", Description: "获取meta信息"},
		{Model: global.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/delSysHistory", Method: "POST", ApiGroup: "代码生成器", Description: "删除回滚记录"},
		// 系统
		{Model: global.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getSystemConfig", Method: "POST", ApiGroup: "系统", Description: "获取配置文件内容"},
		{Model: global.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/setSystemConfig", Method: "POST", ApiGroup: "系统", Description: "设置配置文件内容"},
		{Model: global.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getServerInfo", Method: "POST", ApiGroup: "系统", Description: "获取服务器信息"},
		// 文件上传下载
		{Model: global.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/upload", Method: "POST", ApiGroup: "文件上传下载", Description: "文件上传示例"},
		{Model: global.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/getFileList", Method: "POST", ApiGroup: "文件上传下载", Description: "获取上传文件列表"},
		{Model: global.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/deleteFile", Method: "POST", ApiGroup: "文件上传下载", Description: "删除文件"},
		// 断点续传
		{Model: global.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/breakpointContinue", Method: "POST", ApiGroup: "断点续传", Description: "断点续传"},
		{Model: global.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/findFile", Method: "POST", ApiGroup: "断点续传", Description: "查询当前文件成功的切片"},
		{Model: global.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/breakpointContinueFinish", Method: "POST", ApiGroup: "断点续传", Description: "断点续传完成"},
		{Model: global.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/removeChunk", Method: "POST", ApiGroup: "断点续传", Description: "移除断点续传区块"},
		// 断点续传(插件版)
		{Model: global.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/upload", Method: "POST", ApiGroup: "断点续传(插件版)", Description: "插件版分片上传"},
		{Model: global.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/checkFileMd5", Method: "GET", ApiGroup: "断点续传(插件版)", Description: "文件完整度验证"},
		{Model: global.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/mergeFileMd5", Method: "GET", ApiGroup: "断点续传(插件版)", Description: "上传完成合并文件"},
		// 邮件
		{Model: global.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/email/emailTest", Method: "POST", ApiGroup: "邮件", Description: "发送测试邮件"},
		// 客户
		{Model: global.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Method: "POST", ApiGroup: "客户", Description: "创建客户"},
		{Model: global.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Method: "PUT", ApiGroup: "客户", Description: "更新客户"},
		{Model: global.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Method: "DELETE", ApiGroup: "客户", Description: "删除客户"},
		{Model: global.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Method: "GET", ApiGroup: "客户", Description: "获取单一客户"},
		{Model: global.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customerList", Method: "GET", ApiGroup: "客户", Description: "获取客户列表"},
	}

	var entity model.Api
	if errors.Is(global.Db.Where("id = ?", 74).First(&entity).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		fmt.Printf(fmt.InitDataExist, "postgres", entity.TableName())
		return nil
	}

	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return _errors.Wrap(err, entity.TableName()+"表数据初始化失败!")
	}
	fmt.Printf(fmt.InitDataSuccess, "postgres", entity.TableName())

	return nil
}
