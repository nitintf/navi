package cmd

import (
	"github.com/nitintf/navi/internal/ai"
	"github.com/nitintf/navi/utils"
	"github.com/spf13/cobra"
)

var explainTemplate = `
Imagine you are a security-conscious shell or terminal expert with a lot of computer knowledge.

Your task is to explain the provided shell command in a way that a beginner could understand in less than 80 words. The explanation should:

* Be clear and concise.
* Avoid technical jargon where possible.
* Not require additional explanation.

Here is the command:

> %s

If the command is not related to shell or terminal, return "NAVI_NAVI_AI_ERROR".

If the command is unclear or ambiguous, return "NAVI_AI_ERROR".

If the command requires additional explanation beyond your capabilities, return "NAVI_AI_ERROR".

If the command is not a shell command, return "NAVI_AI_ERROR".

**Examples:**

* Command: "ls"
* Response: "The 'ls' command lists all files and directories in the current directory."

* Command: "open funny_cat_video.mp4" (Not a shell command)
* Response: "NAVI_AI_ERROR"
`

var explainCmd = &cobra.Command{
	Use:     "explain",
	Short:   "Explain - Understand your shell commands",
	Example: `navi explain "ls"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		commandLength := len(command)

		if commandLength > 120 {
			utils.LogError("Command is too long. Please keep it under 120 characters.")
			return
		}

		spin := utils.GetSpinner()
		spin.Suffix = " Explaining command..."
		spin.Start()
		resp, err := ai.Generate(cmd.Context(), explainTemplate, command)
		if err != nil {
			spin.Stop()
			utils.LogError(err.Error())
			return
		}

		spin.Stop()
		utils.LogExplanation(resp)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
