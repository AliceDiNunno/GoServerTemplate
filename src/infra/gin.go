package infra

func GinMode() string {
	return requireEnvVariable("GIN_MODE")
}

func ListenAddress() string {
	return requireEnvVariable("LISTEN_ADDRESS")
}

func AllowedIpAddress() string {
	return requireEnvVariable("ALLOWED_IP_ADDRESS")
}
