import { schema } from 'normalizr'

export const memberSchema = new schema.Entity('members')
export const membersSchema = new schema.Array(memberSchema)
