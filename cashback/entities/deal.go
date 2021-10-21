package entities

import "github.com/google/uuid"

type Deal struct {
	ID              uuid.UUID `json:"deal_id"`
	RetailerID      uuid.UUID `json:"retailer_id"`
	RetailerName    string    `json:"retailer_name"`
	RetailerDomains []string  `json:"retailer_domains"`
	Type            string    `json:"deal_type"`
	Amount          float64   `json:"deal_amount"`
}
