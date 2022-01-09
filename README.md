# sylas - 批量获取 fofa 数据

根据不同城市进行聚合获取更多数据  

## 安装条件和使用条件

因是动态爬虫**需要 chrome 浏览器**

保证语法正确，并且不出问题。尤其注意 [win下转义](https://fofax.xiecat.fun/faq/#windows-%E7%B3%BB%E7%BB%9F%E7%9A%84%E4%BD%BF%E7%94%A8%E9%97%AE%E9%A2%98)。为保证正确可以现在 fofax 下测试通过, 再查询。fofa 查询语句尽量少。不要带无关字段

### Unix/Linux 下

```shell
sylas -q 'app="宝塔-Linux控制面板"'>result.txt
```

### win下未测试

cmd 下

```shell
sylas.exe -q app=\"APACHE-Solr\" >result.txt
```

PowerShell 下

```shell
sylas.exe -q -q 'app=\"APACHE-Solr\"' >result.txt
```

