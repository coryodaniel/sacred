package cmd

type AuthConfiguration struct {
	Token string
	Domain string
}

type DocConfiguration struct {
	Name      string
	Notice 		string
	Files     []string
	SpaceId   string
	ContentId string
}

type Configuration struct {
	Auth AuthConfiguration
	Docs []DocConfiguration
}
