export namespace types {
	
	export class JSResp {
	    success: boolean;
	    msg: string;
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new JSResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}
	export class Paginated {
	    limit: number;
	    offset: number;
	
	    static createFrom(source: any = {}) {
	        return new Paginated(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.limit = source["limit"];
	        this.offset = source["offset"];
	    }
	}
	export class PreferencesGeneral {
	    theme: string;
	    language: string;
	    font: string;
	    fontFamily: string[];
	    fontSize: number;
	    checkUpdate: boolean;
	    scanSize: number;
	    maxConcurrency: number;
	    maxDepth: number;
	    appName: string;
	    timeFormat: string;
	    timeZone: string;
	    tlds: string[];
	    ignoreLocal: boolean;
	    proxyEnabled: boolean;
	    proxyURL: string;
	    userAgent: string;
	    enableProcessing: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PreferencesGeneral(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.language = source["language"];
	        this.font = source["font"];
	        this.fontFamily = source["fontFamily"];
	        this.fontSize = source["fontSize"];
	        this.checkUpdate = source["checkUpdate"];
	        this.scanSize = source["scanSize"];
	        this.maxConcurrency = source["maxConcurrency"];
	        this.maxDepth = source["maxDepth"];
	        this.appName = source["appName"];
	        this.timeFormat = source["timeFormat"];
	        this.timeZone = source["timeZone"];
	        this.tlds = source["tlds"];
	        this.ignoreLocal = source["ignoreLocal"];
	        this.proxyEnabled = source["proxyEnabled"];
	        this.proxyURL = source["proxyURL"];
	        this.userAgent = source["userAgent"];
	        this.enableProcessing = source["enableProcessing"];
	    }
	}
	export class PreferencesBehavior {
	    welcomed: boolean;
	    asideWidth: number;
	    windowWidth: number;
	    windowHeight: number;
	    windowMaximised: boolean;
	    windowPosX: number;
	    windowPosY: number;
	    darkMode: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PreferencesBehavior(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.welcomed = source["welcomed"];
	        this.asideWidth = source["asideWidth"];
	        this.windowWidth = source["windowWidth"];
	        this.windowHeight = source["windowHeight"];
	        this.windowMaximised = source["windowMaximised"];
	        this.windowPosX = source["windowPosX"];
	        this.windowPosY = source["windowPosY"];
	        this.darkMode = source["darkMode"];
	    }
	}
	export class Preferences {
	    behavior: PreferencesBehavior;
	    general: PreferencesGeneral;
	
	    static createFrom(source: any = {}) {
	        return new Preferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.behavior = this.convertValues(source["behavior"], PreferencesBehavior);
	        this.general = this.convertValues(source["general"], PreferencesGeneral);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

