package models

type CfdiRequest struct {
	CfdiType        string            `json:"CfdiType" example:"I"`
	PaymentForm     string            `json:"PaymentForm" example:"01"`
	PaymentMethod   string            `json:"PaymentMethod" example:"PUE"`
	ExpeditionPlace string            `json:"ExpeditionPlace" example:"20160"`
	GlobalInformation *GlobalInformation `json:"GlobalInformation,omitempty"`
	Receiver        Receiver          `json:"Receiver"`
	Items           []Item            `json:"Items"`
}

type GlobalInformation struct {
	Periodicity string `json:"Periodicity" example:"04"` // Mensual
	Months      string `json:"Months" example:"07"`
	Year        int    `json:"Year" example:"2025"`
}

type Receiver struct {
	Rfc          string `json:"Rfc" example:"XAXX010101000"`
	CfdiUse      string `json:"CfdiUse" example:"S01"`
	Name         string `json:"Name" example:"PUBLICO EN GENERAL"`
	FiscalRegime string `json:"FiscalRegime" example:"616"`
	TaxZipCode   string `json:"TaxZipCode" example:"20160"`
}

type Item struct {
	ProductCode string  `json:"ProductCode" example:"31162800"`
	Description string  `json:"Description" example:"Ventas mes de Julio"`
	UnitCode    string  `json:"UnitCode" example:"AS"`
	Unit        string  `json:"Unit" example:"Variedad"`
	Quantity    float64 `json:"Quantity" example:"1.0"`
	UnitPrice   float64 `json:"UnitPrice" example:"8767.24"`
	Subtotal    float64 `json:"Subtotal" example:"8767.24"`
	TaxObject   string  `json:"TaxObject" example:"02"`
	Taxes       []Tax   `json:"Taxes"`
	Total       float64 `json:"Total" example:"10170.00"`
}

type Tax struct {
	Name         string  `json:"Name" example:"IVA"`
	Rate         float64 `json:"Rate" example:"0.16"`
	Base         float64 `json:"Base" example:"8767.24"`
	Total        float64 `json:"Total" example:"1402.76"`
	IsRetention  bool    `json:"IsRetention" example:"false"`
	IsFederalTax bool    `json:"IsFederalTax" example:"true"`
}
