import { createSelector } from 'reselect'

import { currentTeamSelector } from '../current-team'

export const currentTeamMembersSelector = createSelector(
  (state) => currentTeamSelector(state),
  (currentTeam) => {
    if (!currentTeam) {
      return {}
    }

    const mockData = [
      {
        id: 'SuneBear',
        avatarUrl: 'https://avatars.io/twitter/hisunebear'
      },
      {
        id: 'Kumamon',
        avatarUrl: 'https://avatars.io/twitter/55_kumamon_eng'
      },
      {
        id: 'Kumakichi',
        avatarUrl: 'https://avatars.io/twitter/Kumakichi81_bot'
      }
    ]

    return mockData
  }
)
