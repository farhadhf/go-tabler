// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IconChartGridDotsFilled() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-chart-grid-dots-filled\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentColor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M18 2a1 1 0 0 1 1 1v.171a3.008 3.008 0 0 1 1.83 1.83l.17 -.001a1 1 0 0 1 0 2h-.171a3.008 3.008 0 0 1 -1.828 1.829l-.001 2.171h2a1 1 0 0 1 0 2h-2v2.171a3.008 3.008 0 0 1 1.83 1.83l.17 -.001a1 1 0 0 1 0 2h-.171a3.008 3.008 0 0 1 -1.828 1.829l-.001 .171a1 1 0 0 1 -2 0v-.17a3.008 3.008 0 0 1 -1.829 -1.83h-2.171v2a1 1 0 0 1 -2 0v-2h-2.171a3.008 3.008 0 0 1 -1.828 1.829l-.001 .171a1 1 0 0 1 -2 0v-.17a3.008 3.008 0 0 1 -1.829 -1.83h-.171a1 1 0 0 1 0 -2h.17a3.008 3.008 0 0 1 1.83 -1.83v-.34a3.008 3.008 0 0 1 -1.829 -1.83h-.171a1 1 0 0 1 0 -2h.17a3.008 3.008 0 0 1 1.83 -1.83v-2.17h-2a1 1 0 1 1 0 -2h2v-2a1 1 0 1 1 2 0v2h4v-2a1 1 0 0 1 2 0v2h2.17a3.008 3.008 0 0 1 1.83 -1.83v-.17a1 1 0 0 1 1 -1zm-7 11h-2.171a3.008 3.008 0 0 1 -1.828 1.829v.342a3.008 3.008 0 0 1 1.828 1.829h2.171v-4zm6 0h-4v4h2.17a3.008 3.008 0 0 1 1.83 -1.83v-2.17zm-6 -6h-4v2.171a3.008 3.008 0 0 1 1.83 1.83l2.17 -.001v-4zm4.171 0h-2.171v4h4v-2.17a3.008 3.008 0 0 1 -1.829 -1.83z\" stroke-width=\"0\" fill=\"currentColor\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
