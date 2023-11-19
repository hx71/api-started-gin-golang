package migration

import (
	"fmt"
	"log"
)

func RunRollback() {
	if db_conn.Error != nil {
		log.Fatalln(db_conn.Error.Error())
	}

	if exist := db_conn.Migrator().HasTable("roles"); exist {
		err := db_conn.Migrator().DropTable("roles")
		if err == nil {
			fmt.Println("success drop table roles")
		}
	}

	if exist := db_conn.Migrator().HasTable("menus"); exist {
		err := db_conn.Migrator().DropTable("menus")
		if err == nil {
			fmt.Println("success drop table menus")
		}
	}

	if exist := db_conn.Migrator().HasTable("user_menus"); exist {
		err := db_conn.Migrator().DropTable("user_menus")
		if err == nil {
			fmt.Println("success drop table user_menus")
		}
	}

	if exist := db_conn.Migrator().HasTable("users"); exist {
		err := db_conn.Migrator().DropTable("users")
		if err == nil {
			fmt.Println("success drop table users")
		}
	}
}
