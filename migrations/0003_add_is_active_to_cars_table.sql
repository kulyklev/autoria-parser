-- +goose Up
ALTER TABLE "cars" ADD "is_active" BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE "cars" DROP COLUMN "is_active";
