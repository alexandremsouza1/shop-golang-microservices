package queries

import (
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
)

// Ref: https://golangbot.com/inheritance/

type GetMedias struct {
	*utils.ListQuery
}

func NewGetMedias(query *utils.ListQuery) *GetMedias {
	return &GetMedias{ListQuery: query}
}
