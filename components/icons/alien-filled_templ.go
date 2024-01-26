// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IconAlienFilled() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-alien-filled\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentColor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M12.004 2c4.942 0 8.288 2.503 8.85 6.444a12.884 12.884 0 0 1 -2.163 9.308a11.794 11.794 0 0 1 -3.51 3.356c-1.982 1.19 -4.376 1.19 -6.373 -.008a11.763 11.763 0 0 1 -3.489 -3.34a12.808 12.808 0 0 1 -2.171 -9.306c.564 -3.95 3.91 -6.454 8.856 -6.454zm1.913 14.6a1 1 0 0 0 -1.317 -.517l-.146 .055a1.5 1.5 0 0 1 -1.054 -.055l-.11 -.04a1 1 0 0 0 -.69 1.874a3.5 3.5 0 0 0 2.8 0a1 1 0 0 0 .517 -1.317zm-5.304 -6.39a1 1 0 0 0 -1.32 1.497l2 2l.094 .083a1 1 0 0 0 1.32 -1.497l-2 -2zm8.094 .083a1 1 0 0 0 -1.414 0l-2 2l-.083 .094a1 1 0 0 0 1.497 1.32l2 -2l.083 -.094a1 1 0 0 0 -.083 -1.32z\" stroke-width=\"0\" fill=\"currentColor\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
