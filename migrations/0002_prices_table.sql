-- +goose Up
CREATE TABLE prices
(
    id        UUID PRIMARY KEY,
    car_id    UUID  NOT NULL,
    uah       INT  NOT NULL,
    usd       INT  NOT NULL,
    eur       INT  NOT NULL,
    parsed_at DATE NOT NULL,

    FOREIGN KEY (car_id) REFERENCES cars (id)
);

-- +goose Down
DROP TABLE IF EXISTS prices;
