package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Messages []Message `json:"messages"`
	Users  []User    `json:"users"`
}

type RoomHandler struct {
	rooms []Room // Utilise une slice pour stocker les noms des salons
	userHandler *UserHandler
}

func NewRoomHandler(userHandler *UserHandler) *RoomHandler {
	// Initialise et retourne une nouvelle instance de RoomHandler ici
	return &RoomHandler{
		rooms: make([]Room, 0), // Initialise la slice vide
		userHandler: userHandler,
	}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var newRoom Room
	if err := c.BindJSON(&newRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Vérifie si le nom du salon existe déjà
	for _, room := range h.rooms {
		if room.Name == newRoom.Name {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room name already exists"})
			return
		}
	}

	// Initialise la liste des utilisateurs du salon
	newRoom.Users = make([]User, 0)

	// Ajoute le nom du salon à la slice des salons
	newRoom.ID = len(h.rooms)
	newRoom.Messages = make([]Message, 0)
	h.rooms = append(h.rooms, newRoom)

	c.JSON(http.StatusOK, gin.H{"message": "Room created successfully"})
}


func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	roomID := c.Param("id")

	// Convertit l'ID de la salle en entier
	id, err := strconv.Atoi(roomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	// Recherche la salle dans la liste des salles
	index := -1
	for i, room := range h.rooms {
		if room.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// Supprime la salle de la liste des salles
	h.rooms = append(h.rooms[:index], h.rooms[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

func (h *RoomHandler) findRoomByID(roomID string) *Room {
    id, err := strconv.Atoi(roomID)
    if err != nil {
        return nil
    }

    for i := range h.rooms {
        if h.rooms[i].ID == id {
            return &h.rooms[i]
        }
    }
    return nil
}


func (h *RoomHandler) JoinRoom(c *gin.Context) {
	roomID := c.Param("id")
	userID := c.Param("user_id")

	// Recherche le salon correspondant à l'ID fourni
	room := h.findRoomByID(roomID)
	if room == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// Vérifie si l'utilisateur existe dans la slice des utilisateurs
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := h.userHandler.findUserByID(userIDInt)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Ajoute l'utilisateur au salon
	room.Users = append(room.Users, *user)

	c.JSON(http.StatusOK, gin.H{"message": "User joined the room successfully"})
}



func (h *RoomHandler) GetRoomUsers(c *gin.Context) {
	roomID := c.Param("id")

	// Recherche le salon correspondant à l'ID fourni
	room := h.findRoomByID(roomID)
	if room == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// Vous pouvez ajouter la logique appropriée ici pour récupérer la liste des utilisateurs du salon

	// Par exemple, supposons que vous stockiez les utilisateurs dans une map `roomUsers` où la clé est l'ID de l'utilisateur
	roomUsers := make(map[uint]User)
	for _, user := range h.userHandler.users {
		// Vérifie si l'utilisateur fait partie du salon
		if _, ok := roomUsers[user.ID]; ok {
			roomUsers[user.ID] = user
		}
	}

	// Renvoie la liste des utilisateurs en réponse
	users := make([]User, 0, len(roomUsers))
	for _, user := range roomUsers {
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *RoomHandler) ListUsersInRoom(c *gin.Context) {
    roomID := c.Param("id")

    // Recherche le salon correspondant à l'ID fourni
    room := h.findRoomByID(roomID)
    if room == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
        return
    }

    // Renvoie la liste des utilisateurs du salon
    c.JSON(http.StatusOK, gin.H{"users": room.Users})
}