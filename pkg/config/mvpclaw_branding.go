// MClaw Branding Configuration
// This file is for MClaw branding
// To use: Replace domeclaw_branding.go with this file content

package config

const (
	// MVPClawBanner is the ASCII art banner for MClaw
	MVPClawBanner = `
███╗   ███╗    ██████╗██╗      █████╗ ██╗    ██╗
████╗ ████║   ██╔════╝██║     ██╔══██╗██║    ██║
██╔████╔██║   ██║     ██║     ███████║██║ █╗ ██║
██║╚██╔╝██║   ██║     ██║     ██╔══██║██║███╗██║
██║ ╚═╝ ██║   ╚██████╗███████╗██║  ██║╚███╔███╔╝
╚═╝     ╚═╝    ╚═════╝╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝
`

	// MVPClawAppNameDisplay is the display name for the application
	MVPClawAppNameDisplay = "MClaw"

	// MVPClawAppShortDescription is the short description shown in CLI
	MVPClawAppShortDescription = "MClaw - Personal AI Assistant with Wallet & Webhook"

	// MVPClawAppLongDescription is the detailed description
	MVPClawAppLongDescription = "MClaw is a lightweight personal AI assistant with Ethereum wallet integration and webhook channel support."
)

// GetBanner returns the MClaw banner with colors
func MVPClawGetBanner() string {
	return MVPClawBanner
}

// GetAppName returns the application display name
func MVPClawGetAppName() string {
	return MVPClawAppNameDisplay
}
