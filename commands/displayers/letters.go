package displayers

import (
	"encoding/json"
	"io"

	"github.com/gocancel/gocancel-go"
)

type Letters []*gocancel.Letter

var _ Displayable = (*Letters)(nil)

func (l Letters) Cols() []string {
	return []string{
		"ID",
		"AccountID",
		"OrganizationID",
		"OrganizationName",
		"ProductID",
		"ProductName",
		"ProviderID",
		"Locale",
		"State",
		"Parameters",
		"SignatureType",
		"SignatureData",
		"Metadata",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (l Letters) ColMap() map[string]string {
	return map[string]string{
		"ID":               "ID",
		"AccountID":        "Account ID",
		"OrganizationID":   "Organization ID",
		"OrganizationName": "Organization Name",
		"ProductID":        "Product ID",
		"ProductName":      "Product Name",
		"ProviderID":       "Provider ID",
		"Locale":           "Locale",
		"State":            "State",
		"Parameters":       "Parameters",
		"SignatureType":    "Signature Type",
		"SignatureData":    "Signature Data",
		"Metadata":         "Metadata",
		"CreatedAt":        "Created At",
		"UpdatedAt":        "Updated At",
	}
}

func (l Letters) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, len(l))

	for i, letter := range l {
		out[i] = map[string]interface{}{
			"ID":               letter.ID,
			"AccountID":        letter.AccountID,
			"OrganizationID":   letter.OrganizationID,
			"OrganizationName": letter.OrganizationName,
			"ProductID":        letter.ProductID,
			"ProductName":      letter.ProductName,
			"ProviderID":       letter.ProviderID,
			"Locale":           letter.Locale,
			"State":            letter.State,
			"Parameters":       letter.Parameters,
			"SignatureType":    letter.SignatureType,
			"SignatureData":    letter.SignatureData,
			"Metadata":         letter.Metadata,
			"CreatedAt":        letter.CreatedAt,
			"UpdatedAt":        letter.UpdatedAt,
		}
	}

	return out
}

func (l Letters) JSON(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(l)
}
