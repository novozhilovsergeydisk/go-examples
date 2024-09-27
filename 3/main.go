package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    // Подключение к базе данных
    connStr := "user=postgres password=postgres@12345 dbname=breeze sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v\n", err)
    }
    defer db.Close()

    // Проверка подключения
    err = db.Ping()
    if err != nil {
        log.Fatalf("Не удалось установить соединение с базой данных: %v\n", err)
    }
    fmt.Println("Подключение к базе данных установлено")

    // Пример запроса SELECT
    querySelect(db)

    // Пример запроса INSERT
    // queryInsert(db)
}

// Функция для выполнения SELECT-запроса
func querySelect(db *sql.DB) {
    rows, err := db.Query("SELECT id, user_agent FROM sessions")
    if err != nil {
        log.Fatalf("Ошибка выполнения SELECT-запроса: %v\n", err)
    }
    defer rows.Close()

    fmt.Println("Результаты SELECT-запроса:")
    for rows.Next() {
        var id string
        var user_agent string
        err := rows.Scan(&id, &user_agent)
        if err != nil {
            log.Fatalf("Ошибка чтения строки результата: %v\n", err)
        }
        fmt.Printf("ID: %d, Имя: %s\n", id, user_agent)
    }

    // Проверка на ошибки после обработки всех строк
    if err = rows.Err(); err != nil {
        log.Fatalf("Ошибка обработки строк: %v\n", err)
    }
}

// Функция для выполнения INSERT-запроса
func queryInsert(db *sql.DB) {
    // Пример добавления нового пользователя
    var name = "John Doe"
    _, err := db.Exec("INSERT INTO users (name) VALUES ($1)", name)
    if err != nil {
        log.Fatalf("Ошибка выполнения INSERT-запроса: %v\n", err)
    }
    fmt.Printf("Пользователь %s успешно добавлен\n", name)
}
