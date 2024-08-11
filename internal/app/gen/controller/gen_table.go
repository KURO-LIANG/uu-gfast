/*
* @desc:代码生成
 */

package controller

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"uu-gfast/api/v1/common"
	"uu-gfast/api/v1/gen"
	"uu-gfast/internal/app/gen/model/entity"
	"uu-gfast/internal/app/gen/service"
	"strings"
)

type toolsGenTable struct {
}

var ToolsGenTableController = new(toolsGenTable)

// TableList 代码生成页列表数据
func (c *toolsGenTable) TableList(ctx context.Context, req *gen.ToolsGenTableSearchReq) (res *gen.ToolsGenTableSearchRes, err error) {
	res = new(gen.ToolsGenTableSearchRes)
	res.Total, res.List, err = service.ToolsGenTable.SelectListByPage(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// DataList 导入表格页列表数据
func (c *toolsGenTable) DataList(ctx context.Context, req *gen.ToolsGenTableDataListReq) (res *gen.ToolsGenTableSearchRes, err error) {
	res = new(gen.ToolsGenTableSearchRes)
	res.Total, res.List, err = service.ToolsGenTable.SelectDbTableList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// ImportTableSave 导入表结构操作
func (c *toolsGenTable) ImportTableSave(ctx context.Context, req *gen.ToolsImportTableReq) (res *common.NoneRes, err error) {
	res = new(common.NoneRes)
	if req.Tables == "" {
		return res, gerror.New("请选择要导入的表")
	}
	tableArr := strings.Split(req.Tables, ",")
	tableList, err := service.ToolsGenTable.SelectDbTableListByNames(ctx, tableArr)
	if err != nil {
		return res, gerror.New("表信息不存在")
	}
	if tableList == nil {
		return res, gerror.New("表信息不存在")
	}
	err = service.ToolsGenTable.ImportGenTable(ctx, tableList)
	if err != nil {
		return res, err
	}
	return res, nil
}

// ColumnList 表格字段列表数据
func (c *toolsGenTable) ColumnList(ctx context.Context, req *gen.ToolsColumnsReq) (res *gen.ToolsColumnsRes, err error) {
	res = new(gen.ToolsColumnsRes)
	tableId := req.TableId
	if tableId == 0 {
		return res, gerror.New("表不存在")
	}
	list, err := service.ToolsGenTableColumn.SelectGenTableColumnListByTableId(ctx, tableId)
	if err != nil {
		return res, err
	}
	var tableInfo *entity.GenTable
	var tableMap g.Map
	tableInfo, err = service.ToolsGenTable.GetTableInfoByTableId(ctx, tableId)
	if err != nil {
		return res, err
	}
	tableMap = gconv.Map(tableInfo)
	//如果是树表则设置树表配置
	if tableInfo != nil && tableInfo.TplCategory == "tree" {
		options := gjson.New(tableInfo.Options)
		tableMap["treeCode"] = options.Get("treeCode")
		tableMap["treeParentCode"] = options.Get("treeParentCode")
		tableMap["treeName"] = options.Get("treeName")
	}
	res.Rows = list
	res.Info = tableMap
	return
}

// RelationTable 获取可选的关联表
func (c *toolsGenTable) RelationTable(ctx context.Context, req *gen.RelationTableReq) (res *gen.RelationTableRes, err error) {
	res = new(gen.RelationTableRes)
	//获取表数据列表
	_, tableList, err := service.ToolsGenTable.SelectListByPage(ctx, &gen.ToolsGenTableSearchReq{
		PageReq: gen.PageReq{
			PageSize: 1000,
		},
	})
	if err != nil {
		return nil, err
	}
	//获取所有字段
	allColumns, err := service.ToolsGenTableColumn.GetAllTableColumns(ctx)
	if err != nil {
		return nil, err
	}
	tableColumns := make([]*gen.ToolsGenTableColumnsRes, len(tableList))
	for k, v := range tableList {
		tableColumns[k] = &gen.ToolsGenTableColumnsRes{
			GenTable: v,
			Columns:  make([]*entity.GenTableColumn, 0),
		}
		for _, cv := range allColumns {
			if cv.TableId == v.TableId {
				tableColumns[k].Columns = append(tableColumns[k].Columns, cv)
			}
		}
	}
	res.Data = tableColumns
	return
}

// EditSave 编辑表格生成信息
func (c *toolsGenTable) EditSave(ctx context.Context, req *gen.ToolsGenTableEditReq) (res *common.NoneRes, err error) {
	res = new(common.NoneRes)
	err = service.ToolsGenTable.SaveEdit(ctx, req)
	return
}

// Preview 代码预览
func (c *toolsGenTable) Preview(ctx context.Context, req *gen.PreviewReq) (res *gen.PreviewRes, err error) {
	res = new(gen.PreviewRes)
	tableId := req.TableId
	if tableId == 0 {
		return res, gerror.New("参数错误")
	}
	data, _, err := service.ToolsGenTable.GenData(tableId, ctx)
	if err != nil {
		return res, err
	}
	res.Data = data
	return
}

// BatchGenCode 代码生成
func (c *toolsGenTable) BatchGenCode(ctx context.Context, req *gen.BatchGenCodeReq) (res *common.NoneRes, err error) {
	ids := req.Ids
	if len(ids) == 0 {
		return nil, gerror.New("参数错误")
	}
	err = service.ToolsGenTable.GenCode(ctx, ids)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Delete 删除导入的表信息
func (c *toolsGenTable) Delete(ctx context.Context, req *gen.DeleteReq) (res *common.NoneRes, err error) {
	ids := req.Ids
	if len(ids) == 0 {
		return nil, gerror.New("参数错误")
	}
	err = service.ToolsGenTable.Delete(ctx, ids)
	if err != nil {
		return nil, err
	}
	return res, err
}
