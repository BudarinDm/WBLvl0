CREATE TABLE IF NOT EXISTS orders
(
    order_uid          varchar PRIMARY KEY NOT NULL UNIQUE,
    track_number       varchar(128)        NOT NULL,
    entry              varchar(128)        NOT NULL,
    locale             varchar(128)        NOT NULL,
    internal_signature varchar(128)        NOT NULL,
    customer_id        varchar(128)        NOT NULL,
    delivery_service   varchar(128)        NOT NULL,
    shardkey           varchar(128)        NOT NULL,
    sm_id              bigint,
    date_created       timestamp           NOT NULL,
    oof_shard          varchar(128)        NOT NULL
);

CREATE TABLE IF NOT EXISTS deliveries
(
    order_uid varchar references orders (order_uid) on delete cascade NOT NULL,
    name      varchar(128) NOT NULL,
    phone     varchar(128) NOT NULL,
    zip       varchar(128) NOT NULL,
    city      varchar(128) NOT NULL,
    address   varchar(256) NOT NULL,
    region    varchar(256) NOT NULL,
    email     varchar(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS items
(
    chrt_id      bigint PRIMARY KEY NOT NULL,
    track_number varchar(256)       NOT NULL,
    price        bigint             NOT NULL,
    rid          varchar(128)               ,
    name         varchar(128)       NOT NULL,
    sale         int                NOT NULL,
    size         varchar(128)       NOT NULL,
    total_price  bigint             NOT null,
    nm_id        bigint                     ,
    brand        varchar(128)       NOT NULL,
    status       int                NOT NULL,

    order_uid    varchar references orders (order_uid) on delete cascade NOT NULL
);

CREATE TABLE IF NOT EXISTS payments
(
    transaction   varchar PRIMARY KEY NOT NULL UNIQUE,
    request_id    varchar,
    currency      varchar(128)        NOT NULL,
    provider      varchar(128)        NOT NULL,
    amount        bigint              NOT NULL,
    payment_dt    bigint              NOT NULL,
    bank          varchar(128)        NOT NULL,
    delivery_cost bigint              NOT NULL,
    goods_total   bigint              NOT NULL,
    custom_fee    bigint              NOT NULL,

    order_uid     varchar references orders (order_uid) on delete cascade NOT NULL
);