package main

import (
	"os"
	"strings"
)

func InstallPackages() {
	apt_packages := []string{
		"build-essential",
		"curl",
		"libbz2-dev",
		"libffi-dev",
		"liblzma-dev",
		"libncursesw5-dev",
		"libreadline-dev",
		"libsqlite3-dev",
		"libssl-dev",
		"libxml2-dev",
		"libxmlsec1-dev",
		"llvm",
		"make",
		"tk-dev",
		"wget",
		"xz-utils",
		"zlib1g-dev",
	}

	string_command := "apt install -y " + strings.Join(apt_packages, " ")
	RunCommand(string_command)

}

func UpdatePackages() {
	RunCommand("apt update -y")
	RunCommand("apt dist-upgrade -y")
	RunCommand("apt upgrade -y")
	RunCommand("apt autoremove -y")
	RunCommand("apt autoclean -y")
	RunCommand("apt clean -y")
}

func InstallUsefulTools() {
	tools := []string{
		"git",
		"zsh",
		"wget",
		"curl",
		"unzip",
		"zip",
	}
	RunCommand("apt install -y " + strings.Join(tools, " "))
}

func ChangeDefaultShell() {
	RunCommand("chsh -s $(which zsh)")
}

func InstallOhMyZsh() {
	skipMsg := ""

	ohmyzshDir := GetUserHome() + string(os.PathSeparator) + ".oh-my-zsh"

	info, err := os.Stat(ohmyzshDir)
	if err == nil && info.IsDir() {
		skipMsg = "The Oh My Zsh directory already exists. Skipping installation."
	}

	if err != nil && !os.IsNotExist(err) {
		skipMsg = "Error checking Oh My Zsh directory: " + err.Error()
	}

	if skipMsg != "" {
		skipMsg = "\n" + skipMsg
		LogMessagef(skipMsg, Italic, Yellow)
		return
	}

	RunCommand("sh -c \"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)\"")
}

// def setup_zsh_plugins():
//     log("Passo: Configurando plugins do Zsh...", BLUE)
//     home = os.path.expanduser("~")
//     custom_plugins_dir = os.path.join(home, ".oh-my-zsh", "custom", "plugins")
//     os.makedirs(custom_plugins_dir, exist_ok=True)

//     plugins = {
//         "zsh-autosuggestions": "https://github.com/zsh-users/zsh-autosuggestions.git",
//         "zsh-syntax-highlighting": "https://github.com/zsh-users/zsh-syntax-highlighting.git"
//     }

//     for name, repo_url in plugins.items():
//         dest = os.path.join(custom_plugins_dir, name)
//         if os.path.isdir(dest):
//             log(f"Plugin {name} já instalado, pulando.", YELLOW)
//         else:
//             try:
//                 run_command(f"git clone {repo_url} {dest}")
//                 log(f"Plugin {name} instalado com sucesso.", GREEN)
//             except subprocess.CalledProcessError:
//                 log(f"Erro ao clonar o plugin {name}.", RED)

// WIP
func DownloadingZshPlugins() {
	path := GetUserHome()
	customPluginsDir := path + "/.oh-my-zsh/custom/plugins"
	os.MkdirAll(customPluginsDir, os.ModePerm)
	plugins := map[string]string{
		"zsh-autosuggestions":     "https://github.com/zsh-users/zsh-autosuggestions.git",
		"zsh-syntax-highlighting": "https://github.com/zsh-users/zsh-syntax-highlighting.git",
	}

	for name, repoURL := range plugins {
		dest := customPluginsDir + "/" + name
		if _, err := os.Stat(dest); err == nil {
			LogMessage(name+" já instalado, pulando.", Italic, Yellow)
			continue
		}

		LogMessage("Installing plugin "+name+"...", Normal, Yellow)
		err := RunCommandErr("git clone " + repoURL + " " + dest)
		if err != nil {
			LogMessage("Error cloning plugin "+name+": "+err.Error(), Bold, Red)
			continue
		}
	}
}
