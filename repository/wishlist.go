package repository

import (
	"database/sql"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"go.uber.org/zap"
)

type RepoWishlist struct {
	Db     *sql.DB
	logger *zap.Logger
}

func NewRepoWishlist(db *sql.DB, logger *zap.Logger) *RepoWishlist {
	return &RepoWishlist{db, logger}
}

func (r *RepoWishlist) Insert(wishlist *model.Wishlist) error {
	query := `insert into "Wishlist"("user_id", "product_id") values ($1, $2) returning id`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, wishlist.UserId, wishlist.ProductId).Scan(&wishlist.Id)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoWishlist) Delete(wishlist *model.Wishlist) error {
	query := `delete from "Wishlist" where "id"=$1 and "user_id"=$2 returning product_id`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, wishlist.Id, wishlist.UserId).Scan(&wishlist.ProductId)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
