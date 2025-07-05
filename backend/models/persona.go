// backend/models/persona.go
package models

import (
	"time"
)

type Persona struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Edad      int       `json:"edad"`
	Genero    string    `json:"genero"` // "masculino", "femenino"
	URLImagen string    `json:"url_imagen"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsLiked   bool      `json:"is_liked"` // <--- ADD THIS FIELD
}