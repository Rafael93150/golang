Pour run le projet :
    - Ouvrez deux terminal:
    Un dans lequel vous lancez la commande go run main.go à la racine du projet, l'autre dans lequel vous allez tester (requetes ci-dessous) 

## Fonctionalités faites pour le projet
1. Création de la structure de base pour l'API de chat en temps réel.
2. Implémentation de la gestion des utilisateurs :
 1. Ajouter un utilisateur
 2. Supprimer un utilisateur
 3. Authentification des utilisateurs avec JWT
3. Implémentation de la gestion des salons :
 1. Créer un salon
 2. Supprimer un salon
4. Implémentation de la gestion des messages :
 1. Envoyer un message à tous les utilisateurs d'un salon
5. Implémentation de l'API de chat en temps réel :
 1. Gérer les utilisateurs (CRUD + Login)
 2. Gérer les salons en utilisant les routes HTTP : /rooms/create, /rooms/delete
 3. Permettre aux utilisateurs de joindre un salon en utilisant la route HTTP : /rooms/join

-----------------------------------------------------------------------------------------------------------
                     Implémentation de la gestion des utilisateurs :
-----------------------------------------------------------------------------------------------------------

---  Implémentation de la gestion des utilisateurs : 

curl -X GET http://localhost:8000/api/users -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODkwNDk2OTcsInN1YiI6ImpvaG4ifQ.ocxIpYYSP6bnyblhPJJq9RTzRRK9XZocSOkTcvIzunY"
pour récupérér la liste des utilisateurs crées

    -- Ajout d'un utilisateur (deux utilisateurs différents doivent avoir des usernames différents)
Exemple de requête :
curl -X POST -H "Content-Type: application/json" -d '{"username":"Asmaa", "password": "test123"}' http://localhost:8000/api/users

    -- Suppression d'un utilisateur
Exemple de requête :
curl -X DELETE http://localhost:8000/api/users/0

    -- Authentification avec JWT :
Pour utiliser l'authentification avec JSON Web Token (JWT), suivez les étapes ci-dessous :

1. Ajoutez un utilisateur à l'API en utilisant la route POST /api/users.
Exemple de requête :
curl -X POST -H "Content-Type: application/json" -d '{"username":"Asmaa", "password":"password123"}' http://localhost:8000/api/users

2. Pour obtenir un token JWT, effectuez une demande de connexion en utilisant la route POST /api/login. Cette requête doit inclure les informations d'identification de l'utilisateur (nom d'utilisateur et mot de passe) dans le corps de la requête.
Exemple de requête :
curl -X POST -H "Content-Type: application/json" -d '{"username":"Asmaa", "password":"password123"}' http://localhost:8000/api/login


3. Si les informations d'identification sont valides, l'API générera un token JWT qui sera renvoyé dans la réponse. Le token sera utilisé pour authentifier les requêtes ultérieures.
 Exemple de réponse :
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODkwNTA4NjIsInN1YiI6IkFzbWFhIn0.7Ttq1VYIzKIMp5kVLAJPfFP9595KiT3d6JBjL80aryo"}

4. Pour tester l'authentification avec le token JWT, effectuez une autre requête vers une route nécessitant une authentification, par exemple, la route GET /api/users. Vous devez inclure le token JWT dans l'en-tête Authorization de la requête avec la valeur "Bearer <token>". 
Exemple de requête :
curl -X GET http://localhost:8000/api/users -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODkwNTA4NjIsInN1YiI6IkFzbWFhIn0.7Ttq1VYIzKIMp5kVLAJPfFP9595KiT3d6JBjL80aryo"

Faut remplacer "<token>" par la valeur réelle du token JWT généré lors de la demande de connexion.

Note : Le token JWT a une durée de validité limitée. Après expiration, il ne sera plus accepté pour l'authentification des requêtes.


-----------------------------------------------------------------------------------------------------------
                    Implémentation de la gestion des salons :
-----------------------------------------------------------------------------------------------------------

Créer un salon : Nous allons ajouter une fonctionnalité pour créer un nouveau salon dans l'API. Cela peut être fait en envoyant une requête HTTP POST à la route "/api/rooms/create" avec les détails du salon dans le corps de la requête. Le serveur doit enregistrer le salon dans une source de données appropriée, comme une base de données.
Exemple de requête :
curl -X POST -H "Content-Type: application/json" -d '{"name":"Asma1 room"}' http://localhost:8000/api/rooms/create

Supprimer un salon : Nous allons également ajouter la possibilité de supprimer un salon existant de l'API. Cela peut être fait en envoyant une requête HTTP DELETE à la route "/api/rooms/{id}", où "{id}" est l'identifiant unique du salon à supprimer. Le serveur doit rechercher le salon correspondant dans la source de données et le supprimer.
Exemple de requête :

curl -X DELETE http://localhost:8000/api/rooms/1


-----------------------------------------------------------------------------------------------------------
                   Implémentation de la gestion des messages.
-----------------------------------------------------------------------------------------------------------

Implémentation de la gestion des messages :
Envoyer un message à tous les utilisateurs d'un salon avec une requête HTTP POST à la route "/api/rooms/{id}/messages", où "{id}" est l'identifiant unique du salon. Le corps de la requête doit contenir les détails du message à envoyer. Le serveur doit rechercher le salon correspondant dans la source de données, puis envoyer le message à tous les utilisateurs de ce salon.
Exemple de requête :
curl -X POST -H "Content-Type: application/json" -d '{"content":"Bonjour à tous !"}' http://localhost:8000/api/rooms/0/messages
Assure-toi de remplacer les valeurs d'ID et les autres détails dans les exemples de requête par les valeurs réelles correspondantes de ton application.


-----------------------------------------------------------------------------------------------------------
                   Rejoindre un salon
-----------------------------------------------------------------------------------------------------------

curl -X POST http://localhost:8000/api/rooms/0/join/1

-- Tester les utilisateurs présents dans un sallon
curl -X GET http://localhost:8000/api/rooms/0/users

reste to do :
- Permettre aux utilisateurs d'envoyer un message dans un salon en utilisant la route HTTP :
/rooms/message
Mettre en place du SSE pour la communication en temps réel entre les utilisateurs et les salons.

