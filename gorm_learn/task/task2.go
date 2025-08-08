package task

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func connectDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}
	return db
}

func queryTechEmployees(db *sqlx.DB) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatalln("查询失败:", err)
	}

	for _, e := range employees {
		fmt.Printf("%+v\n", e)
	}
}
func queryTopSalaryEmployee(db *sqlx.DB) {
	var e Employee
	err := db.Get(&e, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatalln("查询失败:", err)
	}
	fmt.Println("工资最高的员工:", e)
}
func queryExpensiveBooks(db *sqlx.DB) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatalln("查询失败:", err)
	}

	fmt.Println("价格 > 50 的书籍:")
	for _, b := range books {
		fmt.Printf("%+v\n", b)
	}
}
func initTables(db *sqlx.DB) {
	db.MustExec(`
        CREATE TABLE IF NOT EXISTS employees (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            department TEXT,
            salary REAL
        );
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            author TEXT,
            price REAL
        );
    `)
}
func Task2() {
	db := connectDB()
	defer db.Close()

	// 测试 employees 表查询
	queryTechEmployees(db)
	queryTopSalaryEmployee(db)

	// 测试 books 表查询
	queryExpensiveBooks(db)
}
