package migrate

import (
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "migrate sub-commands",
		Subcommands: []*cli.Command{
			upCommand(),
			upByOneCommand(),
			upToCommand(),
			downCommand(),
			downToCommand(),
			redoCommand(),
			resetCommand(),
			statusCommand(),
			versionCommand(),
			fixCommand(),
		},
	}
}

func upCommand() *cli.Command {
	return &cli.Command{
		Name:   "up",
		Usage:  "migrate the database to the most recent version available",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.Up(sqlDb, ".")
		},
	}
}

func upByOneCommand() *cli.Command {
	return &cli.Command{
		Name:   "up-by-one",
		Usage:  "migrate the database up by 1",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.UpByOne(sqlDb, ".")
		},
	}
}

func upToCommand() *cli.Command {
	return &cli.Command{
		Name:  "up-to",
		Usage: "migrate the database to a specific version",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "version",
				Usage:    "specific version",
				Required: true,
			},
		},
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.UpTo(sqlDb, ".", ctx.Int64("version"))
		},
	}
}

func downCommand() *cli.Command {
	return &cli.Command{
		Name:   "down",
		Usage:  "roll back the version by 1",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.Down(sqlDb, ".")
		},
	}
}

func downToCommand() *cli.Command {
	return &cli.Command{
		Name:  "down-to",
		Usage: "roll back to a specific version",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "version",
				Usage:    "specific version",
				Required: true,
			},
		},
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.DownTo(sqlDb, ".", ctx.Int64("version"))
		},
	}
}

func redoCommand() *cli.Command {
	return &cli.Command{
		Name:   "redo",
		Usage:  "re-run the latest migration",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.Redo(sqlDb, ".")
		},
	}
}

func resetCommand() *cli.Command {
	return &cli.Command{
		Name:   "reset",
		Usage:  "roll back all migrations",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.Reset(sqlDb, ".")
		},
	}
}

func statusCommand() *cli.Command {
	return &cli.Command{
		Name:   "status",
		Usage:  "dump the migration status for the current database",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.Status(sqlDb, ".")
		},
	}
}

func versionCommand() *cli.Command {
	return &cli.Command{
		Name:   "version",
		Usage:  "print the current version of the database",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return goose.Version(sqlDb, ".")
		},
	}
}

func fixCommand() *cli.Command {
	return &cli.Command{
		Name:  "fix",
		Usage: "apply sequential ordering to migrations",
		Action: func(ctx *cli.Context) error {
			return goose.Fix(".")
		},
	}
}
