package db

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func InsertUser(name, email string) error {
	sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, name, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
