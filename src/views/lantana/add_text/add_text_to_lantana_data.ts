import generate_uuid from "@/generate_uuid"
import AddToTextLantanaType from "./add_text_to_lantana_type"

export class AddTextToLantanaData {
    type: AddToTextLantanaType = AddToTextLantanaType.kmemo
    content: string = ""
    text_id: string = generate_uuid()
}