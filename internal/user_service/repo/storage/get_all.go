package storage

import (
	"context"

	"github.com/manc88/demo/internal/models"
)

func (p *PgStorage) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	rows, err := p.pool.Query(ctx, getAllQuery)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UID, &user.Name, &user.Email, &user.Age)
		//probably we dont want to break on some erorrs(ex:null values in table)
		if err != nil {
			continue
		}
		users = append(users, &user)
	}

	return users, nil

}
