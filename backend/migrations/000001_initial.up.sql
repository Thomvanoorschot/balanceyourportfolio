CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE fund (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name                varchar(255),
    currency            varchar(10),
    isin                varchar(20) UNIQUE,
    total_holdings      numeric,
    price               numeric,
    provider            varchar(255),
    external_identifier varchar(255) UNIQUE,
    outstanding_shares  numeric,
    effective_date      date
);

CREATE TABLE holding (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    ticker                varchar(255) UNIQUE,
    type                varchar(255),
    name                varchar(255),
    isin                varchar(20),
    sedol                varchar(20),
    cusip                varchar(20),
    sector              varchar(255)
);

CREATE TABLE fund_holding (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    fund_id uuid REFERENCES fund (id),
    holding_id uuid REFERENCES holding (id),
    amount numeric,
    percentage_of_total numeric,
    market_value numeric,
    CONSTRAINT unique_fund_holding UNIQUE (fund_id,holding_id)
);

CREATE TABLE fund_listing (
    fund_id uuid REFERENCES fund (id),
    ticker            varchar(255) UNIQUE
);

CREATE TABLE "user" (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY
);

CREATE TABLE "portfolio" (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid REFERENCES "user" (id),
    name  varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE portfolio_fund (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    portfolio_id uuid REFERENCES "portfolio" (id),
    fund_id uuid REFERENCES "fund" (id),
    amount numeric,
    CONSTRAINT unique_portfolio_fund UNIQUE (portfolio_id,fund_id)
);
