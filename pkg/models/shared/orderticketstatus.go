// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OrderTicketStatus - An enumeration.
type OrderTicketStatus string

const (
	OrderTicketStatusAe5f2f10F46b1410Fd9a0050ba5d6c38                                  OrderTicketStatus = "ae5f2f10-f46b-1410-fd9a-0050ba5d6c38"
	OrderTicketStatusSevene9f1204F46b1410Fb9a0050ba5d6c38                              OrderTicketStatus = "7e9f1204-f46b-1410-fb9a-0050ba5d6c38"
	OrderTicketStatusThreee7f420cF46b1410Fc9a0050ba5d6c38                              OrderTicketStatus = "3e7f420c-f46b-1410-fc9a-0050ba5d6c38"
	OrderTicketStatusAe7f411eF46b1410009b0050ba5d6c38                                  OrderTicketStatus = "ae7f411e-f46b-1410-009b-0050ba5d6c38"
	OrderTicketStatusThreeThousandEightHundredAndFiftyNinec6e7Cbcb486bBa5377808fe6e593 OrderTicketStatus = "3859c6e7-cbcb-486b-ba53-77808fe6e593"
	OrderTicketStatusSixe5f4218F46b1410Fe9a0050ba5d6c38                                OrderTicketStatus = "6e5f4218-f46b-1410-fe9a-0050ba5d6c38"
)

func (e OrderTicketStatus) ToPointer() *OrderTicketStatus {
	return &e
}

func (e *OrderTicketStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ae5f2f10-f46b-1410-fd9a-0050ba5d6c38":
		fallthrough
	case "7e9f1204-f46b-1410-fb9a-0050ba5d6c38":
		fallthrough
	case "3e7f420c-f46b-1410-fc9a-0050ba5d6c38":
		fallthrough
	case "ae7f411e-f46b-1410-009b-0050ba5d6c38":
		fallthrough
	case "3859c6e7-cbcb-486b-ba53-77808fe6e593":
		fallthrough
	case "6e5f4218-f46b-1410-fe9a-0050ba5d6c38":
		*e = OrderTicketStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderTicketStatus: %v", v)
	}
}
