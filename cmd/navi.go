package cmd

import (
	"fmt"

	"github.com/nitintf/navi/internal/ai"
	"github.com/nitintf/navi/utils"
	"github.com/spf13/cobra" // Import Cobra library
)

var commandTemplate = `
Imagine you are a security-conscious shell or terminal expert with a lot of computer knowledge.

Write a single, safe shell command that achieves the desired outcome. The command should:

* Not modify or delete files or folders.
* Be appropriate for a general audience (avoid offensive or harmful commands).
* Not require additional explanation.

Here is the prompt:

> %s

If the prompt is not related to a safe shell command or is not related to shell or commands or terminal, return "NAVI_AI_ERROR".

If the prompt is not appropriate for a general audience, return "NAVI_AI_ERROR".

If the prompt is unclear or ambiguous, return "NAVI_AI_ERROR".

If the prompt requires additional explanation, return "NAVI_AI_ERROR".

If the prompt is not a shell command, return "NAVI_AI_ERROR".

If the prompt is a shell command but is not safe, return "NAVI_AI_ERROR".

**Examples:**

* Prompt: "List all files in the current directory."
* Response: ls

* Prompt: "Delete all files in the current directory." (Unsafe)
* Response: "NAVI_AI_ERROR"

* Prompt: "Show a funny cat video." (Not a shell command)
* Response: "NAVI_AI_ERROR"
`

var rootCmd = &cobra.Command{
	Use:   "navi",
	Short: "Navi - Your AI-powered Shell Guide",
	Long:  `Navi is a command-line tool that uses AI to understand your natural language prompts and generate the corresponding shell commands. Think of it as your guide in the shell world!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]
		promptLength := len(prompt)

		if promptLength > 120 {
			utils.LogError("Prompt is too long. Please keep it under 120 characters.")
			return
		}

		spin := utils.GetSpinner()
		spin.Suffix = " Generating command..."
		spin.Start()
		resp, err := ai.Generate(cmd.Context(), commandTemplate, prompt)

		if err != nil {
			spin.Stop()
			utils.LogError(err.Error())
			return
		}

		spin.Stop()
		utils.LogInfo(resp)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
