package common

import (
	"context"
	"net/http"
	"runtime"
	"sync"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func CarsParsing(ctx context.Context, log *zap.SugaredLogger, dbInst *sqlx.DB, client *http.Client, apiKeys [15]string, carIds []string) {
	grs := runtime.NumCPU() / 2
	log.Infow("parsing cars", "number of goroutines to start", grs)
	var wg sync.WaitGroup
	wg.Add(grs)

	ch := make(chan string, grs)
	errCh := make(chan ErrorFromGoroutine)

	for proc := 0; proc < grs; proc++ {
		go func(processor int) {
			log.Infow("parsing cars", "started processor #", processor)
			defer log.Infow("parsing cars", "received shutdown signal. processor #", processor)

			defer wg.Done()
			var apiKeyIndex int

			for autoId := range ch {
				log.Infow("parsing cars", "processor #", processor, "started processing auto ID", autoId)
				err := ProcessAutoId(ctx, dbInst, client, apiKeys[apiKeyIndex], autoId)
				if err != nil {
					switch {
					case IsTooManyRequestsError(err):
						apiKeyIndex++
						log.Infow("parsing cars", "processor #", processor, "API Key", "changed")
					default:
						//log.Errorw("parsing cars", "processor #", processor, "failed to process auto-id", autoId, "API Key", apiKeys[apiKeyIndex], "ERROR", err)
						log.Errorw("parsing cars", "processor #", processor, "failed to process auto-id", autoId)
					}

					gorErr := ErrorFromGoroutine{
						Err:          err,
						FailedAutoId: autoId,
					}
					errCh <- gorErr

					continue
				}
				log.Infow("parsing cars", "processor #", processor, "finished processing auto ID", autoId)
			}
		}(proc)
	}

	go func(carIds []string) {
		for _, carId := range carIds {
			ch <- carId
		}
		close(ch)
	}(carIds)

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for errFromCh := range errCh {
		err1 := InsertError(ctx, dbInst, errFromCh)
		if err1 != nil {
			log.Errorw("failed to log error", "ERROR", err1)
		}
	}
}
