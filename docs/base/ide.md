# IDE编辑器

工欲善其事必先利其器,命令行工具虽然能够在一定程度上满足基本操作的需求,但实际工作中总不能一直使用命令行工具进行编码操作吧?

学习 `Go` 语言同样如此,为此需要寻找一个强大的 `IDE` 集成环境帮助我们快速开发,据我所知,市面上比较流行的可能有三个选择:

- `LiteIDE X` : [`LiteIDE`](http://liteide.org/cn/) 是一款简单,开源,跨平台的 `Go IDE`.
- `GoLand` : [`GoLand`](https://www.jetbrains.com/go/) is a cross-platform IDE built specially for Go developers.
- `第三方插件` : `Idea`, `Sublime Text`,`VS Code` ,等常见 `IDE` 一般均有 `Go` 的插件.

萝卜青菜各有所爱,选择哪个 `IDE` 都可以,甚至不用任何 `IDE` 也可以,不过还是推荐下 `GoLand` 吧!

![go-base-ide-goland.png](../images/go-base-ide-goland.png)

## `Goland` 下载安装

> 官网地址: [https://www.jetbrains.com/go/](https://www.jetbrains.com/go/),如果无法访问,可能需要**特殊手段**绕过.

`Goland` 并不像 `Idea` 那样分为专业版和社区版,目前只有收费版,提供 `30` 天免费试用,试用到期后可以选择购买正版也可以上淘宝购买激活码或网上寻求破解版等等.

直接点击页面中间的 `Download` 按钮后就会自动识别当前系统进行下载,也可以点击右上角的 `Download` 按钮自行选择目标平台进行下载.

![go-base-ide-goland-download.png](../images/go-base-ide-goland-download.png)

安装过程比较简单,这里就不再赘述,简单的动图一闪而过看下大致过程吧!

![go-base-ide-goland-windows-install.gif](../images/go-base-ide-goland-windows-install.gif)

其中安装位置,默认是 `C` 盘,可以自行选择合适的安装位置.

![go-base-ide-goland-windows-install-location.png](../images/go-base-ide-goland-windows-install-location.png)

如果是 `Mac` 电脑,安装 `Goland` 更为简单,直接下载拖动到 `Application` 分类,连安装目录都不用选择,简单演示如下:

![go-base-ide-goland-mac-install.png](../images/go-base-ide-goland-mac-install.png)

## `Goland` 克隆项目

初次打开 `Goland` 编辑器,界面出现三个选项:

- `New Project` : 新建项目,适合从零开发新项目
- `Open Project` : 打开项目,适合本地已存在 `Go` 项目
- `Check out from Version Control` : 从版本库中检出项目,适合团队合作时直接从线上项目下载到本地.

![go-base-ide-goland-open.png](../images/go-base-ide-goland-open.png)

三种方式分别对应三种不同的场景,这里选择以第三种方式检出版本库为例,目录源码: [https://github.com/snowdreams1006/learn-go](https://github.com/snowdreams1006/learn-go)

选择 `git` 版本库,并填写项目地址,然后点击右侧的 `Test` 按钮,如果提示失败,可能是 `Git` 基本环境没有配置过,请先配置下 `Git`,可以参考 [git 入门教程
](https://snowdreams1006.github.io/git/)

![go-base-ide-goland-git-clone-test-succes.png](../images/go-base-ide-goland-git-clone-test-succes.png)

> 项目地址: `git@github.com:snowdreams1006/learn-go.git` 或者 `https://github.com/snowdreams1006/learn-go.git` 或者 `https://github.com/snowdreams1006/learn-go`

默认情况下,本地目录是 `GolandProjects`,一定要修改成自己的 `GOPATH` 目录,即 `USERPROFILE/go` 目录.

![go-base-ide-goland-git-clone-change-directory.png](../images/go-base-ide-goland-git-clone-change-directory.png)

耐心等待,`Goland` 会自动下载项目相关依赖,右下角的进度条完毕后意味着项目初始化好了,可以正常工作了.

## `Goland` 打开项目

找到 `hello/hello.go` 文件,其中 `main` 方法左侧有个绿色的启动按钮,点击运行.

![go-base-ide-goland-open-project-main.png](../images/go-base-ide-goland-open-project-main.png)

初始运行,提示配置进行命令配置,设置工作目录为 `GOPATH` 环境变量所在的目录.

![go-base-ide-goland-project-configure.png](../images/go-base-ide-goland-project-configure.png)

此时配置页面左下角的红色报错消失了,保存后关闭该窗口,再次运行 `main` 方法,如我们所愿输出了 `Hello Go!` 的逆序.

![go-base-ide-goland-project-run.png](../images/go-base-ide-goland-project-run.png)

现在打开 `strings/reverse_test.go` 文件,同样点击左侧的启动按钮运行测试文件,证明测试运行正常!

![go-base-ide-goland-project-test.png](../images/go-base-ide-goland-project-test.png)

## `Goland` 配置总结

万事开头难,下载 `Goland` 并初始化项目,其实很简单,之所以特意写下这篇文章主要是为了克服陌生的恐惧,迈出第一步就会有第二步,接下来的 `Go` 语言学习之旅就可以顺利开始了,`Go to Work !`

![go-base-ide-work.png](../images/go-base-ide-work.png)
