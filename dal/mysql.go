package dal

import (
	"database/sql"
	"log"

	"github.com/touchps/hackernews/types"
)

type MySQLDAL struct {
	SqlDB *sql.DB
}

func (mysqldal MySQLDAL) SaveLink(link types.Link) (int64, error) {
	stmt, err := mysqldal.SqlDB.Prepare("INSERT INTO Links(Title, Address) VALUES(?,?)")
	if err != nil {
		log.Println("dal: db statement prepare failure")
		return -1, err
	}
	defer stmt.Close()
	// log.Printf("%+v\n", link)
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Println("dal: db statement execution failure")
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("dal: failed to get ID")
	}

	log.Println("Row inserted")

	return id, nil
}

func (mysqldal MySQLDAL) GetLinks() ([]types.Link, error) {
	stmt, err := mysqldal.SqlDB.Prepare("SELECT id, title, address FROM Links")
	if err != nil {
		log.Println("dal: db statement prepare failure")
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []types.Link
	for rows.Next() {
		var link types.Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
