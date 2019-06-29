package database

import (
	"database/sql"
)

type WikuPosts struct {
	Id int64 `json:"id" db:"ID"`
	PostContent string `json:"postContent" db:"post_content"`
	PostTitle string `json:"postTitle" db:"post_title"`
	PostStatus string `json:"postStatus" db:"post_status"`
	PostName string `json:"postName" db:"post_name"`
	PostDate sql.NullString `json:"postDate" db:"post_date"`
}