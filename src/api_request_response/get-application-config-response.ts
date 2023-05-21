// ˅
'use strict';

import { ApplicationConfig } from '../lantana_data/application-config';
import { LantanaResponse } from './lantana-response';

// ˄

export class GetApplicationConfigResponse extends LantanaResponse {
    // ˅

    // ˄

    application_config: ApplicationConfig;

    constructor() {
        // ˅
        super()
        this.application_config = new ApplicationConfig()
        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
