// ˅
'use strict';

import { LantanaResponse } from './lantana-response';
import { Tag } from '../lantana_data/tag';

// ˄

export class GetTagsRelatedLantanaResponse extends LantanaResponse {
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
