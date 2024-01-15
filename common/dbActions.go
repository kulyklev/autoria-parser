package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	uniqueViolation = "23505"
	undefinedTable  = "42P01"
)

// Set of error variables for CRUD operations.
var (
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrUndefinedTable    = errors.New("undefined table")
)

func dbInsert(ctx context.Context, dbInst *sqlx.DB, car CarAd) error {
	ctx, span := AddSpan(ctx, "parse.common.dbActions.dbInsert")
	defer span.End()

	const q = `
	INSERT INTO cars
		(id, user_id, auto_id, manufacturer, model, add_date, update_date, expire_date, sold_date, year, body_style, fuel_type, gearbox_type, drive, main_currency, vin, vin_svg, url, is_active, parsed_at, location)
	VALUES
		(:id, :user_id, :auto_id, :manufacturer, :model, :add_date, :update_date, :expire_date, :sold_date, :year, :body_style, :fuel_type, :gearbox_type, :drive, :main_currency, :vin, :vin_svg, :url, :is_active, :parsed_at, :location)`

	createModel := ToDBCar(car)
	createModel.IsActive = true
	_, err := dbInst.NamedExecContext(ctx, q, createModel)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code {
			case undefinedTable:
				return ErrUndefinedTable
			case uniqueViolation:
				_ = addNewPrice(ctx, dbInst, createModel, car)
				_ = makeCarActive(ctx, dbInst, createModel)
				return nil
			}
		}
		return fmt.Errorf("inserting car: %w", err)
	}

	const q1 = `
	INSERT INTO prices
		(id, car_id, uah, usd, eur, parsed_at)
	VALUES
		(:id, :car_id, :uah, :usd, :eur, :parsed_at)`

	priceModel := ToDBPrice(car)
	priceModel.CarId = createModel.Id

	_, err = dbInst.NamedExecContext(ctx, q1, priceModel)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code {
			case undefinedTable:
				return ErrUndefinedTable
			case uniqueViolation:
				return ErrDBDuplicatedEntry
			}
		}
		return fmt.Errorf("inserting price: %w", err)
	}

	return nil
}

func addNewPrice(ctx context.Context, dbInst *sqlx.DB, createModel DbCar, carAd CarAd) error {
	data := struct {
		AutoID int `db:"auto_id"`
	}{
		AutoID: createModel.AutoId,
	}

	const q = `
		SELECT *
		FROM cars
		WHERE auto_id = :auto_id;
`

	var rows *sqlx.Rows
	var err error

	rows, err = sqlx.NamedQueryContext(ctx, dbInst, q, data)
	if err != nil {
		return fmt.Errorf("failed to get car: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("not found car")
	}

	var car DbCar
	if err := rows.StructScan(&car); err != nil {
		return err
	}

	// Adding price
	// =================================================================================================================

	const q1 = `
	INSERT INTO prices
		(id, car_id, uah, usd, eur, parsed_at)
	VALUES
		(:id, :car_id, :uah, :usd, :eur, :parsed_at)`

	priceModel := ToDBPrice(carAd)
	priceModel.CarId = car.Id

	_, err = dbInst.NamedExecContext(ctx, q1, priceModel)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code {
			case undefinedTable:
				return ErrUndefinedTable
			case uniqueViolation:
				return ErrDBDuplicatedEntry
			}
		}
		return fmt.Errorf("inserting price: %w", err)
	}

	return nil
}

func updateCarLocation(ctx context.Context, dbInst *sqlx.DB, car CarAd) error {
	fmt.Printf("start: update car location with ID: %d\n", car.AutoData.AutoID)

	createModel := ToDBCar(car)
	const q1 = `
		UPDATE cars SET
			location = :location 
		WHERE auto_id = :auto_id`

	_, err := dbInst.NamedExecContext(ctx, q1, createModel)
	if err != nil {
		return fmt.Errorf("updating car with ID: %d err: %w", car.AutoData.AutoID, err)
	}
	fmt.Printf("finish: update car location with ID: %d\n", car.AutoData.AutoID)

	return nil
}

func DbSetCarsNotActive(ctx context.Context, log *zap.SugaredLogger, dbInst *sqlx.DB) error {
	ctx, span := AddSpan(ctx, "parse.common.dbActions.DbSetCarsNotActive")
	defer span.End()

	log.Infow("start: setting all cars NOT ACTIVE")
	defer log.Infow("finish: setting all cars NOT ACTIVE")

	const q = `
		UPDATE cars SET is_active = FALSE`

	_, err := dbInst.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func makeCarActive(ctx context.Context, dbInst *sqlx.DB, createModel DbCar) error {
	data := struct {
		AutoID int `db:"auto_id"`
	}{
		AutoID: createModel.AutoId,
	}

	const q = `
		UPDATE cars 
		SET is_active = TRUE 
		WHERE auto_id = :auto_id;`

	_, err := dbInst.NamedExecContext(ctx, q, data)
	if err != nil {
		return fmt.Errorf("updating is_active: %w", err)
	}

	return nil
}

func InsertError(ctx context.Context, dbInst *sqlx.DB, grErr ErrorFromGoroutine) error {
	ctx, span := AddSpan(ctx, "parse.common.dbActions.InsertError")
	defer span.End()

	errModel := ToDBError(grErr)

	const q1 = `
	INSERT INTO failed_parses
		(id, auto_id, err, created_at)
	VALUES
		(:id, :auto_id, :err, :created_at)`

	_, err := dbInst.NamedExecContext(ctx, q1, errModel)
	if err != nil {
		return fmt.Errorf("inserting error: %w", err)
	}

	return nil
}

func TruncateFailedParses(ctx context.Context, log *zap.SugaredLogger, dbInst *sqlx.DB) error {
	ctx, span := AddSpan(ctx, "parse.common.dbActions.TruncateFailedParses")
	defer span.End()

	log.Infow("start: truncating failed parses")
	defer log.Infow("finish: truncating failed parses")

	const q = `
		TRUNCATE TABLE failed_parses`

	_, err := dbInst.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func CountFailedParses(ctx context.Context, dbInst *sqlx.DB) (int, error) {
	var count int

	err := dbInst.QueryRowContext(ctx, "SELECT COUNT(*) FROM failed_parses").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetFailedAutoIds(ctx context.Context, dbInst *sqlx.DB) ([]string, error) {
	fmt.Println("start: getting failed auto ids")
	const q = `
		SELECT auto_id
		FROM failed_parses`

	var carIds []string
	var err error

	err = dbInst.SelectContext(ctx, &carIds, q)
	if err != nil {
		return nil, fmt.Errorf("failed to get failed car ids: %w", err)
	}

	fmt.Println("finish: getting failed auto ids")

	return carIds, nil
}
