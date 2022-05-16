package repository

import "time"

type Topic struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
}

type Post struct {
	ID          int64     `json:"id"`
	TopicID     int64     `json:"topic_id"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
}
