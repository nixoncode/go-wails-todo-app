export namespace main {
	
	export class Item {
	    ID: number;
	    Name: string;
	    IsDone: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.IsDone = source["IsDone"];
	    }
	}

}

