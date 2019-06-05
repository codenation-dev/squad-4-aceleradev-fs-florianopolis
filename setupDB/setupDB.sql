CREATE DATABASE uati;

\c uati;

CREATE TABLE IF NOT EXISTS public_func (
func_id SERIAL,
func_name TEXT NOT NULL,
func_wage NUMERIC(10,2) DEFAULT 0.00,
place TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS customer	(
cust_id SERIAL,
cust_name TEXT NOT NULL,
cust_wage NUMERIC(10,2) DEFAULT 0.00,
isPublic bit DEFAULT NULL,
sentWarning TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS users(
uid SERIAL,
email TEXT NOT NULL,
password TEXT
);

CREATE TABLE IF NOT EXISTS warnings(
warning_id SERIAL,
dt TIMESTAMP,
message TEXT,
sent_to text,
from_customer TEXT
);
