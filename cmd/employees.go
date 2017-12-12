package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/spf13/cobra"
)

// employeesCmd represents the employees command
var employeesCmd = &cobra.Command{
	Use:   "employees",
	Short: "Displays a list of employees",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		employees, err := database.Employees()
		if err != nil {
			fmt.Println(err)
		}

		table := termtables.CreateTable()
		table.AddHeaders("#", "Company", "Employee")

		for i, e := range employees {
			table.AddRow(i, e.FullName, e.Company.Name)
		}

		fmt.Println(table.Render())
	},
}