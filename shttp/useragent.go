package shttp

var agentType string

func UseUserAgent(s string) {
	agentType = s
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"

func getUserAgent() string {
	switch agentType {
	default:
		return userAgent
	}
}
