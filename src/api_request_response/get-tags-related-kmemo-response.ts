// ˅
'use strict';

import { Tag } from "@/lantana_data/tag";
import { LantanaResponse } from "./lantana-response";


// ˄

export class GetTagsRelatedKmemoResponse extends LantanaResponse {
    // ˅
    
    // ˄

    tags: Array<Tag>;

    constructor() {
        // ˅
        super()
        this.tags = new Array<Tag>()
        // ˄
    }

    // ˅
    
    // ˄
}

// ˅

// ˄
