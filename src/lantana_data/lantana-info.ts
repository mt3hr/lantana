// ˅
'use strict';

import { Kmemo } from './kmemo';
import { Lantana } from './lantana';
import { Tag } from './tag';
import { Text } from './text';

// ˄

export class LantanaInfo {
    // ˅
    
    // ˄

    lantana: Lantana;

    kmemos: Array<Kmemo>;

    tags: Array<Tag>;

    texts: Array<Text>;

    constructor() {
        // ˅
        this.lantana = new Lantana()
        this.kmemos = new Array<Kmemo>()
        this.tags = new Array<Tag>()
        this.texts = new Array<Text>()
        // ˄
    }

    // ˅
    
    // ˄
}

// ˅

// ˄
