package router

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/train-do/project-app-ecommerce-golang-fernando/database"
	"github.com/train-do/project-app-ecommerce-golang-fernando/handler"
	"github.com/train-do/project-app-ecommerce-golang-fernando/middleware"
	"github.com/train-do/project-app-ecommerce-golang-fernando/repository"
	"github.com/train-do/project-app-ecommerce-golang-fernando/service"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
	"go.uber.org/zap"
)

func RouterAPI() (*chi.Mux, *zap.Logger, error) {
	config, err := utils.ReadConfiguration()
	if err != nil {
		return nil, nil, err
	}
	db, err := database.InitDB(config)
	if err != nil {
		return nil, nil, err
	}
	logger := utils.InitLog(config)

	repoUser := repository.NewRepoUser(db, logger)
	serviceUser := service.NewServiceUser(repoUser)
	handlerUser := handler.NewHandlerUser(serviceUser)
	repoProduct := repository.NewRepoProduct(db, logger)
	serviceProduct := service.NewServiceProduct(repoProduct)
	handlerProduct := handler.NewHandlerProduct(serviceProduct)
	repoWishlist := repository.NewRepoWishlist(db, logger)
	serviceWishlist := service.NewServiceWishlist(repoWishlist)
	handlerWishlist := handler.NewHandlerWishlist(serviceWishlist)
	repoCart := repository.NewRepoCart(db, logger)
	serviceCart := service.NewServiceCart(repoCart)
	handlerCart := handler.NewHandlerCart(serviceCart)
	repoOrder := repository.NewRepoOrder(db, logger)
	serviceOrder := service.NewServiceOrder(repoOrder)
	handlerOrder := handler.NewHandlerOrder(serviceOrder)
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.LoggerReq(logger))
		r.Post("/register", handlerUser.Register)
		r.Post("/login", handlerUser.Login)
		r.Get("/product", handlerProduct.GetAll)
		r.Get("/product/{id}", handlerProduct.GetById)
		r.Get("/banner", handlerProduct.GetBanner)
		r.Get("/promo", handlerProduct.GetPromo)
		r.Get("/recommend", handlerProduct.GetRecommend)
		r.Get("/category", handlerProduct.GetAllCategory)
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.Authentication(handlerUser))
			r.Put("/update", handlerUser.UpdateUser)
			r.Post("/logout", handlerUser.Logout)
			r.Route("/address", func(r chi.Router) {
				r.Get("/", handlerUser.GetAddresses)
				r.Post("/add", handlerUser.AddAddress)
				r.Put("/update/{id}", handlerUser.UpdateAddress)
				r.Patch("/default/{id}", handlerUser.SetDefaultAddress)
				r.Delete("/delete/{id}", handlerUser.DeleteAddress)
			})
		})
		r.Group(func(r chi.Router) {
			r.Use(middleware.Authentication(handlerUser))
			r.Post("/wishtlist", handlerWishlist.AddWishlist)
			r.Delete("/wishtlist/{id}", handlerWishlist.AddWishlist)
			r.Get("/cart", handlerCart.GetAll)
			r.Get("/cart/total", handlerCart.GetTotal)
			r.Post("/cart", handlerCart.AddCart)
			r.Put("/cart/plus/{productVariantId}", handlerCart.UpdateIncrementQty)
			r.Put("/cart/minus/{productVariantId}", handlerCart.UpdateDecrementQty)
			r.Delete("/cart/{productVariantId}", handlerCart.DeleteCart)
			r.Post("/order", handlerOrder.Create)
		})
	})
	return router, logger, nil
}
