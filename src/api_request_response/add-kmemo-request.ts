// ˅
'use strict';

import { Kmemo } from '../lantana_data/kmemo';
import { LantanaRequest } from './lantana-request';

// ˄

export class AddKmemoRequest extends LantanaRequest {
    // ˅

    // ˄

    kmemo: Kmemo;

    constructor() {
        // ˅
        super()
        this.kmemo = new Kmemo()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
