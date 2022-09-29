CREATE TABLE IF NOT EXISTS stats
(
    "date" timestamp primary key,
    views int,
    clicks int,
    cost numeric(100, 2),
    cpc float,
    cpm float
);