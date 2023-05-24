// ˅
'use strict';

import { LantanaResponse } from "./lantana-response";

// ˄

export class GetTagNamesResponse extends LantanaResponse {
    // ˅
    
    // ˄

    tag_names: Array<string>;

    constructor() {
        // ˅
        super()
        this.tag_names = new Array<string>()
        // ˄
    }

    // ˅
    
    // ˄
}

// ˅

// ˄
