"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var express_1 = __importDefault(require("express"));
var loginRoute_1 = require("./routers/loginRoute");
var body_parser_1 = __importDefault(require("body-parser"));
var cookie_session_1 = __importDefault(require("cookie-session"));
require("./controllers/LoginControllers");
var AppRouter_1 = require("./AppRouter");
var app = (0, express_1.default)();
app.use(body_parser_1.default.urlencoded({
    extended: true
}));
app.use((0, cookie_session_1.default)({
    keys: ['someHash']
}));
app.use(loginRoute_1.router);
app.use(AppRouter_1.AppRouter.getInstance());
app.listen(3000, function () {
    console.log("listening on port 3000");
});
