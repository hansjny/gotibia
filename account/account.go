package account

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var G_DB *sql.DB;

type Character struct {
	Name string
	WorldId int
}

type Account struct {
	Uid int
	Prem int
	Chars []Character
}

func RequestAccount(accnum string, pwd string) *Account {
	var id int
	var prem int
	err := G_DB.QueryRow("SELECT id, prem FROM accounts WHERE accnum=? AND password=?", accnum, pwd).Scan(&id, &prem)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		fmt.Println(err.Error())
		return nil
	default:
		a := Account{Uid: id, Prem: prem }
		a.getCharacters()
		return &a
	}

}

func (a *Account) getCharacters() {
	fmt.Println("Getting character for acc: ",  a.Uid)
	rows,  err := G_DB.Query("SELECT name, world_id FROM characters WHERE account_id=?", a.Uid)

	if err != nil {
		fmt.Println("Something went wrong,  getCharacters()")
		return 
	}

	defer rows.Close()

	for rows.Next() {
		var charname string
		var world int
		err := rows.Scan(&charname,  &world)
		if err != nil {
			fmt.Println("Something went wrong,  getCharacters()")
			break
		}
		char := Character{Name: charname,  WorldId: world}
		a.Chars = append(a.Chars,  char)
	}
}
