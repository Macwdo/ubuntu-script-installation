package main

import (
	"os"
	"time"
)

func main() {

	welcomeMessage()
	checkRoot()
	updateAndUpgrade()
	installUsefulTools()
	changeDefaultShell()
	installOhMyZsh()
	installZshPlugins()
}

func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func welcomeMessage() {
	ClearScreen()

	Separator()
	LogMessage("Automated Sexo Script for Linux (Ubuntu/Debian)", Bold, Blue)
	LogMessage("This script will be doing the following stuffs:", Bold, White)

	LogMessage("1. Update and upgrade the system packages", Bold, White)
	LogMessage("2. Install Choosed Programming Languages", Bold, White)
	LogMessage("3. Install Choosed Devops Tools", Bold, White)
	Separator()

	LogMessage("Press enter to continue...", Bold, White)

	WaitForKeyPress()
	ClearScreen()

}

func checkRoot() {
	WithDots("Checking if the script is run as root", func() {
		sleep(1)
		if !CheckIsRoot() {
			LogMessage("This script must be run as root.", Bold, Red)
			os.Exit(1)
		}
	})
}

func updateAndUpgrade() {
	WithDots("Installing Packages", InstallPackages)
	WithDots("Updating Packages", UpdatePackages)
}

func installUsefulTools() {
	WithDots("Installing Useful Tools", InstallUsefulTools)
}

func changeDefaultShell() {
	WithDots("Changing default shell to zsh", ChangeDefaultShell)
}

func installOhMyZsh() {
	WithDots("Installing Oh My Zsh", func() {
		sleep(1)
		InstallOhMyZsh()
	})
}

func installZshPlugins() {
	WithDots("Downloading Zsh Plugins", InstallZshPlugins)
}
