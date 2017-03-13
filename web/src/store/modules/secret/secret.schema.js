import { schema } from 'normalizr'

export const secretSchema = new schema.Entity('secrets')
export const secretsSchema = new schema.Array(secretSchema)
