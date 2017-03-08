import { schema } from 'normalizr'

import { memberSchema } from '../member'

export const teamSchema = new schema.Entity('teams', {
  members: [ memberSchema ]
})

export const teamsSchema = new schema.Array(teamSchema)
