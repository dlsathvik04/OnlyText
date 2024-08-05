const pgp = require("pg-promise")();

const conn_url = process.env.DB_URL;

const cn = {
    host: process.env.DB_HOST,
    port: process.env.DB_PORT,
    database: process.env.DB_NAME,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
};

const db = pgp(conn_url || cn);

module.exports = db;
