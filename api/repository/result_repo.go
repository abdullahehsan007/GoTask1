package repository

import (
	"GOTASK/model"

	"github.com/jmoiron/sqlx"
)

func Result(db *sqlx.DB, filename string, value model.Results, timeTaken string) error {
	_, err := db.Exec(`
				INSERT INTO results (
					filename, words, lines, spaces, punctuation, characters,
					sentences, vowels, consonants, digits, paragraphs, combine_time
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		filename,
		value.Words,
		value.Lines,
		value.Spaces,
		value.Punc,
		value.Characters,
		value.Sentences,
		value.Vowels,
		value.Consonants,
		value.Digits,
		value.Para,
		timeTaken,
	)
	return err
}
