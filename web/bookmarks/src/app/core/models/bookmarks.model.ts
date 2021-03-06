
export interface Data{
    title: string,
    type: string,
    url: string,
    site: string,
    siteName: string,
    description: string,
    locale: string,
    image: string,
    content: string,
}

export class Bookmark{
    constructor(
        public id?:string,
        public url?:string,
        public data?: Data, 
    ){}

    isEmpty(): boolean{
        return this.url == '' || this.data == null || this.data == undefined;
    }

    static fromJSON(json:any): Bookmark{
        return new Bookmark(
           json.id,
           json.url,
           json.data,
        );
    }

    public toJSON(): any{
        return {
            id: this.id,
            url: this.url,
            data: this.data,
        }
    }
}

