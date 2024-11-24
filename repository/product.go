package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
	"go.uber.org/zap"
)

type RepoProduct struct {
	Db     *sql.DB
	logger *zap.Logger
}

func NewRepoProduct(db *sql.DB, logger *zap.Logger) *RepoProduct {
	return &RepoProduct{db, logger}
}

func (r *RepoProduct) FindAll(qp model.QueryProduct) ([]model.Product, int, error) {
	var args []any
	var totalItems int
	page, filter := utils.GenerateQuery(qp, &args)
	subQuery := fmt.Sprintf(`with "total_items" as (select count(*) as "total_items" from "Product" p join "Category" c on p.category_id = c.id %s), "ProductGallery" as (select p.id as product_id, array_agg(g.image_url) as image_urls from "Product" p join "Gallery" g on p.id = g.product_id group by p.id)`, filter)
	query := fmt.Sprintf(`%s select p.*, c."name" , array_to_json(array_agg(pv.stock)) AS stock , array_to_json(array_agg(v.description)) AS variant, ti."total_items", array_to_json(pg.image_urls) AS image_url, array_to_json(array_agg(pv.id)) AS "product variant id" from "Product" p join "Category" c on p.category_id = c.id join "ProductVariant" pv on pv.product_id = p.id   join "Variant" v on v.id = pv.variant_id join "total_items" ti on true join "ProductGallery" pg ON p.id = pg.product_id %s group by p.id , c.id, ti.total_items, pg.image_urls %s`, subQuery, filter, page)
	var products []model.Product
	r.logger.Debug(query)
	rows, err := r.Db.Query(query, args...)
	if err != nil {
		// fmt.Println(err, "PRODUCT1")
		r.logger.Error(err.Error())
		return []model.Product{}, 0, err
	}
	for rows.Next() {
		var product model.Product
		var stock, variant, imageUrl, productVariantId string
		err := rows.Scan(&product.Id, &product.CategoryId, &product.Name, &product.Description, &product.Price, &product.Discount, &product.PriceAfterDiscount, &product.Rating, &product.QtySold, &product.IsBestSelling, &product.CreatedAt, &product.Category, &stock, &variant, &totalItems, &imageUrl, &productVariantId)
		product.CategoryId = 0
		if err != nil {
			// fmt.Println(err, "PRODUCT2")
			r.logger.Error(err.Error())
			return []model.Product{}, 0, err
		}
		json.Unmarshal([]byte(stock), &product.Stock)
		json.Unmarshal([]byte(variant), &product.Variant)
		json.Unmarshal([]byte(imageUrl), &product.ImageUrl)
		json.Unmarshal([]byte(productVariantId), &product.ProductVariantId)
		products = append(products, product)
	}
	return products, totalItems, nil
}
func (r *RepoProduct) FindById(product *model.Product) error {
	query := `with "ProductGallery" as (select p.id as product_id, array_agg(g.image_url) as image_urls from "Product" p join "Gallery" g on p.id = g.product_id group by p.id) select p.*, c."name" , array_to_json(array_agg(pv.stock)) AS stock , array_to_json(array_agg(v.description)) AS variant, array_to_json(pg.image_urls) AS image_url from "Product" p join "Category" c on p.category_id = c.id join "ProductVariant" pv on pv.product_id = p.id   join "Variant" v on v.id = pv.variant_id join "ProductGallery" pg ON p.id = pg.product_id where p.id=$1 group by p.id , c.id, pg.image_urls;`
	var stock, variant, imageUrl string
	r.logger.Debug(query)
	err := r.Db.QueryRow(query, product.Id).Scan(&product.Id, &product.CategoryId, &product.Name, &product.Description, &product.Price, &product.Discount, &product.PriceAfterDiscount, &product.Rating, &product.QtySold, &product.IsBestSelling, &product.CreatedAt, &product.Category, &stock, &variant, &imageUrl)
	if err != nil {
		// fmt.Println(err, "PRODUCT")
		r.logger.Error(err.Error())
		return err
	}
	product.CategoryId = 0
	json.Unmarshal([]byte(stock), &product.Stock)
	json.Unmarshal([]byte(variant), &product.Variant)
	json.Unmarshal([]byte(imageUrl), &product.ImageUrl)
	return nil
}
func (r *RepoProduct) FindBanner() ([]model.Banner, error) {
	query := `select id, banner_url, title, subtitle, path_page, TO_CHAR(start_time, 'YYYY-MM-DD') AS time_start, TO_CHAR(end_time, 'YYYY-MM-DD') AS time_end from "Banner" WHERE CURRENT_DATE BETWEEN "start_time" AND "end_time";`
	var banners []model.Banner
	r.logger.Debug(query)
	rows, err := r.Db.Query(query)
	if err != nil {
		// fmt.Println(err, "Banner1")
		r.logger.Error(err.Error())
		return []model.Banner{}, err
	}
	for rows.Next() {
		var banner model.Banner
		err := rows.Scan(&banner.Id, &banner.BannerUrl, &banner.Title, &banner.Subtitle, &banner.PathPage, &banner.TimeStart, &banner.TimeEnd)
		if err != nil {
			// fmt.Println(err, "Banner2")
			r.logger.Error(err.Error())
			return []model.Banner{}, err
		}
		banners = append(banners, banner)
	}
	return banners, nil
}
func (r *RepoProduct) FindPromo() ([]model.ShowcaseProduct, error) {
	query := `select id, banner_url, title, subtitle, product_id, TO_CHAR(start_time, 'YYYY-MM-DD') AS time_start, TO_CHAR(end_time, 'YYYY-MM-DD') AS time_end from "Promo" WHERE CURRENT_DATE BETWEEN "start_time" AND "end_time";`
	var promos []model.ShowcaseProduct
	r.logger.Debug(query)
	rows, err := r.Db.Query(query)
	if err != nil {
		// fmt.Println(err, "Promo1")
		r.logger.Error(err.Error())
		return []model.ShowcaseProduct{}, err
	}
	for rows.Next() {
		var promo model.ShowcaseProduct
		err := rows.Scan(&promo.Id, &promo.BannerUrl, &promo.Title, &promo.Subtitle, &promo.ProductId, &promo.TimeStart, &promo.TimeEnd)
		if err != nil {
			// fmt.Println(err, "Promo2")
			r.logger.Error(err.Error())
			return []model.ShowcaseProduct{}, err
		}
		promos = append(promos, promo)
	}
	return promos, nil
}
func (r *RepoProduct) FindRecommend() ([]model.ShowcaseProduct, error) {
	query := `select id, banner_url, title, subtitle, product_id, TO_CHAR(start_time, 'YYYY-MM-DD') AS time_start, TO_CHAR(end_time, 'YYYY-MM-DD') AS time_end from "Recommend" WHERE CURRENT_DATE BETWEEN "start_time" AND "end_time";`
	var recommends []model.ShowcaseProduct
	r.logger.Debug(query)
	rows, err := r.Db.Query(query)
	if err != nil {
		// fmt.Println(err, "Recommend1")
		r.logger.Error(err.Error())
		return []model.ShowcaseProduct{}, err
	}
	for rows.Next() {
		var recommend model.ShowcaseProduct
		err := rows.Scan(&recommend.Id, &recommend.BannerUrl, &recommend.Title, &recommend.Subtitle, &recommend.ProductId, &recommend.TimeStart, &recommend.TimeEnd)
		if err != nil {
			// fmt.Println(err, "Recommend2")
			r.logger.Error(err.Error())
			return []model.ShowcaseProduct{}, err
		}
		recommends = append(recommends, recommend)
	}
	return recommends, nil
}
func (r *RepoProduct) FindAllCategory() ([]model.Category, error) {
	query := `select * from "Category";`
	var categories []model.Category
	r.logger.Debug(query)
	rows, err := r.Db.Query(query)
	if err != nil {
		// fmt.Println(err, "Category1")
		r.logger.Error(err.Error())
		return []model.Category{}, err
	}
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			// fmt.Println(err, "Category2")
			r.logger.Error(err.Error())
			return []model.Category{}, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
