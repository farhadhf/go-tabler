// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IconSquareF0Filled() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-square-f0-filled\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentColor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M18.333 2c1.96 0 3.56 1.537 3.662 3.472l.005 .195v12.666c0 1.96 -1.537 3.56 -3.472 3.662l-.195 .005h-12.666a3.667 3.667 0 0 1 -3.662 -3.472l-.005 -.195v-12.666c0 -1.96 1.537 -3.56 3.472 -3.662l.195 -.005h12.666zm-3.833 6a2.5 2.5 0 0 0 -2.495 2.336l-.005 .164v3l.005 .164a2.5 2.5 0 0 0 4.99 0l.005 -.164v-3l-.005 -.164a2.5 2.5 0 0 0 -2.495 -2.336zm-4.5 0h-2l-.117 .007a1 1 0 0 0 -.876 .876l-.007 .117v6l.007 .117a1 1 0 0 0 .876 .876l.117 .007l.117 -.007a1 1 0 0 0 .876 -.876l.007 -.117v-2h1l.117 -.007a1 1 0 0 0 0 -1.986l-.117 -.007h-1v-1h1l.117 -.007a1 1 0 0 0 0 -1.986l-.117 -.007zm4.5 2a.5 .5 0 0 1 .492 .41l.008 .09v3l-.008 .09a.5 .5 0 0 1 -.984 0l-.008 -.09v-3l.008 -.09a.5 .5 0 0 1 .492 -.41z\" stroke-width=\"0\" fill=\"currentColor\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
