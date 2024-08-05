const express = require('express')
const validate = require('../../Middleware/Validators/auth.login')

const loginRouter = express.Router()

loginRouter.post("/", [validate], (req, res) => {
    
})