# go的学习代码






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

#### 修改git提交时的config中的usernaem
* `vi ~/.gitconfig`; 然后在文件中直接修改即可