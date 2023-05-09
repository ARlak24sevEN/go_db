package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type cover struct {
	Id   int
	Name string
}

// var db *sql.DB
var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/go_db")

	if err != nil {
		panic(err)
	}

	// cover := cover{8, "cover_arlak"}
	// err = AddCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// cover := cover{8, "cover_nadia"}
	// err = UpdateCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// err = DeleteCover(8)
	// if err != nil {
	// 	panic(err)
	// }

	// covers, err := GetCovers()
	// covers, err := GetCoversX()
	// if err != nil {
	// 	fmt.Println("error : ", err)
	// 	return
	// }
	// for _, cover := range covers {
	// 	fmt.Printf("cover %v\n", cover)
	// }
	// cover, err := GerCover(1)
	cover, err := GetCoverX(1)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", cover)

}
func GetCoversX() ([]cover, error) {
	query := "select id , name from cover"
	covers := []cover{}
	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}
	return covers, nil

}
func GetCovers() ([]cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	query := "select id , name  from cover"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	covers := []cover{}
	for rows.Next() {
		cover := cover{}
		err := rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		// fmt.Println("id : ", cover.Id, " cover : ", cover.Name)
		covers = append(covers, cover)
	}
	return covers, err
}

func GetCoverX(id int) (*cover, error) {
	query := "select id , name from cover where id = ?"
	covers := cover{}
	err := db.Get(&covers, query, id)
	if err != nil {
		return nil, err
	}
	return &covers, nil

}

func GerCover(id int) (*cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	query := "select id , name  from cover where id = ?" //mysql
	// query := "select id , name  from cover where id = @id"  ms sql

	row := db.QueryRow(query, id) // my sql
	// row := db.QueryRow(query, sql.Named("id", id)) mssql
	cover := cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, err
}

func AddCover(cover cover) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "insert into cover (id,name) values(?,?)"

	result, err := tx.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affect <= 0 {
		return errors.New("cannot insert")
	}
	if tx.Commit(); err != nil {
		return err
	}
	return nil
}

func UpdateCover(cover cover) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	query := "update cover set name = ? where id =?"

	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("cannot update")
	}
	return nil
}

func DeleteCover(id int) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	query := "delete from  cover where id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("cannot delete")
	}
	return nil
}
