// ˅
'use strict';

import { LantanaSearchType } from './lantana-search-type';

// ˄

export class LantanaSearchQuery {
    // ˅

    // ˄

    tags: Array<string>;

    words: string;

    mood: number;

    lantana_search_type: LantanaSearchType;

    constructor() {
        // ˅
        this.tags = new Array<string>()
        this.words = ""
        this.mood = 0
        this.lantana_search_type = LantanaSearchType.all
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
