package ewsxml

import (
	"encoding/xml"
)

type ContainmentComparison string

func (s ContainmentComparison) String() string { return string(s) }

type ContainmentMode string

func (s ContainmentMode) String() string { return string(s) }

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	// The comparison is between the full string and the constant. The property value and the supplied constant are precisely the same.
	ContainmentMode_FullString ContainmentMode = "FullString"
	// The comparison is between the string prefix and the constant.
	ContainmentMode_Prefixed ContainmentMode = "Prefixed"
	// The comparison is between a substring of the string and the constant.
	ContainmentMode_Substring ContainmentMode = "Substring"
	// The comparison is between a prefix on individual words in the string and the constant.
	ContainmentMode_PrefixOnWords ContainmentMode = "PrefixOnWords"
	// The comparison is between an exact phrase in the string and the constant.
	ContainmentMode_ExactPhrase ContainmentMode = "ExactPhrase"

	// The comparison must be exact.
	ContainmentComparison_Exact ContainmentComparison = "Exact"
	// The comparison ignores casing.
	ContainmentComparison_IgnoreCase ContainmentComparison = "IgnoreCase"
	// The comparison ignores non-spacing characters.
	ContainmentComparison_IgnoreNonSpacingCharacters ContainmentComparison = "IgnoreNonSpacingCharacters"
	// The comparison ignores casing and non-spacing characters.
	ContainmentComparison_IgnoreCaseAndNonSpacingCharacters ContainmentComparison = "IgnoreCaseAndNonSpacingCharacters"
)

type Restriction struct {
	XMLName                xml.Name `xml:"m:Restriction"`
	Exists                 *SearchExpresionExists
	Excludes               *SearchExpresionExcludes
	IsEqualTo              *SearchExpresionIsEqualTo
	IsNotEqualTo           *SearchExpresionIsNotEqualTo
	IsGreaterThan          *SearchExpresionIsGreaterThan
	IsGreaterThanOrEqualTo *SearchExpresionIsGreaterThanOrEqualTo
	IsLessThan             *SearchExpresionIsLessThan
	IsLessThanOrEqualTo    *SearchExpresionIsLessThanOrEqualTo
	Contains               *SearchExpresionContains
	Not                    *SearchExpresion `xml:"t:Not"`
	And                    *SearchExpresion `xml:"t:And"`
	Or                     *SearchExpresion `xml:"t:Or"`
}

type SearchExpresion struct {
	Exists                 []*SearchExpresionExists
	Excludes               []*SearchExpresionExcludes
	IsEqualTo              []*SearchExpresionIsEqualTo
	IsNotEqualTo           []*SearchExpresionIsNotEqualTo
	IsGreaterThan          []*SearchExpresionIsGreaterThan
	IsGreaterThanOrEqualTo []*SearchExpresionIsGreaterThanOrEqualTo
	IsLessThan             []*SearchExpresionIsLessThan
	IsLessThanOrEqualTo    []*SearchExpresionIsLessThanOrEqualTo
	Contains               []*SearchExpresionContains
	Not                    []*SearchExpresion `xml:"t:Not"`
	And                    []*SearchExpresion `xml:"t:And"`
	Or                     []*SearchExpresion `xml:"t:Or"`
}

type SearchExpresionExists struct {
	XMLName          xml.Name `xml:"t:Exists"`
	FieldURI         *FieldURI
	IndexedFieldURI  *IndexedFieldURI
	ExtendedFieldURI *ExtendedFieldURI
}

type SearchExpresionExcludes struct {
	XMLName          xml.Name `xml:"t:Excludes"`
	FieldURI         *FieldURI
	IndexedFieldURI  *IndexedFieldURI
	ExtendedFieldURI *ExtendedFieldURI
	Bitmask          Bitmask
}

type SearchExpresionIsEqualTo struct {
	XMLName            xml.Name `xml:"t:IsEqualTo"`
	FieldURI           *FieldURI
	IndexedFieldURI    *IndexedFieldURI
	ExtendedFieldURI   *ExtendedFieldURI
	FieldURIOrConstant FieldURIOrConstant
}

type SearchExpresionIsNotEqualTo struct {
	XMLName            xml.Name `xml:"t:IsNotEqualTo"`
	FieldURI           *FieldURI
	IndexedFieldURI    *IndexedFieldURI
	ExtendedFieldURI   *ExtendedFieldURI
	FieldURIOrConstant FieldURIOrConstant
}

type SearchExpresionIsGreaterThan struct {
	XMLName            xml.Name `xml:"t:IsGreaterThan"`
	FieldURI           *FieldURI
	IndexedFieldURI    *IndexedFieldURI
	ExtendedFieldURI   *ExtendedFieldURI
	FieldURIOrConstant FieldURIOrConstant
}

type SearchExpresionIsGreaterThanOrEqualTo struct {
	XMLName            xml.Name `xml:"t:IsGreaterThanOrEqualTo"`
	FieldURI           *FieldURI
	IndexedFieldURI    *IndexedFieldURI
	ExtendedFieldURI   *ExtendedFieldURI
	FieldURIOrConstant FieldURIOrConstant
}

type SearchExpresionIsLessThan struct {
	XMLName            xml.Name `xml:"t:IsLessThan"`
	FieldURI           *FieldURI
	IndexedFieldURI    *IndexedFieldURI
	ExtendedFieldURI   *ExtendedFieldURI
	FieldURIOrConstant FieldURIOrConstant
}

type SearchExpresionIsLessThanOrEqualTo struct {
	XMLName            xml.Name `xml:"t:IsLessThanOrEqualTo"`
	FieldURI           *FieldURI
	IndexedFieldURI    *IndexedFieldURI
	ExtendedFieldURI   *ExtendedFieldURI
	FieldURIOrConstant FieldURIOrConstant
}

type SearchExpresionContains struct {
	XMLName               xml.Name `xml:"t:Contains"`
	FieldURI              *FieldURI
	IndexedFieldURI       *IndexedFieldURI
	ExtendedFieldURI      *ExtendedFieldURI
	ContainmentMode       ContainmentMode       `xml:"ContainmentMode,attr,omitempty"`
	ContainmentComparison ContainmentComparison `xml:"ContainmentComparison,attr,omitempty"`
	Constant              Constant
}

type Bitmask struct {
	XMLName xml.Name `xml:"t:Bitmask"`
	Value   string   `xml:"Value,attr"`
}

type Constant struct {
	XMLName xml.Name `xml:"t:Constant"`
	Value   string   `xml:"Value,attr"`
}

type FieldURIOrConstant struct {
	XMLName          xml.Name `xml:"t:FieldURIOrConstant"`
	FieldURI         *FieldURI
	IndexedFieldURI  *IndexedFieldURI
	ExtendedFieldURI *ExtendedFieldURI
	Constant         *Constant
}
