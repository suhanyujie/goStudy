# go的学习代码

## 项目list
- [x] 简单爬虫 [点此前往](/http)
- [ ] 数据结构 [点此前往](/dataStructure/binaryTree)  进行中
- [ ] 排序算法 [点此前往](/studySort)  进行中
- [ ] fastcgi学习 [点此前往](/fastCgiStudy)  进行中


## 遇到的问题

### git提交

#### 使用git pull提示refusing to merge unrelated histories

```html
在执行git pull的时候,提示
`fatal: refusing to merge unrelated histories`
解决方法:
git pull --allow-unrelated-histories
```
* 参考资料 https://www.jianshu.com/p/39b890d6e73d

#### 修改git提交时的config中的username
* `vi ~/.gitconfig`; 然后在文件中直接修改即可