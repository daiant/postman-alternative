export namespace http_types {
	
	export class Response {
	    status: string;
	    statusCode: number;
	    body: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.statusCode = source["statusCode"];
	        this.body = source["body"];
	    }
	}

}

