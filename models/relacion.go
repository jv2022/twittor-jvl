package models

/*
Relacion, modelo para grabar la relación de un usuario con otro.
*/
type Relacion struct {
	// ID del usuario
	UsuarioID string `bson:"usuarioid" json:"usuarioId"`
	// ID del usuario que se está siguiendo
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
