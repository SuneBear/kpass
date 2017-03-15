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

export const isCreator = (memberId, userMe) => {
  return memberId === userMe.id
}

// @TODO: Optimize getUserPermissions, relate to roles
export const getUserPermissions = (user, team) => {
  /* == Conditionals == */
  const $alwaysTrue = true

  const $isMember = isMember(user)
  const $isOwner = isOwner(team, user)

  const $isPublicTeam = isPublicTeam(team)
  const $isFrozenTeam = isFrozenTeam(team)

  const $isNonFrozenTeamOrOwner = !$isFrozenTeam || $isOwner

  /* == Construct Permissions == */
  const basePermissions = {
    roleWeight: -1000
  }

  const memberPermissions = {
    ...basePermissions,

    roleWeight: 0,

    // Entry
    createEntry: $isNonFrozenTeamOrOwner,

    // Team Member
    createTeamMember: $alwaysTrue,
    readTeamMember: $isPublicTeam
  }

  const ownerPermissions = {
    ...memberPermissions,

    roleWeight: 1000,

    // Team Member
    deleteTeamMember: $alwaysTrue,

    // Team Setting
    updateTeamSetting: $alwaysTrue
  }

  /* == Return == */
  if ($isOwner) {
    return ownerPermissions
  } else if ($isMember) {
    return memberPermissions
  } else {
    return basePermissions
  }
}

export const getCreatorPermissions = (memberId, userMe, team) => {
  const $alwaysTrue = true

  const $isCreator = $alwaysTrue // @REPLACE: isCreator(memberId, userMe)
  const $isOwner = isOwner(team, userMe)
  const $isFrozenTeam = isFrozenTeam(team)
  const $isNonFrozenTeamAndCreator = !$isFrozenTeam && $isCreator

  const getFinalPermission = (key) => {
    const conditions = [
      $isNonFrozenTeamAndCreator,
      $isOwner
    ]
    return conditions.some((cond) => cond)
  }

  const permissionKeys = [
    // Entry
    'updateEntry',
    'deleteEntry',

    // Secret
    'createSecret',
    'updateSecret',
    'deleteSecret'
  ]

  // Get Permissions
  const permissions = {}
  permissionKeys.map((key) => {
    const permission = getFinalPermission(key)
    permissions[key] = permission
  })

  return permissions
}
