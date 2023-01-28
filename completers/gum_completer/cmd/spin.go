package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/rsteube/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var spinCmd = &cobra.Command{
	Use:   "spin",
	Short: "Display spinner while running a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(spinCmd).Standalone()

	spinCmd.Flags().StringP("align", "a", "", "Alignment of spinner with regard to the title")
	spinCmd.Flags().Bool("show-output", false, "Show output of command")
	spinCmd.Flags().StringP("spinner", "s", "", "Spinner type")
	spinCmd.Flags().String("spinner.align", "", "Text Alignment")
	spinCmd.Flags().String("spinner.background", "", "Background Color")
	spinCmd.Flags().Bool("spinner.bold", false, "Bold text")
	spinCmd.Flags().String("spinner.border", "", "Border Style")
	spinCmd.Flags().String("spinner.border-background", "", "Border Background Color")
	spinCmd.Flags().String("spinner.border-foreground", "", "Border Foreground Color")
	spinCmd.Flags().Bool("spinner.faint", false, "Faint text")
	spinCmd.Flags().String("spinner.foreground", "", "Foreground Color")
	spinCmd.Flags().String("spinner.height", "", "Text height")
	spinCmd.Flags().Bool("spinner.italic", false, "Italicize text")
	spinCmd.Flags().String("spinner.margin", "", "Text margin")
	spinCmd.Flags().String("spinner.padding", "", "Text padding")
	spinCmd.Flags().Bool("spinner.strikethrough", false, "Strikethrough text")
	spinCmd.Flags().Bool("spinner.underline", false, "Underline text")
	spinCmd.Flags().String("spinner.width", "", "Text width")
	spinCmd.Flags().String("title", "", "Text to display to user while spinning")
	spinCmd.Flags().String("title.align", "", "Text Alignment")
	spinCmd.Flags().String("title.background", "", "Background Color")
	spinCmd.Flags().Bool("title.bold", false, "Bold text")
	spinCmd.Flags().String("title.border", "", "Border Style")
	spinCmd.Flags().String("title.border-background", "", "Border Background Color")
	spinCmd.Flags().String("title.border-foreground", "", "Border Foreground Color")
	spinCmd.Flags().Bool("title.faint", false, "Faint text")
	spinCmd.Flags().String("title.foreground", "", "Foreground Color")
	spinCmd.Flags().String("title.height", "", "Text height")
	spinCmd.Flags().Bool("title.italic", false, "Italicize text")
	spinCmd.Flags().String("title.margin", "", "Text margin")
	spinCmd.Flags().String("title.padding", "", "Text padding")
	spinCmd.Flags().Bool("title.strikethrough", false, "Strikethrough text")
	spinCmd.Flags().Bool("title.underline", false, "Underline text")
	spinCmd.Flags().String("title.width", "", "Text width")

	rootCmd.AddCommand(spinCmd)

	carapace.Gen(spinCmd).FlagCompletion(carapace.ActionMap{
		"align": carapace.ActionValues("left", "right"),
		"spinner": carapace.ActionValuesDescribed(
			"line", "/",
			"dot", "⢿",
			"minidot", "⠋",
			"jump", "⡈",
			"pulse", "░",
			"points", "●",
			"globe", "🌍",
			"moon", "🌗",
			"monkey", "🙊",
			"meter", "▰",
			"hamburger", "☲",
		),
		"spinner.background":        gum.ActionColors(),
		"spinner.border":            gum.ActionBorders(),
		"spinner.border-background": gum.ActionColors(),
		"spinner.border-foreground": gum.ActionColors(),
		"spinner.foreground":        gum.ActionColors(),
		"title.align":               gum.ActionAlignments(),
		"title.background":          gum.ActionColors(),
		"title.border":              gum.ActionBorders(),
		"title.border-background":   gum.ActionColors(),
		"title.border-foreground":   gum.ActionColors(),
		"title.foreground":          gum.ActionColors(),
	})

	carapace.Gen(spinCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if spinCmd.ArgsLenAtDash() > 0 {
				return carapace.ActionValues()
			}

			switch len(c.Args) {
			case 0:
				return carapace.Batch(
					os.ActionPathExecutables(),
					carapace.ActionFiles(),
				).ToA()
			default:
				cmd := c.Args[0]
				c.Args = c.Args[1:]
				return bridge.ActionCarapaceBin(cmd).Invoke(c).ToA()
			}
		}),
	)

	embed.EmbedCarapaceBin(spinCmd)
}
