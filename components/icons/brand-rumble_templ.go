// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IconBrandRumble() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-brand-rumble\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentColor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M19.993 9.108c.383 .4 .687 .863 .893 1.368a4.195 4.195 0 0 1 .006 3.166a4.37 4.37 0 0 1 -.887 1.372a20.233 20.233 0 0 1 -2.208 2a20.615 20.615 0 0 1 -2.495 1.669a21.322 21.322 0 0 1 -5.622 2.202a4.213 4.213 0 0 1 -3.002 -.404a3.98 3.98 0 0 1 -1.163 -.967a3.796 3.796 0 0 1 -.695 -1.312c-1.199 -3.902 -1.022 -8.312 .134 -12.23c.609 -2.057 2.643 -3.349 4.737 -2.874c3.88 .88 7.52 3.147 10.302 6.01z\"></path> <path d=\"M14.044 13.034c.67 -.505 .67 -1.489 0 -2.01a14.824 14.824 0 0 0 -1.498 -1.044a15.783 15.783 0 0 0 -1.62 -.865c-.77 -.35 -1.63 .139 -1.753 .973a15.385 15.385 0 0 0 -.1 3.786a1.232 1.232 0 0 0 1.715 1.027a14.783 14.783 0 0 0 1.694 -.827a14.46 14.46 0 0 0 1.562 -1.035v-.005z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
