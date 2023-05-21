// ˅
'use strict';

import { Kmemo } from './kmemo';
import { Tag } from './tag';
import { Text } from './text';

// ˄

export class KmemoInfo {
    // ˅

    // ˄

    kmemo: Kmemo;

    tags: Array<Tag>;

    texts: Array<Text>;

    constructor() {
        // ˅
        this.kmemo = new Kmemo()
        this.tags = new Array<Tag>()
        this.texts = new Array<Text>()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
