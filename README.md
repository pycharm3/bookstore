# goweb

* 今天遇到一个奇葩问题，同样的代码拿回家就正常在公司就不正常，排查发现mysql存数据出了问题，原来是用到外键另一张表没有本张表外键指定值，还是要细心
---
* 2019/12/7
* 今天遇到个奇葩问题，
---
* 2019/12/9 上午9:42
* 真的是沙雕了，上周做了个东西，去数据库判断有没有这个记录，有就加一，没有就创建，我每次执行都重新创建了一条，百思不得其解，数据库里原有几条一样的记录，
这周来，突然想到会不会已经有两条相同的数据，造成混淆无法正确判断，于是进行测试，果然是数据库数据混淆造成的逻辑处理错误，mysql真是掌握太弱鸡了
* 2019/12/9 下午16:10
* 有返回error的地方尽量不要忽略掉，最好加上，包括打印也是
---
* 2019/12/11 晚上18:15
* 完成功能:
* 实现图书加入购物车，从购物车删除图书
* 遗留bug:
* 登录后翻页再访问index页面由于不带登录请求信息无法判断是否登录，所以翻页后页头显示的是未登录状态，其实已经登录了
---
* 2019/12/13 晚上18:50
* 完成功能:
* 完成发货收货等功能，系统基本上算是完成了
* 遗留bug：
* 翻页后不能显示已登录用户
* 日有所思:
* 11月25号开始学习gohttp编程，今天是12月12号，差不多两周时间，算是完成了一个经典练习项目，前半段跟着教学视频做，后半段在数据结构设计思想成熟基础上按自己思路完成部分功能，也学习到很多编程知识和思想
* 现在已经12月了，我也将在这个月底离职，工作还没有着落，今年这个年也是不太顺利，不过问题不大，总之努力奋斗吧，祝自己越来越好，加油gogogoYe
