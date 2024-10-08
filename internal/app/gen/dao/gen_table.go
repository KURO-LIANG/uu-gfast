// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"uu-gfast/internal/app/gen/dao/internal"
)

// internalGenTableDao is internal type for wrapping internal DAO implements.
type internalGenTableDao = *internal.GenTableDao

// genTableDao is the data access object for table tools_gen_table.
// You can define custom methods on it to extend its functionality as you wish.
type genTableDao struct {
	internalGenTableDao
}

var (
	// GenTable is globally public accessible object for table tools_gen_table operations.
	GenTable = genTableDao{
		internal.NewGenTableDao(),
	}
)

// Fill with you ideas below.
