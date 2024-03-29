# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [common.proto](#common.proto)
    - [CommonResponse](#common.CommonResponse)
    - [ImageInfo](#common.ImageInfo)
    - [Pager](#common.Pager)
    - [SendEmailRequest](#common.SendEmailRequest)
    - [SendEmailRequest.CcEntry](#common.SendEmailRequest.CcEntry)
    - [SendEmailRequest.ReceiverEntry](#common.SendEmailRequest.ReceiverEntry)
    - [SendEmailResponse](#common.SendEmailResponse)
    - [UploadImageRequest](#common.UploadImageRequest)
    - [UploadImageResponse](#common.UploadImageResponse)
  
    - [Gender](#common.Gender)
    - [IsEnable](#common.IsEnable)
    - [ReviewStatus](#common.ReviewStatus)
  
    - [CommonService](#common.CommonService)
  
- [feed.proto](#feed.proto)
    - [Feed](#user.Feed)
    - [FeedCreateRequest](#user.FeedCreateRequest)
    - [FeedCreateResponse](#user.FeedCreateResponse)
    - [FeedDeleteRequest](#user.FeedDeleteRequest)
    - [FeedDeleteResponse](#user.FeedDeleteResponse)
    - [FeedEditRequest](#user.FeedEditRequest)
    - [FeedEditResponse](#user.FeedEditResponse)
    - [FeedFollowRequest](#user.FeedFollowRequest)
    - [FeedFollowResponse](#user.FeedFollowResponse)
    - [FeedGetRequest](#user.FeedGetRequest)
    - [FeedGetResponse](#user.FeedGetResponse)
    - [FeedImage](#user.FeedImage)
    - [FeedLike](#user.FeedLike)
    - [FeedLikeRequest](#user.FeedLikeRequest)
    - [FeedLikeResponse](#user.FeedLikeResponse)
    - [FeedOwnRequest](#user.FeedOwnRequest)
    - [FeedOwnResponse](#user.FeedOwnResponse)
    - [FeedQueryRequest](#user.FeedQueryRequest)
    - [FeedQueryResponse](#user.FeedQueryResponse)
    - [FeedRecommendRequest](#user.FeedRecommendRequest)
    - [FeedRecommendResponse](#user.FeedRecommendResponse)
    - [FeedUnLikeRequest](#user.FeedUnLikeRequest)
    - [FeedUnLikeResponse](#user.FeedUnLikeResponse)
    - [FeedVideo](#user.FeedVideo)
  
    - [FeedType](#user.FeedType)
  
    - [FeedService](#user.FeedService)
  
- [gateway.proto](#gateway.proto)
    - [PushToAllRequest](#message.PushToAllRequest)
    - [PushToAllResponse](#message.PushToAllResponse)
    - [RegisterRequest](#message.RegisterRequest)
    - [RegisterResponse](#message.RegisterResponse)
    - [SendMessageRequest](#message.SendMessageRequest)
    - [SendMessageResponse](#message.SendMessageResponse)
    - [UnRegisterRequest](#message.UnRegisterRequest)
    - [UnRegisterResponse](#message.UnRegisterResponse)
  
    - [GatewayService](#message.GatewayService)
  
- [group.proto](#group.proto)
    - [CreateGroupRequest](#group.CreateGroupRequest)
    - [CreateGroupResponse](#group.CreateGroupResponse)
    - [DeleteGroupRequest](#group.DeleteGroupRequest)
    - [DeleteGroupResponse](#group.DeleteGroupResponse)
    - [GetGroupRequest](#group.GetGroupRequest)
    - [GetGroupResponse](#group.GetGroupResponse)
    - [GetGroupUserRequest](#group.GetGroupUserRequest)
    - [GetGroupUserResponse](#group.GetGroupUserResponse)
    - [Group](#group.Group)
    - [JoinGroupRequest](#group.JoinGroupRequest)
    - [JoinGroupResponse](#group.JoinGroupResponse)
    - [QueryGroupRequest](#group.QueryGroupRequest)
    - [QueryGroupResponse](#group.QueryGroupResponse)
    - [QuitGroupRequest](#group.QuitGroupRequest)
    - [QuitGroupResponse](#group.QuitGroupResponse)
    - [UpdateGroupRequest](#group.UpdateGroupRequest)
    - [UpdateGroupResponse](#group.UpdateGroupResponse)
  
    - [GroupService](#group.GroupService)
  
- [message.proto](#message.proto)
    - [FaceMessage](#message.FaceMessage)
    - [FileMessage](#message.FileMessage)
    - [ImageMessage](#message.ImageMessage)
    - [ImageMessageItem](#message.ImageMessageItem)
    - [LocationMessage](#message.LocationMessage)
    - [ParseMessageRequest](#message.ParseMessageRequest)
    - [ParseMessageResponse](#message.ParseMessageResponse)
    - [SendGroupFaceMessageRequest](#message.SendGroupFaceMessageRequest)
    - [SendGroupFaceMessageResponse](#message.SendGroupFaceMessageResponse)
    - [SendGroupFileMessageRequest](#message.SendGroupFileMessageRequest)
    - [SendGroupFileMessageResponse](#message.SendGroupFileMessageResponse)
    - [SendGroupImageMessageRequest](#message.SendGroupImageMessageRequest)
    - [SendGroupImageMessageResponse](#message.SendGroupImageMessageResponse)
    - [SendGroupLocationMessageRequest](#message.SendGroupLocationMessageRequest)
    - [SendGroupLocationMessageResponse](#message.SendGroupLocationMessageResponse)
    - [SendGroupSoundMessageRequest](#message.SendGroupSoundMessageRequest)
    - [SendGroupSoundMessageResponse](#message.SendGroupSoundMessageResponse)
    - [SendGroupTextMessageRequest](#message.SendGroupTextMessageRequest)
    - [SendGroupTextMessageResponse](#message.SendGroupTextMessageResponse)
    - [SendGroupVideoMessageRequest](#message.SendGroupVideoMessageRequest)
    - [SendGroupVideoMessageResponse](#message.SendGroupVideoMessageResponse)
    - [SendUserFaceMessageRequest](#message.SendUserFaceMessageRequest)
    - [SendUserFaceMessageResponse](#message.SendUserFaceMessageResponse)
    - [SendUserFileMessageRequest](#message.SendUserFileMessageRequest)
    - [SendUserFileMessageResponse](#message.SendUserFileMessageResponse)
    - [SendUserImageMessageRequest](#message.SendUserImageMessageRequest)
    - [SendUserImageMessageResponse](#message.SendUserImageMessageResponse)
    - [SendUserLocationMessageRequest](#message.SendUserLocationMessageRequest)
    - [SendUserLocationMessageResponse](#message.SendUserLocationMessageResponse)
    - [SendUserSoundMessageRequest](#message.SendUserSoundMessageRequest)
    - [SendUserSoundMessageResponse](#message.SendUserSoundMessageResponse)
    - [SendUserTextMessageRequest](#message.SendUserTextMessageRequest)
    - [SendUserTextMessageResponse](#message.SendUserTextMessageResponse)
    - [SendUserVideoMessageRequest](#message.SendUserVideoMessageRequest)
    - [SendUserVideoMessageResponse](#message.SendUserVideoMessageResponse)
    - [SoundMessage](#message.SoundMessage)
    - [TextMessage](#message.TextMessage)
    - [VideoMessage](#message.VideoMessage)
  
    - [IMAGE_FORMAT](#message.IMAGE_FORMAT)
    - [IMAGE_TYPE](#message.IMAGE_TYPE)
    - [MESSAGE_TYPE](#message.MESSAGE_TYPE)
    - [VIDEO_FORMAT](#message.VIDEO_FORMAT)
  
    - [MessageService](#message.MessageService)
  
- [ping.proto](#ping.proto)
    - [PingRequest](#ping.PingRequest)
    - [PingResponse](#ping.PingResponse)
  
    - [PingService](#ping.PingService)
  
- [user.proto](#user.proto)
    - [UpdateAvatarRequest](#user.UpdateAvatarRequest)
    - [UpdateAvatarResponse](#user.UpdateAvatarResponse)
    - [UpdatePasswordRequest](#user.UpdatePasswordRequest)
    - [UpdatePasswordResponse](#user.UpdatePasswordResponse)
    - [User](#user.User)
    - [UserCreateRequest](#user.UserCreateRequest)
    - [UserCreateResponse](#user.UserCreateResponse)
    - [UserGetRequest](#user.UserGetRequest)
    - [UserGetResponse](#user.UserGetResponse)
    - [UserLoginRequest](#user.UserLoginRequest)
    - [UserLoginResponse](#user.UserLoginResponse)
    - [UserQueryRequest](#user.UserQueryRequest)
    - [UserQueryResponse](#user.UserQueryResponse)
  
    - [UserService](#user.UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="common.CommonResponse"></a>

### CommonResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [uint64](#uint64) |  | 错误码 |
| business_code | [uint64](#uint64) |  | 业务码 |
| message | [string](#string) |  | 信息 |






<a name="common.ImageInfo"></a>

### ImageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| type | [string](#string) |  |  |
| url | [string](#string) |  |  |
| width | [uint64](#uint64) |  |  |
| height | [uint64](#uint64) |  |  |






<a name="common.Pager"></a>

### Pager



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [uint64](#uint64) |  |  |
| page_zie | [uint64](#uint64) |  |  |
| total_rows | [uint64](#uint64) |  |  |
| total_pages | [uint64](#uint64) |  |  |






<a name="common.SendEmailRequest"></a>

### SendEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| receiver | [SendEmailRequest.ReceiverEntry](#common.SendEmailRequest.ReceiverEntry) | repeated |  |
| cc | [SendEmailRequest.CcEntry](#common.SendEmailRequest.CcEntry) | repeated |  |
| subject | [string](#string) |  |  |
| content | [string](#string) |  |  |






<a name="common.SendEmailRequest.CcEntry"></a>

### SendEmailRequest.CcEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="common.SendEmailRequest.ReceiverEntry"></a>

### SendEmailRequest.ReceiverEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="common.SendEmailResponse"></a>

### SendEmailResponse







<a name="common.UploadImageRequest"></a>

### UploadImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [bytes](#bytes) |  |  |






<a name="common.UploadImageResponse"></a>

### UploadImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |





 


<a name="common.Gender"></a>

### Gender


| Name | Number | Description |
| ---- | ------ | ----------- |
| Unknown | 0 | 未知 |
| Male | 1 | 男 |
| Female | 2 | 女 |



<a name="common.IsEnable"></a>

### IsEnable


| Name | Number | Description |
| ---- | ------ | ----------- |
| NotEnable | 0 | 未启用 |
| Enable | 1 | 启用 |



<a name="common.ReviewStatus"></a>

### ReviewStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ReviewInit | 0 | 未提交审核 |
| ReviewTo | 1 | 已提交，待审核 |
| ReviewPass | 2 | 审核通过 |
| ReviewFail | 3 | 审核未通过 |


 

 


<a name="common.CommonService"></a>

### CommonService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UploadImage | [UploadImageRequest](#common.UploadImageRequest) stream | [UploadImageResponse](#common.UploadImageResponse) |  |
| SendEmail | [SendEmailRequest](#common.SendEmailRequest) | [SendEmailResponse](#common.SendEmailResponse) |  |

 



<a name="feed.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## feed.proto



<a name="user.Feed"></a>

### Feed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| user_id | [uint64](#uint64) |  |  |
| content | [string](#string) |  |  |
| type | [FeedType](#user.FeedType) |  |  |
| like_count | [uint64](#uint64) |  |  |
| view_count | [uint64](#uint64) |  |  |
| comment_count | [uint64](#uint64) |  |  |
| operator | [uint64](#uint64) |  |  |
| remark | [string](#string) |  |  |
| hot | [uint64](#uint64) |  |  |
| is_enable | [common.IsEnable](#common.IsEnable) |  |  |
| review_status | [common.ReviewStatus](#common.ReviewStatus) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="user.FeedCreateRequest"></a>

### FeedCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [FeedType](#user.FeedType) |  |  |
| content | [string](#string) |  |  |
| resource | [string](#string) | repeated |  |






<a name="user.FeedCreateResponse"></a>

### FeedCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| feed | [Feed](#user.Feed) |  |  |






<a name="user.FeedDeleteRequest"></a>

### FeedDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="user.FeedDeleteResponse"></a>

### FeedDeleteResponse







<a name="user.FeedEditRequest"></a>

### FeedEditRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| type | [FeedType](#user.FeedType) |  |  |
| content | [string](#string) |  |  |
| resource | [string](#string) | repeated |  |






<a name="user.FeedEditResponse"></a>

### FeedEditResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="user.FeedFollowRequest"></a>

### FeedFollowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [uint64](#uint64) |  |  |






<a name="user.FeedFollowResponse"></a>

### FeedFollowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| feed | [Feed](#user.Feed) | repeated |  |






<a name="user.FeedGetRequest"></a>

### FeedGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="user.FeedGetResponse"></a>

### FeedGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| feed | [Feed](#user.Feed) |  |  |






<a name="user.FeedImage"></a>

### FeedImage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| feed_id | [uint64](#uint64) |  |  |
| img_url | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="user.FeedLike"></a>

### FeedLike



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| feed_id | [uint64](#uint64) |  |  |
| user_id | [uint64](#uint64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="user.FeedLikeRequest"></a>

### FeedLikeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="user.FeedLikeResponse"></a>

### FeedLikeResponse







<a name="user.FeedOwnRequest"></a>

### FeedOwnRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keyword | [string](#string) |  |  |
| review_status | [common.ReviewStatus](#common.ReviewStatus) |  |  |
| last_id | [uint64](#uint64) |  |  |
| num | [uint64](#uint64) |  |  |






<a name="user.FeedOwnResponse"></a>

### FeedOwnResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| feed | [Feed](#user.Feed) | repeated |  |






<a name="user.FeedQueryRequest"></a>

### FeedQueryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keyword | [string](#string) |  |  |
| last_id | [uint64](#uint64) |  |  |
| num | [uint64](#uint64) |  |  |






<a name="user.FeedQueryResponse"></a>

### FeedQueryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| feed | [Feed](#user.Feed) | repeated |  |






<a name="user.FeedRecommendRequest"></a>

### FeedRecommendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [uint64](#uint64) |  |  |






<a name="user.FeedRecommendResponse"></a>

### FeedRecommendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| feed | [Feed](#user.Feed) | repeated |  |






<a name="user.FeedUnLikeRequest"></a>

### FeedUnLikeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="user.FeedUnLikeResponse"></a>

### FeedUnLikeResponse







<a name="user.FeedVideo"></a>

### FeedVideo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| feed_id | [uint64](#uint64) |  |  |
| video_url | [string](#string) |  |  |
| cover_url | [string](#string) |  |  |
| width | [uint64](#uint64) |  |  |
| height | [uint64](#uint64) |  |  |
| duration | [double](#double) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 


<a name="user.FeedType"></a>

### FeedType


| Name | Number | Description |
| ---- | ------ | ----------- |
| Txt | 0 | 文本 |
| Image | 1 | 图片 |
| Video | 2 | 视频 |


 

 


<a name="user.FeedService"></a>

### FeedService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| FeedCreate | [FeedCreateRequest](#user.FeedCreateRequest) | [FeedCreateResponse](#user.FeedCreateResponse) |  |
| FeedEdit | [FeedEditRequest](#user.FeedEditRequest) | [FeedEditResponse](#user.FeedEditResponse) |  |
| FeedLike | [FeedLikeRequest](#user.FeedLikeRequest) | [FeedLikeResponse](#user.FeedLikeResponse) |  |
| FeedUnLike | [FeedUnLikeRequest](#user.FeedUnLikeRequest) | [FeedUnLikeResponse](#user.FeedUnLikeResponse) |  |
| FeedGet | [FeedGetRequest](#user.FeedGetRequest) | [FeedGetResponse](#user.FeedGetResponse) |  |
| FeedDelete | [FeedDeleteRequest](#user.FeedDeleteRequest) | [FeedDeleteResponse](#user.FeedDeleteResponse) |  |
| FeedOwn | [FeedOwnRequest](#user.FeedOwnRequest) | [FeedOwnResponse](#user.FeedOwnResponse) |  |
| FeedQuery | [FeedQueryRequest](#user.FeedQueryRequest) | [FeedQueryResponse](#user.FeedQueryResponse) |  |
| FeedFollow | [FeedFollowRequest](#user.FeedFollowRequest) | [FeedFollowResponse](#user.FeedFollowResponse) |  |

 



<a name="gateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gateway.proto



<a name="message.PushToAllRequest"></a>

### PushToAllRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gateway_id | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="message.PushToAllResponse"></a>

### PushToAllResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [uint64](#uint64) |  |  |






<a name="message.RegisterRequest"></a>

### RegisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gateway_id | [string](#string) |  |  |
| client_id | [string](#string) |  |  |






<a name="message.RegisterResponse"></a>

### RegisterResponse







<a name="message.SendMessageRequest"></a>

### SendMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gateway_id | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="message.SendMessageResponse"></a>

### SendMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |






<a name="message.UnRegisterRequest"></a>

### UnRegisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gateway_id | [string](#string) |  |  |
| client_id | [string](#string) |  |  |






<a name="message.UnRegisterResponse"></a>

### UnRegisterResponse






 

 

 


<a name="message.GatewayService"></a>

### GatewayService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterRequest](#message.RegisterRequest) | [RegisterResponse](#message.RegisterResponse) |  |
| UnRegister | [UnRegisterRequest](#message.UnRegisterRequest) | [UnRegisterResponse](#message.UnRegisterResponse) |  |
| SendMessage | [SendMessageRequest](#message.SendMessageRequest) stream | [SendMessageResponse](#message.SendMessageResponse) stream |  |
| PushToAllMessage | [PushToAllRequest](#message.PushToAllRequest) stream | [PushToAllResponse](#message.PushToAllResponse) stream |  |

 



<a name="group.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## group.proto



<a name="group.CreateGroupRequest"></a>

### CreateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| introduction | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |






<a name="group.CreateGroupResponse"></a>

### CreateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [Group](#group.Group) |  |  |






<a name="group.DeleteGroupRequest"></a>

### DeleteGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="group.DeleteGroupResponse"></a>

### DeleteGroupResponse







<a name="group.GetGroupRequest"></a>

### GetGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="group.GetGroupResponse"></a>

### GetGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [Group](#group.Group) |  |  |






<a name="group.GetGroupUserRequest"></a>

### GetGroupUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| page | [uint64](#uint64) |  |  |
| page_size | [uint64](#uint64) |  |  |






<a name="group.GetGroupUserResponse"></a>

### GetGroupUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| user | [user.User](#user.User) | repeated |  |






<a name="group.Group"></a>

### Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| name | [string](#string) |  |  |
| introduction | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |
| extra | [string](#string) |  |  |
| user_id | [uint64](#uint64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="group.JoinGroupRequest"></a>

### JoinGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| user_id | [uint64](#uint64) |  |  |






<a name="group.JoinGroupResponse"></a>

### JoinGroupResponse







<a name="group.QueryGroupRequest"></a>

### QueryGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keyword | [string](#string) |  |  |
| page | [uint64](#uint64) |  |  |
| page_size | [uint64](#uint64) |  |  |






<a name="group.QueryGroupResponse"></a>

### QueryGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| group | [Group](#group.Group) | repeated |  |






<a name="group.QuitGroupRequest"></a>

### QuitGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| user_id | [uint64](#uint64) |  |  |






<a name="group.QuitGroupResponse"></a>

### QuitGroupResponse







<a name="group.UpdateGroupRequest"></a>

### UpdateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| name | [string](#string) |  |  |
| introduction | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |






<a name="group.UpdateGroupResponse"></a>

### UpdateGroupResponse






 

 

 


<a name="group.GroupService"></a>

### GroupService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateGroup | [CreateGroupRequest](#group.CreateGroupRequest) | [CreateGroupResponse](#group.CreateGroupResponse) |  |
| UpdateGroup | [UpdateGroupRequest](#group.UpdateGroupRequest) | [UpdateGroupResponse](#group.UpdateGroupResponse) |  |
| GetGroup | [GetGroupRequest](#group.GetGroupRequest) | [GetGroupResponse](#group.GetGroupResponse) |  |
| QueryGroup | [QueryGroupRequest](#group.QueryGroupRequest) | [QueryGroupResponse](#group.QueryGroupResponse) |  |
| DeleteGroup | [DeleteGroupRequest](#group.DeleteGroupRequest) | [DeleteGroupResponse](#group.DeleteGroupResponse) |  |
| JoinGroup | [JoinGroupRequest](#group.JoinGroupRequest) | [JoinGroupResponse](#group.JoinGroupResponse) |  |
| QuitGroup | [QuitGroupRequest](#group.QuitGroupRequest) | [QuitGroupResponse](#group.QuitGroupResponse) |  |

 



<a name="message.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## message.proto



<a name="message.FaceMessage"></a>

### FaceMessage
表情消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| symbol | [string](#string) |  |  |






<a name="message.FileMessage"></a>

### FileMessage
文件消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| size | [uint64](#uint64) |  | 大小 |
| name | [string](#string) |  | 文件名 |
| url | [string](#string) |  | 链接地址 |






<a name="message.ImageMessage"></a>

### ImageMessage
图像消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| list | [ImageMessageItem](#message.ImageMessageItem) | repeated |  |






<a name="message.ImageMessageItem"></a>

### ImageMessageItem
图片


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [IMAGE_TYPE](#message.IMAGE_TYPE) |  | 类型 |
| format | [IMAGE_FORMAT](#message.IMAGE_FORMAT) |  | 格式 |
| size | [uint64](#uint64) |  | 大小 |
| width | [uint64](#uint64) |  | 宽度 |
| height | [uint64](#uint64) |  | 高度 |
| url | [string](#string) |  | 链接地址 |






<a name="message.LocationMessage"></a>

### LocationMessage
地理位置消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| cover_image | [string](#string) |  |  |
| lat | [double](#double) |  |  |
| lng | [double](#double) |  |  |
| map_link | [string](#string) |  |  |
| desc | [string](#string) |  |  |






<a name="message.ParseMessageRequest"></a>

### ParseMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [bytes](#bytes) |  |  |






<a name="message.ParseMessageResponse"></a>

### ParseMessageResponse







<a name="message.SendGroupFaceMessageRequest"></a>

### SendGroupFaceMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| face | [FaceMessage](#message.FaceMessage) |  |  |






<a name="message.SendGroupFaceMessageResponse"></a>

### SendGroupFaceMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendGroupFileMessageRequest"></a>

### SendGroupFileMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| file | [FileMessage](#message.FileMessage) |  |  |






<a name="message.SendGroupFileMessageResponse"></a>

### SendGroupFileMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendGroupImageMessageRequest"></a>

### SendGroupImageMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| image | [ImageMessageItem](#message.ImageMessageItem) |  |  |






<a name="message.SendGroupImageMessageResponse"></a>

### SendGroupImageMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendGroupLocationMessageRequest"></a>

### SendGroupLocationMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| location | [LocationMessage](#message.LocationMessage) |  |  |






<a name="message.SendGroupLocationMessageResponse"></a>

### SendGroupLocationMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendGroupSoundMessageRequest"></a>

### SendGroupSoundMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| sound | [SoundMessage](#message.SoundMessage) |  |  |






<a name="message.SendGroupSoundMessageResponse"></a>

### SendGroupSoundMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendGroupTextMessageRequest"></a>

### SendGroupTextMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| text | [TextMessage](#message.TextMessage) |  |  |






<a name="message.SendGroupTextMessageResponse"></a>

### SendGroupTextMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendGroupVideoMessageRequest"></a>

### SendGroupVideoMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_group_id | [uint64](#uint64) |  |  |
| at_user_id | [uint64](#uint64) | repeated |  |
| video | [VideoMessage](#message.VideoMessage) |  |  |






<a name="message.SendGroupVideoMessageResponse"></a>

### SendGroupVideoMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserFaceMessageRequest"></a>

### SendUserFaceMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| face | [FaceMessage](#message.FaceMessage) |  |  |






<a name="message.SendUserFaceMessageResponse"></a>

### SendUserFaceMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserFileMessageRequest"></a>

### SendUserFileMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| file | [FileMessage](#message.FileMessage) |  |  |






<a name="message.SendUserFileMessageResponse"></a>

### SendUserFileMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserImageMessageRequest"></a>

### SendUserImageMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| image | [ImageMessageItem](#message.ImageMessageItem) |  |  |






<a name="message.SendUserImageMessageResponse"></a>

### SendUserImageMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserLocationMessageRequest"></a>

### SendUserLocationMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| location | [LocationMessage](#message.LocationMessage) |  |  |






<a name="message.SendUserLocationMessageResponse"></a>

### SendUserLocationMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserSoundMessageRequest"></a>

### SendUserSoundMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| sound | [SoundMessage](#message.SoundMessage) |  |  |






<a name="message.SendUserSoundMessageResponse"></a>

### SendUserSoundMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserTextMessageRequest"></a>

### SendUserTextMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| text | [TextMessage](#message.TextMessage) |  |  |






<a name="message.SendUserTextMessageResponse"></a>

### SendUserTextMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SendUserVideoMessageRequest"></a>

### SendUserVideoMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| to_user_id | [uint64](#uint64) |  |  |
| video | [VideoMessage](#message.VideoMessage) |  |  |






<a name="message.SendUserVideoMessageResponse"></a>

### SendUserVideoMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="message.SoundMessage"></a>

### SoundMessage
语音消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| url | [string](#string) |  |  |
| size | [uint64](#uint64) |  |  |
| seconds | [uint64](#uint64) |  |  |






<a name="message.TextMessage"></a>

### TextMessage
文本消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| content | [string](#string) |  |  |






<a name="message.VideoMessage"></a>

### VideoMessage
视频消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| size | [uint64](#uint64) |  |  |
| seconds | [uint64](#uint64) |  |  |
| url | [string](#string) |  |  |
| format | [string](#string) |  |  |
| thumb_url | [string](#string) |  |  |
| thumb_size | [uint64](#uint64) |  |  |
| thumb_width | [double](#double) |  |  |
| thumb_height | [double](#double) |  |  |
| thumb_format | [IMAGE_FORMAT](#message.IMAGE_FORMAT) |  |  |





 


<a name="message.IMAGE_FORMAT"></a>

### IMAGE_FORMAT
图片格式

| Name | Number | Description |
| ---- | ------ | ----------- |
| JPG | 0 |  |
| GIF | 1 |  |
| PNG | 2 |  |
| BMP | 3 |  |
| OTHER | 255 |  |



<a name="message.IMAGE_TYPE"></a>

### IMAGE_TYPE
图片类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| ORIGIN | 0 | 原图 |
| BIG | 1 | 大图 |
| THUMBNAIL | 2 | 缩略图 |



<a name="message.MESSAGE_TYPE"></a>

### MESSAGE_TYPE
消息类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| TEXT | 0 | 文本消息。 |
| LOCATION | 1 | 地理位置消息。 |
| FACE | 2 | 表情消息。 |
| SOUND | 3 | 语音消息。 |
| IMAGE | 4 | 图像消息。 |
| FILE | 5 | 文件消息。 |
| Video | 6 | 视频消息。 |



<a name="message.VIDEO_FORMAT"></a>

### VIDEO_FORMAT
视频格式

| Name | Number | Description |
| ---- | ------ | ----------- |
| MP4 | 0 |  |
| MOV | 1 |  |
| WMV | 2 |  |
| FLV | 3 |  |
| AVI | 4 |  |
| MKV | 5 |  |
| OTHER_VIDEO_FORMAT | 255 |  |


 

 


<a name="message.MessageService"></a>

### MessageService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendUserTextMessage | [SendUserTextMessageRequest](#message.SendUserTextMessageRequest) | [SendUserTextMessageResponse](#message.SendUserTextMessageResponse) |  |
| SendUserLocationMessage | [SendUserLocationMessageRequest](#message.SendUserLocationMessageRequest) | [SendUserLocationMessageResponse](#message.SendUserLocationMessageResponse) |  |
| SendUserFaceMessage | [SendUserFaceMessageRequest](#message.SendUserFaceMessageRequest) | [SendUserFaceMessageResponse](#message.SendUserFaceMessageResponse) |  |
| SendUserSoundMessage | [SendUserSoundMessageRequest](#message.SendUserSoundMessageRequest) | [SendUserSoundMessageResponse](#message.SendUserSoundMessageResponse) |  |
| SendUserVideoMessage | [SendUserVideoMessageRequest](#message.SendUserVideoMessageRequest) | [SendUserVideoMessageResponse](#message.SendUserVideoMessageResponse) |  |
| SendUserImageMessage | [SendUserImageMessageRequest](#message.SendUserImageMessageRequest) | [SendUserImageMessageResponse](#message.SendUserImageMessageResponse) |  |
| SendUserFileMessage | [SendUserFileMessageRequest](#message.SendUserFileMessageRequest) | [SendUserFileMessageResponse](#message.SendUserFileMessageResponse) |  |
| SendGroupTextMessage | [SendGroupTextMessageRequest](#message.SendGroupTextMessageRequest) | [SendGroupTextMessageResponse](#message.SendGroupTextMessageResponse) |  |
| SendGroupLocationMessage | [SendGroupLocationMessageRequest](#message.SendGroupLocationMessageRequest) | [SendGroupLocationMessageResponse](#message.SendGroupLocationMessageResponse) |  |
| SendGroupFceMessage | [SendGroupFaceMessageRequest](#message.SendGroupFaceMessageRequest) | [SendGroupFaceMessageResponse](#message.SendGroupFaceMessageResponse) |  |
| SendGroupSoundMessage | [SendGroupSoundMessageRequest](#message.SendGroupSoundMessageRequest) | [SendGroupSoundMessageResponse](#message.SendGroupSoundMessageResponse) |  |
| SendGroupVideoMessage | [SendGroupVideoMessageRequest](#message.SendGroupVideoMessageRequest) | [SendGroupVideoMessageResponse](#message.SendGroupVideoMessageResponse) |  |
| SendGroupImageMessage | [SendGroupImageMessageRequest](#message.SendGroupImageMessageRequest) | [SendGroupImageMessageResponse](#message.SendGroupImageMessageResponse) |  |
| SendGroupFileMessage | [SendGroupFileMessageRequest](#message.SendGroupFileMessageRequest) | [SendGroupFileMessageResponse](#message.SendGroupFileMessageResponse) |  |

 



<a name="ping.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ping.proto



<a name="ping.PingRequest"></a>

### PingRequest







<a name="ping.PingResponse"></a>

### PingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |





 

 

 


<a name="ping.PingService"></a>

### PingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Ping | [PingRequest](#ping.PingRequest) | [PingResponse](#ping.PingResponse) |  |

 



<a name="user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="user.UpdateAvatarRequest"></a>

### UpdateAvatarRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="user.UpdateAvatarResponse"></a>

### UpdateAvatarResponse







<a name="user.UpdatePasswordRequest"></a>

### UpdatePasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| password | [string](#string) |  |  |
| password_rpt | [string](#string) |  |  |






<a name="user.UpdatePasswordResponse"></a>

### UpdatePasswordResponse







<a name="user.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| nickname | [string](#string) |  |  |
| gender | [common.Gender](#common.Gender) |  |  |
| avatar_url | [string](#string) |  |  |
| extra | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="user.UserCreateRequest"></a>

### UserCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  |  |
| gender | [common.Gender](#common.Gender) |  |  |
| avatar_url | [string](#string) |  |  |
| password | [string](#string) |  |  |
| password_rpt | [string](#string) |  |  |






<a name="user.UserCreateResponse"></a>

### UserCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user.User) |  |  |






<a name="user.UserGetRequest"></a>

### UserGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |






<a name="user.UserGetResponse"></a>

### UserGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user.User) |  |  |






<a name="user.UserLoginRequest"></a>

### UserLoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="user.UserLoginResponse"></a>

### UserLoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="user.UserQueryRequest"></a>

### UserQueryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keyword | [string](#string) |  |  |
| page | [uint64](#uint64) |  |  |
| page_size | [uint64](#uint64) |  |  |






<a name="user.UserQueryResponse"></a>

### UserQueryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pager | [common.Pager](#common.Pager) |  |  |
| user | [User](#user.User) | repeated |  |





 

 

 


<a name="user.UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [UserCreateRequest](#user.UserCreateRequest) | [UserCreateResponse](#user.UserCreateResponse) |  |
| UserLogin | [UserLoginRequest](#user.UserLoginRequest) | [UserLoginResponse](#user.UserLoginResponse) |  |
| GetUser | [UserGetRequest](#user.UserGetRequest) | [UserGetResponse](#user.UserGetResponse) |  |
| QueryUser | [UserQueryRequest](#user.UserQueryRequest) | [UserQueryResponse](#user.UserQueryResponse) |  |
| UpdatePassword | [UpdatePasswordRequest](#user.UpdatePasswordRequest) | [UpdatePasswordResponse](#user.UpdatePasswordResponse) |  |
| UpdateAvatar | [UpdateAvatarRequest](#user.UpdateAvatarRequest) | [UpdateAvatarResponse](#user.UpdateAvatarResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

