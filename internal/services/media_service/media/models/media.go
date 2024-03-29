package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Media model
type Media struct {
	MediaID       uuid.UUID   `json:"mediaId"`       // ID único para o arquivo de mídia
	AuthorID      uuid.UUID   `json:"authorId"`      // ID do autor do arquivo de mídia
	AuthorizedIDs []uuid.UUID `json:"authorizedIds"` // Lista de IDs dos indivíduos autorizados a utilizar o arquivo de mídia
	FileName      string      `json:"fileName"`      // Nome do arquivo de mídia
	FileType      string      `json:"fileType"`      // Tipo de arquivo (pode ser "foto", "vídeo", etc.)
	FileSize      int64       `json:"fileSize"`      // Tamanho do arquivo em bytes
	FilePath      string      `json:"filePath"`      // Caminho do arquivo (caminho relativo ou URL)
	Thumbnail     string      `json:"thumbnail"`     // Caminho para a miniatura do arquivo de mídia
	Permissions   []string    `json:"permissions"`   // Lista de permissões associadas ao arquivo de mídia
	IsPublic      bool        `json:"isPublic"`      // Indica se o arquivo de mídia é público
	IsCompressed  bool        `json:"isCompressed"`  // Indica se o arquivo de mídia está compactado
	Expiration    time.Time   `json:"expiration"`    // Data de expiração do arquivo
	CreatedAt     time.Time   `json:"createdAt"`     // Data de criação do arquivo
	UpdatedAt     time.Time   `json:"updatedAt"`     // Data da última atualização do arquivo
	DeletedAt     *time.Time  `json:"deletedAt"`     // Data de exclusão do arquivo
}
