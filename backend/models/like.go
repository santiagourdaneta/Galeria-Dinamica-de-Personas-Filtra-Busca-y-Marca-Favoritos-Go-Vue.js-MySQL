// backend/models/like.go
package models

import (
	"time"
)

type Like struct {
	ID         int       `json:"id"`
	PersonaID  int       `json:"persona_id"`
	UserID     int       `json:"user_id"` // Si implementas autenticación de usuario
	TipoAccion string    `json:"tipo_accion"` // "me_gusta", "no_me_gusta"
	CreatedAt  time.Time `json:"created_at"`
}