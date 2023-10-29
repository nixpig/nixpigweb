package queries

import (
	"database/sql"
	"fmt"
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

	query := "select * from post_"

	rows, err := q.Query(query)
	if err != nil {
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		post := models.Post{}

		if err := rows.Scan(&post.Id, &post.Title, &post.SubTitle, &post.Body, &post.Slug, &post.Status, &post.CreatedAt, &post.PublishedAt, &post.UpdatedAt, &post.UserId, &post.CategoryId); err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (q *PostQueries) GetPost(id int) (models.Post, error) {
	post := models.Post{}

	query := "select * from post_ where id = $1"

	row := q.QueryRow(query, id)

	if err := row.Scan(&post.Id, &post.Title, &post.SubTitle, &post.Body, &post.Slug, &post.Status, &post.CreatedAt, &post.PublishedAt, &post.UpdatedAt, &post.UserId, &post.CategoryId); err != nil {
		return post, err
	}

	return post, nil
}

func (q *PostQueries) CreatePost(post *models.NewPost) error {
	query := "insert into post_ ( title_, subtitle_, body_, slug_, status_,  published_at_, updated_at_, user_id_, category_id_) values($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	fmt.Println("creating post...")

	slug := slugify.Slugify(post.Title)

	_, err := q.Exec(query, &post.Title, &post.Subtitle, &post.Body, slug, "published", time.Now(), time.Now(), &post.UserId, &post.CategoryId)
	fmt.Println("error in query:", err)
	if err != nil {
		return err
	}

	return nil
}

func (q *PostQueries) DeletePost(id int) error {
	query := "delete from post_ where id = $1"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *PostQueries) UpdatePost(post *models.Post) error {
	query := "update post_ set title_ = $2, subtitle_ = $3, body_ = $4, slug_ = $5, status_ = $6, updated_at_ = $7, category_id_ = $8 where id = $1"

	slug := slugify.Slugify(post.Title)

	_, err := q.Exec(query, &post.Id, &post.Title, &post.SubTitle, &post.Body, slug, &post.Status, &post.UpdatedAt, &post.CategoryId)
	if err != nil {
		return err
	}

	return nil
}
