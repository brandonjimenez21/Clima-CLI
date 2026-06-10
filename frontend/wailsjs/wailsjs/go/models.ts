export namespace model {
	
	export class ForecastPoint {
	    Time: string;
	    Temp: number;
	
	    static createFrom(source: any = {}) {
	        return new ForecastPoint(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Time = source["Time"];
	        this.Temp = source["Temp"];
	    }
	}
	export class Weather {
	    City: string;
	    Temperature: number;
	    FeelsLike: number;
	    Humidity: number;
	    Description: string;
	    Condition: string;
	    WindSpeed: number;
	    ThermalSensation: string;
	    Forecast: ForecastPoint[];
	    IsNight: boolean;
	    UVIndex: number;
	    Sunrise: string;
	    Sunset: string;
	    RainProb: number;
	    IsGoldenHour: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Weather(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.City = source["City"];
	        this.Temperature = source["Temperature"];
	        this.FeelsLike = source["FeelsLike"];
	        this.Humidity = source["Humidity"];
	        this.Description = source["Description"];
	        this.Condition = source["Condition"];
	        this.WindSpeed = source["WindSpeed"];
	        this.ThermalSensation = source["ThermalSensation"];
	        this.Forecast = this.convertValues(source["Forecast"], ForecastPoint);
	        this.IsNight = source["IsNight"];
	        this.UVIndex = source["UVIndex"];
	        this.Sunrise = source["Sunrise"];
	        this.Sunset = source["Sunset"];
	        this.RainProb = source["RainProb"];
	        this.IsGoldenHour = source["IsGoldenHour"];
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

