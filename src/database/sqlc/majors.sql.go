// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: majors.sql

package database

import (
	"context"
)

const createDummyMajor = `-- name: CreateDummyMajor :many
INSERT INTO majors (
    title,
    "desc",
    "interest_num",
    image_url
) VALUES (
    'Teknik Informatika',
    'The study of computers and their applications',
    1100,
    'https://sample.dummy'
) RETURNING id, title, "desc", image_url, interest_num
`

func (q *Queries) CreateDummyMajor(ctx context.Context) ([]Majors, error) {
	rows, err := q.db.QueryContext(ctx, createDummyMajor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Majors
	for rows.Next() {
		var i Majors
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Desc,
			&i.ImageUrl,
			&i.InterestNum,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMajors = `-- name: GetMajors :many
SELECT id, title, "desc", image_url, interest_num FROM majors
`

func (q *Queries) GetMajors(ctx context.Context) ([]Majors, error) {
	rows, err := q.db.QueryContext(ctx, getMajors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Majors
	for rows.Next() {
		var i Majors
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Desc,
			&i.ImageUrl,
			&i.InterestNum,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
