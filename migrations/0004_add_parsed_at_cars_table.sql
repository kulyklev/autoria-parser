-- +goose Up
ALTER TABLE "cars" ADD "parsed_at" DATE;

-- +goose Down
ALTER TABLE "cars" DROP COLUMN "parsed_at";
