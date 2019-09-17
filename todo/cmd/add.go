package cmd

import "github.com/spf13/cobra"
import "fmt"
import "strings"
import "todo/redis"
import "time"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		notiTask := strings.Join(args, " ")
		r := redis.CreateInstance()
		key := strings.TrimSpace(time.Now().Format("2006-01-02 15:04:05"))
		key = "todo:" + key
		err := r.SaveTask(key, notiTask)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("Added \"%s\" to your task list\n", notiTask)
		}
		defer r.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
