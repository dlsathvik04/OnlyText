const { decryptToken } = require("../Services/Utils/encryption");

const authenticate = (req, res, next) => {
    if (req.headers.authorization) {
        try {
            const decoded = decryptToken(
                req.headers.authorization.split(" ")[0]
            );
            req.decoded = decoded;
            next();
        } catch {
            return res.status(400).json({
                error_message: "invalid token",
            });
        }
    } else {
        return res.status(400).json({
            error_message:
                "no authentication token provided, authentication failed",
        });
    }
};

module.exports = authenticate;
