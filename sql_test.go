package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('amin','Amin','amin77@gmail.com',200000,7.0,'1998-12-06', true)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlSelect(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)

		if err != nil {
			panic(err)
		}
		fmt.Println("======================================================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birth_date.Valid {
			fmt.Println("Birth_date:", birth_date)
		}
		fmt.Println("Married:", married)
		fmt.Println("created_at:", created_at)
		fmt.Println("======================================================")

	}
	defer rows.Close()

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; # "
	password := " or 1=1--+"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND Password = '" + password + "'Limit 1"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
