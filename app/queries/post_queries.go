package queries

import (
	"database/sql"
	"fmt"
	"github.com/mozillazg/go-slugify"
	"time"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type PostQueries struct {
	*sql.DB
}

func (q *PostQueries) GetPosts() ([]models.Post, error) {
	posts := []models.Post{}

	query := "select * from posts"

	rows, err := q.Query(query)
	if err != nil {
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		post := models.Post{}

		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.Views, &post.Slug, &post.Published, &post.PublishedAt, &post.UpdatedAt, &post.UserId); err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (q *PostQueries) CreatePost(post *models.NewPost) error {
	query := "insert into posts (user_id, title, body, slug, views, published, published_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8)"

	slug := slugify.Slugify(post.Title)

	_, err := q.Exec(query, &post.UserId, &post.Title, &post.Body, slug, 0, true, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
