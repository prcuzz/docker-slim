package version

import (
	"github.com/docker-slim/docker-slim/pkg/app/master/commands"
)

// zzc ��versionָ��ע�ᵽָ�commands.CLI�У�
func RegisterCommand() {
	commands.CLI = append(commands.CLI, CLI)
	commands.CommandSuggestions = append(commands.CommandSuggestions, CommandSuggestion)
}
