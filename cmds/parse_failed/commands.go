package parse_failed

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
	"gitlab.com/kulyklev/autoria-parser/common"
	"go.uber.org/zap"
)

func CommandParse(log *zap.SugaredLogger) *cli.Command {
	return &cli.Command{
		Name:   "parse-failed",
		Usage:  "parse-failed auto-ria",
		Before: beforeFuncs(initDbOption),
		Action: func(ctx *cli.Context) error {
			return run(ctx, log)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api",
				Usage: "auto-ria API key",
				EnvVars: []string{
					"API_KEY",
				},
			},
			&cli.StringFlag{
				Name:  "api1",
				Usage: "auto-ria API key1",
				EnvVars: []string{
					"API_KEY_1",
				},
			},
			&cli.StringFlag{
				Name:  "api2",
				Usage: "auto-ria API key2",
				EnvVars: []string{
					"API_KEY_2",
				},
			},
			&cli.StringFlag{
				Name:  "api3",
				Usage: "auto-ria API key3",
				EnvVars: []string{
					"API_KEY_3",
				},
			},
			&cli.StringFlag{
				Name:  "api4",
				Usage: "auto-ria API key4",
				EnvVars: []string{
					"API_KEY_4",
				},
			},
			&cli.StringFlag{
				Name:  "api5",
				Usage: "auto-ria API key5",
				EnvVars: []string{
					"API_KEY_5",
				},
			},
			&cli.StringFlag{
				Name:  "api6",
				Usage: "auto-ria API key6",
				EnvVars: []string{
					"API_KEY_6",
				},
			},
			&cli.StringFlag{
				Name:  "api7",
				Usage: "auto-ria API key7",
				EnvVars: []string{
					"API_KEY_7",
				},
			},
			&cli.StringFlag{
				Name:  "api8",
				Usage: "auto-ria API key8",
				EnvVars: []string{
					"API_KEY_8",
				},
			},
			&cli.StringFlag{
				Name:  "api9",
				Usage: "auto-ria API key9",
				EnvVars: []string{
					"API_KEY_9",
				},
			},
			&cli.StringFlag{
				Name:  "api10",
				Usage: "auto-ria API key9",
				EnvVars: []string{
					"API_KEY_10",
				},
			},
			&cli.StringFlag{
				Name:  "api11",
				Usage: "auto-ria API key10",
				EnvVars: []string{
					"API_KEY_11",
				},
			},
			&cli.StringFlag{
				Name:  "api12",
				Usage: "auto-ria API key10",
				EnvVars: []string{
					"API_KEY_12",
				},
			},
			&cli.StringFlag{
				Name:  "api13",
				Usage: "auto-ria API key10",
				EnvVars: []string{
					"API_KEY_13",
				},
			},
			&cli.StringFlag{
				Name:  "api14",
				Usage: "auto-ria API key10",
				EnvVars: []string{
					"API_KEY_14",
				},
			},
		},
	}
}

func run(cCtx *cli.Context, log *zap.SugaredLogger) error {
	defer common.TimeTrack(time.Now(), "auto-ria parsing")

	ctx := context.Background()
	defer sqlDbInst.Close()

	failedAutoIds, err := common.GetFailedAutoIds(ctx, sqlDbInst)
	if err != nil {
		return fmt.Errorf("error: failed to get failed auto ids, err:%w", err)
	}

	client := http.DefaultClient
	apiKey := cCtx.String("api")
	apiKey1 := cCtx.String("api1")
	apiKey2 := cCtx.String("api2")
	apiKey3 := cCtx.String("api3")
	apiKey4 := cCtx.String("api4")
	apiKey5 := cCtx.String("api5")
	apiKey6 := cCtx.String("api6")
	apiKey7 := cCtx.String("api7")
	apiKey8 := cCtx.String("api8")
	apiKey9 := cCtx.String("api9")
	apiKey10 := cCtx.String("api10")
	apiKey11 := cCtx.String("api11")
	apiKey12 := cCtx.String("api12")
	apiKey13 := cCtx.String("api13")
	apiKey14 := cCtx.String("api14")
	apiKeys := [15]string{
		apiKey,
		apiKey1,
		apiKey2,
		apiKey3,
		apiKey4,
		apiKey5,
		apiKey6,
		apiKey7,
		apiKey8,
		apiKey9,
		apiKey10,
		apiKey11,
		apiKey12,
		apiKey13,
		apiKey14,
	}

	err = common.TruncateFailedParses(ctx, log, sqlDbInst)
	if err != nil {
		return fmt.Errorf("error truncating: %w", err)
	}

	common.CarsParsing(ctx, log, sqlDbInst, client, apiKeys, failedAutoIds)

	failedParses, err := common.CountFailedParses(ctx, sqlDbInst)
	if err != nil {
		log.Errorw("failed to count failed parses", "ERROR", err)
	}

	log.Infow("stats", "number of parsed cars", len(failedAutoIds))
	log.Infow("stats", "requests done", common.ApiRequestCount)
	log.Infow("stats", "number of failed requests", failedParses)

	return nil
}
