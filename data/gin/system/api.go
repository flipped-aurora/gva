package data

import (
	"github.com/flipped-aurora/gva/global"
	model "github.com/flipped-aurora/gva/model/gva/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var Api = new(api)

type api struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	var apis = []model.SysApi{
		{global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", global.I18n.T("UserLogin"), "base", "POST"},
		{global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/register", global.I18n.T("UserRegister"), "user", "POST"},
		{global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", global.I18n.T("CreateApi"), "api", "POST"},
		{global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", global.I18n.T("GetApiList"), "api", "POST"},
		{global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", global.I18n.T("GetApiDetail"), "api", "POST"},
		{global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", global.I18n.T("DeleteApi"), "api", "POST"},
		{global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", global.I18n.T("UpdateApi"), "api", "POST"},
		{global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", global.I18n.T("GetAllApis"), "api", "POST"},
		{global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", global.I18n.T("CreateAuthority"), "authority", "POST"},
		{global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", global.I18n.T("DeleteAuthority"), "authority", "POST"},
		{global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", global.I18n.T("GetAuthorityList"), "authority", "POST"},
		{global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", global.I18n.T("GetMenu"), "menu", "POST"},
		{global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", global.I18n.T("GetMenuList"), "menu", "POST"},
		{global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", global.I18n.T("AddBaseMenu"), "menu", "POST"},
		{global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", global.I18n.T("GetBaseMenuTree"), "menu", "POST"},
		{global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", global.I18n.T("AddMenuAuthority"), "menu", "POST"},
		{global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", global.I18n.T("GetMenuAuthority"), "menu", "POST"},
		{global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", global.I18n.T("DeleteBaseMenu"), "menu", "POST"},
		{global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", global.I18n.T("UpdateBaseMenu"), "menu", "POST"},
		{global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", global.I18n.T("GetBaseMenuById"), "menu", "POST"},
		{global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", global.I18n.T("ChangePassword"), "user", "POST"},
		{global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", global.I18n.T("GetUserList"), "user", "POST"},
		{global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", global.I18n.T("SetUserAuthority"), "user", "POST"},
		{global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/upload", global.I18n.T("UploadFile"), "fileUploadAndDownload", "POST"},
		{global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/getFileList", global.I18n.T("GetFileList"), "fileUploadAndDownload", "POST"},
		{global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", global.I18n.T("UpdateCasbin"), "casbin", "POST"},
		{global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", global.I18n.T("GetPolicyPathByAuthorityId"), "casbin", "POST"},
		{global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/deleteFile", global.I18n.T("DeleteFile"), "fileUploadAndDownload", "POST"},
		{global.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/jwt/jsonInBlacklist", global.I18n.T("JsonInBlacklist"), "jwt", "POST"},
		{global.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", global.I18n.T("SetDataAuthority"), "authority", "POST"},
		{global.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", global.I18n.T("GetSystemConfig"), "system", "POST"},
		{global.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", global.I18n.T("SetSystemConfig"), "system", "POST"},
		{global.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", global.I18n.T("CreateCustomer"), "customer", "POST"},
		{global.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", global.I18n.T("UpdateCustomer"), "customer", "PUT"},
		{global.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", global.I18n.T("DeleteCustomer"), "customer", "DELETE"},
		{global.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", global.I18n.T("GetCustomer"), "customer", "GET"},
		{global.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", global.I18n.T("GetCustomerList"), "customer", "GET"},
		{global.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", global.I18n.T("RESTFULTest"), "casbin", "GET"},
		{global.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/createTemp", global.I18n.T("CreateTemp"), "autoCode", "POST"},
		{global.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", global.I18n.T("UpdateAuthority"), "authority", "PUT"},
		{global.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", global.I18n.T("CopyAuthority"), "authority", "POST"},
		{global.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", global.I18n.T("DeleteUser"), "user", "DELETE"},
		{global.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/createSysDictionaryDetail", global.I18n.T("CreateSysDictionaryDetail"), "sysDictionaryDetail", "POST"},
		{global.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/deleteSysDictionaryDetail", global.I18n.T("DeleteSysDictionaryDetail"), "sysDictionaryDetail", "DELETE"},
		{global.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/updateSysDictionaryDetail", global.I18n.T("UpdateSysDictionaryDetail"), "sysDictionaryDetail", "PUT"},
		{global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/findSysDictionaryDetail", global.I18n.T("FindSysDictionaryDetail"), "sysDictionaryDetail", "GET"},
		{global.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/getSysDictionaryDetailList", global.I18n.T("GetSysDictionaryDetailList"), "sysDictionaryDetail", "GET"},
		{global.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/createSysDictionary", global.I18n.T("CreateSysDictionary"), "sysDictionary", "POST"},
		{global.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/deleteSysDictionary", global.I18n.T("DeleteSysDictionary"), "sysDictionary", "DELETE"},
		{global.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/updateSysDictionary", global.I18n.T("UpdateSysDictionary"), "sysDictionary", "PUT"},
		{global.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/findSysDictionary", global.I18n.T("FindSysDictionary"), "sysDictionary", "GET"},
		{global.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/getSysDictionaryList", global.I18n.T("GetSysDictionaryList"), "sysDictionary", "GET"},
		{global.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", global.I18n.T("CreateSysOperationRecord"), "sysOperationRecord", "POST"},
		{global.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", global.I18n.T("DeleteSysOperationRecord"), "sysOperationRecord", "DELETE"},
		{global.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", global.I18n.T("FindSysOperationRecord"), "sysOperationRecord", "GET"},
		{global.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", global.I18n.T("GetSysOperationRecordList"), "sysOperationRecord", "GET"},
		{global.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getTables", global.I18n.T("GetTables"), "autoCode", "GET"},
		{global.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getDB", global.I18n.T("GetDB"), "autoCode", "GET"},
		{global.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getColumn", global.I18n.T("GetColumn"), "autoCode", "GET"},
		{global.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", global.I18n.T("DeleteSysOperationRecordByIds"), "sysOperationRecord", "DELETE"},
		{global.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/upload", global.I18n.T("SubsectionUpload"), "simpleUploader", "POST"},
		{global.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/checkFileMd5", global.I18n.T("CheckFileMd5"), "simpleUploader", "GET"},
		{global.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/mergeFileMd5", global.I18n.T("MergeFileMd5"), "simpleUploader", "GET"},
		{global.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", global.I18n.T("SetUserInfo"), "user", "PUT"},
		{global.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getServerInfo", global.I18n.T("GetServerInfo"), "system", "POST"},
		{global.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/email/emailTest", global.I18n.T("EmailTest"), "email", "POST"},
		{global.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/createWorkflowProcess", global.I18n.T("CreateWorkflowProcess"), "workflowProcess", "POST"},
		{global.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcess", global.I18n.T("DeleteWorkflowProcess"), "workflowProcess", "DELETE"},
		{global.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcessByIds", global.I18n.T("DeleteWorkflowProcessByIds"), "workflowProcess", "DELETE"},
		{global.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/updateWorkflowProcess", global.I18n.T("UpdateWorkflowProcess"), "workflowProcess", "PUT"},
		{global.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowProcess", global.I18n.T("FindWorkflowProcess"), "workflowProcess", "GET"},
		{global.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowProcessList", global.I18n.T("GetWorkflowProcessList"), "workflowProcess", "GET"},
		{global.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowStep", global.I18n.T("FindWorkflowStep"), "workflowProcess", "GET"},
		{global.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/startWorkflow", global.I18n.T("StartWorkflow"), "workflowProcess", "POST"},
		{global.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyStated", global.I18n.T("GetMyStated"), "workflowProcess", "GET"},
		{global.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyNeed", global.I18n.T("GetMyNeed"), "workflowProcess", "GET"},
		{global.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowMoveByID", global.I18n.T("GetWorkflowMoveByID"), "workflowProcess", "GET"},
		{global.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/completeWorkflowMove", global.I18n.T("CompleteWorkflowMove"), "workflowProcess", "POST"},
		{global.Model{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/preview", global.I18n.T("Preview"), "autoCode", "POST"},
		{global.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/importExcel", global.I18n.T("ImportExcel"), "excel", "POST"},
		{global.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/loadExcel", global.I18n.T("LoadExcel"), "excel", "GET"},
		{global.Model{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/exportExcel", global.I18n.T("ExportExcel"), "excel", "POST"},
		{global.Model{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/downloadTemplate", global.I18n.T("DownloadTemplate"), "excel", "GET"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (a *api) TableName() string {
	return "sys_apis"
}
