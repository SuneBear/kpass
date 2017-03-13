import { schema } from 'normalizr'

import { secretSchema } from '../secret'

export const entrySchema = new schema.Entity('entries', {
  secrets: [ secretSchema ]
})

export const entriesSchema = new schema.Array(entrySchema)
