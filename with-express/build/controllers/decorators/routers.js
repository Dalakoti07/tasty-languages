"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.del = exports.patch = exports.put = exports.post = exports.get = void 0;
require("reflect-metadata");
var Methods_1 = require("./Methods");
function routerBinder(method) {
    return function (path) {
        return function (target, key, desc) {
            Reflect.defineMetadata('path', path, target, key);
            Reflect.defineMetadata('method', method, target, key);
        };
    };
}
exports.get = routerBinder(Methods_1.Methods.get);
exports.post = routerBinder(Methods_1.Methods.post);
exports.put = routerBinder(Methods_1.Methods.put);
exports.patch = routerBinder(Methods_1.Methods.patch);
exports.del = routerBinder(Methods_1.Methods.del);
