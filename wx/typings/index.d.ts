/// <reference path="./types/index.d.ts" />

interface IAppOption {
  globalData: {
    userInfo: Promise<WechatMiniprogram.UserInfo>,//为了保持同步，所以需要将结果使用Promise来进行操作
  }
  userInfoReadyCallback?: WechatMiniprogram.GetUserInfoSuccessCallback,
  resolveUserInfo(userInfo: WechatMiniprogram.UserInfo):void//将此处接口与app里实现的保持统一，app里需要这种，所以在这里定义
  rejectUserInfo(userInfo: WechatMiniprogram.UserInfo):void
}