//1. VSCODE插件一次安装
ctrl+shift+p => Go:Install/Update Tools =>全选17-18个工具

// 2. 自动生成函数的测试用例
#1）选择编写好的函数
#2）ctrl+shift+p => Go:Generate Unit Tests For Function
#2）ctrl+shift+p => Go:Generate Unit Tests For File
#3) ctrl+shift+p => Go:Fill struct    //自动填充用例中的结构体字段

#4) t.Log("LOG INFO") //测试用例中日志
#5) Go插件中，settings.json: go.buildFlags 设置-v参数， 也可以在vscode GO插件中添加


// 3. 自动实现接口
#1）先定义好User接口、对应的方法、注释
#2）定义结构体Student Name和Age
#3）ctrl+shift+p => Go:Generate Interface Stubs => s *Student pkgname.User

// 4. 结构体实现自动tag （json yaml xml 。。）
#1）选中结构体Student
#2）增加tag： ctrl+shift+p => Go:Add Tags to Struct Field
#3）删除tag： ctrl+shift+p => Go:Remove Tags From Struct Field

#4) settings.json: go.addTags 里tags增加json,yaml,xml等
#5)transform: snakecase/camelcase

//5. 查询接口的具体实现类
#1) 选中某个interface，User，右键Go to Implementations 

//6重构字段或方法
#1）选择字段，F2 修改，可实现统一修改

//7 表达式抽取成变量 多行抽取成function
#1) 变量：ctrl+shift+p => Go:Extract to variable
#2) 函数：ctrl+shift+p => Go:Extract to Function

//8 添加第三方包代码到工程
#1) ctrl+shift+p => Go:Add Package to Workspace

//9 保存的时候插件 go test  go format go vet go lint
#1) 在vscode GO插件 输入 save
#2）自定义右键命令settings.json: go.editorContextMenuCommands
