package account

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var G_DB *sql.DB;

type Account struct {
	accnum string;
	password string;
	uid string;
}


func RequestAccount(accnum string, pwd string) int  {
	var id int
	err := G_DB.QueryRow("SELECT id FROM accounts WHERE accnum=? AND password=?", accnum, pwd).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return 0
	case err != nil:
		fmt.Println(err.Error())
		return 0
	default:
		return id
	}

}
