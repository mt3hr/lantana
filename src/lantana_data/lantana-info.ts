// ˅
'use strict';

import { KmemoInfo } from './kmemo-info';
import { Lantana } from './lantana';
import { Tag } from './tag';
import { Text } from './text';

// ˄

export class LantanaInfo {
    // ˅

    // ˄

    lantana: Lantana;

    kmemo_infos: Array<KmemoInfo>;

    tags: Array<Tag>;

    texts: Array<Text>;

    constructor() {
        // ˅
        this.lantana = new Lantana()
        this.kmemo_infos = new Array<KmemoInfo>()
        this.tags = new Array<Tag>()
        this.texts = new Array<Text>()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
