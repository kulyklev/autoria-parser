-- ==============================
-- Select all prices for given car
-- ==============================
SELECT c.auto_id, p.usd, p.parsed_at
FROM cars c
         JOIN prices p ON c.id = p.car_id
WHERE c.auto_id in (
                    '33719029',
                    '33728046'
    )
GROUP BY c.auto_id, p.parsed_at, p.usd
ORDER BY c.auto_id, p.parsed_at;

-- ==============================
-- Count car prices
-- ==============================
SELECT c.auto_id, c.user_id, c.drive, c.gearbox_type, c.main_currency, url, count(p.id) AS p_count
FROM cars c
         JOIN prices p ON c.id = p.car_id
WHERE manufacturer = 'Skoda'
GROUP BY c.auto_id, c.user_id, c.drive, c.gearbox_type, c.main_currency, url
ORDER BY p_count;

-- ==============================
-- Select latest prices
-- ==============================
SELECT c.id,
       c.auto_id,
       c.user_id,
       c.drive,
       c.gearbox_type,
       c.main_currency,
       c.url,
       c.is_active,
       c.parsed_at,
       c.model,
       p1.usd,
       c.attitude
FROM cars c
         JOIN prices p1 ON (c.id = p1.car_id)
         LEFT OUTER JOIN prices p2 ON (c.id = p2.car_id AND
                                       (p1.parsed_at < p2.parsed_at OR (p1.parsed_at = p2.parsed_at AND p1.id < p2.id)))
WHERE p2.id IS NULL
  AND c.manufacturer in ('Volkswagen', 'Skoda')
--   AND c.drive in ('AWD', '')
--   AND c.drive in ('FWD')
  AND c.attitude = 'I WOULD BUY IT'
--   AND c.attitude = 'MEH'
--   AND c.url NOT LIKE '%scout%'
  AND c.is_active = true
-- ORDER BY c.parsed_at DESC;
ORDER BY c.model, p1.usd;

-- ==============================
-- Count cars per city
-- ==============================
SELECT location ->> 'Name' AS "location_name",
       count(*)            AS "cars_in_location"
FROM cars
WHERE is_active = true
  AND attitude IN ('I LIKE IT', 'I WOULD BUY IT')
GROUP BY location ->> 'Name'
ORDER BY cars_in_location DESC;

-- ==============================
-- Select cars and show price delta
-- ==============================
SELECT distinct auto_id, delta, parsed_at
FROM (SELECT c.id,
             c.auto_id,
             c.parsed_at,
             p1.usd,
             p2.usd,
             p2.usd - p1.usd as delta
      FROM cars c
               JOIN prices p1 ON (c.id = p1.car_id)
               LEFT OUTER JOIN prices p2 ON (c.id = p2.car_id)
      WHERE p2.usd > p1.usd
        AND c.manufacturer = 'Skoda'
        AND c.drive in ('AWD', '')
        AND c.is_active = true
      GROUP BY c.id, p1.usd, p2.usd, c.auto_id, p1.usd, c.id
      HAVING count(distinct c.auto_id) = 1
      ORDER BY p1.usd) x;

-- ==============================
-- Select cars by region
-- ==============================
SELECT *
FROM cars
WHERE attitude IN ('I WOULD BUY IT', 'I LIKE IT')
  AND is_active = true
  AND location ->> 'StateID' IN ('3', '5', '10', '18', '22');
--   Ternopil - '3', Lviv - '5', Kyiv - '10', Lutsk - '18', Uzhhorod - '22'

-- ==============================
-- Select cars with price changes
-- ==============================
SELECT car_id, parsed_at, COUNT(*)
FROM prices
GROUP BY car_id, parsed_at
HAVING COUNT(*) > 1;

-- ==============================
-- Select and delete duplicate prices
-- ==============================
SELECT E.id,
       E.car_id,
       E.parsed_at,
       T.rank
FROM prices E
         INNER JOIN
     (SELECT *,
             RANK() OVER (PARTITION BY car_id,
                 parsed_at
                 ORDER BY id) rank
      FROM prices) T ON E.ID = t.ID;

SELECT id
FROM (SELECT *,
             row_number() OVER (PARTITION BY car_id,
                 parsed_at
                 ORDER BY id) AS row_num
      FROM prices) t
WHERE row_num > 1;

DELETE
FROM prices
WHERE id in (SELECT id
             FROM (SELECT *,
                          row_number() OVER (PARTITION BY car_id,
                              parsed_at
                              ORDER BY id) AS row_num
                   FROM prices) t
             WHERE row_num > 1);

-- ==============================
-- Set personal attitude
-- ==============================
-- 'MEH', 'I LIKE IT', 'I WOULD BUY IT'
UPDATE cars
SET attitude = 'MEH'
WHERE auto_id IN (
                  '',
                  ''
    );

UPDATE cars
SET attitude = 'I LIKE IT'
WHERE auto_id IN (
                  '',
                  ''
    );

UPDATE cars
SET attitude = 'I WOULD BUY IT'
WHERE auto_id IN (
                  '',
                  ''
    );
