今天遇到一个奇葩问题，同样的代码拿回家就正常在公司就不正常，排查发现mysql存数据出了问题，原来是用到外键另一张表没有本张表外键指定值，还是要细心