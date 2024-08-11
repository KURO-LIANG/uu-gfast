package service

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"os"
	"runtime"
	"uu-gfast/api/v1/gen"
	"uu-gfast/internal/app/gen/dao"
	"uu-gfast/internal/app/gen/model/entity"
	"uu-gfast/internal/app/system/consts"
	"strings"
)

type toolsGenTable struct{}

var ToolsGenTable = new(toolsGenTable)

// SelectListByPage 查询已导入的数据表
func (s *toolsGenTable) SelectListByPage(ctx context.Context, param *gen.ToolsGenTableSearchReq) (total int, list []*entity.GenTable, err error) {
	model := dao.GenTable.Ctx(ctx)
	if param != nil {
		if param.TableName != "" {
			model = model.Where(dao.GenTable.Columns().TableName+" like ?", "%"+param.TableName+"%")
		}
		if param.TableComment != "" {
			model = model.Where(dao.GenTable.Columns().TableComment+"like ?", "%"+param.TableComment+"%")
		}
		if param.BeginTime != "" {
			model = model.Where(dao.GenTable.Columns().CreateTime+" >= ", param.BeginTime)
		}
		if param.EndTime != "" {
			model = model.Where(dao.GenTable.Columns().CreateTime+" <= ", param.EndTime)
		}
		total, err = model.Count()
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取总行数失败")
			return
		}
		if param.PageNum == 0 {
			param.PageNum = 1
		}
		if param.PageSize == 0 {
			param.PageSize = consts.PageSize
		}
		err = model.Page(param.PageNum, param.PageSize).Order(dao.GenTable.Columns().TableId + " asc").Scan(&list)
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取数据失败")
		}
	}
	return
}

// SelectDbTableList 查询据库表
func (s *toolsGenTable) SelectDbTableList(ctx context.Context, param *gen.ToolsGenTableDataListReq) (total int, list []*entity.GenTable, err error) {
	db := g.DB()
	if s.getDbDriver() != "mysql" {
		err = gerror.New("代码生成暂时只支持mysql数据库")
		return
	}
	sql := " from information_schema.tables where table_schema = (select database())" +
		" and table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%' and table_name NOT IN (select table_name from " + dao.GenTable.Table() + ") "
	if param != nil {
		if param.TableName != "" {
			sql += gdb.FormatSqlWithArgs(" and lower(table_name) like lower(?)", []interface{}{"%" + param.TableName + "%"})
		}

		if param.TableComment != "" {
			sql += gdb.FormatSqlWithArgs(" and lower(table_comment) like lower(?)", []interface{}{"%" + param.TableComment + "%"})
		}

		if param.BeginTime != "" {
			sql += gdb.FormatSqlWithArgs(" and date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", []interface{}{param.BeginTime})
		}

		if param.EndTime != "" {
			sql += gdb.FormatSqlWithArgs(" and date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", []interface{}{param.EndTime})
		}
	}
	countSql := "select count(1) " + sql
	total, err = db.GetCount(ctx, countSql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取总表数失败")
		return
	}
	sql = "table_name, table_comment, create_time, update_time " + sql

	if param.PageNum == 0 {
		param.PageNum = 1
	}

	if param.PageSize == 0 {
		param.PageSize = consts.PageSize
	}
	page := (param.PageNum - 1) * param.PageSize
	sql += " order by create_time desc,table_name asc limit  " + gconv.String(page) + "," + gconv.String(param.PageSize)
	err = db.GetScan(ctx, &list, "select "+sql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取数据失败")
	}
	return
}

// 获取数据库驱动类型
func (s *toolsGenTable) getDbDriver() string {
	config := g.DB().GetConfig()
	return gstr.ToLower(config.Type)
}

// SelectDbTableListByNames 查询数据库中对应的表数据
func (s *toolsGenTable) SelectDbTableListByNames(ctx context.Context, tableNames []string) ([]*entity.GenTable, error) {
	if s.getDbDriver() != "mysql" {
		return nil, gerror.New("代码生成只支持mysql数据库")
	}
	db := g.DB()
	sql := "select * from information_schema.tables where table_name NOT LIKE 'qrtz_%' and table_name NOT LIKE 'gen_%' " +
		" and table_schema = (select database()) "
	if len(tableNames) > 0 {
		in := gstr.TrimRight(gstr.Repeat("?,", len(tableNames)), ",")
		sql += " and " + gdb.FormatSqlWithArgs("table_name in ("+in+")", gconv.SliceAny(tableNames))
	}
	var result []*entity.GenTable
	err := db.GetScan(ctx, &result, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取表格信息失败")
	}
	return result, err
}

// ImportGenTable 导入表结构
func (s *toolsGenTable) ImportGenTable(ctx context.Context, tableList []*entity.GenTable) error {
	if tableList != nil {
		tx, err := g.DB().Begin(ctx)
		if err != nil {
			return err
		}
		for _, table := range tableList {
			tableName := table.TableName
			s.InitTable(ctx, table)
			tmpId, err := tx.Model(dao.GenTable.Table()).InsertAndGetId(table)
			if err != nil {
				return err
			}
			if err != nil || tmpId <= 0 {
				tx.Rollback()
				return gerror.New("保存数据失败")
			}

			table.TableId = tmpId

			// 保存列信息
			genTableColumns, err := ToolsGenTableColumn.SelectDbTableColumnsByName(ctx, tableName)

			if err != nil || len(genTableColumns) <= 0 {
				tx.Rollback()
				return gerror.New("获取列数据失败")
			}
			for _, column := range genTableColumns {
				ToolsGenTableColumn.InitColumnField(ctx, column, table)
				_, err = tx.Model(dao.GenTableColumn.Table()).Insert(column)
				if err != nil {
					tx.Rollback()
					return gerror.New("保存列数据失败")
				}
			}
		}
		return tx.Commit()
	} else {
		return gerror.New("参数错误")
	}
}

// InitTable 初始化表信息
func (s *toolsGenTable) InitTable(ctx context.Context, table *entity.GenTable) {
	table.ClassName = s.ConvertClassName(ctx, table.TableName)
	table.PackageName = g.Cfg().MustGet(ctx, "gen.packageName").String()
	table.ModuleName = g.Cfg().MustGet(ctx, "gen.moduleName").String()
	table.BusinessName = s.GetBusinessName(ctx, table.TableName)
	table.FunctionName = strings.ReplaceAll(table.TableComment, "表", "")
	table.FunctionAuthor = g.Cfg().MustGet(ctx, "gen.author").String()
	table.TplCategory = "crud"
	table.CreateTime = gtime.Now()
	table.UpdateTime = table.CreateTime
}

// ConvertClassName 表名转换成类名
func (s *toolsGenTable) ConvertClassName(ctx context.Context, tableName string) string {
	return gstr.CaseCamel(s.removeTablePrefix(ctx, tableName))
}

// GetBusinessName 获取业务名
func (s *toolsGenTable) GetBusinessName(ctx context.Context, tableName string) string {
	return s.removeTablePrefix(ctx, tableName)
}

// 删除表前缀
func (s *toolsGenTable) removeTablePrefix(ctx context.Context, tableName string) string {
	autoRemovePre := g.Cfg().MustGet(ctx, "gen.autoRemovePre").Bool()
	tablePrefix := g.Cfg().MustGet(ctx, "gen.tablePrefix").String()
	if autoRemovePre && tablePrefix != "" {
		searchList := strings.Split(tablePrefix, ",")
		for _, str := range searchList {
			if strings.HasPrefix(tableName, str) {
				tableName = strings.Replace(tableName, str, "", 1) //注意，只替换一次
			}
		}
	}
	return tableName
}

// Delete 删除表信息
func (s *toolsGenTable) Delete(ctx context.Context, ids []int) error {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("开启删除事务出错")
	}
	_, err = tx.Model(dao.GenTable.Table()).Where(dao.GenTable.Columns().TableId+" in(?)", ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		tx.Rollback()
		return gerror.New("删除表格数据失败")
	}
	_, err = tx.Model(dao.GenTableColumn.Table()).Where(dao.GenTableColumn.Columns().TableId+" in(?)", ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		tx.Rollback()
		return gerror.New("删除表格字段数据失败")
	}
	tx.Commit()
	return nil
}

// GetTableInfoByTableId 获取表格信息
func (s *toolsGenTable) GetTableInfoByTableId(ctx context.Context, tableId int64) (info *entity.GenTable, err error) {
	err = dao.GenTable.Ctx(ctx).WherePri(tableId).Scan(&info)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取表格信息出错")
	}
	return
}

// SaveEdit 更新表及字段生成信息
func (s *toolsGenTable) SaveEdit(ctx context.Context, req *gen.ToolsGenTableEditReq) (err error) {
	if req == nil {
		err = gerror.New("参数错误")
		return
	}
	var table *entity.GenTable
	err = dao.GenTable.Ctx(ctx).Where("table_id=?", req.TableId).Scan(&table)
	if err != nil || table == nil {
		err = gerror.New("数据不存在")
		return
	}
	if req.TableName != "" {
		table.TableName = req.TableName
	}
	if req.TableComment != "" {
		table.TableComment = req.TableComment
	}
	if req.BusinessName != "" {
		table.BusinessName = req.BusinessName
	}
	if req.ClassName != "" {
		table.ClassName = req.ClassName
	}
	if req.FunctionAuthor != "" {
		table.FunctionAuthor = req.FunctionAuthor
	}
	if req.FunctionName != "" {
		table.FunctionName = req.FunctionName
	}
	if req.ModuleName != "" {
		table.ModuleName = req.ModuleName
	}
	if req.PackageName != "" {
		table.PackageName = req.PackageName
	}
	if req.Remark != "" {
		table.Remark = req.Remark
	}
	if req.TplCategory != "" {
		table.TplCategory = req.TplCategory
	}
	if req.Params != "" {
		table.Options = req.Params
	}
	table.UpdateTime = gtime.Now()
	var options g.Map
	if req.TplCategory == "tree" {
		//树表设置options
		options = g.Map{
			"treeCode":       req.TreeCode,
			"treeParentCode": req.TreeParentCode,
			"treeName":       req.TreeName,
		}
		table.Options = gconv.String(options)
	} else {
		table.Options = ""
	}

	var tx gdb.TX
	tx, err = g.DB().Begin(ctx)
	if err != nil {
		return
	}
	_, err = tx.Model(dao.GenTable.Table()).Save(table)
	if err != nil {
		tx.Rollback()
		return err
	}

	//保存列数据
	if req.Columns != nil {
		for _, column := range req.Columns {
			if column.ColumnId > 0 {
				var dbColumn *entity.GenTableColumn
				err = dao.GenTableColumn.Ctx(ctx).Where("column_id=?", column.ColumnId).Scan(&dbColumn)
				if dbColumn != nil {
					dbColumn.ColumnComment = column.ColumnComment
					dbColumn.GoType = column.GoType
					dbColumn.HtmlType = column.HtmlType
					dbColumn.HtmlField = column.HtmlField
					dbColumn.QueryType = column.QueryType
					dbColumn.GoField = column.GoField
					dbColumn.DictType = column.DictType
					dbColumn.IsInsert = column.IsInsert
					dbColumn.IsEdit = column.IsEdit
					dbColumn.IsList = column.IsList
					dbColumn.IsQuery = column.IsQuery
					dbColumn.IsRequired = column.IsRequired
					if tc, e := options["treeParentCode"]; options != nil && e && tc != "" && tc == dbColumn.HtmlField {
						dbColumn.IsQuery = "0"
						dbColumn.IsList = "0"
						dbColumn.HtmlType = "select"
					}
					//获取字段关联表信息
					if column.LinkLabelName != "" {
						dbColumn.LinkTableName = column.LinkTableName
						dbColumn.LinkLabelId = column.LinkLabelId
						dbColumn.LinkLabelName = column.LinkLabelName
						var linkTable *entity.GenTable
						err = dao.GenTable.Ctx(ctx).Where("table_name =?", column.LinkTableName).Scan(&linkTable)
						if err != nil {
							tx.Rollback()
							return
						}
						dbColumn.LinkTableClass = linkTable.ClassName
						dbColumn.LinkTablePackage = linkTable.PackageName
					} else {
						dbColumn.LinkTableName = ""
						dbColumn.LinkTableClass = ""
						dbColumn.LinkTablePackage = ""
						dbColumn.LinkLabelId = ""
						dbColumn.LinkLabelName = ""
					}
					_, err = tx.Model(dao.GenTableColumn.Table()).Save(dbColumn)
					if err != nil {
						tx.Rollback()
						return
					}
				}
			}
		}
	}
	tx.Commit()
	return
}

func (s *toolsGenTable) SelectRecordById(ctx context.Context, tableId int64) (entityExtend *gen.ToolsGenTableExtend, err error) {
	var table *entity.GenTable
	table, err = s.GetTableInfoByTableId(ctx, tableId)
	if err != nil {
		return
	}
	m := gconv.Map(table)
	gconv.Struct(m, &entityExtend)
	if entityExtend.TplCategory == "tree" {
		opt := gjson.New(entityExtend.Options)
		entityExtend.TreeParentCode = opt.Get("treeParentCode").String()
		entityExtend.TreeCode = opt.Get("treeCode").String()
		entityExtend.TreeName = opt.Get("treeName").String()
	}
	//表字段数据
	var columns []*entity.GenTableColumn
	columns, err = ToolsGenTableColumn.SelectGenTableColumnListByTableId(ctx, tableId)
	if err != nil {
		return
	}
	entityExtend.Columns = columns
	return
}

func (s *toolsGenTable) GenCode(ctx context.Context, ids []int) (err error) {
	// 判断运行环境是mac还是windows
	var backDirName = "gen.backDirForWindows"
	var frontDirName = "gen.frontDirForWindows"
	if runtime.GOOS != "windows" {
		backDirName = "gen.backDirForMac"
		frontDirName = "gen.frontDirForMac"
	}

	//获取当前运行时目录
	curDir := g.Cfg().MustGet(ctx, backDirName).String()
	if !gfile.IsDir(curDir) {
		err = gerror.New("项目后端路径不存在，请检查是否已在配置文件中配置！")
		return
	}
	frontDir := g.Cfg().MustGet(ctx, frontDirName).String()
	if !gfile.IsDir(frontDir) {
		err = gerror.New("项目前端路径不存在，请检查是否已在配置文件中配置！")
		return
	}
	for _, id := range ids {
		var genData g.MapStrStr
		var extendData *gen.ToolsGenTableExtend
		genData, extendData, err = s.GenData(gconv.Int64(id), ctx)
		idx := gstr.Pos(extendData.PackageName, "/")
		var packageName = extendData.PackageName
		if idx > -1 {
			packageName = gstr.SubStr(extendData.PackageName, idx)
		}
		businessName := gstr.CaseCamelLower(extendData.BusinessName)
		for key, code := range genData {
			switch key {
			case "controller":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/controller/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, false)
			case "dao":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/dao/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, false)
			case "dao_internal":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/dao/internal/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, true)
			case "do":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/dao/do/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, false)
			case "model":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/model/entity/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, true)
			case "router":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/router/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, false)
			case "service":
				path := strings.Join([]string{curDir, "/internal/app/", packageName, "/service/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, false)
			case "apiReq":
				path := strings.Join([]string{curDir, "/api/v1/", packageName, "/", extendData.TableName, ".go"}, "")
				err = s.createFile(path, code, false)
			case "sql":
				path := strings.Join([]string{curDir, "/", "/gen_sql/", packageName, "/", extendData.TableName, ".sql"}, "")
				hasSql := gfile.Exists(path)
				err = s.createFile(path, code, false)
				if !hasSql {
					////第一次生成则向数据库写入菜单数据
					//err = s.writeDb(path, ctx)
					//if err != nil {
					//	return
					//}
					////清除菜单缓存
					//comService.Cache.New().Remove(global.SysAuthMenu)
				}

			case "vue":
				path := strings.Join([]string{frontDir, "/vue/src/views/", extendData.ModuleName, "/", businessName, "/index.vue"}, "")
				if gstr.ContainsI(extendData.PackageName, "plugins") {
					path = strings.Join([]string{frontDir, "/vue/src/views/plugins/", extendData.ModuleName, "/", businessName, "/index.vue"}, "")
				}
				err = s.createFile(path, code, false)
			case "vueEdit":
				path := strings.Join([]string{frontDir, "/vue/src/views/", extendData.ModuleName, "/", businessName, "/component/edit" + strings.ToUpper(businessName[:1]) + businessName[1:] + ".vue"}, "")
				if gstr.ContainsI(extendData.PackageName, "plugins") {
					path = strings.Join([]string{frontDir, "/vue/src/views/plugins/", extendData.ModuleName, "/", businessName, "/component/edit" + strings.ToUpper(businessName[:1]) + businessName[1:] + ".vue"}, "")
				}
				err = s.createFile(path, code, false)
			case "jsApi":
				path := strings.Join([]string{frontDir, "/vue/src/api/", extendData.ModuleName, "/", businessName, "/", "index.ts"}, "")
				if gstr.ContainsI(extendData.PackageName, "plugins") {
					path = strings.Join([]string{frontDir, "/vue/src/api/plugins/", extendData.ModuleName, "/", businessName, "/", "index.ts"}, "")
				}
				err = s.createFile(path, code, false)
			}
		}
		////生成对应的模块路由
		//err = s.genModuleRouter(curDir, extendData.ModuleName, extendData.PackageName)
	}
	return
}

// createFile 创建文件
func (s *toolsGenTable) createFile(fileName, data string, cover bool) (err error) {
	if !gfile.Exists(fileName) || cover {
		var f *os.File
		f, err = gfile.Create(fileName)
		if err == nil {
			f.WriteString(data)
		}
		f.Close()
	}
	return
}

// 写入菜单数据
func (s *toolsGenTable) writeDb(path string, ctx context.Context) (err error) {
	isAnnotation := false
	var fi *os.File
	fi, err = os.Open(path)
	if err != nil {
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var sqlStr []string
	now := gtime.Now()
	var res sql.Result
	var id int64
	var tx gdb.TX
	tx, err = g.DB().Ctx(ctx).Begin(ctx)
	if err != nil {
		return
	}
	for {
		bytes, e := br.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		str := gstr.Trim(string(bytes))

		if str == "" {
			continue
		}

		if strings.Contains(str, "/*") {
			isAnnotation = true
		}

		if isAnnotation {
			if strings.Contains(str, "*/") {
				isAnnotation = false
			}
			continue
		}

		if str == "" || strings.HasPrefix(str, "--") || strings.HasPrefix(str, "#") {
			continue
		}
		if strings.HasSuffix(str, ";") {
			if gstr.ContainsI(str, "select") {
				if gstr.ContainsI(str, "@now") {
					continue
				}
				if gstr.ContainsI(str, "@parentId") {
					id, err = res.LastInsertId()
				}
			}
			sqlStr = append(sqlStr, str)
			sql := strings.Join(sqlStr, "")
			gstr.ReplaceByArray(sql, []string{"@parentId", gconv.String(id), "@now", now.Format("Y-m-d H:i:s")})
			//插入业务
			res, err = tx.Exec(sql)
			if err != nil {
				tx.Rollback()
				return
			}
			sqlStr = nil
		} else {
			sqlStr = []string{str}
		}
	}
	tx.Commit()
	return
}

// GenData 获取生成数据
func (s *toolsGenTable) GenData(tableId int64, ctx context.Context) (data g.MapStrStr, extendData *gen.ToolsGenTableExtend, err error) {
	extendData, err = ToolsGenTable.SelectRecordById(ctx, tableId)
	if err != nil {
		return
	}
	if extendData == nil {
		err = gerror.New("表格数据不存在")
		return
	}
	ToolsGenTableColumn.SetPkColumn(extendData, extendData.Columns)
	view := gview.New()
	view.SetConfigWithMap(g.Map{
		"Paths":      g.Cfg().MustGet(ctx, "gen.templatePath").String(),
		"Delimiters": []string{"{{", "}}"},
	})
	view.BindFuncMap(g.Map{
		"UcFirst": func(str string) string {
			return gstr.UcFirst(str)
		},
		"Sum": func(a, b int) int {
			return a + b
		},
		"CaseCamelLower": gstr.CaseCamelLower, //首字母小写驼峰
		"CaseCamel":      gstr.CaseCamel,      //首字母大写驼峰
		"HasSuffix":      gstr.HasSuffix,      //是否存在后缀
		"ContainsI":      gstr.ContainsI,      //是否包含子字符串
		"VueTag": func(t string) string {
			return t
		},
	})

	//树形菜单选项
	tplData := g.Map{"table": extendData}
	daoKey := "dao"
	daoValue := ""
	var tmpDao string
	if tmpDao, err = view.Parse(ctx, "go/dao.template", tplData); err == nil {
		daoValue = tmpDao
		daoValue, err = s.trimBreak(daoValue)
	} else {
		return
	}
	daoInternalKey := "dao_internal"
	daoInternalValue := ""
	var tmpInternalDao string
	if tmpInternalDao, err = view.Parse(ctx, "go/dao_internal.template", tplData); err == nil {
		daoInternalValue = tmpInternalDao
		daoInternalValue, err = s.trimBreak(daoInternalValue)
	} else {
		return
	}
	modelKey := "model"
	modelValue := ""
	var tmpModel string
	if tmpModel, err = view.Parse(ctx, "go/model.template", tplData); err == nil {
		modelValue = tmpModel
		modelValue, err = s.trimBreak(modelValue)
	} else {
		return
	}
	controllerKey := "controller"
	controllerValue := ""
	var tmpController string
	if tmpController, err = view.Parse(ctx, "go/controller.template", tplData); err == nil {
		controllerValue = tmpController
		controllerValue, err = s.trimBreak(controllerValue)
	} else {
		return
	}

	serviceKey := "service"
	serviceValue := ""
	var tmpService string
	if tmpService, err = view.Parse(ctx, "go/service.template", tplData); err == nil {
		serviceValue = tmpService
		serviceValue, err = s.trimBreak(serviceValue)
	} else {
		return
	}

	apiReqKey := "apiReq"
	apiReqValue := ""
	var tmpApiReq string
	if tmpApiReq, err = view.Parse(ctx, "go/apiReq.template", tplData); err == nil {
		apiReqValue = tmpApiReq
		apiReqValue, err = s.trimBreak(apiReqValue)
	} else {
		return
	}

	doKey := "do"
	doValue := ""
	var tmpDoReq string
	if tmpDoReq, err = view.Parse(ctx, "go/do.template", tplData); err == nil {
		doValue = tmpDoReq
		doValue, err = s.trimBreak(doValue)
	} else {
		return
	}

	//routerKey := "router"
	//routerValue := ""
	//var tmpRouter string
	//if tmpRouter, err = view.Parse(ctx, "go/router.template", tplData); err == nil {
	//	routerValue = tmpRouter
	//	routerValue, err = s.trimBreak(routerValue)
	//} else {
	//	return
	//}

	sqlKey := "sql"
	sqlValue := ""
	var tmpSql string
	if tmpSql, err = view.Parse(ctx, "sql/sql.template", tplData); err == nil {
		sqlValue = tmpSql
		sqlValue, err = s.trimBreak(sqlValue)
	} else {
		return
	}

	jsApiKey := "jsApi"
	jsApiValue := ""
	var tmpJsApi string
	if tmpJsApi, err = view.Parse(ctx, "js/api.template", tplData); err == nil {
		jsApiValue = tmpJsApi
		jsApiValue, err = s.trimBreak(jsApiValue)
	} else {
		return
	}

	vueKey := "vue"
	vueValue := ""
	var tmpVue string
	tmpFile := "vue/list.template"
	if extendData.TplCategory == "tree" {
		//树表
		tmpFile = "vue/tree-vue.template"
	}
	if tmpVue, err = view.Parse(ctx, tmpFile, tplData); err == nil {
		vueValue = tmpVue
		vueValue, err = s.trimBreak(vueValue)
	} else {
		return
	}

	editKey := "vueEdit"
	editValue := ""
	var tmpEdit string
	var tmpEditFile = "vue/edit.template"
	if tmpEdit, err = view.Parse(ctx, tmpEditFile, tplData); err == nil {
		editValue = tmpEdit
		editValue, err = s.trimBreak(editValue)
	} else {
		return
	}

	data = g.MapStrStr{
		daoKey:         daoValue,
		daoInternalKey: daoInternalValue,
		modelKey:       modelValue,
		controllerKey:  controllerValue,
		serviceKey:     serviceValue,
		apiReqKey:      apiReqValue,
		doKey:          doValue,
		//routerKey:      routerValue,
		sqlKey:   sqlValue,
		jsApiKey: jsApiValue,
		vueKey:   vueValue,
		editKey:  editValue,
	}
	return
}

// GenModuleRouter 生成模块路由
func (s *toolsGenTable) genModuleRouter(curDir, moduleName, packageName string) (err error) {
	if gstr.CaseSnake(moduleName) != "system" {
		routerFilePath := strings.Join([]string{curDir, "/router/", gstr.CaseSnake(moduleName), ".go"}, "")
		if gstr.ContainsI(packageName, "plugins") {
			routerFilePath = strings.Join([]string{curDir, "/plugins/router/", gstr.CaseSnake(moduleName), ".go"}, "")
		}
		code := fmt.Sprintf(`package router%simport _ "%s/router"`, "\n", packageName)
		err = s.createFile(routerFilePath, code, false)
	}
	return
}

// 剔除多余的换行
func (s *toolsGenTable) trimBreak(str string) (rStr string, err error) {
	var b []byte
	if b, err = gregex.Replace("(([\\s\t]*)\r?\n){2,}", []byte("$2\n"), []byte(str)); err != nil {
		return
	}
	if b, err = gregex.Replace("(([\\s\t]*)/{4}\r?\n)", []byte("$2\n\n"), b); err == nil {
		rStr = gconv.String(b)
	}
	return
}
