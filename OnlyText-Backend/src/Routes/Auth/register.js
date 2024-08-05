const express = require("express");
const validate = require("../../Middleware/Validators/auth.register");

const registerRouter = express.Router();

registerRouter.post("/", [validate], (req, res) => {
    
});
