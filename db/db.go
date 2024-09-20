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
	// 首先检查条目是否已存在
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM rss_items WHERE guid = ?)", item.GUID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		// 条目已存在，不需要插入
		return nil
	}

	// 条目不存在，执行插入操作
	_, err = DB.Exec(`
		INSERT INTO rss_items (title, link, description, author, pub_date, enclosure_url, enclosure_length, enclosure_type, guid, category)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, item.Title, item.Link, item.Disc, item.Author, formattedPubDate, item.Enclosure.URL, item.Enclosure.Length, item.Enclosure.Type, item.GUID, joinCategories(item.Category))
	return err
}

func joinCategories(categories []string) string {
	return strings.Join(categories, ",")
}

func GetLatestPubDate() (time.Time, error) {
	var latestPubDate []byte
	err := DB.QueryRow("SELECT MAX(pub_date) FROM rss_items").Scan(&latestPubDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, nil
		}
		return time.Time{}, err
	}

	if latestPubDate == nil {
		return time.Time{}, nil
	}

	// 将 []byte 转换为字符串，然后解析为 time.Time
	return time.Parse("2006-01-02 15:04:05", string(latestPubDate))
}
