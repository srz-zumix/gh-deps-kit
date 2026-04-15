package main

import (
	"embed"

	"github.com/srz-zumix/gh-deps-kit/cmd"
)

//go:embed skills
var skillsFS embed.FS

func main() {
	cmd.RegisterSkillsCmd(skillsFS)
	cmd.Execute()
}
