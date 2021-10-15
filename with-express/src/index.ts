import express,{Request, Response} from 'express';
import {router} from './routers/loginRoute';
import bodyParser from 'body-parser';
import cookieSession from 'cookie-session';
import './controllers/LoginControllers';
import { AppRouter } from './AppRouter';

const app =express();

app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(cookieSession({
    keys: ['someHash']
}));
app.use(router);
app.use(AppRouter.getInstance());

app.listen(3000,()=>{
    console.log("listening on port 3000");
})