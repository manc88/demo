package storage

const (
	createQuery = `
	INSERT INTO 
	users (name,email,age)
	VALUES ($1,$2,$3)
	RETURNING uid`

	deleteQuery = `
	UPDATE users
	SET deleted = true
	WHERE uid = $1
	RETURNING uid
	`
	getAllQuery = `
	SELECT uid,name,email,age
	FROM users
	WHERE not deleted`
)
