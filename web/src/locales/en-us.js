/* eslint-disable */

export const enUS = {
  /* == Global == */
  base: {
    more: 'More'
  },

  pageType: {
    secret: 'Secret',
    team: 'Team'
  },

  workspace: {
    personal: 'Personal'
  },

  action: {
    add: 'Add',
    create: 'Create',
    edit: 'Edit',
    save: 'Save'
  },

  member: {
    add: 'Add Member',
    name: 'Name',
    avatar: 'Avatar'
  },

  role: {
    owner: 'Owner',
    member: 'Member'
  },

  /* == Pages & Modals == */
  account: {
    username: 'Username',
    password: 'Password',
    passwordRetype: 'Retype password',
    unauthorized: 'Token experied, please sign in again.',
    signIn: 'Sign in',
    signInTip: 'Already have an account? Sign in',
    signInFailed: 'Incorrect username or password.',
    signUp: 'Sign up',
    signUpTip: 'Dont\' have an account? Sign up',
    signUpFailed: 'Failed to sign up.',
    signUpUserExisted: 'User already exists.',
    signOut: 'Sign Out',
    settings: 'Account Settings'
  },

  accountSettings: {
    publicProfile: 'Public profile',
    uploadAvatar: 'Upload new avatar',
    uploadAvatarSucceed: 'Avatar has been updated.',
    uploadAvatarFailed: 'Failed to upload new avatar.'
  },

  entries: {
    title: 'Entries',
    placeholderTitle: 'Start your secret life here.',
    placeholderDescription: 'Keep the KPassword safe, keep everything safe.'
  },

  entry: {
    menu: 'Entry Menu',
    new: 'New Entry',
    createSucceed: 'Entry created successfully',
    edit: 'Edit Entry',
    editSucceed: 'Entry edited successfully.',
    delete: 'Delete Entry',
    deleteSucceed: 'Entry deleted successfully.',
    entryName: 'Entry name',
    entryCategory: 'Entry category',
    category: 'Category',
    updatedDate: 'Updated date'
  },

  entryCategory: {
    'Login': 'Login',
    'Network': 'Network',
    'Software License': 'Software License',
    'Secure Note': 'Secure Note',
    'Server': 'Server',
    'Wallet': 'Wallet'
  },

  secrets: {
    placeholderTitle: 'Obliviate yourself when entry and secret meet.'
  },

  secret: {
    menu: 'Secret Menu',
    new: 'Add a secret',
    newTitle: 'Add a secret to "%{entryName}"',
    createSucceed: 'Secret added successfully',
    edit: 'Edit Secret',
    editTitle: 'Edit secret "%{secretName}"',
    editSucceed: 'Secret edited successfully.',
    delete: 'Delete Secret',
    deleteSucceed: 'Secret deleted successfully.',
    secretName: 'Name',
    secretNamePlaceholder: 'Name, username, email, etc identity',
    secretPassword: 'Key',
    secretPasswordCopy: 'Copy password',
    secretPasswordCopySucceed: 'Password copied successfully',
    secretPasswordPlaceholder: 'Key / Password',
    secretUrl: 'URL',
    secretUrlGoTo: 'Open link in new tab',
    secretUrlPlaceholder: 'URL / IP',
    secretNote: 'Note',
    secretNotePlaceholder: 'Note, support markdown syntax'
  },

  team: {
    new: 'New Team',
    teamName: 'Team name',
    isFrozen: 'No permission or team was frozen.',
    frozenLabel: 'Frozen',
    createSucceed: 'Team created successfully.',
    joinSucceed: 'Team joined successfully.',
    joinFailed: 'Invalid or expired invite link, please try another one.',
  },

  teamMembers: {
    title: 'Members',
    remove: 'Remove from team',
    removeSucceed: 'Member removed successfully.',
    leave: 'Leave the team',
    leaveSucceed: 'Team left successfully.',
    leaveAndDisband: 'Leave & Disband the team',
    invite: 'Invite to team',
    inviteGenerate: 'Generate invite link',
    inviteGenerateSucceed: 'Invite link generated successfully.',
    inviteLink: 'Invite link',
    inviteLinkCopy: 'Copy invite link',
    inviteLinkCopySucceed: 'Invite link copied successfully.',
    inviteLinkDescription: 'Invite link which will be expired after %{validMinutes} minutes is only valid for this user: ',
    inviteRepeated: 'The user is already a member of the team.',
    inviteFailed: 'Invalid username, please try another one.',
    joining: 'Joining a workspace...'
  },

  teamSettings: {
    title: 'Settings',
    rename: 'Rename team',
    renameDescription: 'Think of an ingenious name for the team that let members know what they can find.',
    freeze: 'Freeze team',
    freezeDescription: 'When the team has been frozen, all non-owners get in read-only mode that can\'t modify any entries.',
    updateSucceed: 'Team updated successfully.'
  },

  zPlaceholder: 'Hello World'
}
