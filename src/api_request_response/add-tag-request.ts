// ˅
'use strict';

import { LantanaRequest } from './lantana-request';
import { Tag } from '../lantana_data/tag';

// ˄

export class AddTagRequest extends LantanaRequest {
    // ˅

    // ˄

    tag: Tag;

    constructor() {
        // ˅
        super()
        this.tag = new Tag()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
