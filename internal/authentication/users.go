package authentication

import "database/sql"

var users map[int][]string

func GetUserMap(db *sql.DB) {
	rows, err := db.Query("SELECT id, username, fullname, email, country, city FROM Users")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id                                       int
			fullname, username, email, country, city string
		)
		if err := rows.Scan(&id, &fullname, &username, &email, &country, &city); err != nil {
			return
		}
		userInfo := []string{fullname, username, email, country, city}
		users[id] = userInfo
	}

	if err := rows.Err(); err != nil {
		return
	}
}
