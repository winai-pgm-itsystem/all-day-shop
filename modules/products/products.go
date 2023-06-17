package products

import (
	"github.com/winai-pgm-itsystem/all-day-shop/modules/appinfo"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/entities"
)

type Product struct {
	Id          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Category    *appinfo.Category `json:"category"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	Price       float64           `json:"price"`
	Images      []*entities.Image `json:"images"`
}
