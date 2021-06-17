package warehouse

import (
	"github.com/flosch/pongo2/v4"
	"github.com/segmentio/ksuid"

	_ "github.com/flosch/pongo2-addons"
)

/*
tagNodeKSUID implements the pongo2 interface for tag, named "ksuid".
*/
type tagNodeKSUID struct {
	position *pongo2.Token
}

/*
Execute generates a new KSUID written into the pongo2 template writer.
*/
func (n *tagNodeKSUID) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	writer.WriteString(ksuid.New().String())
	return nil
}

/*
tagKSUID adds the "ksuid" tag to pongo2 templates.
*/
func tagKSUID(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	n := &tagNodeKSUID{
		position: start,
	}

	if arguments.Remaining() > 0 {
		return nil, arguments.Error("Malformed ksuid-tag arguments.", nil)
	}

	return n, nil
}

/*
init adds the filters and tags to pongo2.
*/
func init() {
	pongo2.RegisterTag("ksuid", tagKSUID)
}
