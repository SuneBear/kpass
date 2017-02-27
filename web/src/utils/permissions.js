// @TODO: getPermissions

export const isPublicTeam = (team) => {
  return team && team.visibility !== 'private'
}

export const isOwner = (team, member) => {
  if (!team || !member) {
    return false
  }

  return team.userID === member.id
}

export const isMe = (member, userMe) => {
  return member.id === userMe.id
}
