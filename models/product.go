package models

type Product struct {
	Name           string         `json:"name"`
	Price          string         `json:"price"`
	CurrentPrice   string         `json:"current_price"`
	FullPrice      string         `json:"full_price"`
	ProductId      int64          `json:"product_id"`
	HasViewerSaved bool           `json:"has_viewer_saved"`
	Description    string         `json:"description"`
	MainImage      ProductImage   `json:"main_image"`
	ThumbnailImage ProductImage   `json:"thumbnail_image"`
	ProductImages  []ProductImage `json:"product_images"`
	ExternalUrl    string         `json:"external_url"`
	CheckoutStyle  string         `json:"checkout_style"`
	ReviewStatus   string         `json:"review_status"`
}
