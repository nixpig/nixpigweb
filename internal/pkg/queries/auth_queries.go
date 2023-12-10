package queries

import (
	"fmt"

	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

func GetUserByUsername(username string) (models.User, error) {
	query := `select id_, username_, email_, password_, is_admin_ from users_ where username_ = $1`

	var user models.User

	row := database.DB.QueryRow(query, username)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	query := `select id_, username_, email_, password_, is_admin_ from users_ where email_ = $1`

	var user models.User

	row := database.DB.QueryRow(query, email)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
		return user, err
	}

	return user, nil
}

func SaveSession(session models.Session) (int64, error) {
	query := `insert into sessions_ (token_, expires_at_, issued_at_, user_id_) values ($1, $2, $3, $4)`

	res, err := database.DB.Exec(query, &session.Token, &session.ExpiresAt, &session.IssuedAt, &session.UserId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteSessionsByUserId(userId int) (int64, error) {
	query := `delete from sessions_ where user_id = $1`

	res, err := database.DB.Exec(query, userId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteSessionsBySessionId(sessionId int) (int64, error) {
	query := `delete from sessions_ where id_ = $1`

	res, err := database.DB.Exec(query, sessionId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func GetSessionByToken(token string) (models.Session, error) {
	query := `select id_, token_, expires_at_, issued_at_, user_id_ from sessions_ where token_ = $1`

	var session models.Session

	row := database.DB.QueryRow(query, token)

	if err := row.Scan(&session.Id, &session.Token, &session.ExpiresAt, &session.IssuedAt, &session.UserId); err != nil {
		return session, err
	}

	return session, nil
}

func GetSessionByUserId(userId int) (models.Session, error) {
	query := `select id_, token_, expires_at_, issued_at_, user_id_ from sessions_ where user_id_ = $1`

	var session models.Session

	row := database.DB.QueryRow(query, userId)

	if err := row.Scan(&session.Id, &session.Token, &session.ExpiresAt, &session.IssuedAt, &session.UserId); err != nil {
		return session, err
	}

	return session, nil
}

func ChangePassword(username string, newPassword string) (bool, error) {
	query := `update users_ set password_ = $2 where username_ = $1`

	fmt.Println("username: ", username)
	fmt.Println("new password: ", newPassword)

	res, err := database.DB.Exec(query, username, newPassword)
	if err != nil {
		fmt.Println("ERROR: failed to execute query to update password\n", err)
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("ERROR: failed to fetch rows affected by update password query\n", err)
		return false, err
	}

	if rowsAffected == 0 {
		fmt.Println("ERROR: no rows affected by update password query")

		return false, err
	}

	return true, nil
}
