package common

// StartUp the service
func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize private/public keys for JWT auth
	initKeys()
	// Start a MongoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
