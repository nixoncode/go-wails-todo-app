export namespace main {
	
	export class Item {
	    id: number;
	    name: string;
	    is_done: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.is_done = source["is_done"];
	    }
	}

}

