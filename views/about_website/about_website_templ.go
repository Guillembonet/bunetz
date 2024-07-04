// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package about_website

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AboutWebsite() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container text-xl px-6 md:mx-auto mt-4 dark:text-white\"><h1 class=\"text-4xl mb-6\">About this website</h1><h2 class=\"text-3xl mb-6\">The website</h2><p class=\"mb-4\">This website is server-side rendered using Go and the templ library, which is a Go package that allows you to create websites using templates. I also used HTMX to make the website blazing fast and more interactive. I find it is a very interesting project which takes web development back to its roots, emphasizing on simplicity, and getting rid of overly complicated frontend frameworks, which are good, but oftentimes really unnecessary. It pairs perfectly with Go, which is focused on simplicity as well, as all things in life should be.</p><p class=\"mb-4\">This is why loading this website for the first time costed you less than 100kB of data, and the next time you visit it, it will be even less (&lt10kB) thanks to the cache. What if you now switch to  <a href=\"/about-me\" hx-boost=\"true\" hx-target=\"#content\" hx-swap=\"innerHTML show:window:top\" class=\"text-sky-600 dark:text-cyan-400\">About me</a>? It will only get the necessary fragment and will take around 1kB. I think you get the idea...</p><p class=\"mb-6\">You can find the source code of this website on <a href=\"https://github.com/guillembonet/bunetz\" class=\"text-sky-600 dark:text-cyan-400 inline-block\">Github</a>, for now it is private, but I will make it public when I am done with the first version.</p><h2 class=\"text-3xl mb-6\">Hosting it</h2><p class=\"mb-4\">The hosting of this website is also a bit special. This website is hosted on my 2-Raspberry Pi cluster (k3s) at home, but the connection is proxied through a VPS to hide my home IP and avoid port-forwards on my router. I wrote a blog post about this whole setup which you can find in <a href=\"/blog/posts/how-i-over-engineered-my-cluster-part-1\" hx-boost=\"true\" hx-target=\"#content\" hx-swap=\"innerHTML show:window:top\" class=\"text-sky-600 dark:text-cyan-400\">my blog</a> (available soon).</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
