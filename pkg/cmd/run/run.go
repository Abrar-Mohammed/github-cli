package run

import (
<<<<<<< HEAD
	cmdList "github.com/cli/cli/pkg/cmd/run/list"
	cmdView "github.com/cli/cli/pkg/cmd/run/view"
=======
	cmdDownload "github.com/cli/cli/pkg/cmd/run/download"
	cmdList "github.com/cli/cli/pkg/cmd/run/list"
	cmdRerun "github.com/cli/cli/pkg/cmd/run/rerun"
	cmdView "github.com/cli/cli/pkg/cmd/run/view"
	cmdWatch "github.com/cli/cli/pkg/cmd/run/watch"
>>>>>>> origin/bad-branch
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRun(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
<<<<<<< HEAD
		Use:   "run <command>",
		Short: "View details about workflow runs",
		Long:  "List, view, and watch recent workflow runs from GitHub Actions.",
		// TODO i'd like to have all the actions commands sorted into their own zone which i think will
		// require a new annotation
=======
		Use:    "run <command>",
		Short:  "View details about workflow runs",
		Long:   "List, view, and watch recent workflow runs from GitHub Actions.",
		Hidden: true,
		Annotations: map[string]string{
			"IsActions": "true",
		},
>>>>>>> origin/bad-branch
	}
	cmdutil.EnableRepoOverride(cmd, f)

	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdView.NewCmdView(f, nil))
<<<<<<< HEAD
=======
	cmd.AddCommand(cmdRerun.NewCmdRerun(f, nil))
	cmd.AddCommand(cmdDownload.NewCmdDownload(f, nil))
	cmd.AddCommand(cmdWatch.NewCmdWatch(f, nil))
>>>>>>> origin/bad-branch

	return cmd
}
