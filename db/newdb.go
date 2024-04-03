package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
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
	// db.Query("")
	fmt.Printf("Connected to DB!")
}

func InsertChatLogEntry() (int, error) {
	result, err := db.Exec("INSERT INTO [dbo].[ChatLog] (firstMessageTime) VALUES (" + strconv.Itoa(int(time.Now().Unix())) + ");")
	if err != nil {
		log.Print(err)
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Print(err)
		return 0, err
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
		return 0, errors.New("Expected 1 row to be affected")
	}
	log.Println("According to all known laws in aviation")
	var ID int
	row, err := db.Query("SELECT MAX(ChatID) FROM [dbo].[ChatLog]")
	for row.Next() {
		err = row.Scan(&ID)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		fmt.Println(ID)
		err = row.Err()
		if err != nil {
			log.Println(err)
		}
		if err != nil {
			log.Print(err)
			return 0, err
		}
	}
	return ID, err
}

func InsertBotConversation(ChatID int, Message string, IsBot bool) error {
	if IsBot {
		result, err := db.Exec(fmt.Sprintf("INSERT INTO [dbo].[ChatEntries] (chatID, entryTime, interaction, isBot) VALUES (%d, %d, '%s', 1)", ChatID, time.Now().Unix(), Message))
		if err != nil {
			log.Print(err)
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			log.Print(err)
			return err
		}
		if rows != 1 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
			return errors.New("Expected 1 row to be affected")
		}
		log.Println("Uploaded to DB")
		return err
	} else {
		result, err := db.Exec(fmt.Sprintf("INSERT INTO [dbo].[ChatEntries] (chatID, entryTime, interaction, isBot) VALUES (%d, %d, '%s', 0)", ChatID, time.Now().Unix(), Message))
		if err != nil {
			log.Print(err)
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			log.Print(err)
			return err
		}
		if rows != 1 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
			return errors.New("Expected 1 row to be affected")
		}
		log.Println("Uploaded to DB")
		return err
	}
}
