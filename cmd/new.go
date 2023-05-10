package cmd

import (
	"fmt"
	"github.com/nsecho/mtemp/internal/temp"
	"github.com/spf13/cobra"
	"time"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new email",
	RunE: func(cmd *cobra.Command, args []string) error {
		one := &temp.OneSecMail{}
		mail, err := createNewTemp(one)
		if err != nil {
			return err
		}
		fmt.Printf("Created new \"%s\" mail\n", mail)

		mon, err := cmd.Flags().GetBool("monitor")
		if err != nil {
			return err
		}

		if !mon {
			return nil
		}

		done := make(chan error)
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				go func() {
					mails, err := one.Check()
					if err != nil {
						done <- err
						return
					}
					for _, mail := range mails {
						fmt.Printf("New mail %s from %s\n",
							mail.Subject, mail.From)
						if err := one.Read(mail.ID, &mail); err != nil {
							done <- err
							return
						}
						fmt.Printf("Content: \n%s\n", mail.Body)
					}
				}()
			case <-done:
				return err
			}
		}
		return nil
	},
}

func init() {
	newCmd.Flags().BoolP("monitor", "m", false, "monitor mailbox for new mails")
	rootCmd.AddCommand(newCmd)
}

func createNewTemp(tmp temp.Temper) (string, error) {
	return tmp.Fetch()
}
