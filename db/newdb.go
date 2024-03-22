package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB
var server = "xln-xephyer.database.windows.net"
var port = 1433
var user = "XLNGroupBDB"
var password = "Xephyer2024@cantor2020"
var database = "XLNGroup"

func init() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Query("")
	fmt.Printf("Connected to DB!")
}

func InsertChatLogEntry() (int, error) {
	result, err := db.ExecContext(context.Background(), "INSERT INTO [dbo].[ChatLog] (firstMessageTime) VALUES (?); SELECT @@IDENTITY", int32(time.Now().Unix()))
	if err != nil {
		log.Print(err)
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Print(err)
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return 0, err
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
		return 0, errors.New("Expected 1 row to be affected")
	}
	return int(lastInsertID), err
}

func InsertBotConversation(BotSaid string) {
	db.Exec("INSERT INTO [dbo].[ChatEntries] (chatID, entryTime, interaction, isBot) VALUES ({chatID}, {time}, {interaction}, 1))")
}
