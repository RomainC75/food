import { v7 as uuidv7 } from 'uuid';

export class UuidGen{
    constructor(){
    }

    now(): string{
        return uuidv7();
    }
}