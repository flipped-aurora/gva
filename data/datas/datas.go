package datas

import (
	"time"

	"github.com/flipped-aurora/gva/data/model"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitExaCustomer(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.ExaCustomer{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CustomerName: "测试客户", CustomerPhoneData: "1761111111", SysUserID: 1, SysUserAuthorityID: "888"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitCasbinModel(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable("casbin_rule") {
		if err := db.Migrator().CreateTable(&gormadapter.CasbinRule{}); err != nil {
			return err
		}
	}
	tx := db.Begin() // 开始事务
	insert := []gormadapter.CasbinRule{
		{PType: "p", V0: "888", V1: "/base/login", V2: "POST"},
		{PType: "p", V0: "888", V1: "/base/register", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{PType: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{PType: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
		{PType: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{PType: "p", V0: "888", V1: "/casbin/casbinTest/:pathParam", V2: "GET"},
		{PType: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{PType: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
		{PType: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
		{PType: "p", V0: "888", V1: "/customer/customer", V2: "POST"},
		{PType: "p", V0: "888", V1: "/customer/customer", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/customer/customer", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/customer/customer", V2: "GET"},
		{PType: "p", V0: "888", V1: "/customer/customerList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
		{PType: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
		{PType: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
		{PType: "p", V0: "888", V1: "/autoCode/getColume", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
		{PType: "p", V0: "888", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
		{PType: "p", V0: "888", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/email/emailTest", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/base/login", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/base/register", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/api/createApi", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/api/getApiList", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/api/getApiById", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/api/deleteApi", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/api/updateApi", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/api/getAllApis", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/authority/createAuthority", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/authority/deleteAuthority", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/authority/getAuthorityList", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/authority/setDataAuthority", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/getMenu", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/getMenuList", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/addBaseMenu", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/addMenuAuthority", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/getMenuAuthority", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/updateBaseMenu", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/menu/getBaseMenuById", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/user/changePassword", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/user/getUserList", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/user/setUserAuthority", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/casbin/updateCasbin", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/system/getSystemConfig", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/system/setSystemConfig", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/customer/customer", V2: "POST"},
		{PType: "p", V0: "8881", V1: "/customer/customer", V2: "PUT"},
		{PType: "p", V0: "8881", V1: "/customer/customer", V2: "DELETE"},
		{PType: "p", V0: "8881", V1: "/customer/customer", V2: "GET"},
		{PType: "p", V0: "8881", V1: "/customer/customerList", V2: "GET"},
		{PType: "p", V0: "9528", V1: "/base/login", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/base/register", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/customer/customer", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/customer/customer", V2: "PUT"},
		{PType: "p", V0: "9528", V1: "/customer/customer", V2: "DELETE"},
		{PType: "p", V0: "9528", V1: "/customer/customer", V2: "GET"},
		{PType: "p", V0: "9528", V1: "/customer/customerList", V2: "GET"},
		{PType: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
	}
	if tx.Table("casbin_rule").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysAuthority(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.SysAuthority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysBaseMenus(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.SysBaseMenu{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{Title: "关于我们", Icon: "info"}},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: model.Meta{Title: "示例文件", Icon: "s-management"}},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "table", Name: "table", Component: "view/example/table/table.vue", Sort: 1, Meta: model.Meta{Title: "表格示例", Icon: "s-order"}},
		{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "form", Name: "form", Component: "view/example/form/form.vue", Sort: 2, Meta: model.Meta{Title: "表单示例", Icon: "document"}},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "rte", Name: "rte", Component: "view/example/rte/rte.vue", Sort: 3, Meta: model.Meta{Title: "富文本编辑器", Icon: "reading"}},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: model.Meta{Title: "excel导入导出", Icon: "s-marketing"}},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: model.Meta{Title: "上传下载", Icon: "upload"}},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: model.Meta{Title: "断点续传", Icon: "upload"}},
		{Model: gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: model.Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "17", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "iconList", Name: "iconList", Component: "view/iconList/index.vue", Sort: 2, Meta: model.Meta{Title: "图标集合", Icon: "star-on"}},
		{Model: gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "notebook-2"}},
		{Model: gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "s-order"}},
		{Model: gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
		{Model: gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: model.Meta{Title: "断点续传（插件版）", Icon: "upload"}},
		{Model: gorm.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: false, Component: "/", Sort: 0, Meta: model.Meta{Title: "官方网站", Icon: "s-home"}},
		{Model: gorm.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitAuthorityMenu(db *gorm.DB) (err error) {
	return db.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`, `sys_base_menus`.`updated_at` AS `updated_at`, `sys_base_menus`.`deleted_at` AS `deleted_at`, `sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`, `sys_base_menus`.`title`  AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on ((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)))").Error
}

func InitSysDictionary(db *gorm.DB) (err error) {
	status := new(bool)
	*status = true
	tx := db.Begin() // 开始事务
	insert := []model.SysDictionary{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysAuthorityMenus(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.SysAuthorityMenus{
		{"888", 1},
		{"888", 2},
		{"888", 3},
		{"888", 4},
		{"888", 5},
		{"888", 6},
		{"888", 7},
		{"888", 8},
		{"888", 9},
		{"888", 10},
		{"888", 11},
		{"888", 12},
		{"888", 13},
		{"888", 14},
		{"888", 15},
		{"888", 16},
		{"888", 17},
		{"888", 18},
		{"888", 19},
		{"888", 20},
		{"888", 21},
		{"888", 22},
		{"888", 23},
		{"888", 24},
		{"888", 25},
		{"888", 26},
		{"888", 27},
		{"8881", 1},
		{"8881", 2},
		{"8881", 8},
		{"8881", 17},
		{"8881", 18},
		{"8881", 19},
		{"8881", 20},
		{"9528", 1},
		{"9528", 2},
		{"9528", 3},
		{"9528", 4},
		{"9528", 5},
		{"9528", 6},
		{"9528", 7},
		{"9528", 8},
		{"9528", 9},
		{"9528", 10},
		{"9528", 11},
		{"9528", 12},
		{"9528", 13},
		{"9528", 14},
		{"9528", 15},
		{"9528", 17},
		{"9528", 18},
		{"9528", 19},
		{"9528", 20},
	}
	if tx.Table("sys_authority_menus").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDataAuthorityId(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.SysDataAuthorityId{
		{"888", "888"},
		{"888", "8881"},
		{"888", "9528"},
		{"9528", "8881"},
		{"9528", "9528"},
	}
	if tx.Migrator().HasTable("sys_data_authority_ids") {
		if tx.Table("sys_data_authority_ids").Create(&insert).Error != nil { // 遇到错误时回滚事务
			tx.Rollback()
		}
		return tx.Commit().Error
	}
	if tx.Table("sys_data_authority_id").Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDictionaryDetail(db *gorm.DB) (err error) {
	status := new(bool)
	*status = true
	tx := db.Begin() // 开始事务
	insert := []model.SysDictionaryDetail{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "smallint", 1, status, 1, 2},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumint", 2, status, 2, 2},
		{gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "int", 3, status, 3, 2},
		{gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "bigint", 4, status, 4, 2},
		{gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "data", 0, status, 0, 3},
		{gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "time", 1, status, 1, 3},
		{gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "year", 2, status, 2, 3},
		{gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "datetime", 3, status, 3, 3},
		{gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "timestamp", 5, status, 5, 3},
		{gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "float", 0, status, 0, 4},
		{gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "double", 1, status, 1, 4},
		{gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "decimal", 2, status, 2, 4},
		{gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "char", 0, status, 0, 5},
		{gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "varchar", 1, status, 1, 5},
		{gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyblob", 2, status, 2, 5},
		{gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinytext", 3, status, 3, 5},
		{gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "text", 4, status, 4, 5},
		{gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "blob", 5, status, 5, 5},
		{gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumblob", 6, status, 6, 5},
		{gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumtext", 7, status, 7, 5},
		{gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longblob", 8, status, 8, 5},
		{gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longtext", 9, status, 9, 5},
		{gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyint", 0, status, 0, 6},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitExaFileUploadAndDownload(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.ExaFileUploadAndDownload{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}
