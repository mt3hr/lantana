// ˅
'use strict';

import { LantanaRequest } from './lantana-request';

// ˄

export class DeleteTagRequest extends LantanaRequest {
    // ˅

    // ˄

    tag_id: string;

    constructor() {
        // ˅
        super()
        this.tag_id = ""
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
