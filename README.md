# ppim

gateway
	分配客户端连接节点
	
客户端连接节点
	无额外状态
	用户->连接
	

	
逻辑处理节点
	虚拟节点 动态扩容
	发消息
	

group
	id
	type 讨论组 公开群 聊天室 直播间
	name
	content
	owner
	insertTime
	modifyTime
	
groupUser
	id
	groupId
	userId 
	dnd
	isOwner
	isManager
	
user
	id
	name
	
userBlacklist
	id
	userId1
	userId2
	
	
friend
	id
	userId1
	userId2
	
message
	id
	type
	fromType
	fromId
	toType
	toId
	toUserId
	actUserId
	
groupNotify
	type 加群 退群 成员资料变更
	
notice
	type 
	
	
	
	

用户设置
	黑名单
	全局免打扰
	
	群免打扰
	临时聊天
	基础资料
		头像
		昵称
		名字
		性别
	
用户关系
	好友
		单向好友
		双向好友
		临时会话
	群组