# synk

synk是一个基于lorca，gin，react构建的一个局域网互传文件应用

> 在你的隔空投送因为莫名原因不好用的时候，又不想借助qq/微信等聊天软件，那么不如自己实现一个隔空投送

### 演示

> 主页

![image-20220820180008406](http://octane.oss-cn-beijing.aliyuncs.com/img/image-20220820180008406.png)

> 选择电脑的ip

![image-20220820180435463](http://octane.oss-cn-beijing.aliyuncs.com/img/image-20220820180435463.png)

> 选择一张图片

![image-20220820180810971](http://octane.oss-cn-beijing.aliyuncs.com/img/image-20220820180810971.png)

> 使用手机扫码

![image-20220820180447488](http://octane.oss-cn-beijing.aliyuncs.com/img/image-20220820180447488.png)

> 在手机上下载这张图片

<img src="http://octane.oss-cn-beijing.aliyuncs.com/img/image-20220820180703996.png" alt="image-20220820180703996" style="zoom:30%;" />

### 使用

1. 将本项目克隆到本地

```sh
git clone https://github.com/Octane0411/synk.git
或
git clone git@github.com:Octane0411/synk.git
```

2. 自行编译

```sh
cd synk
CGO_ENABLED=0 GOOS=「windows/darwin/linux」 GOARCH=「amd64/arm64」 go build main.go
cd server/frontend
yarn
yarn build
cd -
```

3. 运行

```sh
./synk
```

### 实现思路

#### 使用lorca创建出一个窗口

> 为什么不用electron? 因为我不会

我了解到 Go 的如下库可以实现窗口：

1. [lorca](https://github.com/zserge/lorca) - 调用系统现有的 Chromium 实现简单的窗口，UI 通过前端技术实现
2. [webview](https://github.com/webview/webview) - 比 lorca 功能更强，实现 UI 的思路差不多
3. [fyne](https://github.com/fyne-io/fyne) - 使用 Canvas 绘制的 UI 框架，性能不错
4. [qt](https://github.com/therecipe/qt) - 更复杂更强大的 UI 框架

只有lorca没什么学习成本，故使用lorca

#### 用 HTML/CSS/JS 制作 UI

我用 React + ReactRouter 来实现页面结构，文件上传和对话框是使用原生 JS 写的，没有依赖其他 UI 组件库。

lorca 的主要功能就是展示我写出来的 index.html。

#### 用gin实现后端接口

index.html 中的 JS 用到了五个接口，我使用 gin 来实现：

```
router.GET("/uploads/:path", controllers.UploadsController)              
router.GET("/api/v1/addresses", controllers.AddressesController) 
router.GET("/api/v1/qrcodes", controllers.QrcodesController)   
router.POST("/api/v1/files", controllers.FilesController)      
router.POST("/api/v1/texts", controllers.TextsController)
```

其中的二维码生成用到了 [go-qrcode](https://github.com/skip2/go-qrcode)。

#### 用 [gorilla/websocket](https://github.com/gorilla/websocket) 实现手机通知 PC

这个库提供了几个example，看完几个example就大概明白该怎么做了

#### 整体思路

总得来说：

1. 用 Lorca 搞出一个窗口
2. 用 HTML 制作界面，用 JS 调用后台接口
3. 用 Gin 实现后台接口
4. 上传的文件都放到 uploads 文件夹中，为文件生成uuid以区分

