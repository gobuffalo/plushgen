package plushgen

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

// Transformer will plushify any file that has a ".plush" extension
func Transformer(ctx *plush.Context) genny.Transformer {
	t := genny.NewTransformer(".plush", func(f genny.File) (genny.File, error) {
		s, err := plush.RenderR(f, ctx)
		if err != nil {
			return f, errors.Wrap(err, f.Name())
		}
		return genny.NewFileS(f.Name(), s), nil
	})
	t.StripExt = true
	return t
}
