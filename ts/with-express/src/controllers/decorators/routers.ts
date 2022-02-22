import 'reflect-metadata';
import {Methods} from './Methods';

function routerBinder(method: string){
    return function(path: string){
        return function(target: any, key: string, desc: PropertyDescriptor){
            Reflect.defineMetadata('path',path, target, key);
            Reflect.defineMetadata('method',method, target, key);
        }
    }
}

export const get = routerBinder(Methods.get);
export const post = routerBinder(Methods.post);
export const put = routerBinder(Methods.put);
export const patch = routerBinder(Methods.patch);
export const del = routerBinder(Methods.del);
