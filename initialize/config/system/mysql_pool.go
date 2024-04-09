package system

import (
	"gorm.io/gorm"
	"loopy-manager/initialize/global"
	"sync"
)

func MysqlPoolInit() {
	global.UserTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.UserTableSlave0,
			global.UserTableSlave1,
		},
	}

	global.RoleTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.RoleTableSlave0,
			global.RoleTableSlave1,
		},
	}

	global.ApiTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.ApiTableSlave0,
			global.ApiTableSlave1,
		},
	}

	global.UserRoleTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.UserRoleTableSlave0,
			global.UserRoleTableSlave1,
		},
	}

	global.RoleApiTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.RoleApiTableSlave0,
			global.RoleApiTableSlave1,
		},
	}

	global.LogTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.LogTableSlave0,
			global.LogTableSlave1,
		},
	}

	global.CommentTableSlave = global.TableSlave{Mutex: new(sync.Mutex), CurrIndex: 0,
		Slave: []*gorm.DB{
			global.CommentTableSlave0,
			global.CommentTableSlave1,
		},
	}
}
