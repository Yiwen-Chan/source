---
title: springboot连接phoenix的几种方法
date: 2021-07-15 14:23:48
top_img: https://pic3.zhimg.com/80/v2-5a6b5833d6e69a064f1b6cd48d19d4ec_720w.png
cover: https://pic3.zhimg.com/80/v2-5a6b5833d6e69a064f1b6cd48d19d4ec_720w.png
tags: 
    - springboot
    - phoenix
    - mybatis
---

虽说是介绍`springboot`连接phoenix的方法，其实因为用的是`jdbc`连接方式，其实只要是通过`jdbc`方式连接的数据库都可以通用，比如 `hive`。

### 1.原生方法连接数据库

自己手写 `Connection` ，`Statement`，`PreparedStatement`，`ResultSet`，手写数据库连接和关闭，自己提交sql语句。

这种方法是最开始学习数据库连接的方法，后面连接数据库的方法都是在这样的基础上进行了封装和扩充。优点是能够了解数据库连接过程，但是缺陷更加明显，比如何时建立连接何时关闭，是否自动提交，还要自己维护线程池，非常麻烦。但是假如你需要在项目中`动态连接`多个数据库，大概就要采用这种方式了。除此之外，不建议采用这种纯原生的连接方式。

示例：[https://github.com/gitriver/alad-phoenix](https://github.com/gitriver/alad-phoenix)    想学习这种方法的可以看看这个。


### 2.JdbcTemplate
通过配置类中加载`DataSource`、`JdbcTemplate`来操作`phoenix`，实际上`JdbcTemplate`是`spring`对方法（1）的封装，但是可以省事，何乐而不为呢？

配置类：

```java
@Configuration
public class PhoenixDataSource {

    @Autowired
    private Environment env;

    @Bean(name = "phoenixJdbcDataSource")
    @Qualifier("phoenixJdbcDataSource")
    public DataSource dataSource() {
        DruidDataSource dataSource = new DruidDataSource();
        dataSource.setUrl(env.getProperty("phoenix.url"));
        dataSource.setDriverClassName(env.getProperty("phoenix.driver-class-name"));
        dataSource.setUsername(env.getProperty("phoenix.username"));//phoenix的用户名默认为空
        dataSource.setPassword(env.getProperty("phoenix.password"));//phoenix的密码默认为空
        dataSource.setDefaultAutoCommit(Boolean.valueOf(env.getProperty("phoenix.default-auto-commit")));
        dataSource.setConnectionProperties("phoenix.schema.isNamespaceMappingEnabled=true");

        return dataSource;
    }

    @Bean(name = "phoenixJdbcTemplate")
    public JdbcTemplate phoenixJdbcTemplate(@Qualifier("phoenixJdbcDataSource") DataSource dataSource) {
        return new JdbcTemplate(dataSource);
    }
}


```

Service层：

```java
@Service
public class CRUDServiceImp implements CRUDService {

    @Autowired
    @Qualifier("phoenixJdbcTemplate")
    JdbcTemplate phoenixJdbcTemplate;
    public Result add(){

        phoenixJdbcTemplate.update("upsert into data_provision.company(id,name,address) values('20','xuxiao','德国柏林')");

        return new Result(true,"数据添加成功");

    }

    public Result update(){
        int res = phoenixJdbcTemplate.update("upsert into data_provision.company(id,name) values('20','yyggg')");
        return new Result(true,"数据更新成功");

    }

    public Result delete(){
        phoenixJdbcTemplate.update("delete from data_provision.company where id ='20'");
        return new Result(true,"数据删除成功");

    }


    public List<Map<String, Object>> query(){
        return phoenixJdbcTemplate.queryForList("select * from data_provision.company");

    }

    public List<Map<String, Object>> tSql(String sql){
        return phoenixJdbcTemplate.queryForList(sql);
    }

    public int updateSql(String sql){
        return phoenixJdbcTemplate.update(sql);
    }
}


```

优点：方便，代码量不大。通过这种方式不需要实现dao层，假如不需要映射到实体类的话可以说是上选。

缺点：由于没有orm，所以想直接将数据库映射到实体类是做不到的。当然，你要自己实现也未尝不可，但是为什么要自己造轮子呢？

[JdbcTemplate实体映射](https://www.cnblogs.com/bener/p/10617065.html)

示例：[https://github.com/Gyoliu/phoenix-hbase](https://github.com/Gyoliu/phoenix-hbase)

###3.orm框架

通过现有的`orm`框架来进行连接，`hibernate`和`mybatis`都可以。

`hibernate`能够通过数据库类型来将hql语句转变成对应数据库的方言。虽然`hibernate`没有`phoenix`的方言，不过在github上有人制作了。我没试过，希望尝试过的人能和我分享一下使用心得。

[https://github.com/jruesga/phoenix-hibernate-dialect](https://github.com/jruesga/phoenix-hibernate-dialect)  （hibernate的phoenix方言制作）

`mybatis`，通过`jdbc`连接到`phoenix`，指定`phoenix`的驱动jar，就可以连接到了。`jdbc`连接既可以是`zookeeper`的地址( jdbc:phoenix:zookeeper )，也可以是`phoenix`的thin-connect( jdbc:phoenix:thin:url )。在使用轻连接之前需要先打开phoenix的轻连接。

通过`mybatis`连接不能指定数据库，连接的是默认数据库。所以在指定数据表的时候需要加上数据库，如（db.table）。也可以通过mybatis-plus的@TableName注解来指定数据表（如@TableName("DEV.TEST")），假如需要小写的话可以用双引号限定小写（如@TableName("\\"dev.test2\\"")）。

```
datasource:
    #数据库连接信息
    url: jdbc:phoenix:slaver01-robin,slaver02-robin,master-robin:2181
    username:
    password:
    driver-class-name: org.apache.phoenix.jdbc.PhoenixDriver     #驱动
    # 如果不想配置对数据库连接池做特殊配置的话,以下关于连接池的配置就不是必须的
    # spring-boot 2.X 默认采用高性能的 Hikari 作为连接池 更多配置可以参考 https://github.com/brettwooldridge/HikariCP#configuration-knobs-baby
    type: com.zaxxer.hikari.HikariDataSource


```

示例：[https://gitee.com/zhengshunzi/springboot-phoenix](https://gitee.com/zhengshunzi/springboot-phoenix)

[https://github.com/mlwise/springboot-mybatis-phoenix-demo](https://github.com/mlwise/springboot-mybatis-phoenix-demo)

优点和缺点都是`orm`框架自身的缺点，毕竟是可以商用的框架，在简单、易用、稳定、可移植性上比原生代码强太多了。在连接方式上可以说是最好的选择了。



-   连接异常

1.客户端命名空间映射未启用

具体报错为：

```
java.sql.SQLException: ERROR 726 (43M10):  Inconsistent namespace mapping properties. Cannot initiate connection as SYSTEM:CATALOG is found but client does not have phoenix.schema.isNamespaceMappingEnabled enabled


```

单看字面意思是说你的`hbase`配置中没有以下配置，很多博客也说是这个原因，确实也有可能，不过我看网上博客在分享如何安装`phoenix`的时候，没有一篇会把这个配置给漏掉，所以我认为因为`hbase`配置问题才连不上的是少数中的少数。

```
<property>
   <name>phoenix.schema.isNamespaceMappingEnabled</name>
   <value>true</value>
</property>


```

上面报错说是客户端没有将`phoenix.schema.isNamespaceMappingEnabled`设置为true。这里的客户端，指的是你的springboot项目,而不是hbase的客户端配置。

![](https://oscimg.oschina.net/oscnet/up-2adc027b056e80091f2687eb32e39175997.png)

假如你用的是`cdh` `hadoop`，在CM平台上把客户端高级配置加上命名空间映射是没用的。正确的解决办法有两种：

1）在连接池的配置中把`phoenix.schema.isNamespaceMappingEnabled`设置为true

这个操作可以在配置文件中完成，也可以在手动加载配置类中完成。但是通过这种方式，你的连接池必须为`Druid`而不是Springboot2默认的`Hikari`。

在`application.properties`中增加配置：

```markdown
spring.datasource.connectionProperties=phoenix.schema.isNamespaceMappingEnabled=true
```

或，手动加载配置类：

```java
@Bean(name = "phoenixJdbcDataSource")
    @Qualifier("phoenixJdbcDataSource")
    public DataSource dataSource() {
        DruidDataSource dataSource = new DruidDataSource();
        dataSource.setUrl(env.getProperty("phoenix.url"));
        dataSource.setDriverClassName(env.getProperty("phoenix.driver-class-name"));
        dataSource.setUsername(env.getProperty("phoenix.username"));//phoenix的用户名默认为空
        dataSource.setPassword(env.getProperty("phoenix.password"));//phoenix的密码默认为空
        dataSource.setDefaultAutoCommit(Boolean.valueOf(env.getProperty("phoenix.default-auto-commit")));
        dataSource.setConnectionProperties("phoenix.schema.isNamespaceMappingEnabled=true");

        return dataSource;
    }


```

配置文件：

```
# phoenix 数据源自定义配置
phoenix.enable= true
phoenix.url=jdbc:phoenix:192.168.49.180,192.168.49.181:2181
phoenix.type=com.alibaba.druid.pool.DruidDataSource
phoenix.driver-class-name=org.apache.phoenix.jdbc.PhoenixDriver
phoenix.username=
phoenix.password=
phoenix.default-auto-commit=true
phoenix.schema.isNamespaceMappingEnabled=true


```

2）在配置文件中新增一个`hbase-site.xml`，在这里加载配置（推荐）

![](https://oscimg.oschina.net/oscnet/up-0a9ba4a5b9e33223bf3a636d3d09b3374b5.png)

只需要加载这一个配置就可以了，没有必要和`hbase`的同步。这样也能用`Hikari`连接池。

另外附一个不那么正经的解决办法，就是通过`phoenix`的轻连接来连接`phoenix`，不会报命名空间映射未启用的错误。


```
spring.datasource.url=jdbc:phoenix:thin:url=http://phoenix:8765;serialization=PROTOBUF


```

[2.com/google/protobuf/LiteralByteString类问题](http://2.com/google/protobuf/LiteralByteString%E7%B1%BB%E9%97%AE%E9%A2%98)

这个问题大概只有我自己遇到了...不过也写给大家分享一下。最开始报错是java.lang.NoClassDefFoundError: com/google/protobuf/LiteralByteString，没有找到这个类。我选择通过idea寻找maven添加，就是这一步让我上了大当。idea给我的搜索结果是：

![](https://oscimg.oschina.net/oscnet/up-d694766e637a5d73ae1865c844890f8a9fc.png)

添加org.apache.hive:hive-exec:3.1.0。看上去没什么问题，添加上去之后也能找到这个类，但是运行起来报了另外一个错误：

```
VerifyError: class com.google.protobuf.LiteralByteString overrides final met


```

具体报错找不到了，大概意思是说类型转换错误。

思考了很久，通过idea的Dependency Analyzer对比jar包，才发现网上能够正常运行的项目，com.google.protobuf：protobuf-java这个包是2.5.0版本，而我的则是3.1.0。在pom中指定版本之后，错误就消失了。