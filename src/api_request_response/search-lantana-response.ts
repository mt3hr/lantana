// ˅
'use strict';

import { Lantana } from '../lantana_data/lantana';
import { LantanaResponse } from './lantana-response';

// ˄

export class SearchLantanaResponse extends LantanaResponse {
    // ˅

    // ˄

    lantanas: Array<Lantana>;

    constructor() {
        // ˅
        super()
        this.lantanas = new Array<Lantana>()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
