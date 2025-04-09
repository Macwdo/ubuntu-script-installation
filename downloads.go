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

	ohmyzshDir := GetUserHome() + "/" + ".oh-my-zsh"

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

// def install_powerlevel10k():
//     log("Passo: Instalando Powerlevel10k...", BLUE)
//     home = os.path.expanduser("~")
//     omz_custom = os.path.join(home, ".oh-my-zsh", "custom")
//     themes_dir = os.path.join(omz_custom, "themes")
//     os.makedirs(themes_dir, exist_ok=True)
//     dest = os.path.join(themes_dir, "powerlevel10k")
//     clone_repo("https://github.com/romkatv/powerlevel10k.git --depth=1", dest)

func InstallZshPlugins() {
	path := GetUserHome()

	omzCustomDir := path + "/.oh-my-zsh/custom/"
	customPluginsDir := omzCustomDir + "plugins"

	os.MkdirAll(customPluginsDir, os.ModePerm)
	plugins := map[string]string{
		"zsh-autosuggestions":     "https://github.com/zsh-users/zsh-autosuggestions.git",
		"zsh-syntax-highlighting": "https://github.com/zsh-users/zsh-syntax-highlighting.git",
	}

	for name, repoURL := range plugins {
		dest := customPluginsDir + "/" + name

		LogMessage("Installing "+name+" plugin", Normal, Grey)

		if CheckDir(dest) {
			LogMessage(name+" already exists", Italic, Yellow)
			continue
		}

		err := RunCommandErr("git clone " + repoURL + " " + dest)
		if err != nil {
			LogMessage("Error cloning plugin "+name+": "+err.Error(), Bold, Red)
			continue
		}
	}

	LogMessage("Installing P10K", Normal, Grey)
	customThemesDir := omzCustomDir + "themes"
	if CheckDir(customThemesDir) {
		LogMessage("P10k Already Exists", Italic, Yellow)
	}

	repoUrl := "https://github.com/romkatv/powerlevel10k.git --depth=1"
	err := RunCommandErr("git clone " + repoUrl + " " + customThemesDir)

	if err != nil {
		LogMessage("Error trying to install P10k", Bold, Red)
	}

}
