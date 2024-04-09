package global

import (
	"gorm.io/gorm"
	"sync"
)

type TableSlave struct {
	Mutex     *sync.Mutex
	CurrIndex int
	Slave     []*gorm.DB
}

func (t *TableSlave) ChooseSlave() *gorm.DB {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	db := t.Slave[t.CurrIndex]
	t.CurrIndex = (t.CurrIndex + 1) % len(t.Slave)
	return db
}

var (
	//主库
	MysqlClientMaster *gorm.DB
	MysqlClientSlave0 *gorm.DB
	MysqlClientSlave1 *gorm.DB

	UserTableMaster *gorm.DB
	UserTableSlave  TableSlave
	UserTableSlave0 *gorm.DB
	UserTableSlave1 *gorm.DB

	RoleTableMaster *gorm.DB
	RoleTableSlave  TableSlave
	RoleTableSlave0 *gorm.DB
	RoleTableSlave1 *gorm.DB

	ApiTableMaster *gorm.DB
	ApiTableSlave  TableSlave
	ApiTableSlave0 *gorm.DB
	ApiTableSlave1 *gorm.DB

	UserRoleTableMaster *gorm.DB
	UserRoleTableSlave  TableSlave
	UserRoleTableSlave0 *gorm.DB
	UserRoleTableSlave1 *gorm.DB

	RoleApiTableMaster *gorm.DB
	RoleApiTableSlave  TableSlave
	RoleApiTableSlave0 *gorm.DB
	RoleApiTableSlave1 *gorm.DB

	LogTableMaster *gorm.DB
	LogTableSlave  TableSlave
	LogTableSlave0 *gorm.DB
	LogTableSlave1 *gorm.DB

	CommentTableMaster *gorm.DB
	CommentTableSlave  TableSlave
	CommentTableSlave0 *gorm.DB
	CommentTableSlave1 *gorm.DB

	MomentTableMaster *gorm.DB
	MomentTableSlave  TableSlave
	MomentTableSlave0 *gorm.DB
	MomentTableSlave1 *gorm.DB
)
