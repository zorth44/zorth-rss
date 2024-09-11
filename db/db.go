package db

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zorth/zorth-rss/model"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "user:passwd@tcp(127.0.0.1:3306)/rss_db")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS rss_items (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			link TEXT NOT NULL,
			description TEXT,
			author VARCHAR(255),
			pub_date DATETIME NOT NULL,
			enclosure_url TEXT,
			enclosure_length VARCHAR(255),
			enclosure_type VARCHAR(255),
			guid VARCHAR(255) UNIQUE,
			category TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertRSSItem(item model.RssItemDMHY, formattedPubDate string) error {
	_, err := DB.Exec(`
		INSERT INTO rss_items (title, link, description, author, pub_date, enclosure_url, enclosure_length, enclosure_type, guid, category)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, item.Title, item.Link, item.Disc, item.Author, formattedPubDate, item.Enclosure.URL, item.Enclosure.Length, item.Enclosure.Type, item.GUID, joinCategories(item.Category))
	return err
}

func joinCategories(categories []string) string {
	return strings.Join(categories, ",")
}

func GetLatestPubDate() (time.Time, error) {
	var latestPubDate sql.NullTime
	err := DB.QueryRow("SELECT MAX(pub_date) FROM rss_items").Scan(&latestPubDate)
	if err != nil {
		if err == sql.ErrNoRows {
			// 如果没有记录，返回零值时间和 sql.ErrNoRows 错误
			return time.Time{}, sql.ErrNoRows
		}
		return time.Time{}, err
	}

	if !latestPubDate.Valid {
		// 如果 MAX(pub_date) 返回 NULL，返回零值时间
		return time.Time{}, nil
	}

	return latestPubDate.Time, nil
}
