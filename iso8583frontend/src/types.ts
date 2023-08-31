export interface Message {
    fields:      { [key: string]: Field };
    mti:         Mti;
    bitmap:      string;
    raw:         string;
    logFileName: string;
    lineNumber:  string;
}

export interface Field {
    length: string;
    value:  string;
    raw:    string;
}

export interface Mti {
    version:  number;
    class:    number;
    function: number;
    origin:   number;
}

export interface IParseResult {
    id: number;
    name: string;
    date: Date;
    status: 'downloading' | 'parsing' | 'done' | 'error';
    type: 'ssh' | 'upload'
    messages: Message[];
}

export interface Search {
    mtiVersion?: string;
    mtiClass?: string;
    mtiFunction?: string;
    mtiOrigin?: string;
    bitmap?: string;
    logFiles?: string[];
    fields?: { [key: string]: string };
}