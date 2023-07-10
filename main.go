package main

import (
	"golang-chat-api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// utiliser une structure pour stocker les connexions clients ou d'autres données nécessaires
	clients = make(map[*websocket.Conn]bool)
)

func main() {
	r := gin.Default()


	// Route pour gérer les connexions WebSocket
	r.GET("/ws", handleWebSocket)

	// Crée une instance des gestionnaires
	userHandler := handlers.NewUserHandler()
	roomHandler := handlers.NewRoomHandler(userHandler)
	messageHandler := handlers.NewMessageHandler()

	api := r.Group("/api")


	// Routes pour la gestion des utilisateurs
	api.POST("/login", userHandler.Login)
	api.POST("/users", userHandler.AddUser)
	api.DELETE("/users/:id", userHandler.DeleteUser)
	api.GET("/users", userHandler.GetUsers)
	api.GET("/users/count", userHandler.GetUserCount)
	api.GET("/rooms/:id/users", roomHandler.ListUsersInRoom)



	// Routes pour la gestion des salons
	api.POST("/rooms/create", roomHandler.CreateRoom)
	api.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	// Route pour la gestion des messages
	api.POST("/rooms/:id/messages", messageHandler.SendMessageToRoom)

	// Routes pour rejoindre un salon
	api.POST("/rooms/:id/join/:user_id", roomHandler.JoinRoom)
	r.GET("/rooms/:id/users", roomHandler.GetRoomUsers)


	r.Run(":8000")
}

func handleWebSocket(c *gin.Context) {
	// Mise à niveau de la connexion HTTP en WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// Gestion de l'erreur
		return
	}

	// Fermez la connexion lorsque cette fonction se termine
	defer conn.Close()

	// Ajoutez la connexion à la liste des clients
	clients[conn] = true

	// Boucle de lecture des messages du client
	for {
		// Lire le message du client
		_, message, err := conn.ReadMessage()
		if err != nil {
			// Gestion de l'erreur ou de la déconnexion du client
			delete(clients, conn)
			break
		}

		// Faites quelque chose avec le message reçu, par exemple, diffusez-le à tous les autres clients connectés
		broadcast(message)
	}
}

func broadcast(message []byte) {
	// Parcourir tous les clients connectés et envoyer le message à chacun
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			// Gestion de l'erreur ou de la déconnexion du client
			client.Close()
			delete(clients, client)
		}
	}
}