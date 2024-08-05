const validate = (req, res, next) => {
    if (req.body.username) {
        if (req.body.password) {
            next()
        } else {
            return res.staus(400).json({
                error_message: "no password provided",
            });
        }
    } else {
        return res.staus(400).json({
            error_message : "no username provided"
        })
    }
}

module.exports = validate