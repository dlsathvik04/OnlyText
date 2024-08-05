const validate = (req, res, next) => {
    if (req.body.firstname == null) {
        return res.status(400).json({
            error_message: "first name is required",
        });
    } else if (req.body.username == null) {
        return res.status(400).json({
            error_message: "username is required",
        });
    } else if (req.body.password == null) {
        return res.status(400).json({
            error_message: "password is required",
        })
    } else if (req.body.email == null) {
        return res.status(400).json({
            error_message: "email is required"
        })
    } else {
        next();
    }
};

module.exports = validate;
