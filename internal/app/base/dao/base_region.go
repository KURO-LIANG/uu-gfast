// 功能：行政区域省市区县 dao
package dao

import (
	"uu-gfast/internal/app/base/dao/internal"
)

// internalRegionDao is internal type for wrapping internal DAO implements.
type internalRegionDao = *internal.RegionDao

// regionDao is the data access object for table base_region.
// You can define custom methods on it to extend its functionality as you wish.
type regionDao struct {
	internalRegionDao
}

var (
	// Region is globally public accessible object for table base_region operations.
	Region = regionDao{
		internal.NewRegionDao(),
	}
)

// Fill with you ideas below.
