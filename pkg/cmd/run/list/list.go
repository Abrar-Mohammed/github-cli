package list

import (
	"fmt"
	"net/http"

	"github.com/cli/cli/api"
	"github.com/cli/cli/internal/ghrepo"
	"github.com/cli/cli/pkg/cmd/run/shared"
<<<<<<< HEAD
=======
	workflowShared "github.com/cli/cli/pkg/cmd/workflow/shared"
>>>>>>> origin/bad-branch
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/utils"
	"github.com/spf13/cobra"
)

const (
	defaultLimit = 10
)

type ListOptions struct {
	IO         *iostreams.IOStreams
	HttpClient func() (*http.Client, error)
	BaseRepo   func() (ghrepo.Interface, error)

<<<<<<< HEAD
	ShowProgress bool
	PlainOutput  bool

	Limit int
}

// TODO filters
// --state=(pending,pass,fail,etc)
// --active - pending
// --workflow - filter by workflow name

=======
	PlainOutput bool

	Limit            int
	WorkflowSelector string
}

>>>>>>> origin/bad-branch
func NewCmdList(f *cmdutil.Factory, runF func(*ListOptions) error) *cobra.Command {
	opts := &ListOptions{
		IO:         f.IOStreams,
		HttpClient: f.HttpClient,
	}

	cmd := &cobra.Command{
<<<<<<< HEAD
		Use:   "list",
		Short: "List recent workflow runs",
		Args:  cobra.NoArgs,
=======
		Use:    "list",
		Short:  "List recent workflow runs",
		Args:   cobra.NoArgs,
		Hidden: true,
>>>>>>> origin/bad-branch
		RunE: func(cmd *cobra.Command, args []string) error {
			// support `-R, --repo` override
			opts.BaseRepo = f.BaseRepo

			terminal := opts.IO.IsStdoutTTY() && opts.IO.IsStdinTTY()
<<<<<<< HEAD
			opts.ShowProgress = terminal
=======
>>>>>>> origin/bad-branch
			opts.PlainOutput = !terminal

			if opts.Limit < 1 {
				return &cmdutil.FlagError{Err: fmt.Errorf("invalid limit: %v", opts.Limit)}
			}

			if runF != nil {
				return runF(opts)
			}

			return listRun(opts)
		},
	}

	cmd.Flags().IntVarP(&opts.Limit, "limit", "L", defaultLimit, "Maximum number of runs to fetch")
<<<<<<< HEAD
=======
	cmd.Flags().StringVarP(&opts.WorkflowSelector, "workflow", "w", "", "Filter runs by workflow")
>>>>>>> origin/bad-branch

	return cmd
}

func listRun(opts *ListOptions) error {
<<<<<<< HEAD
	if opts.ShowProgress {
		opts.IO.StartProgressIndicator()
	}
	baseRepo, err := opts.BaseRepo()
	if err != nil {
		// TODO better err handle
		return err
=======
	baseRepo, err := opts.BaseRepo()
	if err != nil {
		return fmt.Errorf("failed to determine base repo: %w", err)
>>>>>>> origin/bad-branch
	}

	c, err := opts.HttpClient()
	if err != nil {
<<<<<<< HEAD
		// TODO better error handle
		return err
	}
	client := api.NewClientFromHTTP(c)

	runs, err := shared.GetRuns(client, baseRepo, opts.Limit)
	if err != nil {
		// TODO better error handle
		return err
=======
		return fmt.Errorf("failed to create http client: %w", err)
	}
	client := api.NewClientFromHTTP(c)

	var runs []shared.Run
	var workflow *workflowShared.Workflow

	opts.IO.StartProgressIndicator()
	if opts.WorkflowSelector != "" {
		states := []workflowShared.WorkflowState{workflowShared.Active}
		workflow, err = workflowShared.ResolveWorkflow(
			opts.IO, client, baseRepo, false, opts.WorkflowSelector, states)
		if err == nil {
			runs, err = shared.GetRunsByWorkflow(client, baseRepo, opts.Limit, workflow.ID)
		}
	} else {
		runs, err = shared.GetRuns(client, baseRepo, opts.Limit)
	}
	opts.IO.StopProgressIndicator()
	if err != nil {
		return fmt.Errorf("failed to get runs: %w", err)
>>>>>>> origin/bad-branch
	}

	tp := utils.NewTablePrinter(opts.IO)

	cs := opts.IO.ColorScheme()

<<<<<<< HEAD
	if opts.ShowProgress {
		opts.IO.StopProgressIndicator()
	}
	for _, run := range runs {
		//idStr := cs.Cyanf("%d", run.ID)
=======
	if len(runs) == 0 {
		if !opts.PlainOutput {
			fmt.Fprintln(opts.IO.ErrOut, "No runs found")
		}
		return nil
	}

	out := opts.IO.Out

	for _, run := range runs {
>>>>>>> origin/bad-branch
		if opts.PlainOutput {
			tp.AddField(string(run.Status), nil, nil)
			tp.AddField(string(run.Conclusion), nil, nil)
		} else {
<<<<<<< HEAD
			tp.AddField(shared.Symbol(cs, run.Status, run.Conclusion), nil, nil)
=======
			symbol, symbolColor := shared.Symbol(cs, run.Status, run.Conclusion)
			tp.AddField(symbol, nil, symbolColor)
>>>>>>> origin/bad-branch
		}

		tp.AddField(run.CommitMsg(), nil, cs.Bold)

		tp.AddField(run.Name, nil, nil)
		tp.AddField(run.HeadBranch, nil, cs.Bold)
		tp.AddField(string(run.Event), nil, nil)

		if opts.PlainOutput {
			elapsed := run.UpdatedAt.Sub(run.CreatedAt)
<<<<<<< HEAD
=======
			if elapsed < 0 {
				elapsed = 0
			}
>>>>>>> origin/bad-branch
			tp.AddField(elapsed.String(), nil, nil)
		}

		tp.AddField(fmt.Sprintf("%d", run.ID), nil, cs.Cyan)

		tp.EndRow()
	}

	err = tp.Render()
	if err != nil {
<<<<<<< HEAD
		// TODO better error handle
		return err
	}

	fmt.Fprintln(opts.IO.Out)
	fmt.Fprintln(opts.IO.Out, "For details on a run, try: gh run view <run-id>")
=======
		return err
	}

	if !opts.PlainOutput {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "For details on a run, try: gh run view <run-id>")
	}
>>>>>>> origin/bad-branch

	return nil
}
