package gen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "uu-gfast/api/v1/common"
	"uu-gfast/internal/app/gen/model/entity"
)

type PageReq struct {
	BeginTime string          `p:"beginTime"` //开始时间
	EndTime   string          `p:"endTime"`   //结束时间
	PageNum   int             `p:"pageNum"`   //当前页码
	PageSize  int             `p:"pageSize"`  //每页数
	Ctx       context.Context `swaggerignore:"true"`
	OrderBy   string          //排序方式
}

// ToolsGenTableSearchReq 分页请求参数
type ToolsGenTableSearchReq struct {
	g.Meta       `path:"/tableList" tags:"代码生成" method:"get" summary:"表列表"`
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	PageReq
}

type ToolsGenTableDataListReq struct {
	g.Meta       `path:"/dataList" tags:"代码生成" method:"get" summary:"已导入表数据"`
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	PageReq
}

// ToolsImportTableReq 导入表结构操作
type ToolsImportTableReq struct {
	g.Meta `path:"/importTableSave" tags:"代码生成" method:"post" summary:"表导入操作"`
	Tables string `p:"tables"` //表名称
}

// ToolsColumnsReq 表的列
type ToolsColumnsReq struct {
	g.Meta  `path:"/columnList" tags:"代码生成" method:"get" summary:"表的列"`
	TableId int64 `p:"tableId"` //表Id
}

// ToolsColumnsRes 表几列返回
type ToolsColumnsRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.GenTableColumn `json:"rows"` //表Id
	Info   g.Map                    `json:"info"` //表信息
}

// ToolsGenTableSearchRes 查询列表返回
type ToolsGenTableSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.GenTable `json:"list"`
	commonApi.ListRes
}

// RelationTableReq 关联表查询
type RelationTableReq struct {
	g.Meta `path:"/relationTable" tags:"代码生成" method:"get" summary:"关联表列表"`
}

// RelationTableRes 关联表查询返回
type RelationTableRes struct {
	g.Meta `mime:"application/json"`
	Data   []*ToolsGenTableColumnsRes `json:"data"`
}

// PreviewReq 预览代码请求
type PreviewReq struct {
	g.Meta  `path:"/preview" tags:"代码生成" method:"get" summary:"预览代码"`
	TableId int64 `p:"tableId"` //表Id
}

// PreviewRes 预览代码返回
type PreviewRes struct {
	g.Meta `mime:"application/json"`
	Data   g.MapStrStr `json:"data"`
}

// BatchGenCodeReq 批量生成代码 请求
type BatchGenCodeReq struct {
	g.Meta `path:"/batchGenCode" tags:"代码生成" method:"put" summary:"生成/批量生成代码"`
	Ids    []int `p:"ids"` //表Id
}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/delete" tags:"代码生成" method:"delete" summary:"删除导入的表"`
	Ids    []int `p:"ids"` //表Id
}

// ToolsGenTableColumnsRes 表与字段组合数据
type ToolsGenTableColumnsRes struct {
	g.Meta `mime:"application/json"`
	*entity.GenTable
	Columns []*entity.GenTableColumn `json:"columns"`
}

// ToolsGenTableEditReq 生成信息修改参数
type ToolsGenTableEditReq struct {
	g.Meta         `path:"/editSave" tags:"代码生成" method:"put" summary:"生成修改参数"`
	TableId        int64                    `p:"tableId" v:"required#主键ID不能为空"`
	TableName      string                   `p:"tableName"  v:"required#表名称不能为空"`
	TableComment   string                   `p:"tableComment"  v:"required#表描述不能为空"`
	ClassName      string                   `p:"className" v:"required#实体类名称不能为空"`
	FunctionAuthor string                   `p:"functionAuthor"  v:"required#作者不能为空"`
	TplCategory    string                   `p:"tplCategory"`
	PackageName    string                   `p:"packageName" v:"required#生成包路径不能为空"`
	ModuleName     string                   `p:"moduleName" v:"required#生成模块名不能为空"`
	BusinessName   string                   `p:"businessName" v:"required#生成业务名不能为空"`
	FunctionName   string                   `p:"functionName" v:"required#生成功能名不能为空"`
	Remark         string                   `p:"remark"`
	Params         string                   `p:"params"`
	Columns        []*entity.GenTableColumn `p:"columns"`
	TreeCode       string                   `p:"tree_code"`
	TreeParentCode string                   `p:"tree_parent_code"`
	TreeName       string                   `p:"tree_name"`
	UserName       string
}

// ToolsGenTableExtend 实体扩展
type ToolsGenTableExtend struct {
	TableId        int64                    `orm:"table_id,primary" json:"table_id"`        // 编号
	TableName      string                   `orm:"table_name"       json:"table_name"`      // 表名称
	TableComment   string                   `orm:"table_comment"    json:"table_comment"`   // 表描述
	ClassName      string                   `orm:"class_name"       json:"class_name"`      // 实体类名称
	TplCategory    string                   `orm:"tpl_category"     json:"tpl_category"`    // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string                   `orm:"package_name"     json:"package_name"`    // 生成包路径
	ModuleName     string                   `orm:"module_name"      json:"module_name"`     // 生成模块名
	BusinessName   string                   `orm:"business_name"    json:"business_name"`   // 生成业务名
	FunctionName   string                   `orm:"function_name"    json:"function_name"`   // 生成功能名
	FunctionAuthor string                   `orm:"function_author"  json:"function_author"` // 生成功能作者
	Options        string                   `orm:"options"          json:"options"`         // 其它生成选项
	CreateBy       string                   `orm:"create_by"        json:"create_by"`       // 创建者
	CreateTime     *gtime.Time              `orm:"create_time"      json:"create_time"`     // 创建时间
	UpdateBy       string                   `orm:"update_by"        json:"update_by"`       // 更新者
	UpdateTime     *gtime.Time              `orm:"update_time"      json:"update_time"`     // 更新时间
	Remark         string                   `orm:"remark"           json:"remark"`          // 备注
	TreeCode       string                   `json:"tree_code"`                              // 树编码字段
	TreeParentCode string                   `json:"tree_parent_code"`                       // 树父编码字段
	TreeName       string                   `json:"tree_name"`                              // 树名称字段
	Columns        []*entity.GenTableColumn `json:"columns"`                                // 表列信息
	PkColumn       *entity.GenTableColumn   `json:"pkColumn"`                               // 主键列信息
}
