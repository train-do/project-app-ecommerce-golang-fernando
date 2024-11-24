package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"go.uber.org/zap"
)

type RepoOrder struct {
	Db     *sql.DB
	logger *zap.Logger
}

func NewRepoOrder(db *sql.DB, logger *zap.Logger) *RepoOrder {
	return &RepoOrder{db, logger}
}

func (r *RepoOrder) Insert(userId int, orderProduct model.OrderProduct) error {
	var order model.Order
	var carts []model.Cart
	fmt.Printf("%d %+v\n", userId, orderProduct)
	query := `select c.id, p.name , pv.id  , p.price_after_discount , c.qty from "Cart" c join "ProductVariant" pv on pv.id = c.product_variant_id join "Product" p on p.id = pv.product_id join "Variant" v on v.id = pv.variant_id where c.user_id = $1 and c.id = $2;`
	r.logger.Debug(query)
	tx, _ := r.Db.Begin()
	for _, v := range orderProduct.CartId {
		var cart model.Cart
		err := tx.QueryRow(query, userId, v).Scan(&cart.Id, &cart.Name, &cart.ProductVariantId, &cart.Price, &cart.Qty)
		if err != nil {
			r.logger.Error(err.Error())
			tx.Rollback()
			return fmt.Errorf(" No Cart or Invalid CartId")
		}
		carts = append(carts, cart)
	}
	query = `insert into "OrderProduct"("product_variant_id", "qty", "sub_total", "created_at") values ($1, $2, $3, $4) returning id;`
	var orderProductId []int
	var totalPrice float32
	r.logger.Debug(query)
	for _, v := range carts {
		var id int
		err := tx.QueryRow(query, v.ProductVariantId, v.Qty, float32(v.Qty)*v.Price, time.Now()).Scan(&id)
		if err != nil {
			r.logger.Error(err.Error())
			tx.Rollback()
			return err
		}
		order.Name = append(order.Name, v.Name)
		order.SubTotal = append(order.SubTotal, float32(v.Qty)*v.Price)
		orderProductId = append(orderProductId, id)
		totalPrice += (float32(v.Qty) * v.Price)
	}
	var address int
	query = `select id from "Address" where "user_id"=$1 and "is_default"=$2;`
	r.logger.Debug(query)
	err := tx.QueryRow(query, userId, true).Scan(&address)
	if err != nil {
		r.logger.Error("No address or default address has not been set")
		tx.Rollback()
		return fmt.Errorf(" No address or default address has not been set")
	}
	var orderId int
	query = `insert into "Order"("user_id", "total_price", "address_id", "created_at") values ($1, $2, $3, $4) returning id;`
	r.logger.Debug(query)
	err = tx.QueryRow(query, userId, totalPrice, address, time.Now()).Scan(&orderId)
	if err != nil {
		r.logger.Error(err.Error())
		tx.Rollback()
		return err
	}
	query = `update "OrderProduct" set "order_id"=$1 where "id"=$2;`
	r.logger.Debug(query)
	for _, v := range orderProductId {
		_, err := tx.Exec(query, orderId, v)
		if err != nil {
			r.logger.Error(err.Error())
			tx.Rollback()
			return err
		}
	}
	for _, v := range orderProduct.CartId {
		query := `delete from "Cart" where "id"=$1 and "user_id" = $2;`
		r.logger.Debug(query)
		_, err := tx.Exec(query, v, userId)
		if err != nil {
			r.logger.Error(err.Error())
			return err
		}
	}
	for _, v := range carts {
		query := `update "ProductVariant" set "stock"= "stock" - $1 where id=$2;`
		r.logger.Debug(query)
		_, err := tx.Exec(query, v.Qty, v.ProductVariantId)
		if err != nil {
			r.logger.Error(err.Error())
			return err
		}
	}
	tx.Commit()
	return nil
}
