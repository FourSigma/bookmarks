
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
        public id:string,
        public url:string,
        public hash:string,
        public data: Data, 
    ){}
}