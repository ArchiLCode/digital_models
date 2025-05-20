package middleware

import (
	"gateway/utils"
	"net/http"
)

// Middleware для проверки наличия и валидности токена
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Ожидается, что токен будет в формате "Bearer <token>"
		token, err := utils.ExtractToken(authHeader)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		// Проверяем валидность токена
		if !utils.IsValidToken(token) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Продолжаем выполнение следующего обработчика
		next.ServeHTTP(w, r)
	})
}
