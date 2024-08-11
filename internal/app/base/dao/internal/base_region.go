// 功能：行政区域省市区县 internal
package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type RegionDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns RegionColumns // columns contains all the column names of Table for convenient usage.
}
type RegionColumns struct {
	Id               string //
	ParentId         string //父ID
	RegionName       string //名称
	MergerName       string //全称
	ShortName        string //简称
	MergerShortName  string //简称合并
	Level            string //层级，1是省份，2是城市，3是区县
	CityCode         string //城市代码
	ZipCode          string //邮编号码
	FullPinyin       string //全拼
	SimplifiedPinyin string //简拼
	FirstChar        string //第一个字
	Longitude        string //纬度
	Latitude         string //经度
}

// regionColumns holds the columns for table base_region.
var regionColumns = RegionColumns{
	Id:               "id",
	ParentId:         "parent_id",
	RegionName:       "region_name",
	MergerName:       "merger_name",
	ShortName:        "short_name",
	MergerShortName:  "merger_short_name",
	Level:            "level",
	CityCode:         "city_code",
	ZipCode:          "zip_code",
	FullPinyin:       "full_pinyin",
	SimplifiedPinyin: "simplified_pinyin",
	FirstChar:        "first_char",
	Longitude:        "longitude",
	Latitude:         "latitude",
}

// NewRegionDao creates and returns a new DAO object for table data access.
func NewRegionDao() *RegionDao {
	return &RegionDao{
		group:   "default",
		table:   "base_region",
		columns: regionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RegionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RegionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RegionDao) Columns() RegionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RegionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RegionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RegionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
