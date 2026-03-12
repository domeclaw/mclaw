// MClaw Branding Configuration
// This is the main branding file for MClaw (in mclaw branch)

package config

const (
	// DomeClawBanner is the ASCII art banner for MClaw
	DomeClawBanner = `
███╗   ███╗    ██████╗██╗      █████╗ ██╗    ██╗
████╗ ████║   ██╔════╝██║     ██╔══██╗██║    ██║
██╔████╔██║   ██║     ██║     ███████║██║ █╗ ██║
██║╚██╔╝██║   ██║     ██║     ██╔══██║██║███╗██║
██║ ╚═╝ ██║   ╚██████╗███████╗██║  ██║╚███╔███╔╝
╚═╝     ╚═╝    ╚═════╝╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝
`

	// AppNameDisplay is the display name for the application
	AppNameDisplay = "MClaw"

	// AppShortDescription is the short description shown in CLI
	AppShortDescription = "MClaw - Personal AI Assistant with Wallet & Webhook"

	// AppLongDescription is the detailed description
	AppLongDescription = "MClaw is a lightweight personal AI assistant with Ethereum wallet integration and webhook channel support."
)

// GetBanner returns the MClaw banner with colors
func GetBanner() string {
	return DomeClawBanner
}

// GetAppName returns the application display name
func GetAppName() string {
	return AppNameDisplay
}
