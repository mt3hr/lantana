// ˅
'use strict';

// ˄

export class ApplicationConfig {
[x: string]: any;
    // ˅
    
    // ˄

    hidden_tags: Array<string>;

    uncheck_tags: Array<string>;

    tag_struct: any;

    constructor() {
        // ˅
        this.hidden_tags = new Array<string>()
        this.uncheck_tags = new Array<string>()
        this.tag_struct = {}
        // ˄
    }

    // ˅
    
    // ˄
}

// ˅

// ˄
