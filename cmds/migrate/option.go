package migrate

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
	"gitlab.com/kulyklev/autoria-parser/migrations"
)

var (
	sqlDb *sql.DB
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
	driver := "postgres"
	dbUrl := ctx.String("url")

	goose.SetBaseFS(migrations.EmbedMigrations)

	db, err := goose.OpenDBWithDriver(driver, dbUrl)
	if err != nil {
		return err
	}

	sqlDb = db
	return nil
}
