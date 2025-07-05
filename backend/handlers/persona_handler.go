// hot-or-not/backend/handlers/persona_handler.go
package handlers

import (
	"net/http"
    "hot-or-not/backend/config"
    "hot-or-not/backend/models"
    "log"
    "strings" // <--- MAKE SURE THIS IS HERE
    "github.com/gin-gonic/gin"
)

// Re-using DUMMY_USER_ID from like_handler.go or define it globally if shared across handlers
const DUMMY_USER_ID_FOR_QUERY = 1 // Use same dummy ID as in like_handler.go

// GetPersonas obtiene una lista de personas con filtros y búsqueda, y el estado de like
func GetPersonas(c *gin.Context) {
	genero := c.Query("genero")
	search := c.Query("search")

	// Base query
	query := `
		SELECT
			p.id, p.nombre, p.edad, p.genero, p.url_imagen, p.created_at, p.updated_at,
			CASE WHEN ul.user_id IS NOT NULL THEN TRUE ELSE FALSE END AS is_liked
		FROM
			personas p
		LEFT JOIN
			user_likes ul ON p.id = ul.persona_id AND ul.user_id = ?
	` // <<< KEY CHANGE: LEFT JOIN and CASE WHEN for is_liked

	args := []interface{}{DUMMY_USER_ID_FOR_QUERY} // Add user ID to args for the JOIN condition
	whereClauses := []string{}

	if genero != "" {
		whereClauses = append(whereClauses, "p.genero = ?")
		args = append(args, genero)
	}

	if search != "" {
		whereClauses = append(whereClauses, "LOWER(p.nombre) LIKE ?") // Using LOWER for case-insensitive search
		args = append(args, "%"+strings.ToLower(search)+"%")
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " ORDER BY p.id ASC" // Use p.id to specify column from personas table

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener personas", "details": err.Error()})
		return
	}
	defer rows.Close()

	var personas []models.Persona
	for rows.Next() {
		var p models.Persona
		// Scan into the new IsLiked field as well
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Edad, &p.Genero, &p.URLImagen, &p.CreatedAt, &p.UpdatedAt, &p.IsLiked); err != nil {
			log.Printf("Error al escanear fila de persona: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear persona", "details": err.Error()})
			return
		}
		personas = append(personas, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error durante la iteración de filas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar resultados", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personas)
}
// LikePersona registra un "me gusta" para una persona
func LikePersona(c *gin.Context) {
	personaID := c.Param("id")

	// Implementar lógica para guardar en la tabla 'likes'
	// Necesitarías el ID del usuario (si hay autenticación)
	// Ejemplo simplificado:
	stmt, err := config.DB.Prepare("INSERT INTO likes (persona_id, tipo_accion, user_id) VALUES (?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparando la declaración SQL", "details": err.Error()})
		return
	}
	defer stmt.Close()

	// Asume user_id 1 por ahora, reemplaza con el ID del usuario autenticado
	_, err = stmt.Exec(personaID, "me_gusta", 1) // Using a placeholder user ID for now
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registrando 'me gusta'", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Me gusta registrado exitosamente"})
}

// DislikePersona registra un "no me gusta" para una persona
func DislikePersona(c *gin.Context) {
	personaID := c.Param("id")

	stmt, err := config.DB.Prepare("INSERT INTO likes (persona_id, tipo_accion, user_id) VALUES (?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparando la declaración SQL", "details": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(personaID, "no_me_gusta", 1) // Using a placeholder user ID for now
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registrando 'no me gusta'", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "No me gusta registrado exitosamente"})
}