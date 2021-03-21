package data

import (
	"github.com/flipped-aurora/gva/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: apis 表数据初始化
func (a *api) Init() error {
	apis := []model.Api{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/login", Description: global.I18n.T("UserLogin"), ApiGroup: "base", Method: "POST"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/register", Description: global.I18n.T("UserRegister"), ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/createApi", Description: global.I18n.T("CreateApi"), ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiList", Description: global.I18n.T("GetApiList"), ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiById", Description: global.I18n.T("GetApiDetail"), ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApi", Description: global.I18n.T("DeleteApi"), ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/updateApi", Description: global.I18n.T("UpdateApi"), ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getAllApis", Description: global.I18n.T("GetAllApis"), ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/createAuthority", Description: global.I18n.T("CreateAuthority"), ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/deleteAuthority", Description: global.I18n.T("DeleteAuthority"), ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/getAuthorityList", Description: global.I18n.T("GetAuthorityList"), ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenu", Description: global.I18n.T("GetMenu"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuList", Description: global.I18n.T("GetMenuList"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addBaseMenu", Description: global.I18n.T("AddBaseMenu"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuTree", Description: global.I18n.T("GetBaseMenuTree"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addMenuAuthority", Description: global.I18n.T("AddMenuAuthority"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuAuthority", Description: global.I18n.T("GetMenuAuthority"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/deleteBaseMenu", Description: global.I18n.T("DeleteBaseMenu"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/updateBaseMenu", Description: global.I18n.T("UpdateBaseMenu"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuById", Description: global.I18n.T("GetBaseMenuById"), ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/changePassword", Description: global.I18n.T("ChangePassword"), ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserList", Description: global.I18n.T("GetUserList"), ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthority", Description: global.I18n.T("SetUserAuthority"), ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/upload", Description: global.I18n.T("UploadFile"), ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/getFileList", Description: global.I18n.T("GetFileList"), ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/updateCasbin", Description: global.I18n.T("UpdateCasbin"), ApiGroup: "casbin", Method: "POST"},
		{Model: global.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/getPolicyPathByAuthorityId", Description: global.I18n.T("GetPolicyPathByAuthorityId"), ApiGroup: "casbin", Method: "POST"},
		{Model: global.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/deleteFile", Description: global.I18n.T("DeleteFile"), ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/jwt/jsonInBlacklist", Description: global.I18n.T("JsonInBlacklist"), ApiGroup: "jwt", Method: "POST"},
		{Model: global.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/setDataAuthority", Description: global.I18n.T("SetDataAuthority"), ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getSystemConfig", Description: global.I18n.T("GetSystemConfig"), ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/setSystemConfig", Description: global.I18n.T("SetSystemConfig"), ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: global.I18n.T("CreateCustomer"), ApiGroup: "customer", Method: "POST"},
		{Model: global.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: global.I18n.T("UpdateCustomer"), ApiGroup: "customer", Method: "PUT"},
		{Model: global.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: global.I18n.T("DeleteCustomer"), ApiGroup: "customer", Method: "DELETE"},
		{Model: global.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: global.I18n.T("GetCustomer"), ApiGroup: "customer", Method: "GET"},
		{Model: global.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customerList", Description: global.I18n.T("GetCustomerList"), ApiGroup: "customer", Method: "GET"},
		{Model: global.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/casbinTest/:pathParam", Description: global.I18n.T("RESTFULTest"), ApiGroup: "casbin", Method: "GET"},
		{Model: global.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/createTemp", Description: global.I18n.T("CreateTemp"), ApiGroup: "autoCode", Method: "POST"},
		{Model: global.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/updateAuthority", Description: global.I18n.T("UpdateAuthority"), ApiGroup: "authority", Method: "PUT"},
		{Model: global.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/copyAuthority", Description: global.I18n.T("CopyAuthority"), ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUser", Description: global.I18n.T("DeleteUser"), ApiGroup: "user", Method: "DELETE"},
		{Model: global.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: global.I18n.T("CreateSysDictionaryDetail"), ApiGroup: "sysDictionaryDetail", Method: "POST"},
		{Model: global.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: global.I18n.T("DeleteSysDictionaryDetail"), ApiGroup: "sysDictionaryDetail", Method: "DELETE"},
		{Model: global.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: global.I18n.T("UpdateSysDictionaryDetail"), ApiGroup: "sysDictionaryDetail", Method: "PUT"},
		{Model: global.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: global.I18n.T("FindSysDictionaryDetail"), ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: global.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: global.I18n.T("GetSysDictionaryDetailList"), ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: global.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/createSysDictionary", Description: global.I18n.T("CreateSysDictionary"), ApiGroup: "sysDictionary", Method: "POST"},
		{Model: global.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/deleteSysDictionary", Description: global.I18n.T("DeleteSysDictionary"), ApiGroup: "sysDictionary", Method: "DELETE"},
		{Model: global.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/updateSysDictionary", Description: global.I18n.T("UpdateSysDictionary"), ApiGroup: "sysDictionary", Method: "PUT"},
		{Model: global.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/findSysDictionary", Description: global.I18n.T("FindSysDictionary"), ApiGroup: "sysDictionary", Method: "GET"},
		{Model: global.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/getSysDictionaryList", Description: global.I18n.T("GetSysDictionaryList"), ApiGroup: "sysDictionary", Method: "GET"},
		{Model: global.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/createSysOperationRecord", Description: global.I18n.T("CreateSysOperationRecord"), ApiGroup: "sysOperationRecord", Method: "POST"},
		{Model: global.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecord", Description: global.I18n.T("DeleteSysOperationRecord"), ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: global.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/findSysOperationRecord", Description: global.I18n.T("FindSysOperationRecord"), ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: global.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/getSysOperationRecordList", Description: global.I18n.T("GetSysOperationRecordList"), ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: global.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getTables", Description: global.I18n.T("GetTables"), ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getDB", Description: global.I18n.T("GetDB"), ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getColumn", Description: global.I18n.T("GetColumn"), ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: global.I18n.T("DeleteSysOperationRecordByIds"), ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: global.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/upload", Description: global.I18n.T("SubsectionUpload"), ApiGroup: "simpleUploader", Method: "POST"},
		{Model: global.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/checkFileMd5", Description: global.I18n.T("CheckFileMd5"), ApiGroup: "simpleUploader", Method: "GET"},
		{Model: global.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUploader/mergeFileMd5", Description: global.I18n.T("MergeFileMd5"), ApiGroup: "simpleUploader", Method: "GET"},
		{Model: global.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserInfo", Description: global.I18n.T("SetUserInfo"), ApiGroup: "user", Method: "PUT"},
		{Model: global.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getServerInfo", Description: global.I18n.T("GetServerInfo"), ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/email/emailTest", Description: global.I18n.T("EmailTest"), ApiGroup: "email", Method: "POST"},
		{Model: global.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/createWorkflowProcess", Description: global.I18n.T("CreateWorkflowProcess"), ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/deleteWorkflowProcess", Description: global.I18n.T("DeleteWorkflowProcess"), ApiGroup: "workflowProcess", Method: "DELETE"},
		{Model: global.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/deleteWorkflowProcessByIds", Description: global.I18n.T("DeleteWorkflowProcessByIds"), ApiGroup: "workflowProcess", Method: "DELETE"},
		{Model: global.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/updateWorkflowProcess", Description: global.I18n.T("UpdateWorkflowProcess"), ApiGroup: "workflowProcess", Method: "PUT"},
		{Model: global.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/findWorkflowProcess", Description: global.I18n.T("FindWorkflowProcess"), ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getWorkflowProcessList", Description: global.I18n.T("GetWorkflowProcessList"), ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/findWorkflowStep", Description: global.I18n.T("FindWorkflowStep"), ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/startWorkflow", Description: global.I18n.T("StartWorkflow"), ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getMyStated", Description: global.I18n.T("GetMyStated"), ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getMyNeed", Description: global.I18n.T("GetMyNeed"), ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/getWorkflowMoveByID", Description: global.I18n.T("GetWorkflowMoveByID"), ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/workflowProcess/completeWorkflowMove", Description: global.I18n.T("CompleteWorkflowMove"), ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/preview", Description: global.I18n.T("Preview"), ApiGroup: "autoCode", Method: "POST"},
		{Model: global.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/importExcel", Description: global.I18n.T("ImportExcel"), ApiGroup: "excel", Method: "POST"},
		{Model: global.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/loadExcel", Description: global.I18n.T("LoadExcel"), ApiGroup: "excel", Method: "GET"},
		{Model: global.Model{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/exportExcel", Description: global.I18n.T("ExportExcel"), ApiGroup: "excel", Method: "POST"},
		{Model: global.Model{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/downloadTemplate", Description: global.I18n.T("DownloadTemplate"), ApiGroup: "excel", Method: "GET"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 84}).Find(&[]model.Api{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> apis 表的初始数据已存在!")
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
	return "apis"
}
