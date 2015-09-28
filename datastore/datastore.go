package datastore

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("postgres", "host=localhost sslmode=disable")
	if err != nil {
		log.Print(err)
		panic(err)
	}
	// connect to DB
}

func CreateReview(title string, author string, gitdiff string) (int, error) {

	var reviewId int

	err := db.QueryRow("SELECT * FROM add_review($1, $2, $3)", title, author, gitdiff).Scan(&reviewId)
	if err != nil {
		return 0, err
	}

	return reviewId, nil
}

func GetReview(reviewId int) (title string, author string, diffs []string, err error) {

	err = db.QueryRow("SELECT title, author FROM reviews WHERE id = $1", reviewId).Scan(&title, &author)
	if err != nil {
		log.Print("1", err)
		return "", "", nil, err
	}

	var rows *sql.Rows
	rows, err = db.Query("SELECT gitdiff FROM review_diffs WHERE review_id = $1", reviewId)
	if err != nil {
		log.Print("2", err)
		return "", "", nil, err
	}

	defer rows.Close()
	for rows.Next() {

		var diff string
		err = rows.Scan(&diff)
		if err != nil {
			return "", "", nil, err
		}

		diffs = append(diffs, diff)
	}

	err = rows.Err()
	if err != nil {
		return "", "", nil, err
	}

	return title, author, diffs, nil
}

func UpdateReview(reviewId int, gitdiff string) error {

	_, err := db.Exec("INSERT INTO review_diffs(review_id, gitdiff) VALUES ($1, $2)", reviewId, gitdiff)
	if err != nil {
		return err
	}

	return nil
}
