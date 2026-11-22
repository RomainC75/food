import { Uuids } from "../domain/uuid";

export class InMemUuidGen implements Uuids{
    nextuuid: string;
    constructor(){
        this.nextuuid = ""
    }

    setNextUuid(next: string){
        this.nextuuid = next
    }

    now(): string{
        return this.nextuuid
    }
}