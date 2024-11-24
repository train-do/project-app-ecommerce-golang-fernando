package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/train-do/project-app-ecommerce-golang-fernando/handler"
	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
	"github.com/train-do/project-app-ecommerce-golang-fernando/utils"
	"go.uber.org/zap"
)

func Authentication(h *handler.HandlerUser) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := r.Cookie("id")
			if err != nil || err == http.ErrNoCookie || id.Value == "" {
				fmt.Println(err, "DARI AUTH MIDDLEWARE - ID")
				response := utils.SetResponse(w, model.Response{}, http.StatusUnauthorized, "No Authentication")
				json.NewEncoder(w).Encode(response)
				return
			}
			token, err := r.Cookie("token")
			if err != nil || err == http.ErrNoCookie || token.Value == "" {
				fmt.Println(err, "DARI AUTH MIDDLEWARE - TOKEN")
				response := utils.SetResponse(w, model.Response{}, http.StatusUnauthorized, "No Authentication")
				json.NewEncoder(w).Encode(response)
				return
			}
			var user model.User
			user.Id = utils.ToInt(id.Value)
			user.Token = &token.Value
			err = h.Service.Repo.FindUser(&user)
			if err != nil {
				response := utils.SetResponse(w, model.Response{}, http.StatusUnauthorized, "No Authentication")
				json.NewEncoder(w).Encode(response)
				return
			}
			fmt.Println(id, token, "MIDDLEWARE")
			next.ServeHTTP(w, r)
		})
	}
}
func LoggerReq(log *zap.Logger) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			handler.ServeHTTP(w, r)

			duration := time.Since(start)

			log.Info("http request middleware", zap.String("url", r.URL.String()), zap.Duration("duration", duration))
		})
	}
}
