package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/train-do/project-app-ecommerce-golang-fernando/model"
)

func SetResponse(w http.ResponseWriter, data model.Response, statusCode int, message string) model.Response {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data.StatusCode = statusCode
	data.Message = message
	return data
}
func SetCookie(w http.ResponseWriter, name string, value string) {
	expiration := time.Now().Add(10 * time.Minute)
	// expirationFormatted := expiration.Format("Mon, 02 Jan 2006 15:04:05 +0700")
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Domain:  "localhost",
		Expires: expiration,
		// MaxAge:   360, // kalau mau pake detik
		HttpOnly: true,
	}
	// w.Header().Set("Expires", expirationFormatted)
	http.SetCookie(w, &cookie)
}
func ToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
func ToBool(str string) bool {
	return strings.ToLower(str) == "true"
}
func GenerateQuery(qp model.QueryProduct, args *[]any) (string, string) {
	page := ``
	filter := `where 1=1`
	count := 0
	if qp.Name != "" {
		count++
		filter += fmt.Sprintf(` and p."name" ilike $%d`, count)
		*args = append(*args, "%"+qp.Name+"%")
	}
	if qp.Category != "" {
		count++
		filter += fmt.Sprintf(` and c."name" ilike $%d`, count)
		*args = append(*args, qp.Category)
	}
	if qp.IsBestSelling {
		count++
		filter += fmt.Sprintf(` and p."is_best_selling" =  $%d`, count)
		*args = append(*args, qp.IsBestSelling)
	}
	if qp.Page == 0 || qp.Page == 1 {
		page = `limit 6 offset 0;`
	} else {
		page = fmt.Sprintf(`limit 6 offset %d;`, ((qp.Page - 1) * 6))
	}
	return page, filter
}
