// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package about_me

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AboutMe() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container text-xl px-6 md:mx-auto mt-4 dark:text-white\"><h1 class=\"text-4xl mb-6\">About me</h1><p class=\"mb-4\">Hi, I'm <b>Guillem Bonet</b> and I am a Software Engineer from Barcelona. I mainly have experience programming in Go, but I also have some frontend knowledge (Vue.js, React, ...).</p><p class=\"mb-4\">I created this blog to learn more about creating and hosting a website on my own, in the <a href=\"/about-this-website\" class=\"text-sky-600 dark:text-cyan-400\" hx-boost=\"true\" hx-target=\"#content\" hx-swap=\"innerHTML show:window:top\">About this website</a> section you can find more information about how it is built and deployed. Here, I intend to share some interesting projects or thoughts which I think someone might find interesting.</p><p class=\"mb-4\">In this website you will not find any ads, trackers or cookies. This is a vintage website, without overly complicated designs, no bloated javascript libraries, just a plain old website.</p><p class=\"mb-4\">To contact me use the form below, message me on <a href=\"https://www.linkedin.com/in/guillembonet\" class=\"text-sky-600 dark:text-cyan-400\">LinkedIn</a>, or send an email to <b><u>hello</u></b> at <b><u>bunetz.dev</u></b>.</p><div class=\"w-full flex justify-center my-20\"><div class=\"lg:w-1/2 lg:mx-6\"><div class=\"w-full px-8 py-10 mx-auto overflow-hidden bg-slate-100 rounded-lg shadow-2xl dark:bg-gray-900 lg:max-w-xl shadow-gray-300/50 dark:shadow-black/50\"><h2 class=\"text-xl font-bold\">Contact me</h2><form hx-post=\"/contact\" class=\"mt-6\"><div class=\"flex-1\"><label class=\"block mb-2 text-md text-gray-600 dark:text-gray-200\">Your Name</label> <input name=\"name\" type=\"text\" class=\"block w-full px-5 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-slate-100 border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40\"></div><div class=\"flex-1 mt-6\"><label class=\"block mb-2 text-md text-gray-600 dark:text-gray-200\">Email (or other contact form)</label> <input name=\"contact\" type=\"text\" class=\"block w-full px-5 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-slate-100 border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40\"></div><div class=\"w-full mt-6\"><label class=\"block mb-2 text-md text-gray-600 dark:text-gray-200\">Message</label> <textarea name=\"message\" class=\"block w-full h-32 px-5 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-slate-100 border border-gray-200 rounded-md md:h-48 dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40\"></textarea></div><button type=\"submit\" class=\"w-full px-6 py-3 mt-6 text-md font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-500 rounded-md hover:bg-blue-400 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-50\">get in touch</button></form></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
