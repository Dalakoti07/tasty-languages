"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.router = void 0;
var express_1 = require("express");
var router = (0, express_1.Router)();
exports.router = router;
router.post('/login', function (req, res) {
    var _a = req.body, email = _a.email, password = _a.password;
    if (email && password && email === "sd@sd.com" && password === "password") {
        req.session = { loggedIn: true };
        res.redirect('/');
    }
    else {
        res.send("invalid creds");
    }
});
function post(routerName) {
    return function (target, key, desc) {
        router.post(routerName, target[key]);
    };
}
function requireAuth(req, res, next) {
    if (req.session && req.session.loggedIn) {
        next();
        return;
    }
    else {
        res.status(403);
        res.send('Not allowed');
    }
}
router.get('/', function (req, res) {
    if (req.session && req.session.loggedIn) {
        res.send("\n            <div>\n                <h3>You are logged in</h3>\n                <a href=\"/logout\">Logout</a>\n            </div>\n        ");
    }
    else {
        res.send("\n        <div>\n            <h3>You are not logged in</h3>\n            <a href=\"/login\">Login</a>\n        </div>\n    ");
    }
});
router.get('/logout', function (req, res) {
    req.session = undefined;
    res.redirect('/');
});
router.get('/protected', requireAuth, function (req, res) {
    res.send('Welcome to protected route');
});
