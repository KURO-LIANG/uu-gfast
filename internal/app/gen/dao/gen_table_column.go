// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"uu-gfast/internal/app/gen/dao/internal"
)

// internalGenTableColumnDao is internal type for wrapping internal DAO implements.
type internalGenTableColumnDao = *internal.GenTableColumnDao

// genTableColumnDao is the data access object for table tools_gen_table_column.
// You can define custom methods on it to extend its functionality as you wish.
type genTableColumnDao struct {
	internalGenTableColumnDao
}

var (
	// GenTableColumn is globally public accessible object for table tools_gen_table_column operations.
	GenTableColumn = genTableColumnDao{
		internal.NewGenTableColumnDao(),
	}
)

// Fill with you ideas below.
