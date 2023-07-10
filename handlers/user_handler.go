package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"time"
	"fmt"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
)

type User struct {
    gorm.Model // Utilisation de gorm.Model pour les champs communs (ID, CreatedAt, UpdatedAt, DeletedAt)
	ID       uint   `gorm:"primaryKey"`
    Username   string `gorm:"column:username"`
    Password   string `gorm:"column:password"`
}

type UserHandler struct {
	users []User // Utilise une slice pour stocker les utilisateurs
}

func NewUserHandler() *UserHandler {
	// Initialise et retourne une nouvelle instance de UserHandler ici
	return &UserHandler{
        users: make([]User, 0), // Initialise la slice vide
    }
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Vérifie si l'utilisateur existe déjà avec le même nom d'utilisateur (username)
	for _, user := range h.users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
	}

	// Ajoute le nouvel utilisateur à la slice
	h.users = append(h.users, newUser)

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}


func (h *UserHandler) DeleteUser(c *gin.Context) {
    userID := c.Param("id")

    // Recherche l'index de l'utilisateur dans la slice
    index := -1
    for i, user := range h.users {
        // Convertit le user.ID en string pour la comparaison
        if strconv.Itoa(int(user.ID)) == userID {
            index = i
            break
        }
    }

    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Supprime l'utilisateur de la slice
    h.users = append(h.users[:index], h.users[index+1:]...)

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}



func (h *UserHandler) Login(c *gin.Context) {
	// Récupère les données de la requête pour l'authentification de l'utilisateur
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Vérifie les informations d'identification de l'utilisateur (par exemple, vérifie si le nom d'utilisateur et le mot de passe correspondent)
	// Si les informations d'identification sont valides, génère un token JWT

	// Exemple de génération d'un token JWT
	claims := &jwt.StandardClaims{
		Subject:   credentials.Username,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Durée de validité du token (1 jour)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret")) // Clé secrète pour la signature du token
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil // Clé secrète utilisée pour la vérification de la signature du token
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	// Vérifie la validité du token JWT
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	_, err := validateToken(strings.TrimPrefix(tokenString, "Bearer "))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Vérifie si l'utilisateur existe
	// Si l'utilisateur n'existe pas, renvoie une réponse d'erreur appropriée

	// Si l'utilisateur existe, récupère les détails des utilisateurs depuis la source de données
	// et les renvoie en réponse
	usernames := make([]string, len(h.users))
	for i, user := range h.users {
		usernames[i] = user.Username
	}

	c.JSON(http.StatusOK, gin.H{"usernames": usernames})
}

// Pour récupérer le nombre d'utilisateurs créés :
func (h *UserHandler) GetUserCount(c *gin.Context) {
	count := len(h.users)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *UserHandler) findUserByID(userID int) *User {
	for _, user := range h.users {
		if int(user.ID) == userID {
			return &user
		}
	}
	return nil
}