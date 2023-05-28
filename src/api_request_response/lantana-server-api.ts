// ˅
'use strict';

import { KmemoInfo } from "@/lantana_data/kmemo-info";
import { LantanaInfo } from "@/lantana_data/lantana-info";
import { AddKmemoRequest } from "./add-kmemo-request";
import { AddKmemoResponse } from "./add-kmemo-response";
import { AddLantanaRequest } from "./add-lantana-request";
import { AddLantanaResponse } from "./add-lantana-response";
import { AddTagRequest } from "./add-tag-request";
import { AddTagResponse } from "./add-tag-response";
import { AddTextRequest } from "./add-text-request";
import { AddTextResponse } from "./add-text-response";
import { DeleteKmemoRequest } from "./delete-kmemo-request";
import { DeleteKmemoResponse } from "./delete-kmemo-response";
import { DeleteLantanaRequest } from "./delete-lantana-request";
import { DeleteLantanaResponse } from "./delete-lantana-response";
import { DeleteTagRequest } from "./delete-tag-request";
import { DeleteTagResponse } from "./delete-tag-response";
import { DeleteTextRequest } from "./delete-text-request";
import { DeleteTextResponse } from "./delete-text-response";
import { GetApplicationConfigRequest } from "./get-application-config-request";
import { GetApplicationConfigResponse } from "./get-application-config-response";
import { GetKmemosRelatedLantanaRequest } from "./get-kmemos-related-lantana-request";
import { GetKmemosRelatedLantanaResponse } from "./get-kmemos-related-lantana-response";
import { GetTagsRelatedLantanaRequest } from "./get-tags-related-lantana-request";
import { GetTagsRelatedLantanaResponse } from "./get-tags-related-lantana-response";
import { GetTextsRelatedLantanaRequest } from "./get-texts-related-lantana-request";
import { SearchLantanaRequest } from "./search-lantana-request";
import { SearchLantanaResponse } from "./search-lantana-response";
import { LantanaAPIAddress } from "./lantana-api-address";
import { GetTextsRelatedLantanaResponse } from "./get-texts-related-lantana-response";
import { Lantana } from "@/lantana_data/lantana";
import { Kmemo } from "@/lantana_data/kmemo";
import { GetTagsRelatedKmemoRequest } from "./get-tags-related-kmemo-request";
import { GetTextsRelatedKmemoRequest } from "./get-texts-related-kmemo-request";
import { GetTagsRelatedKmemoResponse } from "./get-tags-related-kmemo-response";
import { GetTextsRelatedKmemoResponse } from "./get-texts-related-kmemo-response";
import { GetTagNamesRequest } from "./get-tag-names-request";
import { GetTagNamesResponse } from "./get-tag-names-response";

// ˄

export class LantanaServerAPI {
    // ˅

    // ˄

    async search_lantana(request: SearchLantanaRequest): Promise<SearchLantanaResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.search_lantana_address, {
            method: LantanaAPIAddress.search_lantana_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: SearchLantanaResponse = json
        for (let i = 0; i < response.lantanas.length; i++) {
            response.lantanas[i].time = new Date(response.lantanas[i].time)
        }
        return response
        // ˄
    }

    async get_lantana_info(lantana: Lantana): Promise<LantanaInfo> {
        // ˅
        const response = new LantanaInfo()
        response.lantana = lantana

        const get_kmemos_related_lantana_request = new GetKmemosRelatedLantanaRequest()
        get_kmemos_related_lantana_request.lantana_id = lantana.lantana_id
        const get_kmemos_related_lantana_response = await this.get_kmemos_related_lantana(get_kmemos_related_lantana_request)
        for (let i = 0; i < get_kmemos_related_lantana_response.kmemos.length; i++) {
            const kmemo_info = await this.get_kmemo_info(get_kmemos_related_lantana_response.kmemos[i])
            response.kmemo_infos.push(kmemo_info)
        }

        const get_tags_related_lantana_request = new GetTagsRelatedLantanaRequest()
        get_tags_related_lantana_request.lantana_id = lantana.lantana_id
        const get_tags_related_lantana_response = await this.get_tags_related_lantana(get_tags_related_lantana_request)
        response.tags = get_tags_related_lantana_response.tags

        const get_texts_related_lantana_request = new GetTextsRelatedLantanaRequest()
        get_texts_related_lantana_request.lantana_id = lantana.lantana_id
        const get_texts_related_lantana_response = await this.get_texts_related_lantana(get_texts_related_lantana_request)
        response.texts = get_texts_related_lantana_response.texts

        return response
        // ˄
    }

    async add_lantana(request: AddLantanaRequest): Promise<AddLantanaResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.add_lantana_address, {
            method: LantanaAPIAddress.add_lantana_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: AddLantanaResponse = json
        return response
        // ˄
    }

    async delete_lantana(request: DeleteLantanaRequest): Promise<DeleteLantanaResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.delete_lantana_address, {
            method: LantanaAPIAddress.delete_lantana_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: DeleteLantanaResponse = json
        return response
        // ˄
    }

    async get_kmemo_info(kmemo: Kmemo): Promise<KmemoInfo> {
        // ˅
        const response = new KmemoInfo()
        response.kmemo = kmemo

        const get_tags_related_kmemo_request = new GetTagsRelatedKmemoRequest()
        get_tags_related_kmemo_request.kmemo_id = kmemo.id
        const get_tags_related_kmemo_response = await this.get_tags_related_kmemo(get_tags_related_kmemo_request)
        response.tags = get_tags_related_kmemo_response.tags

        const get_texts_related_kmemo_request = new GetTextsRelatedKmemoRequest()
        get_texts_related_kmemo_request.kmemo_id = kmemo.id
        const get_texts_related_kmemo_response = await this.get_texts_related_kmemo(get_texts_related_kmemo_request)
        response.texts = get_texts_related_kmemo_response.texts

        return response
        // ˄
    }

    async add_kmemo(request: AddKmemoRequest): Promise<AddKmemoResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.add_kmemo_address, {
            method: LantanaAPIAddress.add_kmemo_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: AddKmemoResponse = json
        return response
        // ˄
    }

    async delete_kmemo(request: DeleteKmemoRequest): Promise<DeleteKmemoResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.delete_kmemo_address, {
            method: LantanaAPIAddress.delete_kmemo_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: DeleteKmemoResponse = json
        return response
        // ˄
    }

    async add_tag(request: AddTagRequest): Promise<AddTagResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.add_tag_address, {
            method: LantanaAPIAddress.add_tag_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: AddTagResponse = json
        return response
        // ˄
    }

    async delete_tag(request: DeleteTagRequest): Promise<DeleteTagResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.delete_tag_address, {
            method: LantanaAPIAddress.delete_tag_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: DeleteTagResponse = json
        return response
        // ˄
    }

    async add_text(request: AddTextRequest): Promise<AddTextResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.add_text_address, {
            method: LantanaAPIAddress.add_text_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: AddTextResponse = json
        return response
        // ˄
    }

    async delete_text(request: DeleteTextRequest): Promise<DeleteTextResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.delete_text_address, {
            method: LantanaAPIAddress.delete_text_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: DeleteTextResponse = json
        return response
        // ˄
    }

    async get_kmemos_related_lantana(request: GetKmemosRelatedLantanaRequest): Promise<GetKmemosRelatedLantanaResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.get_kmemos_related_lantana_address, {
            method: LantanaAPIAddress.get_kmemos_related_lantana_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetKmemosRelatedLantanaResponse = json
        for (let i = 0; i < response.kmemos.length; i++) {
            response.kmemos[i].time = new Date(response.kmemos[i].time)
        }
        return response
        // ˄
    }

    async get_tags_related_lantana(request: GetTagsRelatedLantanaRequest): Promise<GetTagsRelatedLantanaResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.get_tags_related_lantana_address, {
            method: LantanaAPIAddress.get_tags_related_lantana_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetTagsRelatedLantanaResponse = json
        for (let i = 0; i < response.tags.length; i++) {
            response.tags[i].time = new Date(response.tags[i].time)
        }
        return response
        // ˄
    }

    async get_texts_related_lantana(request: GetTextsRelatedLantanaRequest): Promise<GetTextsRelatedLantanaResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.get_texts_related_lantana_address, {
            method: LantanaAPIAddress.get_texts_related_lantana_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetTextsRelatedLantanaResponse = json
        for (let i = 0; i < response.texts.length; i++) {
            response.texts[i].time = new Date(response.texts[i].time)
        }
        return response
        // ˄
    }

    async get_tags_related_kmemo(request: GetTagsRelatedKmemoRequest): Promise<GetTagsRelatedKmemoResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.get_tags_related_kmemo_address, {
            method: LantanaAPIAddress.get_tags_related_kmemo_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetTagsRelatedKmemoResponse = json
        for (let i = 0; i < response.tags.length; i++) {
            response.tags[i].time = new Date(response.tags[i].time)
        }
        return response
        // ˄
    }

    async get_texts_related_kmemo(request: GetTextsRelatedKmemoRequest): Promise<GetTextsRelatedKmemoResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.get_texts_related_kmemo_address, {
            method: LantanaAPIAddress.get_texts_related_kmemo_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetTextsRelatedKmemoResponse = json
        for (let i = 0; i < response.texts.length; i++) {
            response.texts[i].time = new Date(response.texts[i].time)
        }
        return response
        // ˄
    }

    async get_tag_names(request: GetTagNamesRequest): Promise<GetTagNamesResponse> {
        const res = await fetch(LantanaAPIAddress.get_tag_names_address, {
            method: LantanaAPIAddress.get_tag_names_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetTagNamesResponse= json
        return response
    }

    async get_application_config(request: GetApplicationConfigRequest): Promise<GetApplicationConfigResponse> {
        // ˅
        const res = await fetch(LantanaAPIAddress.get_application_config_address, {
            method: LantanaAPIAddress.get_application_config_method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request),
        })
        const json = await res.json()
        const response: GetApplicationConfigResponse = json
        return response
        // ˄
    }

    constructor() {
        // ˅

        // ˄
    }

    // ˅

    // ˄
}

// ˅

// ˄
