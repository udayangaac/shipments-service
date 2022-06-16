package entity

// Shipment stores the details related to the shipments.
type Shipment struct {
	TrackingID   string `gorm:"primaryKey"`
	OriginPort   *string
	DSTPort      *string
	NumberOfPkgs int
	ConsigneeNum uint
	Consignee    Consignee `gorm:"foreignKey:ConsigneeNum"`
	ShipperNum uint
	Shipper      Shipper   `gorm:"foreignKey:ShipperNum"`
	InvoiceNum uint
	Invoice      Invoice   `gorm:"foreignKey:InvoiceNum"`
}

// Consignee stores the details related to the consignee.
type Consignee struct {
	Num       uint    `gorm:"primaryKey;column:CONSIGNEE_NUM"`
	Name      *string `gorm:"type:varchar(100);column:CONSIGNEE_NAME"`
	Addr      *string `gorm:"type:varchar(300);column:CONSIGNEE_ADDR"`
	ContactNo *string `gorm:"type:varchar(100);column:CONSIGNEE_CONTACT_NO"`
}

// Shipper stores the details related to the shipper.
type Shipper struct {
	Num       uint    `gorm:"primaryKey;column:SHIPPER_NUM"`
	Name      *string `gorm:"type:varchar(100);column:SHIPPER_NAME"`
	Addr      *string `gorm:"type:varchar(300);column:SHIPPER_ADDR"`
	ContactNo *string `gorm:"type:varchar(100);column:SHIPPER_CONTACT_NO"`
}

// Invoice stores the details related to the invoice.
type Invoice struct {
	Num      uint     `gorm:"primaryKey;column:INVOICE_NAME"`
	Amount   float32 `gorm:"column:INVOICE_AMOUNT"`
	Currency string  `gorm:"type:varchar(10);column:INVOICE_CURRENCY"`
}