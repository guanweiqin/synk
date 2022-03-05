## 项目实例

### 手机电脑文件互传

**使用概览：**

`文件运行前提路径下浏览器存在：`

`C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe`

1. 拉取项目
```bash
git clone https://github.com/guanweiqin/synk.git
```

2. 进入 server/frontend 目录
```bash
npm run build
```

3. 进入 synk 根目录
```bash
go build -ldflags "-H windowsgui"
```

## 关于 Synk 说明

此项目根据 [FrankFang/synk](https://github.com/FrankFang/synk) 项目进行学习搭建并做了一些改动。
