// initialize dotenv for environment variable setup
require("dotenv").config();

// require the dependecies here
const express = require("express");
const db = require("./src/Services/Utils/db");
const authenticate = require("./src/Middleware/auth");

// create express application
const app = express();

// configure Middleware here
app.use(express.json());
app.use(authenticate)

// configure root behaviour if need
app.get("/", (req, res) => {
    res.send("This is the root of the application");
});

// configure routers here

// Declare the port number
const portNumber = process.env.PORT_NUMBER || 5001;

// Configurations before starting a server
// Database
db.one("SELECT $1 AS value", 123)
    .then((data) => {
        console.log("Connection to database suceeded....");
    })
    .catch((error) => {
        console.log("ERROR:", error);
    });

// Start listening on the given port number
const server = app.listen(portNumber, () => {
    console.log(`listening on port portNumber ${portNumber}`);
});

// Write the server event managers here
server.on("connection", (client) => {
    console.log(`Connection from ${client.address().address}`);
});
