package mapper

import (
	"database/sql"

	// Postgres Driver
	_ "github.com/lib/pq"
)

// FromDB reads the mapping from the database
func FromDB(dbURL string) (map[string]string, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT key, url FROM public.mapping")
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for rows.Next() {
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}
		result[key] = value
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
