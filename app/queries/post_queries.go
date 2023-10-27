package queries

import (
	"database/sql"
	"time"

	"github.com/mozillazg/go-slugify"

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

func (q *PostQueries) GetPost(id int) (models.Post, error) {
	post := models.Post{}

	query := "select * from posts where id = $1"

	row := q.QueryRow(query, id)

	if err := row.Scan(&post.Id, &post.Title, &post.Body, &post.Views, &post.Slug, &post.Published, &post.PublishedAt, &post.UpdatedAt, &post.UserId); err != nil {
		return post, err
	}

	return post, nil
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

func (q *PostQueries) DeletePost(id int) error {
	query := "delete from posts where id = $1"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *PostQueries) UpdatePost(post *models.Post) error {
	query := "update posts set title = $2, body = $3, slug = $4, published = $5, updated_at = $6 where id = $1"

	_, err := q.Exec(query, &post.Id, &post.Title, &post.Body, &post.Slug, &post.Published, &post.PublishedAt)
	if err != nil {
		return err
	}

	return nil
}
