package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Эта функция устанавливает соединение с PostgreSQL и создает в базе данных
// таблицу с именем users. Эта таблица будет использоваться исключительно
// для тестирования.
func create_table() {
	connStr := "user=postgres password=pass dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	const query = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name TEXT,
			last_name TEXT
	)`

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Close()
}

// удаляет таблицу, созданнуюфункцией create_table()
func drop_table() {
	connStr := "user=postgres password=pass dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Close()
}

// вставляет запись в таблицу PostgreSQL
func insert_record(query string) {
	connStr := "user=postgres password=pass dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Close()
}

// Это первая тестовая функция пакета, и она работает в два этапа: сначала встав-
// ляет три записи в таблицу базы данных; затем проверяет, что эта таблица базы
// данных содержит ровно три записи.
func Test_count(t *testing.T) {
	var count int
	create_table()

	insert_record("INSERT INTO users (first_name, last_name) VALUES ('Epifanios', 'Doe')")
	insert_record("INSERT INTO users (first_name, last_name) VALUES ('Pavel', 'Gazukin')")
	insert_record("INSERT INTO users (first_name, last_name) VALUES ('Pavel', 'Unknown')")

	connStr := "user=postgres password=pass dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	row := db.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&count)
	db.Close()

	if count != 3 {
		t.Errorf("Select query returned %d", count)
	}

	drop_table()
}

// Это еще одна тестовая функция, которая вставляет запись в таблицу базы дан-
// ных и проверяет, что данные были записаны правильно.
func Test_queryDB(t *testing.T) {
	create_table()

	connStr := "user=postgres password=pass dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	query := "INSERT INTO users (first_name, last_name) VALUES ('Random Text', '123456')"
	insert_record(query)

	rows, err := db.Query(`SELECT * FROM users WHERE last_name=$1`, `123456`)
	if err != nil {
		fmt.Println(err)
		return
	}
	var col1 int
	var col2 string
	var col3 string
	for rows.Next() {
		rows.Scan(&col1, &col2, &col3)
	}

	if col2 != "Random Text" {
		t.Errorf("first_name returned %s", col2)
	}
	if col3 != "123456" {
		t.Errorf("last_name returned %s", col3)
	}

	db.Close()
	drop_table()
}

// Функция пакета взаимодействует с веб-сервером и посещает
// URL-адрес /getdata. Затем она проверяет, соответствует ли возвращаемое
// значение тому, что ожидается.
func Test_record(t *testing.T) {
	create_table()
	insert_record("INSERT INTO users (first_name, last_name) VALUES ('John', 'Doe')")

	req, err := http.NewRequest("GET", "/getdata", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned %v", status)
	}

	if rr.Body.String() != "<h3 align=\"center\">1, John, Doe</h3>\n" {
		t.Errorf("Wrong server response")
	}
	drop_table()
}
