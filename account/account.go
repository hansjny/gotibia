package account

import (
	"fmt"
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	accnum string;
	password string;
	uid string;
}


func RequestAccount(accnum string) {

	fmt.Println("Requestin account: ", accnum);

}
