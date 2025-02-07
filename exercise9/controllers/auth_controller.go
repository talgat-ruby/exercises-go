package controllers

import (
	"fmt"
	"net/http"
	"time"

	"alinurmyrzakhanov/models"
	"alinurmyrzakhanov/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("MY_SUPER_SECRET")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
			c.Abort()
			return
		}

		tokenStr := ""
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenStr)
		if err != nil || tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Некорректный заголовок"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неверный метод подписи")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный payload"})
			c.Abort()
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id не найден в токене"})
			c.Abort()
			return
		}
		c.Set("user_id", uint(userIDFloat))

		c.Next()
	}
}

func SignUp(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось захешировать пароль"})
		return
	}
	user := models.User{
		Email:        input.Email,
		PasswordHash: hashed,
	}
	if err := repositories.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Пользователь уже существует или другая ошибка"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно создан",
		"user":    user.Email,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repositories.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	if !CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Вы успешно вышли из системы (токен считать недействительным на клиенте).",
	})
}

func GetUserID(c *gin.Context) uint {
	val, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	userID, _ := val.(uint)
	return userID
}
