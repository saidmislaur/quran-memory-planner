package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// contetx.Context принимаем, для создания подключений и быть в курсе статуса подключения
// Почему принимает connString - для универсальности функции и возможность в любой момент менять подключение
// error - если при подключении к бд появится какая то ошибка, функция выведет эту ошибку 
func NewPostgres(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}