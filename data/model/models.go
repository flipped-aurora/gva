package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
}

type SysApi struct {
	gorm.Model
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:POST" gorm:"comment:方法"`
}

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

type ExaFile struct {
	gorm.Model
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}

type ExaCustomer struct {
	gorm.Model
	CustomerName       string  `json:"customerName" form:"customerName" gorm:"comment:客户名"`
	CustomerPhoneData  string  `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`
	SysUserID          uint    `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"`
	SysUser            SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`
}

type SysBaseMenu struct {
	gorm.Model
	MenuLevel     uint   `json:"-"`
	ParentId      string `json:"parentId" gorm:"comment:父菜单ID"`
	Path          string `json:"path" gorm:"comment:路由path"`
	Name          string `json:"name" gorm:"comment:路由name"`
	Hidden        bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component     string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort          int    `json:"sort" gorm:"comment:排序标记"`
	Meta          `json:"meta" gorm:"comment:附加属性"`
	SysAuthoritys []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter `json:"parameters"`
}

type SysWorkflow struct {
	gorm.Model
	WorkflowNickName    string                `json:"workflowNickName" gorm:"comment:工作流中文名称"`  // 工作流名称
	WorkflowName        string                `json:"workflowName" gorm:"comment:工作流英文名称"`      // 工作流英文id
	WorkflowDescription string                `json:"workflowDescription" gorm:"comment:工作流描述"` // 工作流描述
	WorkflowStepInfo    []SysWorkflowStepInfo `json:"workflowStep" gorm:"comment:工作流步骤"`        // 工作流步骤
}

type ExaFileChunk struct {
	gorm.Model
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}

type SysDictionary struct {
	gorm.Model
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:状态"`
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

type ExaSimpleUploader struct {
	ChunkNumber      string `json:"chunkNumber" gorm:"comment:当前切片标记"`
	CurrentChunkSize string `json:"currentChunkSize" gorm:"comment:当前切片容量"`
	CurrentChunkPath string `json:"currentChunkPath" gorm:"comment:切片本地路径"`
	TotalSize        string `json:"totalSize" gorm:"comment:总容量"`
	Identifier       string `json:"identifier" gorm:"comment:文件标识（md5）"`
	Filename         string `json:"filename" gorm:"comment:文件名"`
	TotalChunks      string `json:"totalChunks" gorm:"comment:切片总数"`
	IsDone           bool   `json:"isDone" gorm:"comment:是否上传完成"`
	FilePath         string `json:"filePath" gorm:"comment:文件本地路径"`
}

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
}

type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string
	DataAuthorityIdAuthorityId string
}

type SysOperationRecord struct {
	gorm.Model
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟"`
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`
	Body         string        `json:"body" form:"body" gorm:"column:body;comment:请求Body"`
	Resp         string        `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User         SysUser       `json:"user"`
}

type SysDictionaryDetail struct {
	gorm.Model
	Label           string `json:"label" form:"label" gorm:"column:label;comment:展示值"`
	Value           int    `json:"value" form:"value" gorm:"column:value;comment:字典值"`
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:关联标记"`
}

type SysWorkflowStepInfo struct {
	gorm.Model
	SysWorkflowID   uint    `json:"workflowID" gorm:"comment:所属工作流ID"`      // 所属工作流ID
	IsStart         bool    `json:"isStart" gorm:"comment:是否是开始流节点"`        // 是否是开始流节点
	StepName        string  `json:"stepName" gorm:"comment:工作流节点名称"`        // 工作流名称
	StepNo          float64 `json:"stepNo" gorm:"comment:步骤id （第几步）"`       // 步骤id （第几步）
	StepAuthorityID string  `json:"stepAuthorityID" gorm:"comment:操作者级别id"` // 操作者级别id
	IsEnd           bool    `json:"isEnd" gorm:"comment:是否是完结流节点"`          // 是否是完结流节点
}

type SysBaseMenuParameter struct {
	gorm.Model
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"commit:'地址栏携带参数为params还是query'"`
	Key           string `json:"key" gorm:"commit:'地址栏携带参数的key'"`
	Value         string `json:"value" gorm:"commit:'地址栏携带参数的值'"`
}

type ExaFileUploadAndDownload struct {
	gorm.Model
	Name string `json:"name" gorm:"comment:文件名"`
	Url  string `json:"url" gorm:"comment:文件地址"`
	Tag  string `json:"tag" gorm:"comment:文件标签"`
	Key  string `json:"key" gorm:"comment:编号"`
}

type SysDictionaryToPostgresql struct {
	gorm.Model
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:状态"`
	Description          string                `json:"description" form:"description" gorm:"column:description;comment:'描述'"`
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}
