package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure of table tools_gen_table for DAO operations like Where/Data.
type GenTable struct {
	g.Meta         `orm:"table:tools_gen_table, do:true"`
	TableId        interface{} // 编号
	TableName      interface{} // 表名称
	TableComment   interface{} // 表描述
	ClassName      interface{} // 实体类名称
	TplCategory    interface{} // 使用的模板（crud单表操作 tree树表操作）
	PackageName    interface{} // 生成包路径
	ModuleName     interface{} // 生成模块名
	BusinessName   interface{} // 生成业务名
	FunctionName   interface{} // 生成功能名
	FunctionAuthor interface{} // 生成功能作者
	Options        interface{} // 其它生成选项
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
	Remark         interface{} // 备注
}
