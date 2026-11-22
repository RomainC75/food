import { Uuids } from "../domain/uuid";

export class InMemUuidGen implements Uuids{
    nextuuids: string[];
    index=0;
    constructor(){
        this.nextuuids = ["a","b","c","d","e","f","g","h","i"]
    }

    now(): string{
        const nowId = this.nextuuids[this.index];
        if (this.index < this.nextuuids.length-1){
            this.index++;
        } else {
            this.index = 0;
        }
        
        return nowId;
    }
}