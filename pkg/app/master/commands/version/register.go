package version

import (
	"github.com/docker-slim/docker-slim/pkg/app/master/commands"
)

// zzc 将version指令注册到指令集commands.CLI中？
func RegisterCommand() {
	commands.CLI = append(commands.CLI, CLI)
	commands.CommandSuggestions = append(commands.CommandSuggestions, CommandSuggestion)
}
