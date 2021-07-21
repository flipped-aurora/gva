package model

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin/system"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct {}

// TableName 初始化表名
// Author [SliverHorn](https://github.com/SliverHorn)
func (c *casbin) TableName() string {
	var entity model.Casbin
	return entity.TableName()
}

// Initialize system_authorities_menus表数据初始化
// Author [SliverHorn](https://github.com/SliverHorn)
func (c *casbin) Initialize() error {
	entities := []model.Casbin{
		{PType: "p", AuthorityId: "888", Path: "/base/login", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/register", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/createApi", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/getApiList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/getApiById", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/deleteApi", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/updateApi", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/api/getAllApis", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/createAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/deleteAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/getAuthorityList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/setDataAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/authority/updateAuthority", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/authority/copyAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getMenuList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/addBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getBaseMenuTree", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/addMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/deleteBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/updateBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/menu/getBaseMenuById", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/changePassword", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/getUserList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/setUserAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/user/deleteUser", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/getFileList", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/fileUploadAndDownload/deleteFile", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/casbin/updateCasbin", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/casbin/casbinTest/:pathParam", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/system/getSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/system/setSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/system/getServerInfo", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customer", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/customer/customerList", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/createTemp", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/preview", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getTables", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getDB", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getColumn", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/createSysDictionary", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/deleteSysDictionary", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/updateSysDictionary", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/findSysDictionary", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysDictionary/getSysDictionaryList", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/createSysOperationRecord", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/deleteSysOperationRecord", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/updateSysOperationRecord", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/findSysOperationRecord", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/getSysOperationRecordList", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/user/setUserInfo", Method: "PUT"},
		{PType: "p", AuthorityId: "888", Path: "/email/emailTest", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/simpleUploader/upload", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/simpleUploader/checkFileMd5", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/simpleUploader/mergeFileMd5", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/excel/importExcel", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/excel/loadExcel", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/excel/exportExcel", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/excel/downloadTemplate", Method: "GET"},
		{PType: "p", AuthorityId: "888", Path: "/api/deleteApisByIds", Method: "DELETE"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getSysHistory", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/rollback", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/getMeta", Method: "POST"},
		{PType: "p", AuthorityId: "888", Path: "/autoCode/delSysHistory", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/base/login", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/register", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/createApi", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/getApiList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/getApiById", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/deleteApi", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/updateApi", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/api/getAllApis", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/createAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/deleteAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/getAuthorityList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/authority/setDataAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getMenuList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/addBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getBaseMenuTree", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/addMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/deleteBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/updateBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/menu/getBaseMenuById", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/changePassword", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/getUserList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/user/setUserAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/fileUploadAndDownload/getFileList", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/fileUploadAndDownload/deleteFile", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/casbin/updateCasbin", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/system/getSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/system/setSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "POST"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "PUT"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "DELETE"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customer", Method: "GET"},
		{PType: "p", AuthorityId: "8881", Path: "/customer/customerList", Method: "GET"},
		{PType: "p", AuthorityId: "9528", Path: "/base/login", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/register", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/createApi", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/getApiList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/getApiById", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/deleteApi", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/updateApi", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/api/getAllApis", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/createAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/deleteAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/getAuthorityList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/authority/setDataAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getMenuList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/addBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getBaseMenuTree", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/addMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getMenuAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/deleteBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/updateBaseMenu", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/menu/getBaseMenuById", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/changePassword", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/getUserList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/user/setUserAuthority", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/fileUploadAndDownload/getFileList", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/fileUploadAndDownload/deleteFile", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/casbin/updateCasbin", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/casbin/getPolicyPathByAuthorityId", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/system/getSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/system/setSystemConfig", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "POST"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "PUT"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "DELETE"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customer", Method: "GET"},
		{PType: "p", AuthorityId: "9528", Path: "/customer/customerList", Method: "GET"},
		{PType: "p", AuthorityId: "9528", Path: "/autoCode/createTemp", Method: "POST"},
	}

	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.Authority 初始化数据
		return _errors.Wrap(err, c.TableName()+"表数据初始化失败!")
	}
	return nil
}

// CheckDataExist 检查重复数据,防止重复初始化数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (c *casbin) CheckDataExist() bool {
	if _errors.Is(global.Db.Where("ptype = ? AND v0 = ? AND v1 = ? AND v2 = ?", "p", "9528", "/autoCode/createTemp", "POST").First(&model.Casbin{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
