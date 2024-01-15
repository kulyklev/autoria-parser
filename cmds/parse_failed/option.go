package parse_failed

import (
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli/v2"
	"gitlab.com/kulyklev/autoria-parser/database"
)

var (
	sqlDbInst *sqlx.DB
)

func beforeFuncs(opts ...cli.BeforeFunc) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		for _, opt := range opts {
			err := opt(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func initDbOption(ctx *cli.Context) error {
	var err error

	sqlDbInst, err = database.Open(database.Config{
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		Name:       "postgres",
		DisableTLS: true,
	})
	if err != nil {
		return err
	}

	return nil
}
