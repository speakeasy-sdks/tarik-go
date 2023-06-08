// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Customer - Represents info about order customer
type Customer struct {
	Email string `json:"email"`
	// This field could be numeric string
	ErpID     int64  `json:"erp_id"`
	FirstName string `json:"first_name"`
	// This field could be numeric string
	HumanID int64 `json:"human_id"`
	// This field could be numeric string
	ID         int64  `json:"id"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
}
