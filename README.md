# sylas - 批量获取 fofa 数据

> 根据多个不同地区进行聚合查询以获取更多 fofa 数据量的查询。  

## 0x00 预置条件

因为是动态爬虫，所以在使用前需要已经拥有 Chrome 浏览器。

## 0x01 注意事项

一定要确保语法正确，Windows 系统上使用尤其要注意转义的问题，具体可以参见 [Windows 系统的使用问题](https://fofax.xiecat.fun/faq/#windows-%E7%B3%BB%E7%BB%9F%E7%9A%84%E4%BD%BF%E7%94%A8%E9%97%AE%E9%A2%98)。为保证正确使用，可以先在 [fofax](https://github.com/xiecat/fofax/) 下测试通过，再使用 sylas 查询。

fofa 查询语句尽量少，不要带无关字段。

## 0x02 下载地址

点击 [Releases下载链接](https://github.com/xiecat/sylas/releases) ，按照自己的系统架构选择相应的发行版本下载。

## 0x03 使用方法

### 帮助信息

```bash
sylas -h

   _____       __
  / ___/__  __/ /___ ______
  \__ \/ / / / / __  / ___/
 ___/ / /_/ / / /_/ (__  )
/____/\__, /_/\__,_/____/
     /____/  0.0.1

	 sylas

Usage of sylas:
  -debug
    	debug
  -h	show help
  -p string
    	chrome path
  -q string
    	fofa query
```

### macOS/Linux

使用 `-q` 参数指定查询语句， 查询的结果会返回在终端（如下部分结果被省略）。

```bash
sylas -q 'app="宝塔-Linux控制面板"'

   _____       __
  / ___/__  __/ /___ ______
  \__ \/ / / / / __  / ___/
 ___/ / /_/ / / /_/ (__  )
/____/\__, /_/\__,_/____/
     /____/  0.0.1

	 sylas

2022/01/09 12:20:52 query:app="宝塔-Linux控制面板",base64:YXBwPSLlrp3loZQtTGludXjmjqfliLbpnaLmnb8i
2022/01/09 12:20:52 fofa Query URL:https://fofa.so/result?qbase64=YXBwPSLlrp3loZQtTGludXjmjqfliLbpnaLmnb8i
2022/01/09 12:20:56 request fofa query :app="宝塔-Linux控制面板" stat data
2022/01/09 12:20:56 get fofa query :app="宝塔-Linux控制面板" stat data
2022/01/09 12:21:04 query:app="宝塔-Linux控制面板" && country="GB",base64:YXBwPSLlrp3loZQtTGludXjmjqfliLbpnaLmnb8iICYmIGNvdW50cnk9IkdCIg==
2022/01/09 12:21:04 fofa Query URL:https://fofa.so/result?qbase64=YXBwPSLlrp3loZQtTGludXjmjqfliLbpnaLmnb8iICYmIGNvdW50cnk9IkdCIg==
2022/01/09 12:21:06 request fofa query :app="宝塔-Linux控制面板" && country="GB" stat data
2022/01/09 12:21:06 get fofa query :app="宝塔-Linux控制面板" && country="GB" stat data
2022/01/09 12:21:07 query:app="宝塔-Linux控制面板" && country="DE",base64:YXBwPSLlrp3loZQtTGludXjmjqfliLbpnaLmnb8iICYmIGNvdW50cnk9IkRFIg==
2022/01/09 12:21:07 fofa Query URL:https://fofa.so/result?qbase64=YXBwPSLlrp3loZQtTGludXjmjqfliLbpnaLmnb8iICYmIGNvdW50cnk9IkRFIg==
2022/01/09 12:21:09 request fofa query :app="宝塔-Linux控制面板" && country="DE" stat data
2022/01/09 12:21:11 get fofa query :app="宝塔-Linux控制面板" && country="DE" stat data
(region="HK" && app="宝塔-Linux控制面板" && region="Sha Tin")||(country="US" && app="宝塔-Linux控制面板" && region="Virginia")||(country="CN" && app="宝塔-Linux控制面板" && region="Shaanxi")||(app="宝塔-Linux控制面板" && country="AE")
country="US" && app="宝塔-Linux控制面板" && region="California"
country="US" && app="宝塔-Linux控制面板" && region="North Carolina"
country="CN" && app="宝塔-Linux控制面板" && region="Zhejiang"
......
2022/01/09 12:21:11

The number of queries in the app="宝塔-Linux控制面板" is  354829
```

也可以使用重定向将查询结果输入至一个文本文件，后期便可以使用 fofax 通过指定文本的方式进行查询。

```bash
sylas -q 'app="宝塔-Linux控制面板"' > bts.txt
```

查看此文本，大约有 56 条语句

```bash
wc -l bts.txt
56 bts.txt
```

使用 fofax 进行查询。

```bash
fofax -qf bts.txt -fs 10000 > bts_ip.txt
```

### Windows

Windows 系统对引号不太友好，需要转义。

#### CMD

```cmd
sylas.exe -q app=\"APACHE-Solr\"
```

#### PowerShell

```powershell
./sylas.exe -q 'app=\"APACHE-Solr\"'
```
