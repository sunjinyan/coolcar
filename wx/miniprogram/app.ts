import { getUserInfo, getUserSetting } from "./utils/util"
// app.ts

let  resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void
let  rejectUserInfo: (reason?: any) => void
App<IAppOption>({
  globalData: {
    userInfo: new Promise<WechatMiniprogram.UserInfo>((resolve,reject)=>{//为了保持网络信息同步返回，页面与所需数据顺序加载，所以需要将结果使用Promise来进行操作
      resolveUserInfo = resolve
      rejectUserInfo = reject
    })
    //起初最开始的思路
    // userInfo: new Promise((resolve,reject)=>{//为了保持网络信息同步返回，页面与所需数据顺序加载，所以需要将结果使用Promise来进行操作
    //   getUserSetting().then(res=>{
    //     if (res.authSetting['scope.userInfo']){
    //       return getUserInfo()
    //     }
    //     return Promise.resolve(undefined)
    //   }).then(res => {
    //     if (!res) {
    //       return
    //     }
    //     // this.globalData.userInfo = res?.userInfo
    //     // if (this.userInfoReadyCallback) {
    //     //   this.userInfoReadyCallback(res)
    //     // }
    //     // resolve(res.userInfo)
    //     return resolve
    //   }).catch(reject)
    // })
  },
  async onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })


    //将最初的globalData中的userInfo思路修改为使用 resolveUserInfo = resolve，rejectUserInfo = reject暴露给外部去处理
    try {
      const setting = await  getUserSetting()
      if (setting.authSetting['scope.userInfo']) {
        const  userInfoRes = await getUserInfo()
        resolveUserInfo(userInfoRes.userInfo)
      }
    } catch (error) {
      rejectUserInfo(error)
    }
    // getUserSetting().then(res=>{
    //   if (res.authSetting['scope.userInfo']){
    //     return getUserInfo()
    //   }
    //   return Promise.resolve(undefined)
    // }).then(res => {
    //   if (!res) {
    //     return
    //   }
    //   // this.globalData.userInfo = res?.userInfo
    //   // if (this.userInfoReadyCallback) {
    //   //   this.userInfoReadyCallback(res)
    //   // }
    //   // resolve(res.userInfo)
    //   return resolveUserInfo
    // }).catch(rejectUserInfo)
    
  },
  resolveUserInfo(userInfo: WechatMiniprogram.UserInfo){
    //不希望使用 export将resolveUserInfo暴露出去，提供函数调用的方式
    resolveUserInfo(userInfo)
  },
  rejectUserInfo(userInfo: WechatMiniprogram.UserInfo){
    //不希望使用 export将resolveUserInfo暴露出去，提供函数调用的方式
    resolveUserInfo(userInfo)
  },
})