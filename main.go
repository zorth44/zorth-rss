package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zorth/zorth-rss/db"
	"github.com/zorth/zorth-rss/feed"
	"github.com/zorth/zorth-rss/model"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		updateRSS()
	}
}

func updateRSS() {
	var rssFeed model.RSSFeedDMHY
	var err error

	if os.Getenv("ENVIRONMENT") == "dev" {
		rssFeed, err = feed.UseVpnUrlToRSSFeedDMHY("https://share.dmhy.org/topics/rss/rss.xml")
	} else {
		rssFeed, err = feed.UseVpnUrlToRSSFeedDMHY("https://share.dmhy.org/topics/rss/rss.xml")
	}

	if err != nil {
		log.Println("Error fetching RSS feed:", err)
		return
	}

	_, err = db.GetLatestPubDate()
	if err != nil {
		// 如果是因为没有记录导致的错误，我们可以继续处理
		if err == sql.ErrNoRows {
			// No need to set latestPubDate if it's not used
			// latestPubDate = time.Time{} // 使用零值时间
		} else {
			log.Println("Error getting latest pub date:", err)
			return
		}
	}

	for _, item := range rssFeed.Channel.Item {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("Error parsing pub date:", err)
			continue
		}

		// 将 pubDate 转换为 MySQL 兼容的格式
		formattedPubDate := pubDate.Format("2006-01-02 15:04:05")

		// 使用格式化后的日期插入数据库
		err = db.InsertRSSItem(item, formattedPubDate)
		if err != nil {
			log.Println("Error inserting RSS item:", err)
		} else {
			fmt.Printf("Inserted new item: %s\n", item.Title)
		}
	}
}
