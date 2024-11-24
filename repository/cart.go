package repository

import (
	"database/sql"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"go.uber.org/zap"
)

type RepoCart struct {
	Db     *sql.DB
	logger *zap.Logger
}

func NewRepoCart(db *sql.DB, logger *zap.Logger) *RepoCart {
	return &RepoCart{db, logger}
}

func (r *RepoCart) FindAll(userId int) ([]model.Cart, error) {
	var carts []model.Cart
	query := `select c.id, p."name" , v.description , p.price_after_discount , c.qty from "Cart" c join "ProductVariant" pv on pv.id = c.product_variant_id join "Product" p on p.id = pv.product_id join "Variant" v on v.id = pv.variant_id where c.user_id = $1;`
	r.logger.Debug(query)
	rows, err := r.Db.Query(query, userId)
	if err != nil {
		r.logger.Error(err.Error())
		return []model.Cart{}, err
	}
	for rows.Next() {
		var cart model.Cart
		rows.Scan(&cart.Id, &cart.Name, &cart.Variant, &cart.Price, &cart.Qty)
		if err != nil {
			r.logger.Error(err.Error())
			return []model.Cart{}, err
		}
		carts = append(carts, cart)
		// fmt.Println(userId, cart, query)
	}
	return carts, nil
}
func (r *RepoCart) Insert(Cart *model.Cart) error {
	query := `insert into "Cart"("user_id", "product_variant_id", "qty") select $1, $2, $3 where not exists (select 1 from "Cart" where "user_id"=$1 and "product_variant_id"=$2) returning id;`
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, Cart.UserId, Cart.ProductVariantId, 1).Scan(&Cart.Id)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoCart) UpdateIncrementQty(id int, userId int) error {
	// fmt.Println("MASUK INCREASE", id, userId)
	query := `update "Cart" set "qty"= "qty" + $1 where "product_variant_id"=$2 and "user_id"=$3`
	r.logger.Debug(query)
	_, err := r.Db.Exec(query, 1, id, userId)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoCart) UpdateDecrementQty(id int, userId int) error {
	// fmt.Println("MASUK DECREASE", id, userId)
	query := `update "Cart" set "qty"= "qty" - $1 where "product_variant_id"=$2 and "user_id"=$3`
	r.logger.Debug(query)
	_, err := r.Db.Exec(query, 1, id, userId)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoCart) Delete(id int, userId int) error {
	query := `delete from "Cart" where "id"=$1 and "user_id" = $2`
	r.logger.Debug(query)
	_, err := r.Db.Exec(query, id, userId)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
func (r *RepoCart) FindTotal(userId int) (int, error) {
	query := `select sum(qty) from "Cart" c where c.user_id = $1;`
	r.logger.Debug(query)
	var totalCart int
	err := r.Db.QueryRow(query, userId).Scan(&totalCart)
	if err != nil {
		r.logger.Error(err.Error())
		return 0, err
	}
	return totalCart, nil
}
