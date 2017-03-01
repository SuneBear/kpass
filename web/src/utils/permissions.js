export const isPublicTeam = (team) => {
  return team && team.visibility !== 'private'
}

export const isFrozenTeam = (team) => {
  const conditions = [
    () => isPublicTeam(team),
    () => team && team.isFrozen
  ]

  return conditions.every((fn) => fn())
}

export const isMember = (member) => {
  return member && member.id
}

export const isOwner = (team, member) => {
  if (!team || !isMember(member)) {
    return false
  }

  return team.userID === member.id
}

export const isMe = (member, userMe) => {
  return member.id === userMe.id
}

// @TODO: Optimize getMemberPermissions, relate to roles
export const getMemberPermissions = (team, member) => {
  /* == Conditionals == */
  const $alwaysTrue = true

  const $isMember = isMember(member)
  const $isOwner = isOwner(team, member)

  const $isPublicTeam = isPublicTeam(team)
  const $isFrozenTeam = isFrozenTeam(team)

  const $isNonFrozenTeamOrOwner = !$isFrozenTeam || $isOwner

  /* == Construct Permissions == */
  const basePermissions = {
    roleWeight: -1000
  }

  const memberPermissions = Object.assign({}, basePermissions, {
    roleWeight: 0,

    // Entry
    createEntry: $isNonFrozenTeamOrOwner,

    // Team Member
    createTeamMember: $alwaysTrue,
    readTeamMember: $isPublicTeam
  })

  const ownerPermissions = Object.assign({}, memberPermissions, {
    roleWeight: 1000,

    // Team Member
    deleteTeamMember: $alwaysTrue,

    // Team Setting
    updateTeamSetting: $alwaysTrue
  })

  /* == Return == */
  if ($isOwner) {
    return ownerPermissions
  } else if ($isMember) {
    return memberPermissions
  } else {
    return basePermissions
  }
}
