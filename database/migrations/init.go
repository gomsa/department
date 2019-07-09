package migrations

import (
	departmentPB "github.com/gomsa/department/proto/department"
	db "github.com/gomsa/department/providers/database"
)

func init() {
	department()
}

// department 权限数据迁移
func department() {
	department := &departmentPB.Department{}
	if !db.DB.HasTable(&department) {
		db.DB.Exec(`
			CREATE TABLE departments (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			parent int(11) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			phone varchar(128) DEFAULT NULL,
			manages varchar(255) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}
