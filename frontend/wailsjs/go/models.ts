export namespace types {
	
	export class Item {
	    name: string;
	    description: string;
	    url: string;
	    thumbnail: string;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.url = source["url"];
	        this.thumbnail = source["thumbnail"];
	    }
	}

}

