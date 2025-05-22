package productmodel

import (
	"github.com/ezzycreative1/toko-kue-system/internal/core/models/helperModel"
)

type Products struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	MarginPercent int    `json:"margin_percent"`
	helperModel.DateAuditModel
	helperModel.UserAuditModel
}

func (p *Products) TableName() string {
	return "public.tmst_products"
}
