// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IconBrandVivaldi() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-brand-vivaldi\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentColor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M21.648 6.808c-2.468 4.28 -4.937 8.56 -7.408 12.836c-.397 .777 -1.366 1.301 -2.24 1.356c-.962 .102 -1.7 -.402 -2.154 -1.254c-1.563 -2.684 -3.106 -5.374 -4.66 -8.064c-.943 -1.633 -1.891 -3.266 -2.83 -4.905a2.47 2.47 0 0 1 -.06 -2.45a2.493 2.493 0 0 1 2.085 -1.307c.951 -.065 1.85 .438 2.287 1.281c.697 1.19 2.043 3.83 2.55 4.682a3.919 3.919 0 0 0 3.282 2.017c2.126 .133 3.974 -.95 4.21 -3.058c0 -.164 .228 -3.178 .846 -3.962c.619 -.784 1.64 -1.155 2.606 -.893a2.484 2.484 0 0 1 1.814 2.062c.08 .581 -.041 1.171 -.343 1.674\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
