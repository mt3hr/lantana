// ˅
'use strict';

import { Lantana } from '../lantana_data/lantana';
import { LantanaRequest } from './lantana-request';

// ˄

export class AddLantanaRequest extends LantanaRequest {
    // ˅

    // ˄

    lantana: Lantana;

    constructor() {
        // ˅
        super()
        this.lantana = new Lantana()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
