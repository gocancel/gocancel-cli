package displayers

import (
	"encoding/json"
	"io"

	"github.com/gocancel/gocancel-go"
)

type Organizations []*gocancel.Organization

var _ Displayable = (*Organizations)(nil)

func (o Organizations) Cols() []string {
	return []string{
		"ID",
		"Name",
		"Slug",
		"Email",
		"URL",
		"Phone",
		"Fax",
		"CategoryID",
		"RequiresConsent",
		"RequiresProofOfID",
		"Metadata",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (o Organizations) ColMap() map[string]string {
	return map[string]string{
		"ID":                "ID",
		"Name":              "Name",
		"Slug":              "Slug",
		"Email":             "Email",
		"URL":               "URL",
		"Phone":             "Phone",
		"Fax":               "Fax",
		"CategoryID":        "Category ID",
		"RequiresConsent":   "Requires Consent",
		"RequiresProofOfID": "Requires Proof Of ID",
		"Metadata":          "Metadata",
		"CreatedAt":         "Created At",
		"UpdatedAt":         "Updated At",
	}
}

func (o Organizations) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, len(o))

	for i, organization := range o {
		out[i] = map[string]interface{}{
			"ID":                organization.ID,
			"Name":              organization.Name,
			"Slug":              organization.Slug,
			"Email":             organization.Email,
			"URL":               organization.URL,
			"Phone":             organization.Phone,
			"Fax":               organization.Fax,
			"CategoryID":        organization.CategoryID,
			"RequiresConsent":   organization.RequiresConsent,
			"RequiresProofOfID": organization.RequiresProofOfID,
			"Metadata":          organization.Metadata,
			"CreatedAt":         organization.CreatedAt,
			"UpdatedAt":         organization.UpdatedAt,
		}
	}

	return out
}

func (o Organizations) JSON(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(o)
}
