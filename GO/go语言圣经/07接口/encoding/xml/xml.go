package xml
import "io"

type Name struct {
	Local string //e.g., "Title" or "id"
}

type Attr struct  {
	Name Name
	Value string
}


// A Token includes StartElement, EndElement, CharData,
// and Comment, plus a few esoteric types (not shown).
type Token interface{}
type StartElement struct { // e.g., <name>
	Name Name
	Attr []Attr
}
type EndElement struct { Name Name } // e.g., </name>
type CharData []byte                 // e.g., <p>CharData</p>
type Comment []byte                  // e.g., <!-- Comment -->

type Decoder struct{ /* ... */ }

func NewDecoder(io.Reader) *Decoder

func (*Decoder) Token() (Token, error) // returns next Token in sequence

