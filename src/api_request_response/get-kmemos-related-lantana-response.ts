// ˅
'use strict';

import { Kmemo } from '../lantana_data/kmemo';
import { LantanaResponse } from './lantana-response';

// ˄

export class GetKmemosRelatedLantanaResponse extends LantanaResponse {
    // ˅

    // ˄

    kmemos: Array<Kmemo>;

    constructor() {
        // ˅
        super()
        this.kmemos = new Array<Kmemo>()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
