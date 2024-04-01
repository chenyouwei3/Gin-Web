package global

import "gorm.io/gorm"

var (
	MysqlClientMaster   *gorm.DB
	UserTableMaster     *gorm.DB
	RoleTableMaster     *gorm.DB
	ApiTableMaster      *gorm.DB
	UserRoleTableMaster *gorm.DB
	RoleApiTableMaster  *gorm.DB
	LogTableMaster      *gorm.DB
)

var (
	MysqlClientSlave0   *gorm.DB
	UserTableSlave0     *gorm.DB
	RoleTableSlave0     *gorm.DB
	ApiTableSlave0      *gorm.DB
	UserRoleTableSlave0 *gorm.DB
	RoleApiTableSlave0  *gorm.DB
	LogTableSlave0      *gorm.DB
)
