-- +goose Up
CREATE TYPE attitude AS ENUM ('MEH', 'I LIKE IT', 'I WOULD BUY IT');
ALTER TABLE "cars" ADD "attitude" attitude DEFAULT 'MEH';

-- +goose Down
ALTER TABLE "cars" DROP COLUMN "attitude";
