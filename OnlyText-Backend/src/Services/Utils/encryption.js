const bcrypt = require("bcrypt");
const jwt = require("jsonwebtoken");

const hashPassword = async (password) => {
    let salt = await bcrypt.genSalt(10);
    let hashedPassword = await bcrypt.hash(password, salt);
    return hashedPassword;
};

const verifyPassword = async (password, hashedPassword) => {
    let isValid = await bcrypt.compare(password, hashedPassword);
    return isValid;
};

const decryptToken = (token) => {
    return jwt.decode(token);
};

const generateToken = (payload) => {
    return jwt.sign(
        payload,
        process.env.SERVER_SECRET,
        process.env.TOKEN_EXP_PERIOD || null
    );
};
module.exports = { hashPassword, verifyPassword, decryptToken, generateToken };
