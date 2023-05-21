// ˅
'use strict';

import { LantanaRequest } from './lantana-request';
import { LantanaSearchQuery } from '../lantana_data/lantana-search-query';

// ˄

export class SearchLantanaRequest extends LantanaRequest {
    // ˅

    // ˄

    query: LantanaSearchQuery;

    constructor() {
        // ˅
        super()
        this.query = new LantanaSearchQuery();

        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
