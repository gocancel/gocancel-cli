package displayers

import (
	"encoding/json"
	"io"

	"github.com/gocancel/gocancel-go"
)

type Products []*gocancel.Product

var _ Displayable = (*Products)(nil)

func (o Products) Cols() []string {
	return []string{
		"ID",
		"Name",
		"Slug",
		"Email",
		"URL",
		"Phone",
		"Fax",
		"OrganizationID",
		"RequiresConsent",
		"RequiresProofOfID",
		"Metadata",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (o Products) ColMap() map[string]string {
	return map[string]string{
		"ID":                "ID",
		"Name":              "Name",
		"Slug":              "Slug",
		"Email":             "Email",
		"URL":               "URL",
		"Phone":             "Phone",
		"Fax":               "Fax",
		"OrganizationID":    "Organization ID",
		"RequiresConsent":   "Requires Consent",
		"RequiresProofOfID": "Requires Proof Of ID",
		"Metadata":          "Metadata",
		"CreatedAt":         "Created At",
		"UpdatedAt":         "Updated At",
	}
}

func (o Products) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, len(o))

	for i, product := range o {
		out[i] = map[string]interface{}{
			"ID":                *product.ID,
			"Name":              *product.Name,
			"Slug":              *product.Slug,
			"Email":             *product.Email,
			"URL":               *product.URL,
			"Phone":             *product.Phone,
			"Fax":               *product.Fax,
			"OrganizationID":    *product.OrganizationID,
			"RequiresConsent":   *product.RequiresConsent,
			"RequiresProofOfID": *product.RequiresProofOfID,
			"Metadata":          *product.Metadata,
			"CreatedAt":         product.CreatedAt,
			"UpdatedAt":         product.UpdatedAt,
		}
	}

	return out
}

func (o Products) JSON(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(o)
}
