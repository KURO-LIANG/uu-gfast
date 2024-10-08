package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GenTableColumn is the golang structure of table tools_gen_table_column for DAO operations like Where/Data.
type GenTableColumn struct {
	g.Meta           `orm:"table:tools_gen_table_column, do:true"`
	ColumnId         interface{} // 编号
	TableId          interface{} // 归属表编号
	ColumnName       interface{} // 列名称
	ColumnComment    interface{} // 列描述
	ColumnType       interface{} // 列类型
	GoType           interface{} // Go类型
	GoField          interface{} // Go字段名
	HtmlField        interface{} // html字段名
	IsPk             interface{} // 是否主键（1是）
	IsIncrement      interface{} // 是否自增（1是）
	IsRequired       interface{} // 是否必填（1是）
	IsInsert         interface{} // 是否为插入字段（1是）
	IsEdit           interface{} // 是否编辑字段（1是）
	IsList           interface{} // 是否列表字段（1是）
	IsQuery          interface{} // 是否查询字段（1是）
	QueryType        interface{} // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType         interface{} // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType         interface{} // 字典类型
	Sort             interface{} // 排序
	LinkTableName    interface{} // 关联表名
	LinkTableClass   interface{} // 关联表类名
	LinkTablePackage interface{} // 关联表包名
	LinkLabelId      interface{} // 关联表键名
	LinkLabelName    interface{} // 关联表字段值
}
