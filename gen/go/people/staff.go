package people

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"
)

type Staff interface {
	proto.Message

	Print_Something() string
}

func (c *OfficeStaff) Print_Something() string {
	return fmt.Sprintf("Working days: %v", c.WorkingDays)
}

func (c *SaleStaff) Print_Something() string {
	return fmt.Sprintf("Product quantity sold: %v", c.ProductQuantity)
}
