CREATE TABLE IF NOT EXISTS stats
(
    stats_date timestamp primary key,
    views int,
    clicks int,
    cost float
);