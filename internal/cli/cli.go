package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mholm/wlog/internal/storage"
	"github.com/mholm/wlog/internal/task"
	"github.com/spf13/cobra"
)

// CLI represents the command-line interface
type CLI struct {
	rootCmd *cobra.Command
	storage *storage.Storage
}

// NewCLI creates a new CLI instance
func NewCLI() (*CLI, error) {
	storage, err := storage.NewStorage()
	if err != nil {
		return nil, err
	}

	cli := &CLI{
		storage: storage,
	}

	cli.rootCmd = &cobra.Command{
		Use:   "wlog",
		Short: "A terminal-based work logging application",
		Long: `wlog helps you track and manage your work tasks with importance levels from 1 (least important) to 5 (most important).

Tasks are stored in: ~/.wlog/tasks.json`,
	}

	cli.rootCmd.AddCommand(cli.newAddCmd())
	cli.rootCmd.AddCommand(cli.newListCmd())
	cli.rootCmd.AddCommand(cli.newHelpCmd())

	return cli, nil
}

// Execute runs the CLI
func (c *CLI) Execute() error {
	return c.rootCmd.Execute()
}

// newAddCmd creates the add command
func (c *CLI) newAddCmd() *cobra.Command {
	var importance int

	cmd := &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := c.storage.LoadTasks()
			if err != nil {
				return err
			}

			newTask := task.NewTask(args[0], task.Importance(importance))
			tasks = append(tasks, newTask)

			if err := c.storage.SaveTasks(tasks); err != nil {
				return err
			}

			importanceColor := c.getImportanceColor(newTask.Importance)
			fmt.Printf("Added task: %s %s\n", 
				importanceColor.Sprintf("[%d]", newTask.Importance),
				newTask.Description,
			)
			return nil
		},
	}

	cmd.Flags().IntVarP(&importance, "importance", "i", 3, "Task importance (1-5)")

	return cmd
}

// newListCmd creates the list command
func (c *CLI) newListCmd() *cobra.Command {
	var (
		period string
		top    int
		minImportance int
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := c.storage.LoadTasks()
			if err != nil {
				return err
			}

			// Sort tasks by importance and date
			task.SortByImportanceAndDate(tasks)

			// Filter by minimum importance if specified
			if minImportance > 0 {
				filteredTasks := make([]*task.Task, 0)
				for _, t := range tasks {
					if int(t.Importance) >= minImportance {
						filteredTasks = append(filteredTasks, t)
					}
				}
				tasks = filteredTasks
			}

			// Apply top N filter if specified
			if top > 0 && top < len(tasks) {
				tasks = tasks[:top]
			}

			// Print tasks with colors
			for _, t := range tasks {
				importanceColor := c.getImportanceColor(t.Importance)
				dateColor := color.New(color.FgHiBlack)
				fmt.Printf("%s %s %s\n", 
					importanceColor.Sprintf("[%d]", t.Importance),
					t.Description,
					dateColor.Sprintf("(%s)", t.CreatedAt.Format("2006-01-02 15:04")),
				)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&period, "period", "p", "all", "Time period (day, week, month, quarter)")
	cmd.Flags().IntVarP(&top, "top", "t", 0, "Show only top N tasks")
	cmd.Flags().IntVarP(&minImportance, "min-importance", "m", 0, "Show only tasks with at least this importance (1-5)")

	return cmd
}

// getImportanceColor returns a color based on the task's importance
func (c *CLI) getImportanceColor(importance task.Importance) *color.Color {
	switch importance {
	case 5:
		return color.New(color.FgRed, color.Bold)
	case 4:
		return color.New(color.FgYellow, color.Bold)
	case 3:
		return color.New(color.FgGreen)
	case 2:
		return color.New(color.FgCyan)
	case 1:
		return color.New(color.FgBlue)
	default:
		return color.New(color.FgWhite)
	}
}

// newHelpCmd creates the help command
func (c *CLI) newHelpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Show detailed help information",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			tasksFile := filepath.Join(homeDir, ".wlog", "tasks.json")

			fmt.Println("wlog - Work Logging Application")
			fmt.Println("\nAvailable Commands:")
			fmt.Println("  add [description] -i <importance>")
			fmt.Println("    Add a new task with importance level (1-5)")
			fmt.Println("    Example: wlog add \"Important meeting\" -i 5")
			fmt.Println("\n  list [options]")
			fmt.Println("    List tasks with optional filters:")
			fmt.Println("    --min-importance, -m: Show tasks with at least this importance (1-5)")
			fmt.Println("    --period, -p: Filter by time period (day, week, month, quarter)")
			fmt.Println("    --top, -t: Show only top N tasks")
			fmt.Println("    Example: wlog list -m 4 -p week")
			fmt.Println("\nTasks are stored in:", tasksFile)
			return nil
		},
	}
} 