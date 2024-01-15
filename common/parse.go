package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/jmoiron/sqlx"
)

var ApiRequestCount uint32

func ProcessAutoId(ctx context.Context, dbInst *sqlx.DB, client *http.Client, apiKey, autoId string) error {
	ctx, span := AddSpan(ctx, "parse.common.parse.ProcessAutoId")
	defer span.End()

	requestURL := fmt.Sprintf("https://developers.ria.com/auto/info?api_key=%s&auto_id=%s", apiKey, autoId)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return fmt.Errorf("client: could not create request with autoId: %s\nerror: %w\n", autoId, err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client: error making http request with autoId: %s\nerror: %w\n", autoId, err)
	}

	if err = checkResponse(res); err != nil {
		return err
	}

	atomic.AddUint32(&ApiRequestCount, 1)

	var carAd CarAd
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&carAd)
	if err != nil {
		return fmt.Errorf("client: could not decode response body with autoId: %s\nerror: %w\n", autoId, err)
	}

	err = dbInsert(ctx, dbInst, carAd)
	//err = updateCarLocation(ctx, dbInst, carAd)
	if err != nil {
		return fmt.Errorf("saving info with autoId: %s\nerror: %w\n", autoId, err)
	}

	return nil
}

func GetCarIds(ctx context.Context, client *http.Client, apiKey string) ([]string, error) {
	ctx, span := AddSpan(ctx, "parse.common.parse.GetCarIds")
	defer span.End()

	hasMorePages := true
	var pageId int
	var carIds []string

	for hasMorePages {
		requestURL := fmt.Sprintf("https://developers.ria.com/auto/search?api_key=%s&marka_id[0]=70&model_id[0]=652&s_yers[0]=2013&po_yers[0]=2020&marka_id[1]=70&model_id[1]=3009&s_yers[1]=2013&po_yers[1]=2020&marka_id[2]=70&model_id[2]=3167&s_yers[2]=2013&po_yers[2]=2020&marka_id[3]=70&model_id[3]=63961&s_yers[3]=2013&po_yers[3]=2020&marka_id[4]=84&model_id[4]=1653&s_yers[4]=2010&po_yers[4]=2016&marka_id[5]=84&model_id[5]=32103&s_yers[5]=2010&po_yers[5]=2016&marka_id[6]=84&model_id[6]=3155&s_yers[6]=2010&po_yers[6]=2016&marka_id[7]=84&model_id[7]=35449&s_yers[7]=2014&po_yers[7]=2020&marka_id[8]=84&model_id[8]=59942&s_yers[8]=2014&po_yers[8]=2020&marka_id[9]=84&model_id[9]=63351&s_yers[9]=2014&po_yers[9]=2020&marka_id[10]=84&model_id[10]=64469&s_yers[10]=2014&po_yers[10]=2020&marka_id[11]=84&model_id[11]=45343&s_yers[11]=2014&po_yers[11]=2020&marka_id[12]=84&model_id[12]=39690&s_yers[12]=2014&po_yers[12]=2020&marka_id[13]=84&model_id[13]=2805&s_yers[13]=2014&po_yers[13]=2020&marka_id[14]=84&model_id[14]=2093&s_yers[14]=2014&po_yers[14]=2020&engineVolumeFrom=2&countpage=100&with_photo=1&page=%d", apiKey, pageId)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		if err = checkResponse(res); err != nil {
			return nil, err
		}

		ApiRequestCount++

		var searchResult CarsSearchResult
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&searchResult)
		if err != nil {
			return nil, err
		}

		carIds = append(carIds, searchResult.Result.SearchResult.Ids...)

		if len(searchResult.Result.SearchResult.Ids) == 0 {
			hasMorePages = false
		}

		pageId++
	}

	return carIds, nil
}

func checkResponse(res *http.Response) error {
	switch res.StatusCode {
	case 200:
		return nil
	case 429:
		return NewTooManyRequestsError(
			fmt.Errorf("client error: too many requests"),
			http.StatusTooManyRequests,
		)
	default:
		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("request error: status code: %d, status: %s, body: FAILED TO DECODE", res.StatusCode, res.Status)
		}

		return NewRequestError(fmt.Errorf("request error"), res.StatusCode, res.Status, string(b))
	}
}
