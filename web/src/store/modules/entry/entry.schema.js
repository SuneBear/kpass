import { schema } from 'normalizr'

export const entrySchema = new schema.Entity('entries')
export const entriesSchema = new schema.Array(entrySchema)
