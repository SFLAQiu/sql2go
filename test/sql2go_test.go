package test

import (
	"database/sql"
	"github.com/SFLAQiu/sql2go"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestFromFile1(t *testing.T) {
	args := sql2go.NewConvertArgs().SetGenJson(true)
	code, err := sql2go.FromFile("./1.sql", args)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(code))
}
func TestFromFile2(t *testing.T) {
	args := sql2go.NewConvertArgs().
		SetGenXorm(true).
		SetColPrefix("f_").
		SetTablePrefix("t_")

	code, err := sql2go.FromFile("./2.sql", args)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(code))
}

func TestFromSql1(t *testing.T) {
	sql := `
CREATE TABLE IF NOT EXISTS t_person (
  f_age INT(11) NULL,
  f_id INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  f_name VARCHAR(30) NOT NULL,
  f_sex VARCHAR(2) NULL,
  f_test TEXT
  ) ENGINE=InnoDB;
`
	args := sql2go.NewConvertArgs().
		SetGenJson(true).
		SetGenXorm(true).
		SetColPrefix("f_").
		SetTablePrefix("t_").
		SetOtherTags("db,json xlsx")

	code, err := sql2go.FromSql(sql, args)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(code))
}

func TestFromSql2(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/Demo?charset=utf8")
	rows, _ := db.Query("show create table student")
	var tableName string
	var createTableSql string
	for rows.Next() {
		err = rows.Scan(&tableName, &createTableSql)
	}
	args := sql2go.NewConvertArgs().
		SetGenJson(true).
		SetGenGorm(true)

	code, err := sql2go.FromSql(createTableSql, args)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(code))
}
