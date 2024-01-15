-- +goose Up
CREATE TABLE cars
(
    id            UUID PRIMARY KEY,
    user_id       INT         NOT NULL,
    auto_id       INT         NOT NULL UNIQUE,
    manufacturer  VARCHAR(36) NOT NULL,
    model         VARCHAR(36) NOT NULL,
    add_date      DATE        NOT NULL,
    update_date   DATE        NOT NULL,
    expire_date   DATE        NOT NULL,
    sold_date     DATE        NOT NULL,
    year          int         NOT NULL,
    body_style    VARCHAR(36) NOT NULL,
    fuel_type     VARCHAR(16) NOT NULL,
    gearbox_type  VARCHAR(16) NOT NULL,
    drive         VARCHAR(4)  NOT NULL,
    main_currency VARCHAR(4)  NOT NULL,
    vin           VARCHAR(32) NOT NULL,
    vin_svg       VARCHAR     NOT NULL,
    url           VARCHAR     NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS cars;
