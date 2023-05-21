// ˅
'use strict';

import { LantanaRequest } from './lantana-request';
import { Text } from '../lantana_data/text';

// ˄

export class AddTextRequest extends LantanaRequest {
    // ˅

    // ˄

    text: Text;

    constructor() {
        // ˅
        super()
        this.text = new Text()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
