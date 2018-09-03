# gorocket
[![Build Status](https://travis-ci.org/detached/gorocket.svg?branch=master)](https://travis-ci.org/detached/gorocket)
[![Coverage Status](https://coveralls.io/repos/github/detached/gorocket/badge.svg?branch=master)](https://coveralls.io/github/detached/gorocket?branch=master)

RocketChat client for golang. Compatible to the rest API of version 0.48.2.

The tests are failing because the library is not fully compatible to the newest version of RocketChat.
I will not update the lib because I am not using RocketChat any more.

RocketChat provides a rest and a realtime interface. This library provides clients for both.

```
go get github.com/detached/gorocket/rest
go get github.com/detached/gorocket/realtime
```

For more information checkout the [rest-godoc](https://godoc.org/github.com/detached/gorocket/rest) and [realtime-godoc](https://godoc.org/github.com/detached/gorocket/realtime), the test files or the examples.


### REST API functionality (TODO)
- Miscellaneous
    - [ ] info
    - [ ] directory
    - [ ] spotlight
    - [ ] statistics
    - [ ] statistics.list
- Assets
    - [ ] assets.setAsset
    - [ ] assets.unsetAsset
- Authentication
    - [X] login
    - [X] logout
    - [ ] me
- Users
    - [ ] users.create
    - [ ] users.createToken
    - [ ] users.delete
    - [ ] users.deleteOwnAccount
    - [ ] users.forgotPassword
    - [ ] users.generatePersonalAccessToken
    - [ ] users.getAvatar
    - [ ] users.getPersonalAccessTokens
    - [ ] users.getPreferences
    - [ ] users.getPresence
    - [ ] users.getUsernameSuggestion
    - [ ] users.info
    - [ ] users.list
    - [ ] users.regeneratePersonalAccessToken
    - [ ] users.register
    - [ ] users.removePersonalAccessToken
    - [ ] users.resetAvatar
    - [ ] users.setAvatar
    - [ ] users.setPreferences
    - [ ] users.update
    - [ ] users.updateOwnBasicInfo
- Channels
	- [ ] channels.addAll
	- [ ] channels.archive
	- [ ] channels.cleanHistory
	- [ ] channels.close
	- [ ] channels.counters
	- [ ] channels.create
	- [ ] channels.files
	- [ ] channels.getAllUserMentionsByChannel
	- [ ] channels.getIntegrations
	- [ ] channels.history
	- [ ] channels.info
	- [ ] channels.invite
	- [ ] channels.kick
	- [ ] channels.leave
	- [ ] channels.list
	- [ ] channels.list.joined
	- [ ] channels.members
	- [ ] channels.open
	- [ ] channels.rename
	- [ ] channels.roles
	- [ ] channels.setCustomFields
	- [ ] channels.setAnnouncement
	- [ ] channels.setDefault
	- [ ] channels.setDescription
	- [ ] channels.setJoinCode
	- [ ] channels.setPurpose
	- [ ] channels.setReadOnly
	- [ ] channels.setTopic
	- [ ] channels.setType
	- [ ] channels.unarchive
- Groups
	- [ ] groups.archive
	- [ ] groups.close
	- [ ] groups.counters
	- [ ] groups.create
	- [ ] groups.files
	- [ ] groups.history
	- [ ] groups.info
	- [ ] groups.invite
	- [ ] groups.kick
	- [ ] groups.leave
	- [ ] groups.list
	- [ ] groups.listAll
	- [ ] groups.open
	- [ ] groups.rename
	- [ ] groups.roles
	- [ ] groups.setCustomFields
	- [ ] groups.setDescription
	- [ ] groups.setPurpose
	- [ ] groups.setReadOnly
	- [ ] groups.setTopic
	- [ ] groups.setType
	- [ ] groups.unarchive
- Chat
	- [ ] chat.delete
	- [ ] chat.getMessage
	- [ ] chat.pinMessage
	- [ ] chat.postMessage
	- [ ] chat.react
	- [ ] chat.reportMessage
	- [ ] chat.search
	- [ ] chat.starMessage
	- [ ] chat.sendMessage
	- [ ] chat.unPinMessage
	- [ ] chat.unStarMessage
	- [ ] chat.update
	- [ ] chat.getMessageReadReceipts
- IM
	- [ ] im.close
	- [ ] im.counters
	- [ ] im.create
	- [ ] im.history
	- [ ] im.files
	- [ ] im.members
	- [ ] im.messages.others
	- [ ] im.list
	- [ ] im.list.everyone
	- [ ] im.open
	- [ ] im.setTopic
- Permissions
    - [ ] permissions.list
    - [ ] permissions.update
- Rooms
	- [ ] rooms.cleanHistory
	- [ ] rooms.favorite
	- [ ] rooms.get
	- [ ] rooms.saveNotification
	- [ ] rooms.upload/:rid
- Command Methods
    - [ ] commands.get
    - [ ] commands.list
    - [ ] commands.run
- Emoji Custom
    - [ ] emoji-custom
- Messages
    - [ ] messages/types
- Settings
    - [ ] settings
	- [ ] settings.public
	- [ ] settings.oauth
	- [ ] service.configurations
    - [ ] settings/:_id
- Subscriptions
	- [ ] subscriptions.get
	- [ ] subscriptions.getOne
	- [ ] subscriptions.read
	- [ ] subscriptions.unread