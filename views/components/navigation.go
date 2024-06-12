package components

type NavigationLink struct {
	Text string
	Href string
}

var NavigationLinks = []NavigationLink{
	{Text: "Blog", Href: "/blog"},
	{Text: "About me", Href: "/about-me"},
	{Text: "About this website", Href: "/about-this-website"},
	{Text: "Echo", Href: "/echo"},
}
