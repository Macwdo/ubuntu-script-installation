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

// def set_zsh_plugins():
//     log("Passo: Configurando plugins no .zshrc...", BLUE)
//     home = os.path.expanduser("~")
//     zshrc_path = os.path.join(home, ".zshrc")
//     plugins_line = "plugins=( git docker docker-compose python celery zsh-autosuggestions zsh-syntax-highlighting )"
//     try:
//         if os.path.exists(zshrc_path):
//             with open(zshrc_path, "r") as f:
//                 lines = f.readlines()
//             found = False
//             for i, line in enumerate(lines):
//                 if line.startswith("plugins="):
//                     lines[i] = plugins_line + "\n"
//                     found = True
//                     break
//             if not found:
//                 lines.append("\n" + plugins_line + "\n")
//             with open(zshrc_path, "w") as f:
//                 f.writelines(lines)
//             log("Plugins configurados no .zshrc com sucesso.", GREEN)
//         else:
//             with open(zshrc_path, "w") as f:
//                 f.write(plugins_line + "\n")
//             log("Arquivo .zshrc criado e plugins configurados com sucesso.", GREEN)
//     except Exception as e:
//         log(f"Erro ao configurar plugins no .zshrc: {e}", RED)
