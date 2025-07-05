package handlers

import (
	
	"net/http"
	"strconv"
	"log"

	"hot-or-not/backend/config"
	"github.com/gin-gonic/gin"
)

// Define a dummy user ID for now. In a real app, this would come from authentication.
const DUMMY_USER_ID = 1

// ToggleLike handles liking and unliking a persona
func ToggleLike(c *gin.Context) {
	personaIDStr := c.Param("id")
	personaID, err := strconv.Atoi(personaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de persona inválido"})
		return
	}

	// For now, use a dummy user ID. In a real app, you'd get this from context (e.g., JWT).
	userID := DUMMY_USER_ID

	var exists bool
	// Check if the like already exists
	err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_likes WHERE user_id = ? AND persona_id = ?)", userID, personaID).Scan(&exists)
	if err != nil {
		log.Printf("Error al verificar like existente: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	if exists {
		// If like exists, remove it (unlike)
		_, err = config.DB.Exec("DELETE FROM user_likes WHERE user_id = ? AND persona_id = ?", userID, personaID)
		if err != nil {
			log.Printf("Error al eliminar like: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar 'Me gusta'"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Like eliminado", "is_liked": false})
	} else {
		// If like does not exist, add it (like)
		_, err = config.DB.Exec("INSERT INTO user_likes (user_id, persona_id) VALUES (?, ?)", userID, personaID)
		if err != nil {
			log.Printf("Error al añadir like: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al añadir 'Me gusta'"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Like añadido", "is_liked": true})
	}
}