package repository

import (
	"database/sql"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"go.uber.org/zap"
)

type RepoUser struct {
	Db     *sql.DB
	logger *zap.Logger
}

func NewRepoUser(db *sql.DB, logger *zap.Logger) *RepoUser {
	return &RepoUser{db, logger}
}

func (r *RepoUser) Login(user *model.User) error {
	query := `update "User" set "token" = $1 where ("email" ilike $2 or "phone" = $3) and "password" = $4 returning *`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, *user.Token, user.Email, user.Phone, user.Password).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, user.Token)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoUser) FindUser(user *model.User) error {
	query := `select 1 from "User" where "id"=$1 and "token"=$2`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, user.Id, user.Token).Scan(&user.Password)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoUser) InsertUser(user *model.User) error {
	query := `insert into "User"("name", "email", "phone", "password") values ($1, $2, $3, $4)`
	r.logger.Debug(query)
	_, err := r.Db.Exec(query, user.Name, user.Email, user.Phone, user.Password)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	// fmt.Printf("REGISTER USER %+v\n", result)
	return nil
}
func (r *RepoUser) UpdateUser(user *model.User) error {
	query := `update "User" set "name" = $1, "email" = $2, "password" = $3 where "id"=$4 and "token" = $5 returning *`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, user.Name, user.Email, user.Password, user.Id, *user.Token).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, user.Token)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoUser) FindAddressesByUserId(userId int) ([]model.Address, error) {
	var addresses []model.Address
	query := `select * from "Address" where "user_id"=$1 order by "id"`
	r.logger.Debug(query)
	rows, err := r.Db.Query(query, userId)
	if err != nil {
		return []model.Address{}, err
	}
	for rows.Next() {
		var address model.Address
		err := rows.Scan(&address.Id, &address.UserId, &address.Name, &address.Email, &address.Address, &address.IsDefault)
		if err != nil {
			r.logger.Error(err.Error())
			return []model.Address{}, err
		}
		addresses = append(addresses, address)
	}
	return addresses, err
}
func (r *RepoUser) InsertAddress(address *model.Address) error {
	query := `insert into "Address"("user_id", "name", "email", "address") values ($1, $2, $3, $4) returning *`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, address.UserId, address.Name, address.Email, address.Address).Scan(&address.Id, &address.UserId, &address.Name, &address.Email, &address.Address, &address.IsDefault)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoUser) UpdateAddress(address *model.Address) error {
	query := `update "Address" set "name" = $1, "email" = $2, "address" = $3 where "id"=$4 and "user_id" = $5 returning *`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, address.Name, address.Email, address.Address, address.Id, address.UserId).Scan(&address.Id, &address.UserId, &address.Name, &address.Email, &address.Address, &address.IsDefault)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoUser) UpdateDefaultAddress(address *model.Address) error {
	query := `update "Address" set "is_default"=$1 where "user_id" = $2 and "is_default"=$3`
	// r.logger.Debug(query)
	for {
		err := r.Db.QueryRow(query, false, address.UserId, true).Scan()
		if err != nil {
			// fmt.Println("SET FALSE")
			break
		}
	}
	query = `update "Address" set "is_default"=$1 where "id"=$2 and "user_id" = $3 returning *`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, true, address.Id, address.UserId).Scan(&address.Id, &address.UserId, &address.Name, &address.Email, &address.Address, &address.IsDefault)
	if err != nil {
		// fmt.Println("SET TRUE")
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoUser) DeleteAddress(address *model.Address) error {
	query := `delete from "Address" where "id"=$1 and "user_id" = $2 returning *`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, address.Id, address.UserId).Scan(&address.Id, &address.UserId, &address.Name, &address.Email, &address.Address, &address.IsDefault)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
