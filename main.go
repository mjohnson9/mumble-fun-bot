package main

import (
	"os"

	"github.com/nightexcessive/hal"
	_ "github.com/nightexcessive/hal-mumble"
	_ "github.com/nightexcessive/hal/store/memory"
)

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(

		&hal.Handler{
			Method:  hal.HEAR,
			Pattern: `tableflip`,
			Run: func(res *hal.Response) error {
				return res.Send(`(╯°□°）╯︵ ┻━┻`)
			},
		},

		hal.Respond(`speak (.*)`, func(res *hal.Response) error {
			return res.Play(res.Match[1])
		}),

		hal.Enter(func(res *hal.Response) error {
			return res.Send("/nom en-US")
		}),
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
