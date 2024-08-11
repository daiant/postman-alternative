export namespace http_types {
	
	export class Request {
	    url: string;
	    method: string;
	    body: string;
	    params: string[][];
	    headers: string[][];
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.method = source["method"];
	        this.body = source["body"];
	        this.params = source["params"];
	        this.headers = source["headers"];
	    }
	}
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

