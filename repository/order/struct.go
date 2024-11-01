package order

import (
	"github.com/funmi4194/instashop/enum"
	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel `bun:"table:orders" rsf:"false"`

	ID     string `bun:"id,pk" json:"id"`
	UserID string `bun:"user_id" json:"user_id"`

	// items that make up the total amount
	Price      float64            `bun:"price" json:"price"`
	Stock      int64              `bun:"stock" json:"stock"`
	ProductUrl string             `bun:"product_url" json:"product_url"`
	Status     enum.ProductStatus `bun:"status" json:"status" rsfr:"false"`
	Currency   enum.Currency      `bun:"currency" json:"currency"`
	Paid       bool               `bun:"paid" json:"paid"`
	PaidAt     bun.NullTime       `bun:"paid_at" json:"paid_at"`
	Failed     bool               `bun:"failed" json:"failed"`
	FailedAt   bun.NullTime       `bun:"failed_at" json:"failed_at"`
	// this will most not be used (actually, don't use it)
	Cancelled   bool         `bun:"cancelled" json:"cancelled"`
	CancelledAt bun.NullTime `bun:"cancelled_at" json:"cancelled_at"`
	// SHA256 of the transaction.Invoice (prevents tampering & duplication)
	Checksum string `bun:"checksum" json:"checksum"`

	// the total amount to be paid (including all possible fees)
	Amount float64 `bun:"amount" json:"amount"`

	Remark string `bun:"remark" json:"remark"`

	AssetID string `bun:"asset_id" json:"asset_id"`

	Description string       `bun:"description" json:"description"`
	CreatedAt   bun.NullTime `bun:"created_at" json:"created_at" rsfr:"false"`
	UpdatedAt   bun.NullTime `bun:"updated_at" json:"updated_at" rsfr:"false"`
}

// schematic representation of an item in a order's invoice
type Item struct {
	// the key of the item
	Key string `json:"key"`

	// the name of the item
	Name string `json:"name"`

	// the amount of the item
	Amount float64 `json:"amount"`

	// the quantity of the item
	Quantity int `json:"quantity"`

	// the metadata for the item (can also be a json string)
	Metadata string `json:"metadata"`
}
