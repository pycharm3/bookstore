# goweb

* 今天遇到一个奇葩问题，同样的代码拿回家就正常在公司就不正常，排查发现mysql存数据出了问题，原来是用到外键另一张表没有本张表外键指定值，还是要细心
---
* 2019/12/7
* 今天遇到个奇葩问题，
---
* 2019/12/9 上午9：42
* 真的是沙雕了，上周做了个东西，去数据库判断有没有这个记录，有就加一，没有就创建，我每次执行都重新创建了一条，百思不得其解，数据库里原有几条一样的记录，
这周来，突然想到会不会已经有两条相同的数据，造成混淆无法正确判断，于是进行测试，果然是数据库数据混淆造成的逻辑处理错误，mysql真是掌握太弱鸡了
* 2019/12/9 下午
* 打印一定要加不要
