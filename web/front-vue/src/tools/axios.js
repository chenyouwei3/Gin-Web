const VueAxios = {
    vm:{},//存储vue实例

    //app(vue3实例) install(传入的axios实例)
    install(app, instance) {
      // 如果插件已经安装，则返回
      if (this.install){
        return
      }
      this.install=true//标记插件已安装
       // 如果没有传入 axios 实例，打印错误信息
      if (!instance) {
        console.error('You have to install axios')
        return
      }
    // 使用 provide 注入 axios 实例到 Vue 组件中
    //app.config.globalProperties 是vue3注册全局变量的方式
    app.config.globalProperties.axios = instance; // 将 axios 实例注入到 Vue 组件中
    // 通过 this.$http 访问 axios 实例
    app.provide('$http', instance);
    }
  }
  
export { VueAxios };   //暴露组件