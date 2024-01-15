-- +goose Up
ALTER TABLE "cars" ADD "location" json;

-- +goose Down
ALTER TABLE "cars" DROP COLUMN "location";
