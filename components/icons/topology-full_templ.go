// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IconTopologyFull() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-topology-full\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentColor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M20 18a2 2 0 1 0 -4 0a2 2 0 0 0 4 0z\"></path> <path d=\"M8 18a2 2 0 1 0 -4 0a2 2 0 0 0 4 0z\"></path> <path d=\"M8 6a2 2 0 1 0 -4 0a2 2 0 0 0 4 0z\"></path> <path d=\"M20 6a2 2 0 1 0 -4 0a2 2 0 0 0 4 0z\"></path> <path d=\"M6 8v8\"></path> <path d=\"M18 16v-8\"></path> <path d=\"M8 6h8\"></path> <path d=\"M16 18h-8\"></path> <path d=\"M7.5 7.5l9 9\"></path> <path d=\"M7.5 16.5l9 -9\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
