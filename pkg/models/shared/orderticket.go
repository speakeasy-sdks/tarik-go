// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type OrderTicket struct {
	// Decision
	Decision string `json:"decision"`
	// Marketplace Order ID
	OrderID int64 `json:"order_id"`
	// Format YYYY-MM-DD[T]HH:MM
	SolutionTime time.Time `json:"solution_time"`
	// Status
	Status OrderTicketStatus `json:"status"`
	// BPM ticket ID
	TicketID string `json:"ticket_id"`
}
