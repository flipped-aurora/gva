package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

func (c *casbin) TableName() string {
	var entity system.Casbin
	return entity.TableName()
}

func (c *casbin) Initialize() error {
	entities := []system.Casbin{
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/base/login"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/user/register"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/api/createApi"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/api/getApiList"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/api/getApiById"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/api/deleteApi"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/api/updateApi"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/api/getAllApis"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/api/deleteApisByIds"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/authority/copyAuthority"},
		{PType: "p", AuthorityId: "888", Method: "PUT", Path: "/authority/updateAuthority"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/authority/createAuthority"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/authority/deleteAuthority"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/authority/getAuthorityList"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/authority/setDataAuthority"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/getMenu"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/getMenuList"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/addBaseMenu"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/getBaseMenuTree"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/addMenuAuthority"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/getMenuAuthority"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/deleteBaseMenu"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/updateBaseMenu"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/menu/getBaseMenuById"},

		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/user/getUserInfo"},
		{PType: "p", AuthorityId: "888", Method: "PUT", Path: "/user/setUserInfo"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/user/getUserList"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/user/deleteUser"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/user/changePassword"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/user/setUserAuthority"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/user/setUserAuthorities"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/fileUploadAndDownload/upload"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/fileUploadAndDownload/deleteFile"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/fileUploadAndDownload/getFileList"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/casbin/updateCasbin"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/jwt/jsonInBlacklist"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/system/getSystemConfig"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/system/setSystemConfig"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/system/getServerInfo"},

		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "888", Method: "PUT", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/customer/customerList"},

		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/autoCode/getDB"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/autoCode/getMeta"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/autoCode/preview"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/autoCode/getTables"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/autoCode/getColumn"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/autoCode/rollback"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/autoCode/createTemp"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/autoCode/delSysHistory"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/autoCode/getSysHistory"},

		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail"},
		{PType: "p", AuthorityId: "888", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail"},

		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/sysDictionary/findSysDictionary"},
		{PType: "p", AuthorityId: "888", Method: "PUT", Path: "/sysDictionary/updateSysDictionary"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/sysDictionary/getSysDictionaryList"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/sysDictionary/createSysDictionary"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary"},

		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord"},
		{PType: "p", AuthorityId: "888", Method: "PUT", Path: "/sysOperationRecord/updateSysOperationRecord"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord"},
		{PType: "p", AuthorityId: "888", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/email/emailTest"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/simpleUploader/upload"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/simpleUploader/checkFileMd5"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/simpleUploader/mergeFileMd5"},

		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/excel/importExcel"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/excel/loadExcel"},
		{PType: "p", AuthorityId: "888", Method: "POST", Path: "/excel/exportExcel"},
		{PType: "p", AuthorityId: "888", Method: "GET", Path: "/excel/downloadTemplate"},

		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/base/login"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/user/register"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/api/createApi"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/api/getApiList"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/api/getApiById"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/api/deleteApi"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/api/updateApi"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/api/getAllApis"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/authority/createAuthority"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/authority/deleteAuthority"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/authority/getAuthorityList"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/authority/setDataAuthority"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/getMenu"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/getMenuList"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/addBaseMenu"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/getBaseMenuTree"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/addMenuAuthority"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/getMenuAuthority"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/deleteBaseMenu"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/updateBaseMenu"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/menu/getBaseMenuById"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/user/changePassword"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/user/getUserList"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/user/setUserAuthority"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/fileUploadAndDownload/upload"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/fileUploadAndDownload/getFileList"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/fileUploadAndDownload/deleteFile"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/casbin/updateCasbin"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/jwt/jsonInBlacklist"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/system/getSystemConfig"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/system/setSystemConfig"},
		{PType: "p", AuthorityId: "8881", Method: "POST", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "8881", Method: "PUT", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "8881", Method: "DELETE", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "8881", Method: "GET", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "8881", Method: "GET", Path: "/customer/customerList"},
		{PType: "p", AuthorityId: "8881", Method: "GET", Path: "/user/getUserInfo"},

		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/base/login"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/user/register"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/api/createApi"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/api/getApiList"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/api/getApiById"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/api/deleteApi"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/api/updateApi"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/api/getAllApis"},

		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/authority/createAuthority"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/authority/deleteAuthority"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/authority/getAuthorityList"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/authority/setDataAuthority"},

		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/getMenu"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/getMenuList"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/addBaseMenu"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/getBaseMenuTree"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/addMenuAuthority"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/getMenuAuthority"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/deleteBaseMenu"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/updateBaseMenu"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/menu/getBaseMenuById"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/user/changePassword"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/user/getUserList"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/user/setUserAuthority"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/fileUploadAndDownload/upload"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/fileUploadAndDownload/getFileList"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/fileUploadAndDownload/deleteFile"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/casbin/updateCasbin"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/jwt/jsonInBlacklist"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/system/getSystemConfig"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/system/setSystemConfig"},
		{PType: "p", AuthorityId: "9528", Method: "PUT", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "9528", Method: "GET", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "9528", Method: "DELETE", Path: "/customer/customer"},
		{PType: "p", AuthorityId: "9528", Method: "GET", Path: "/customer/customerList"},
		{PType: "p", AuthorityId: "9528", Method: "POST", Path: "/autoCode/createTemp"},
		{PType: "p", AuthorityId: "9528", Method: "GET", Path: "/user/getUserInfo"},
	}
	if err := global.Db.Create(&entities).Error; err != nil {
		return errors.Wrap(err, c.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (c *casbin) CheckDataExist() bool {
	if errors.Is(global.Db.Where(system.Casbin{PType: "p", AuthorityId: "9528", Method: "GET", Path: "/user/getUserInfo"}).First(&system.Casbin{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
