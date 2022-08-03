/* Package metadata implements types for working with the Chia NFT Metadata Specification
(CHIP-0007: https://github.com/Chia-Network/chips/blob/metadata-schema/CHIPs/chip-will-riches-metadata.md),
in Go (https://golang.org).

This package is part of the broader set of packages at https://github.com/Jsewill/chia/
*/
package metadata

// Metadata implements types and methods for working with the chia NFT metadata specification draft.
type Metadata struct {
	Format           string            `json:"format"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	MintingTool      string            `json:"minting_tool"`
	SensitiveContent *SensitiveContent `json:"sensitive_content"`
	EditionNumber    uint              `json:"edition_number"` // Not in CHIP-0007, but in chia 1.4.0 as "series_number", rather than "edition_total", due to bug. This was fixed in chia 1.5.0.
	EditionTotal     uint              `json:"edition_total"`  // Not in CHIP-0007, but in chia 1.4.0 as "series_total", rather than "edition_total", due to bug. This was fixed in chia 1.5.0.
	SeriesNumber     uint              `json:"series_number"`  // In CHIP-0007, but not in chia 1.4.0 as "series_number", due to bug. This has been deprecated as of chia 1.5.0, likely until NFT2.
	SeriesTotal      uint              `json:"series_total"`   // In CHIP-0007, but not in chia 1.4.0 as "series_total", due to bug.. This has been deprecated as of chia 1.5.0, likely until NFT2.
	Attributes       []*Attribute      `json:"attributes"`
	Collection       *Collection       `json:"collection"`
	Data             any               `json:"data"`
}

// Attribute retrives an attribute by type, from the Attribute slice. Returns an empty *Attribute if not found.
func (m *Metadata) Attribute(t string) *Attribute {
	for _, a := range m.Attributes {
		if a.Type == t {
			return a
		}
	}
	return &Attribute{}
}
