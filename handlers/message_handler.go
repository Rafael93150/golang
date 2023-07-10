package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type MessageHandler struct {
	rooms []Room // Utilise une slice pour stocker les salons
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		rooms: make([]Room, 0), // Initialise la slice des salons vide
	}
}

// Méthode auxiliaire findRoomByID pour rechercher un salon par ID :
func (h *MessageHandler) findRoomByID(id int) *Room {
	for i := range h.rooms {
		if h.rooms[i].ID == id {
			return &h.rooms[i]
		}
	}
	return nil
}

func (h *MessageHandler) SendMessageToRoom(c *gin.Context) {
	roomID := c.Param("id")

	// Convertit l'ID de la salle en entier
	id, err := strconv.Atoi(roomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	// Récupère le message depuis la requête
	var newMessage Message
	if err := c.BindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Recherche le salon correspondant à l'ID fourni
	room := h.findRoomByID(id)
	if room == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// Ajoute le message au salon
	room.Messages = append(room.Messages, newMessage)

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}



