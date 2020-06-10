package common

type ClientParameters struct {
	SharedCredentialsFile string
	Profile               string
}

type Credentials struct {
	ServiceAccountID string `json:"serviceAccountId"`
	AccessKey        string `json:"accessKey"`
	SecretKey        string `json:"secretKey"`
}

type Client struct {
	AccessToken string
}

type Constants struct {
	BaseURL string `yaml:"base_url"`
	Routes  struct {
		Auth   string `yaml:"authentication"`
		Policy struct {
			Create string `yaml:"create"`
		} `yaml:"policy"`
	} `yaml:"routes"`
}
