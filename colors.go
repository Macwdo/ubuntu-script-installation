package main

type Color string
type Style string

const (
	Bold      Style = "\033[1m"
	Underline Style = "\033[4m"
	Italic    Style = "\033[3m"
	Reversed  Style = "\033[7m"
	Normal    Style = "\033[0m"
)

const (
	Red    Color = "\033[31m"
	Blue   Color = "\033[34m"
	Green  Color = "\033[32m"
	Grey   Color = "\033[37m"
	Yellow Color = "\033[33m"
	Reset  Color = "\033[0m"
	White  Color = "\033[97m"
)
