// ˅
'use strict';

import { LantanaResponse } from './lantana-response';
import { Text } from '../lantana_data/text';

// ˄

export class GetTextsRelatedLantanaResponse extends LantanaResponse {
    // ˅

    // ˄

    texts: Array<Text>;

    constructor() {
        // ˅
        super()
        this.texts = new Array<Text>()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
