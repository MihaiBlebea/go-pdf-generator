package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/MihaiBlebea/go-diploma/diploma"
	"github.com/MihaiBlebea/go-diploma/user"
	"github.com/gosuri/uiprogress"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	generateCmd.Flags().IntP("count", "c", 0, "The number of users to generate")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate diplomas for users.",
	Long:  "Generate diplomas for users.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmd.Flags().GetInt("count")
		if err != nil {
			log.Fatal(err)
		}

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		store := user.NewStore(l)
		store.GenerateUsers(c)

		d := diploma.New(l)

		users := store.Users()

		l.Info("Start generating diplomas in concurret mode")

		uiprogress.Start()
		bar := uiprogress.AddBar(len(users)).AppendCompleted().PrependElapsed()

		type result struct {
			index uint
			user  user.User
			err   error
		}

		resChan := make(chan result)

		for i, usr := range users {
			go func(usr user.User, index uint) {
				err := d.GeneratePDFNoHTML(usr.ID, usr)
				resChan <- result{index, usr, err}
			}(usr, uint(i))
		}

		var count uint
		bar.PrependFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("progress: %d/%d", count, len(users))
		})

		// for i := 0; i < len(users); i++ {
		// 	select {
		// 	case res := <-resChan:
		// 		if res.err != nil {
		// 			l.Info(res.err.Error())
		// 			break
		// 		}
		// 		count++
		// 		bar.Incr()
		// 	}
		// }

		for res := range resChan {
			// fmt.Println(elem)
			if res.err != nil {
				l.Info(res.err.Error())
				break
			}
			count++
			bar.Incr()

		}

		l.Info("Job completed")

		return nil
	},
}
